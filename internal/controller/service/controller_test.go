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

package service

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	virtualmachinev1 "kubevirt.io/kubevirtbmc/api/v1alpha1"
	ctlvirtualmachinebmc "kubevirt.io/kubevirtbmc/internal/controller/virtualmachinebmc"
)

var _ = Describe("Service Controller", func() {
	const (
		testVirtualMachineBMCName      = "default-test-vm"
		testVirtualMachineBMCNamespace = "kubevirtbmc-system"
		testUsername                   = "test-username"
		testPassword                   = "test-password"
		testVMName                     = "test-vm"
		testVMNamespace                = "default"
		testClusterIP                  = "10.0.0.100"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When clusterIP for the Service is ready", func() {
		It("Should update the VirtualMachineBMC status", func() {
			ctx := context.Background()

			// we need to create the namespace in the cluster first
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{Name: testVirtualMachineBMCNamespace},
			}
			Expect(k8sClient.Create(ctx, ns)).Should(Succeed())

			By("Creating a new VirtualMachineBMC")
			virtualMachineBMC := &virtualmachinev1.VirtualMachineBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testVirtualMachineBMCName,
					Namespace: testVirtualMachineBMCNamespace,
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "virtualmachine.kubevirt.io/v1alpha1",
					Kind:       "VirtualMachineBMC",
				},
				Spec: virtualmachinev1.VirtualMachineBMCSpec{
					Username:                testUsername,
					Password:                testPassword,
					VirtualMachineNamespace: testVMNamespace,
					VirtualMachineName:      testVMName,
				},
			}
			Expect(k8sClient.Create(ctx, virtualMachineBMC)).To(Succeed())

			By("Creating a new Service")
			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						ctlvirtualmachinebmc.VirtualMachineBMCNameLabel: testVirtualMachineBMCName,
					},
					Name:      testVirtualMachineBMCName,
					Namespace: testVirtualMachineBMCNamespace,
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "v1",
					Kind:       "Service",
				},
				Spec: corev1.ServiceSpec{
					ClusterIP: testClusterIP,
					Ports: []corev1.ServicePort{
						{
							Port: ctlvirtualmachinebmc.IPMISvcPort,
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, svc)).To(Succeed())

			By("Checking that the VirtualMachineBMC has ServiceIP reflected")
			virtualMachineBMCLookupKey := types.NamespacedName{Name: virtualMachineBMC.Name, Namespace: virtualMachineBMC.Namespace}
			updatedVirtualMachineBMC := &virtualmachinev1.VirtualMachineBMC{}

			Eventually(func() bool {
				if err := k8sClient.Get(ctx, virtualMachineBMCLookupKey, updatedVirtualMachineBMC); err != nil {
					return false
				}
				return updatedVirtualMachineBMC.Status.ServiceIP == testClusterIP
			}, timeout, interval).Should(BeTrue())
		})
	})
})
