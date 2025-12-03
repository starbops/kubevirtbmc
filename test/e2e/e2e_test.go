package e2e

import (
	"context"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	bmcv1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

const (
	vmName      = "test-vm"
	vmNamespace = "default"
	bmcName     = "test-bmc"
)

var _ = Describe("KubeVirtBMC controller manager", Ordered, func() {

	serviceAccountName := fmt.Sprintf("%s-virtbmc", vmName)
	roleBindingName := fmt.Sprintf("%s-virtbmc-rolebinding", vmName)

	It("should run successfully", func() {
		By("validating the controller-manager pod exists")
		var podList corev1.PodList
		sets := labels.Set{
			"control-plane": "controller-manager",
		}
		err := k8sClient.List(context.TODO(), &podList, &client.ListOptions{
			LabelSelector: labels.SelectorFromSet(sets),
		})
		Expect(err).ToNot(HaveOccurred())
		Expect(podList.Items).To(HaveLen(1), "expected one controller-manager pod exists")

		By("validating the controller-manager pod is ready to serve")
		var pod corev1.Pod
		podLookupKey := types.NamespacedName{
			Namespace: podList.Items[0].Namespace,
			Name:      podList.Items[0].Name,
		}
		Eventually(func() bool {
			err := k8sClient.Get(context.TODO(), podLookupKey, &pod)
			if err != nil {
				return false
			}
			for _, condition := range pod.Status.Conditions {
				if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionTrue {
					return true
				}
			}
			return false
		}, timeout, interval).Should(BeTrue())
	})

	Context("initially", func() {
		It("should have no VirtualMachineBMCs", func() {
			var vmBMCList bmcv1.VirtualMachineBMCList
			err := k8sClient.List(context.TODO(), &vmBMCList, &client.ListOptions{})
			Expect(err).ToNot(HaveOccurred())
			Expect(vmBMCList.Items).To(HaveLen(0))
		})
	})

	Context("a new virtual machine", func() {
		var (
			vm           kubevirtv1.VirtualMachine
			secret       corev1.Secret
			createdVMBMC *bmcv1.VirtualMachineBMC
		)

		Specify("a VirtualMachine", func() {
			vm = kubevirtv1.VirtualMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      vmName,
					Namespace: vmNamespace,
				},
				Spec: kubevirtv1.VirtualMachineSpec{
					Running: func(b bool) *bool { return &b }(true),
					Template: &kubevirtv1.VirtualMachineInstanceTemplateSpec{
						Spec: kubevirtv1.VirtualMachineInstanceSpec{
							Domain: kubevirtv1.DomainSpec{
								Devices: kubevirtv1.Devices{
									Disks:      []kubevirtv1.Disk{},
									Interfaces: []kubevirtv1.Interface{},
								},
							},
						},
					},
				},
			}
			err := k8sClient.Create(context.TODO(), &vm)
			Expect(err).ToNot(HaveOccurred())
		})

		Specify("a Secret for BMC credentials", func() {
			secret = corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "bmc-credentials-secret",
					Namespace: "default",
				},
				Data: map[string][]byte{
					"username": []byte("admin"),
					"password": []byte("password"),
				},
			}
			err := k8sClient.Create(context.TODO(), &secret)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should allow the user to create a VirtualMachineBMC in the same namespace", func() {
			createdVMBMC = &bmcv1.VirtualMachineBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      bmcName,
					Namespace: vm.Namespace,
				},
				Spec: bmcv1.VirtualMachineBMCSpec{
					VirtualMachineRef: &corev1.LocalObjectReference{
						Name: vm.Name,
					},
					AuthSecretRef: &corev1.LocalObjectReference{
						Name: secret.Name,
					},
				},
			}
			err := k8sClient.Create(context.TODO(), createdVMBMC)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should create ServiceAccount and RoleBinding in the same namespace", func() {
			saLookupKey := types.NamespacedName{
				Name:      serviceAccountName,
				Namespace: createdVMBMC.Namespace,
			}
			sa := &corev1.ServiceAccount{}
			Eventually(func() bool {
				err := k8sClient.Get(context.TODO(), saLookupKey, sa)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			rbLookupKey := types.NamespacedName{
				Name:      roleBindingName,
				Namespace: createdVMBMC.Namespace,
			}
			rb := &rbacv1.RoleBinding{}
			Eventually(func() bool {
				err := k8sClient.Get(context.TODO(), rbLookupKey, rb)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})

		It("should create an agent Pod in the same namespace", func() {
			agentPodLookupKey := types.NamespacedName{
				Name:      strings.Join([]string{createdVMBMC.Spec.VirtualMachineRef.Name, "virtbmc"}, "-"),
				Namespace: createdVMBMC.Namespace,
			}
			pod := &corev1.Pod{}
			Eventually(func() bool {
				err := k8sClient.Get(context.TODO(), agentPodLookupKey, pod)
				if err != nil {
					return false
				}
				for _, condition := range pod.Status.Conditions {
					if condition.Type == corev1.PodReady && condition.Status == corev1.ConditionTrue {
						return true
					}
				}
				return false
			}, timeout, interval).Should(BeTrue())
		})

		It("should create an agent Service in the same namespace", func() {
			agentSvcLookupKey := types.NamespacedName{
				Name:      strings.Join([]string{createdVMBMC.Spec.VirtualMachineRef.Name, "virtbmc"}, "-"),
				Namespace: createdVMBMC.Namespace,
			}
			svc := &corev1.Service{}
			Eventually(func() bool {
				err := k8sClient.Get(context.TODO(), agentSvcLookupKey, svc)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
