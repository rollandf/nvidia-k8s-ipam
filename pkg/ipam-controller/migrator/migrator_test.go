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

package migrator_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Mellanox/nvidia-k8s-ipam/pkg/ipam-controller/config"
	"github.com/Mellanox/nvidia-k8s-ipam/pkg/ipam-controller/migrator"
	"github.com/Mellanox/nvidia-k8s-ipam/pkg/pool"

	ipamv1alpha1 "github.com/Mellanox/nvidia-k8s-ipam/api/v1alpha1"
)

const (
	pool1Name = "pool1"
	pool2Name = "pool2"
)

func updateConfigMap(data string) {
	d := map[string]string{config.ConfigMapKey: data}
	err := k8sClient.Create(ctx, &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{Name: TestConfigMapName, Namespace: TestNamespace},
		Data:       d,
	})
	if err == nil {
		return
	}
	if apiErrors.IsAlreadyExists(err) {
		configMap := &corev1.ConfigMap{}
		Expect(k8sClient.Get(
			ctx, types.NamespacedName{Name: TestConfigMapName, Namespace: TestNamespace}, configMap)).NotTo(HaveOccurred())
		configMap.Data = d
		Expect(k8sClient.Update(
			ctx, configMap)).NotTo(HaveOccurred())
	} else {
		Expect(err).NotTo(HaveOccurred())
	}
}

var validConfig = fmt.Sprintf(`
    {
      "pools": {
        "%s": { "subnet": "192.168.0.0/16", "perNodeBlockSize": 10 , "gateway": "192.168.0.1"},
        "%s": { "subnet": "172.16.0.0/16", "perNodeBlockSize": 50 , "gateway": "172.16.0.1"}
      }
    }
`, pool1Name, pool2Name)

func createNode(name string) *corev1.Node {
	node := &corev1.Node{ObjectMeta: metav1.ObjectMeta{Name: name}}
	Expect(k8sClient.Create(ctx, node)).NotTo(HaveOccurred())
	return node
}

func getNode(name string) *corev1.Node {
	node := &corev1.Node{}
	Expect(k8sClient.Get(ctx, types.NamespacedName{Name: name}, node)).NotTo(HaveOccurred())
	return node
}

func updateNode(node *corev1.Node) *corev1.Node {
	Expect(k8sClient.Update(ctx, node)).NotTo(HaveOccurred())
	return node
}

func getRangeFromNode(nodeName string) map[string]*pool.IPPool {
	node := getNode(nodeName)
	poolCfg, err := pool.NewConfigReader(node)
	if err != nil {
		return nil
	}
	return poolCfg.GetPools()
}

// WaitAndCheckForStability wait for condition and then check it is stable for 1 second
func WaitAndCheckForStability(check func(g Gomega), wait interface{}, stability interface{}) {
	Eventually(func(g Gomega) { check(g) }, wait).Should(Succeed())
	Consistently(func(g Gomega) { check(g) }, stability).Should(Succeed())
}

var _ = Describe("Controller Migrator", func() {
	It("Basic tests", func() {
		testNode1 := "node1"
		testNode2 := "node2"

		By("Create valid cfg1")
		updateConfigMap(validConfig)

		By("Set annotation with valid ranges for node1")
		node1 := createNode(testNode1)
		node1InitialRanges := map[string]*pool.IPPool{pool1Name: {
			Name:    pool1Name,
			Subnet:  "192.168.0.0/16",
			StartIP: "192.168.0.11",
			EndIP:   "192.168.0.20",
			Gateway: "192.168.0.1",
		}, pool2Name: {
			Name:    pool2Name,
			Subnet:  "172.16.0.0/16",
			StartIP: "172.16.0.1",
			EndIP:   "172.16.0.50",
			Gateway: "172.16.0.1",
		}}
		Expect(pool.SetIPBlockAnnotation(node1, node1InitialRanges)).NotTo(HaveOccurred())
		Expect(updateNode(node1))

		By("Set annotation with valid ranges for node2")
		node2 := createNode(testNode2)
		node2InitialRanges := map[string]*pool.IPPool{pool1Name: {
			Name:    pool1Name,
			Subnet:  "192.168.0.0/16",
			StartIP: "192.168.0.21",
			EndIP:   "192.168.0.30",
			Gateway: "192.168.0.1",
		}, pool2Name: {
			Name:    pool2Name,
			Subnet:  "172.16.0.0/16",
			StartIP: "172.16.0.51",
			EndIP:   "172.16.0.100",
			Gateway: "172.16.0.1",
		}}
		Expect(pool.SetIPBlockAnnotation(node2, node2InitialRanges)).NotTo(HaveOccurred())
		Expect(updateNode(node2))

		By("Run migrator")
		Expect(migrator.Migrate(ctx, k8sClient, TestConfigMapName, TestNamespace, TestNamespace)).NotTo(HaveOccurred())

		By("Verify Pool1 Spec")
		pool1 := &ipamv1alpha1.IPPool{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{Namespace: TestNamespace, Name: pool1Name}, pool1)).NotTo(HaveOccurred())
		Expect(pool1.Spec.Gateway == "192.168.0.1" &&
			pool1.Spec.Subnet == "192.168.0.0/16" && pool1.Spec.PerNodeBlockSize == 10).To(BeTrue())

		By("Verify Pool1 Allocations")
		expectedAllocations := []ipamv1alpha1.Allocation{{NodeName: testNode1, StartIP: "192.168.0.11", EndIP: "192.168.0.20"},
			{NodeName: testNode2, StartIP: "192.168.0.21", EndIP: "192.168.0.30"}}
		Expect(expectedAllocations).To(BeEquivalentTo(pool1.Status.Allocations))

		By("Verify Pool2 Spec")
		pool2 := &ipamv1alpha1.IPPool{}
		Expect(k8sClient.Get(ctx, types.NamespacedName{Namespace: TestNamespace, Name: pool2Name}, pool2)).NotTo(HaveOccurred())
		Expect(pool2.Spec.Gateway == "172.16.0.1" &&
			pool2.Spec.Subnet == "172.16.0.0/16" && pool2.Spec.PerNodeBlockSize == 50).To(BeTrue())

		By("Verify Pool2 Allocations")
		expectedAllocations = []ipamv1alpha1.Allocation{{NodeName: testNode1, StartIP: "192.168.0.11", EndIP: "192.168.0.20"},
			{NodeName: testNode2, StartIP: "192.168.0.21", EndIP: "192.168.0.30"}}
		Expect(expectedAllocations).To(BeEquivalentTo(pool1.Status.Allocations))

		By("Verify Nodes annotations are removed")
		Expect(getRangeFromNode(testNode1)).To(BeEmpty())
		Expect(getRangeFromNode(testNode2)).To(BeEmpty())

		By("Verify Config Map is deleted")
		configMap := &corev1.ConfigMap{}
		Expect(k8sClient.Get(
			ctx, types.NamespacedName{Name: TestConfigMapName, Namespace: TestNamespace}, configMap)).To(HaveOccurred())
	})
})
