package bmc_test

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	. "kubevirt.io/kubevirtbmc/internal/controller/bmc"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("ReconcileDeployment", func() {
	var (
		ctx        context.Context
		reconciler *VirtualMachineBMCReconciler
		log        logr.Logger
		virtBMC    *bmcv1beta1.VirtualMachineBMC
	)

	BeforeEach(func() {
		ctx = context.TODO()
		log = funcr.New(func(prefix, args string) {
			_, _ = GinkgoWriter.Write([]byte(args))
		}, funcr.Options{})

		scheme := runtime.NewScheme()
		Expect(appsv1.AddToScheme(scheme)).To(Succeed())
		Expect(corev1.AddToScheme(scheme)).To(Succeed())
		Expect(bmcv1beta1.AddToScheme(scheme)).To(Succeed())

		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "default",
			},
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				VirtualMachineRef: &corev1.LocalObjectReference{
					Name: "vm-test",
				},
				AuthSecretRef: &corev1.LocalObjectReference{
					Name: "test-secret",
				},
			},
		}

		reconciler = &VirtualMachineBMCReconciler{
			Client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			Scheme: scheme,
			AgentImageName: AgentImageConfig{
				ImageName:     "test-image",
				ImageTag:      "latest",
				ContainerName: "virtbmc",
				FullImage:     "test-image:latest",
			},
		}
	})

	It("should create Deployment if not found", func() {
		_, err := reconciler.ReconcileDeployment(ctx, virtBMC, log)
		Expect(err).To(BeNil())

		dep := &appsv1.Deployment{}
		err = reconciler.Get(ctx, types.NamespacedName{
			Name:      virtBMC.Name + BMCDeploymentNameSuffix,
			Namespace: virtBMC.Namespace,
		}, dep)
		Expect(err).To(BeNil())
		Expect(*dep.Spec.Replicas).To(Equal(int32(DefaultReplicas)))
	})

	It("should update Deployment if replica count mismatches", func() {
		dep := reconciler.NewDeployment(virtBMC)
		dep.Spec.Replicas = new(int32)
		*dep.Spec.Replicas = 2

		err := reconciler.Client.Create(ctx, dep)
		Expect(err).To(BeNil())

		res, err := reconciler.ReconcileDeployment(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue())

		updatedDep := &appsv1.Deployment{}
		err = reconciler.Get(ctx, types.NamespacedName{
			Name:      virtBMC.Name + BMCDeploymentNameSuffix,
			Namespace: virtBMC.Namespace,
		}, updatedDep)
		Expect(err).To(BeNil())
		Expect(*updatedDep.Spec.Replicas).To(Equal(int32(DefaultReplicas)))
	})

	It("should do nothing if deployment is up to date", func() {
		dep := reconciler.NewDeployment(virtBMC)
		dep.Spec.Replicas = new(int32)
		*dep.Spec.Replicas = DefaultReplicas

		err := reconciler.Client.Create(ctx, dep)
		Expect(err).To(BeNil())

		res, err := reconciler.ReconcileDeployment(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeFalse())
	})

})
