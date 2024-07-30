/*
Copyright 2024.

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

package virtualmachine

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"

	virtualmachinev1 "kubevirt.io/kubevirtbmc/api/v1"
)

var _ = Describe("VirtualMachine Controller", func() {
	const (
		testVMName                     = "test-vm"
		testVMNamespace                = "default"
		testVirtualMachineBMCName      = "default-test-vm"
		testVirtualMachineBMCNamespace = "kubevirtbmc-system"
		testUsername                   = "test-username"
		testPassword                   = "test-password"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating a VirtualMachine", func() {
		It("Should create a VirtualMachineBMC", func() {
			ctx := context.Background()

			// we need to create the namespace in the cluster first
			ns := corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{Name: testVirtualMachineBMCNamespace},
			}
			Expect(k8sClient.Create(ctx, &ns)).Should(Succeed())

			By("Creating a new VirtualMachine")
			vm := &kubevirtv1.VirtualMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testVMName,
					Namespace: testVMNamespace,
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "kubevirt.io/v1",
					Kind:       "VirtualMachine",
				},
				Spec: kubevirtv1.VirtualMachineSpec{
					Template: &kubevirtv1.VirtualMachineInstanceTemplateSpec{},
				},
			}
			Expect(k8sClient.Create(ctx, vm)).To(Succeed())

			By("Checking that the VirtualMachineBMC is created")
			virtualMachineBMCLookupKey := types.NamespacedName{Name: testVirtualMachineBMCName, Namespace: testVirtualMachineBMCNamespace}
			createdVirtualMachineBMC := &virtualmachinev1.VirtualMachineBMC{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, virtualMachineBMCLookupKey, createdVirtualMachineBMC)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
