package e2e

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	kubevirtv1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	virtualmachinev1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

const (
	kubeVirtBMCNamespace = "kubevirtbmc-system"
	testVMName           = "test-vm"
	testSecretName       = "test-secret"
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
			Namespace:     kubeVirtBMCNamespace,
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

	Context("a new VirtualMachineBMC setup", func() {

		BeforeAll(func() {
			By("creating the Secret for BMC auth")
			secret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testSecretName,
					Namespace: kubeVirtBMCNamespace,
				},
				Data: map[string][]byte{
					"username": []byte("admin"),
					"password": []byte("password"),
				},
			}
			Expect(k8sClient.Create(context.TODO(), secret)).To(Succeed())

			By("creating the VirtualMachine")
			vm := &kubevirtv1.VirtualMachine{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testVMName,
					Namespace: kubeVirtBMCNamespace,
				},
				Spec: kubevirtv1.VirtualMachineSpec{
					Running: ptr.To(true),
					Template: &kubevirtv1.VirtualMachineInstanceTemplateSpec{
						Spec: kubevirtv1.VirtualMachineInstanceSpec{
							Domain: kubevirtv1.DomainSpec{},
						},
					},
				},
			}
			Expect(k8sClient.Create(context.TODO(), vm)).To(Succeed())

			By("creating the VirtualMachineBMC")
			bmc := &virtualmachinev1.VirtualMachineBMC{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testVMName,
					Namespace: kubeVirtBMCNamespace,
				},
				Spec: virtualmachinev1.VirtualMachineBMCSpec{
					VirtualMachineRef: &corev1.LocalObjectReference{
						Name: testVMName,
					},
					AuthSecretRef: &corev1.LocalObjectReference{
						Name: testSecretName,
					},
				},
			}
			Expect(k8sClient.Create(context.TODO(), bmc)).To(Succeed())
		})

		It("should create a Deployment", func() {
			Eventually(func() error {
				return k8sClient.Get(context.TODO(), types.NamespacedName{
					Name:      testVMName + "-bmc-deployment",
					Namespace: kubeVirtBMCNamespace,
				}, &appsv1.Deployment{})
			}, timeout, interval).Should(Succeed())
		})

		It("should create a Service", func() {
			Eventually(func() error {
				return k8sClient.Get(context.TODO(), types.NamespacedName{
					Name:      testVMName + "-bmc-service",
					Namespace: kubeVirtBMCNamespace,
				}, &corev1.Service{})
			}, timeout, interval).Should(Succeed())
		})
	})

	Context("when the VirtualMachineBMC is deleted", func() {
		It("should delete the associated Deployment and Service", func() {
			By("deleting the VirtualMachineBMC resource")
			bmc := &virtualmachinev1.VirtualMachineBMC{}
			bmcKey := types.NamespacedName{
				Name:      testVMName,
				Namespace: kubeVirtBMCNamespace,
			}
			Expect(k8sClient.Get(context.TODO(), bmcKey, bmc)).To(Succeed())
			Expect(k8sClient.Delete(context.TODO(), bmc)).To(Succeed())

			By("verifying the Deployment is deleted")
			Eventually(func() bool {
				deploy := &appsv1.Deployment{}
				err := k8sClient.Get(context.TODO(), types.NamespacedName{
					Name:      testVMName + "-bmc-deployment",
					Namespace: kubeVirtBMCNamespace,
				}, deploy)
				return apierrors.IsNotFound(err)
			}, timeout, interval).Should(BeTrue())

			By("verifying the Service is deleted")
			Eventually(func() bool {
				svc := &corev1.Service{}
				err := k8sClient.Get(context.TODO(), types.NamespacedName{
					Name:      testVMName + "-bmc-service",
					Namespace: kubeVirtBMCNamespace,
				}, svc)
				return apierrors.IsNotFound(err)
			}, timeout, interval).Should(BeTrue())
		})
	})

})
