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

package kubebmc

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	virtualmachinev1 "zespre.com/kubebmc/api/v1"
)

var _ = Describe("KubeBMC Controller", func() {
	const (
		testKubeBMCName      = "default-test-vm"
		testKubeBMCNamespace = "kubebmc-system"
		testUsername         = "test-username"
		testPassword         = "test-password"
		testVMName           = "test-vm"
		testVMNamespace      = "default"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating an KubeBMC", func() {
		It("Should create a Pod and a Service", func() {
			ctx := context.Background()

			// we need to create the namespace in the cluster first
			ns := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{Name: testKubeBMCNamespace},
			}
			Expect(k8sClient.Create(ctx, ns)).Should(Succeed())

			By("Creating a new KubeBMC")
			kubeBMC := &virtualmachinev1.KubeBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testKubeBMCName,
					Namespace: testKubeBMCNamespace,
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "zespre.com/v1",
					Kind:       "KubeBMC",
				},
				Spec: virtualmachinev1.KubeBMCSpec{
					Username:                testUsername,
					Password:                testPassword,
					VirtualMachineNamespace: testVMNamespace,
					VirtualMachineName:      testVMName,
				},
			}
			Expect(k8sClient.Create(ctx, kubeBMC)).To(Succeed())

			By("Checking that the Pod is created")
			podLookupKey := types.NamespacedName{Name: kubeBMC.Name + "-kbmc", Namespace: kubeBMC.Namespace}
			createdPod := &corev1.Pod{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, podLookupKey, createdPod)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			By("Checking that the Service is created")
			svcLookupKey := types.NamespacedName{Name: kubeBMC.Name + "-kbmc", Namespace: kubeBMC.Namespace}
			createdSvc := &corev1.Service{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, svcLookupKey, createdSvc)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
