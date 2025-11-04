package e2e

import (
	"context"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	virtualmachinev1 "kubevirt.io/kubevirtbmc/api/virtualmachine/v1alpha1"
)

const (
	kubeVirtBMCNamespace = "kubevirtbmc-system"
)

var _ = Describe("KubeVirtBMC controller manager", Ordered, func() {

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
			var vmBMCList virtualmachinev1.VirtualMachineBMCList
			err := k8sClient.List(context.TODO(), &vmBMCList, &client.ListOptions{})
			Expect(err).ToNot(HaveOccurred())
			Expect(vmBMCList.Items).To(HaveLen(0))
		})
	})

	Context("a new virtual machine", func() {
		var (
			vm           kubevirtv1.VirtualMachine
			createdVMBMC *virtualmachinev1.VirtualMachineBMC
		)

		Specify("a VirtualMachine", func() {
			vm = kubevirtv1.VirtualMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-vm",
					Namespace: "default",
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

		It("should create a VirtualMachineBMC", func() {
			vmBMCLookupKey := types.NamespacedName{
				Name:      strings.Join([]string{vm.Namespace, vm.Name}, "-"),
				Namespace: kubeVirtBMCNamespace,
			}
			createdVMBMC = &virtualmachinev1.VirtualMachineBMC{}
			Eventually(func() bool {
				err := k8sClient.Get(context.TODO(), vmBMCLookupKey, createdVMBMC)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})

		It("should create a agent Pod", func() {
			agentPodLookupKey := types.NamespacedName{
				Name:      strings.Join([]string{createdVMBMC.Name, "virtbmc"}, "-"),
				Namespace: kubeVirtBMCNamespace,
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

		It("should create a agent Service", func() {
			agentSvcLookupKey := types.NamespacedName{
				Name:      strings.Join([]string{createdVMBMC.Name, "virtbmc"}, "-"),
				Namespace: kubeVirtBMCNamespace,
			}
			svc := &corev1.Service{}
			Eventually(func() bool {
				err := k8sClient.Get(context.TODO(), agentSvcLookupKey, svc)
				return err == nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
