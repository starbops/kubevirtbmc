package bmc_test

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	. "kubevirt.io/kubevirtbmc/internal/controller/bmc"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("ReconcileService", func() {
	var (
		ctx        context.Context
		reconciler *VirtualMachineBMCReconciler
		virtBMC    *bmcv1beta1.VirtualMachineBMC
	)

	BeforeEach(func() {
		ctx = context.TODO()

		scheme := runtime.NewScheme()
		Expect(bmcv1beta1.AddToScheme(scheme)).To(Succeed())
		Expect(corev1.AddToScheme(scheme)).To(Succeed())

		reconciler = &VirtualMachineBMCReconciler{
			Client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			Scheme: scheme,
		}

		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "default",
			},
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				VirtualMachineRef: &corev1.LocalObjectReference{
					Name: "vm-test",
				},
			},
		}
	})

	It("should create a Service if not found", func() {
		_, res, err := reconciler.ReconcileService(ctx, virtBMC)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue()) // requeue expected after creation

		service := &corev1.Service{}
		err = reconciler.Get(ctx, types.NamespacedName{
			Name:      virtBMC.Name + BMCServiceNameSuffix,
			Namespace: virtBMC.Namespace,
		}, service)
		Expect(err).To(BeNil())
		Expect(service.Labels).To(Equal(LabelsForBMC(virtBMC.Name, virtBMC.Spec.VirtualMachineRef.Name)))
		Expect(service.Spec.Selector).To(Equal(LabelsForBMC(virtBMC.Name, virtBMC.Spec.VirtualMachineRef.Name)))
		Expect(service.Spec.Selector["app.kubernetes.io/instance"]).To(Equal(virtBMC.Name))

	})

	It("should do nothing if the Service already exists", func() {
		service := reconciler.NewService(virtBMC)
		Expect(reconciler.Create(ctx, service)).To(Succeed())

		_, res, err := reconciler.ReconcileService(ctx, virtBMC)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeFalse())
	})
})
