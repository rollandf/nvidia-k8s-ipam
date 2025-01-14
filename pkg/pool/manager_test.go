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

package pool_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"

	"github.com/Mellanox/nvidia-k8s-ipam/pkg/pool"
)

var _ = Describe("Manager", func() {
	It("Update pool data", func() {
		testPools := make(map[string]*pool.IPPool)
		testPoolName := "my-pool-1"
		testPools[testPoolName] = &pool.IPPool{
			Name:    "my-pool-1",
			Subnet:  "192.168.0.0/16",
			StartIP: "192.168.0.2",
			EndIP:   "192.168.0.254",
			Gateway: "192.168.0.1",
		}
		node := &corev1.Node{}
		Expect(pool.SetIPBlockAnnotation(node, testPools)).NotTo(HaveOccurred())

		mgr := pool.NewManager()
		Expect(mgr.GetPoolByName(testPoolName)).To(BeNil())
		Expect(mgr.Update(node)).NotTo(HaveOccurred())
		Expect(mgr.GetPoolByName(testPoolName)).NotTo(BeNil())
		Expect(mgr.GetPools()).To(HaveLen(1))
		mgr.Reset()
		Expect(mgr.GetPoolByName(testPoolName)).To(BeNil())
	})
})
