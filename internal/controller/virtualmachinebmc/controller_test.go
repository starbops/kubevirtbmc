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

package virtualmachinebmc

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"

	bmcv1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

var _ = Describe("VirtualMachineBMC Controller", func() {
	const (
		testVirtualMachineBMCName      = "test-vmbmc"
		testVirtualMachineBMCNamespace = "default"
		testVMName                     = "test-vm"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When creating a VirtualMachineBMC", func() {
		serviceAccountName := fmt.Sprintf("%s-virtbmc", testVMName)
		roleBindingName := fmt.Sprintf("%s-virtbmc-rolebinding", testVMName)

		It("Should create RBAC resources, Pod and Service", func() {
			ctx := context.Background()

			By("Creating the referenced VirtualMachine first")
			vm := &kubevirtv1.VirtualMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testVMName,
					Namespace: testVirtualMachineBMCNamespace,
				},
				Spec: kubevirtv1.VirtualMachineSpec{
					Running: boolPtr(false),
					Template: &kubevirtv1.VirtualMachineInstanceTemplateSpec{
						Spec: kubevirtv1.VirtualMachineInstanceSpec{
							Domain: kubevirtv1.DomainSpec{},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, vm)).Should(Succeed())

			By("Creating a new VirtualMachineBMC")
			virtualMachineBMC := &bmcv1.VirtualMachineBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testVirtualMachineBMCName,
					Namespace: testVirtualMachineBMCNamespace,
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "bmc.kubevirt.io/v1beta1",
					Kind:       "VirtualMachineBMC",
				},
				Spec: bmcv1.VirtualMachineBMCSpec{
					VirtualMachineRef: &corev1.LocalObjectReference{
						Name: testVMName,
					},
				},
			}
			Expect(k8sClient.Create(ctx, virtualMachineBMC)).To(Succeed())

			By("Checking that the ServiceAccount is created")
			saLookupKey := types.NamespacedName{
				Name:      serviceAccountName,
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdSA := &corev1.ServiceAccount{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, saLookupKey, createdSA)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			By("Checking that the RoleBinding is created")
			rbLookupKey := types.NamespacedName{
				Name:      roleBindingName,
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdRB := &rbacv1.RoleBinding{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, rbLookupKey, createdRB)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			Expect(createdRB.RoleRef.Name).To(Equal(clusterRoleName))
			Expect(createdRB.Subjects).To(HaveLen(1))
			Expect(createdRB.Subjects[0].Name).To(Equal(serviceAccountName))

			By("Checking that the Pod is created with correct name")
			// Pod name is based on VM name, not BMC name
			podLookupKey := types.NamespacedName{
				Name:      testVMName + "-virtbmc",
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdPod := &corev1.Pod{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, podLookupKey, createdPod)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			Expect(createdPod.Labels).To(HaveKeyWithValue(VirtualMachineBMCNameLabel, testVirtualMachineBMCName))
			Expect(createdPod.Labels).To(HaveKeyWithValue(VMNameLabel, testVMName))
			Expect(createdPod.Spec.ServiceAccountName).To(Equal(serviceAccountName))
			Expect(createdPod.Spec.Containers).To(HaveLen(1))
			Expect(createdPod.Spec.Containers[0].Name).To(Equal(virtBMCContainerName))

			By("Checking that the Service is created with correct name")
			// Service name is also based on VM name, not BMC name
			svcLookupKey := types.NamespacedName{
				Name:      testVMName + "-virtbmc",
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdSvc := &corev1.Service{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, svcLookupKey, createdSvc)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			Expect(createdSvc.Labels).To(HaveKeyWithValue(VirtualMachineBMCNameLabel, testVirtualMachineBMCName))
			Expect(createdSvc.Labels).To(HaveKeyWithValue(VMNameLabel, testVMName))
			Expect(createdSvc.Spec.Selector).To(HaveKeyWithValue(VirtualMachineBMCNameLabel, testVirtualMachineBMCName))
			Expect(createdSvc.Spec.Ports).To(HaveLen(2))
		})

		It("Should set status condition when VirtualMachine does not exist", func() {
			ctx := context.Background()

			By("Creating a VirtualMachineBMC without a referenced VM")
			virtualMachineBMC := &bmcv1.VirtualMachineBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vmbmc-no-vm",
					Namespace: testVirtualMachineBMCNamespace,
				},
				Spec: bmcv1.VirtualMachineBMCSpec{
					VirtualMachineRef: &corev1.LocalObjectReference{
						Name: "non-existent-vm",
					},
				},
			}
			Expect(k8sClient.Create(ctx, virtualMachineBMC)).To(Succeed())

			By("Checking that the status condition is set")
			bmcLookupKey := types.NamespacedName{
				Name:      "test-vmbmc-no-vm",
				Namespace: testVirtualMachineBMCNamespace,
			}
			Eventually(func() bool {
				var bmc bmcv1.VirtualMachineBMC
				if err := k8sClient.Get(ctx, bmcLookupKey, &bmc); err != nil {
					return false
				}
				for _, cond := range bmc.Status.Conditions {
					if cond.Type == bmcv1.ConditionReady &&
						cond.Status == metav1.ConditionFalse &&
						cond.Reason == "VirtualMachineNotFound" {
						return true
					}
				}
				return false
			}, timeout, interval).Should(BeTrue())
		})

		It("Should reconcile and create resources when VirtualMachine is created after VirtualMachineBMC", func() {
			ctx := context.Background()

			By("Creating a VirtualMachineBMC first, before the VM exists")
			virtualMachineBMC := &bmcv1.VirtualMachineBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vmbmc-late-vm",
					Namespace: testVirtualMachineBMCNamespace,
				},
				Spec: bmcv1.VirtualMachineBMCSpec{
					VirtualMachineRef: &corev1.LocalObjectReference{
						Name: "late-created-vm",
					},
				},
			}
			Expect(k8sClient.Create(ctx, virtualMachineBMC)).To(Succeed())

			By("Verifying that the status condition shows VM not found")
			bmcLookupKey := types.NamespacedName{
				Name:      "test-vmbmc-late-vm",
				Namespace: testVirtualMachineBMCNamespace,
			}
			Eventually(func() bool {
				var bmc bmcv1.VirtualMachineBMC
				if err := k8sClient.Get(ctx, bmcLookupKey, &bmc); err != nil {
					return false
				}
				for _, cond := range bmc.Status.Conditions {
					if cond.Type == bmcv1.ConditionReady &&
						cond.Status == metav1.ConditionFalse &&
						cond.Reason == "VirtualMachineNotFound" {
						return true
					}
				}
				return false
			}, timeout, interval).Should(BeTrue())

			By("Verifying that Pod and Service are NOT created yet")
			podLookupKey := types.NamespacedName{
				Name:      "late-created-vm-virtbmc",
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdPod := &corev1.Pod{}
			Consistently(func() bool {
				err := k8sClient.Get(ctx, podLookupKey, createdPod)
				return err != nil // Should not exist
			}, time.Second*2, interval).Should(BeTrue())

			By("Now creating the VirtualMachine in the same namespace")
			vm := &kubevirtv1.VirtualMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "late-created-vm",
					Namespace: testVirtualMachineBMCNamespace,
				},
				Spec: kubevirtv1.VirtualMachineSpec{
					Running: boolPtr(false),
					Template: &kubevirtv1.VirtualMachineInstanceTemplateSpec{
						Spec: kubevirtv1.VirtualMachineInstanceSpec{
							Domain: kubevirtv1.DomainSpec{},
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, vm)).Should(Succeed())

			By("Verifying that the controller reconciles and creates RBAC resources")
			saLookupKey := types.NamespacedName{
				Name:      "late-created-vm-virtbmc",
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdSA := &corev1.ServiceAccount{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, saLookupKey, createdSA)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			rbLookupKey := types.NamespacedName{
				Name:      "late-created-vm-virtbmc-rolebinding",
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdRB := &rbacv1.RoleBinding{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, rbLookupKey, createdRB)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			By("Verifying that the Pod is now created")
			Eventually(func() bool {
				err := k8sClient.Get(ctx, podLookupKey, createdPod)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			Expect(createdPod.Labels).To(HaveKeyWithValue(VirtualMachineBMCNameLabel, "test-vmbmc-late-vm"))
			Expect(createdPod.Labels).To(HaveKeyWithValue(VMNameLabel, "late-created-vm"))
			Expect(createdPod.Spec.ServiceAccountName).To(Equal("late-created-vm-virtbmc"))

			By("Verifying that the Service is now created")
			svcLookupKey := types.NamespacedName{
				Name:      "late-created-vm-virtbmc",
				Namespace: testVirtualMachineBMCNamespace,
			}
			createdSvc := &corev1.Service{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, svcLookupKey, createdSvc)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			Expect(createdSvc.Labels).To(HaveKeyWithValue(VirtualMachineBMCNameLabel, "test-vmbmc-late-vm"))
			Expect(createdSvc.Labels).To(HaveKeyWithValue(VMNameLabel, "late-created-vm"))
			Expect(createdSvc.Spec.Ports).To(HaveLen(2))
		})
	})
})

func boolPtr(b bool) *bool {
	return &b
}
