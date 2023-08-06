/*
 Copyright 2023, NVIDIA CORPORATION & AFFILIATES
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sort"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	corev1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	v1helper "k8s.io/component-helpers/scheduling/corev1"

	ipamv1alpha1 "github.com/Mellanox/nvidia-k8s-ipam/api/v1alpha1"
	"github.com/Mellanox/nvidia-k8s-ipam/pkg/ipam-controller/allocator"
	"github.com/Mellanox/nvidia-k8s-ipam/pkg/ipam-controller/config"
)

// PoolReconciler reconciles Pool objects
type PoolReconciler struct {
	Allocator      *allocator.Allocator
	PoolsNamespace string
	NodeEventCh    chan event.GenericEvent
	client.Client
	Scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile contains logic to sync Node objects
func (r *PoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLog := log.FromContext(ctx)
	if req.Namespace != r.PoolsNamespace {
		// this should never happen because of the watcher configuration of the manager from controller-runtime pkg
		return ctrl.Result{}, nil
	}

	pool := &ipamv1alpha1.IPPool{}
	err := r.Client.Get(ctx, req.NamespacedName, pool)
	if err != nil {
		if apiErrors.IsNotFound(err) {
			reqLog.Info("Pool not found, removing from Allocator")
			r.Allocator.RemovePool(ctx, req.Name)
			return ctrl.Result{}, nil
		}
		reqLog.Error(err, "failed to get Pool object from the cache")
		return ctrl.Result{}, err
	}
	reqLog.Info("Notification on Pool", "name", pool.Name)

	err = config.ValidatePool(pool.Name, pool.Spec.Subnet, pool.Spec.Gateway, pool.Spec.PerNodeBlockSize)
	if err != nil {
		return r.handleInvalidSpec(ctx, err, pool)
	}

	// already validated by Validate function
	_, subnet, _ := net.ParseCIDR(pool.Spec.Subnet)
	gateway := net.ParseIP(pool.Spec.Gateway)
	allocatorConfig := allocator.AllocationConfig{
		PoolName:         pool.Name,
		Subnet:           subnet,
		Gateway:          gateway,
		PerNodeBlockSize: pool.Spec.PerNodeBlockSize,
		NodeSelector:     pool.Spec.NodeSelector,
	}
	if !r.Allocator.IsPoolLoaded(pool.Name) {
		r.Allocator.ConfigureAndLoadAllocations(ctx, allocatorConfig, pool)
	} else {
		r.Allocator.Configure(ctx, allocatorConfig)
	}

	nodeList := &corev1.NodeList{}
	if err := r.Client.List(ctx, nodeList); err != nil {
		reqLog.Error(err, "failed to list Nodes")
		return ctrl.Result{}, err
	}
	selectedNodes := make([]*corev1.Node, 0)
	for i := range nodeList.Items {
		node := nodeList.Items[i]
		if pool.Spec.NodeSelector != nil {
			match, err := v1helper.MatchNodeSelectorTerms(&node, pool.Spec.NodeSelector)
			if err != nil {
				reqLog.Error(err, "failed to match Node", "node", node.Name)
				continue
			}
			if match {
				selectedNodes = append(selectedNodes, &node)
			} else {
				reqLog.Info("node doesn't match selector, ensure range is not allocated", "node", node.Name)
				r.Allocator.DeallocateFromPool(ctx, pool.Name, node.Name)
			}
		} else {
			selectedNodes = append(selectedNodes, &node)
		}
	}
	sort.Slice(selectedNodes, func(i, j int) bool {
		return selectedNodes[i].Name < selectedNodes[j].Name
	})

	allocations := make([]ipamv1alpha1.Allocation, 0)
	for i := range selectedNodes {
		node := selectedNodes[i]
		a, err := r.Allocator.AllocateFromPool(ctx, pool.Name, node.Name)
		if err != nil {
			if errors.Is(allocator.ErrNoFreeRanges, err) {
				msg := fmt.Sprintf("failed to allocate IPs on Node: %s", node.Name)
				reqLog.Error(err, msg)
				r.recorder.Event(pool, "Warning", "NoFreeRanges", msg)
				continue
			}
			return ctrl.Result{}, err
		}
		alloc := ipamv1alpha1.Allocation{
			NodeName: node.Name,
			StartIP:  a.StartIP,
			EndIP:    a.EndIP,
		}
		allocations = append(allocations, alloc)
	}
	pool.Status.Allocations = allocations
	if err := r.Status().Update(ctx, pool); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *PoolReconciler) handleInvalidSpec(ctx context.Context,
	err error, pool *ipamv1alpha1.IPPool) (reconcile.Result, error) {
	reqLog := log.FromContext(ctx)
	reqLog.Error(err, "invalid Pool Spec, clearing Status")
	pool.Status.Allocations = make([]ipamv1alpha1.Allocation, 0)
	if err2 := r.Status().Update(ctx, pool); err2 != nil {
		return ctrl.Result{}, err2
	}
	r.recorder.Event(pool, "Warning", "InvalidSpec", err.Error())
	return ctrl.Result{Requeue: false}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.recorder = mgr.GetEventRecorderFor("IPPoolController")
	return ctrl.NewControllerManagedBy(mgr).
		For(&ipamv1alpha1.IPPool{}).
		// catch notifications received through chan from Node controller
		Watches(&source.Channel{Source: r.NodeEventCh}, handler.Funcs{
			GenericFunc: func(e event.GenericEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Namespace: e.Object.GetNamespace(),
					Name:      e.Object.GetName(),
				}})
			}}).
		Complete(r)
}
