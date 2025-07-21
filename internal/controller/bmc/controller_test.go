package bmc_test

import (
	"context"
	"testing"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	. "kubevirt.io/kubevirtbmc/internal/controller/bmc"

	corev1 "k8s.io/api/core/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestBMCController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BMC Controller Unit Tests")
}

var _ = Describe("ValidateRefs", func() {
	var (
		ctx        context.Context
		client     client.Client
		reconciler *VirtualMachineBMCReconciler
		virtBMC    *bmcv1beta1.VirtualMachineBMC
	)

	BeforeEach(func() {
		ctx = context.TODO()

		scheme := runtime.NewScheme()
		Expect(bmcv1beta1.AddToScheme(scheme)).To(Succeed())
		Expect(kubevirtv1.AddToScheme(scheme)).To(Succeed())
		Expect(corev1.AddToScheme(scheme)).To(Succeed())

		client = fake.NewClientBuilder().WithScheme(scheme).Build()

		reconciler = &VirtualMachineBMCReconciler{
			Client: client,
			Scheme: scheme,
			Log:    ctrl.Log.WithName("test"),
		}
	})

	It("should return error when VirtualMachineBMC is nil", func() {
		res, err := reconciler.ValidateRefs(ctx, nil)
		Expect(err).To(Equal(ErrVirtualMachineBMCNil))
		Expect(res).To(Equal(ctrl.Result{}))
	})

	It("should return error when VirtualMachineRef is missing", func() {
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				AuthSecretRef: &corev1.LocalObjectReference{Name: "my-secret"},
			},
		}
		res, err := reconciler.ValidateRefs(ctx, virtBMC)
		Expect(err).To(Equal(ErrVirtualMachineRefNil))
		Expect(res).To(Equal(ctrl.Result{}))
	})

	It("should return error when AuthSecretRef is missing", func() {
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				VirtualMachineRef: &corev1.LocalObjectReference{Name: "my-vm"},
			},
		}
		vm := &kubevirtv1.VirtualMachine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-vm",
				Namespace: "default",
			},
		}
		client = fake.NewClientBuilder().
			WithScheme(reconciler.Scheme).
			WithObjects(vm).
			Build()
		reconciler.Client = client

		res, err := reconciler.ValidateRefs(ctx, virtBMC)
		Expect(err).To(Equal(ErrAuthSecretRefNil))
		Expect(res).To(Equal(ctrl.Result{}))
	})

	It("should succeed when both VM and Secret exist", func() {
		vm := &kubevirtv1.VirtualMachine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-vm",
				Namespace: "default",
			},
		}
		secret := &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-secret",
				Namespace: "default",
			},
		}
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				VirtualMachineRef: &corev1.LocalObjectReference{Name: "my-vm"},
				AuthSecretRef:     &corev1.LocalObjectReference{Name: "my-secret"},
			},
		}

		client = fake.NewClientBuilder().
			WithScheme(reconciler.Scheme).
			WithObjects(vm, secret).
			Build()
		reconciler.Client = client

		res, err := reconciler.ValidateRefs(ctx, virtBMC)
		Expect(err).To(BeNil())
		Expect(res).To(Equal(ctrl.Result{}))
	})
})
