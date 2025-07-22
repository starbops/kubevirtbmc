package bmc_test

import (
	"context"
	"testing"
	"time"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	. "kubevirt.io/kubevirtbmc/internal/controller/bmc"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	kubevirtv1 "kubevirt.io/api/core/v1"

	"github.com/go-logr/logr"
	"github.com/go-logr/logr/funcr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
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
		log        logr.Logger
	)

	BeforeEach(func() {
		ctx = context.TODO()
		log = funcr.New(func(prefix, args string) { GinkgoWriter.Write([]byte(args)) }, funcr.Options{})

		scheme := runtime.NewScheme()
		Expect(bmcv1beta1.AddToScheme(scheme)).To(Succeed())
		Expect(kubevirtv1.AddToScheme(scheme)).To(Succeed())
		Expect(corev1.AddToScheme(scheme)).To(Succeed())

		client = fake.NewClientBuilder().WithScheme(scheme).Build()

		reconciler = &VirtualMachineBMCReconciler{
			Client: client,
			Scheme: scheme,
		}
	})

	It("should requeue when VirtualMachineBMC is nil", func() {
		res, err := reconciler.ValidateRefs(ctx, nil, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue())
	})

	It("should requeue when VirtualMachineRef is missing", func() {
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				AuthSecretRef: &corev1.LocalObjectReference{Name: "my-secret"},
			},
		}
		res, err := reconciler.ValidateRefs(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue())
	})

	It("should requeue when AuthSecretRef is missing", func() {
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				VirtualMachineRef: &corev1.LocalObjectReference{Name: "my-vm"},
			},
		}
		res, err := reconciler.ValidateRefs(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue())
	})

	It("should requeue when referenced VirtualMachine does not exist", func() {
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			Spec: bmcv1beta1.VirtualMachineBMCSpec{
				VirtualMachineRef: &corev1.LocalObjectReference{Name: "my-vm"},
				AuthSecretRef:     &corev1.LocalObjectReference{Name: "my-secret"},
			},
		}
		res, err := reconciler.ValidateRefs(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue())
	})

	It("should requeue when referenced Secret does not exist", func() {
		vm := &kubevirtv1.VirtualMachine{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-vm",
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
			WithObjects(vm).
			Build()

		reconciler.Client = client

		res, err := reconciler.ValidateRefs(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeTrue())
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

		res, err := reconciler.ValidateRefs(ctx, virtBMC, log)
		Expect(err).To(BeNil())
		Expect(res.Requeue).To(BeFalse())
	})
})

var _ = Describe("HandleFinalizer", func() {
	var (
		ctx        context.Context
		reconciler *VirtualMachineBMCReconciler
		log        logr.Logger
		virtBMC    *bmcv1beta1.VirtualMachineBMC
		scheme     *runtime.Scheme
	)

	BeforeEach(func() {
		ctx = context.TODO()
		log = funcr.New(func(prefix, args string) { GinkgoWriter.Write([]byte(args)) }, funcr.Options{})
		scheme = runtime.NewScheme()
		Expect(bmcv1beta1.AddToScheme(scheme)).To(Succeed())
		Expect(appsv1.AddToScheme(scheme)).To(Succeed())
		Expect(corev1.AddToScheme(scheme)).To(Succeed())

		reconciler = &VirtualMachineBMCReconciler{
			Client: fake.NewClientBuilder().WithScheme(scheme).Build(),
			Scheme: scheme,
		}
	})

	It("should add finalizer if not present", func() {
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test",
				Namespace: "default",
			},
		}

		Expect(reconciler.Create(ctx, virtBMC)).To(Succeed())

		_, err := reconciler.HandleFinalizer(ctx, virtBMC, log)
		Expect(err).To(BeNil())

		updated := &bmcv1beta1.VirtualMachineBMC{}
		Expect(reconciler.Get(ctx, client.ObjectKeyFromObject(virtBMC), updated)).To(Succeed())
		Expect(controllerutil.ContainsFinalizer(updated, BMCFinalizer)).To(BeTrue())
	})

	It("should remove finalizer and delete resources when marked for deletion", func() {
		// Setup with finalizer
		virtBMC = &bmcv1beta1.VirtualMachineBMC{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test",
				Namespace:         "default",
				Finalizers:        []string{BMCFinalizer},
				DeletionTimestamp: &metav1.Time{Time: time.Now()},
			},
		}

		// Create deployment and service to be deleted
		dep := &appsv1.Deployment{
			ObjectMeta: MetaForBMC(virtBMC.Name, virtBMC.Namespace, BMCDeploymentNameSuffix),
		}
		svc := &corev1.Service{
			ObjectMeta: MetaForBMC(virtBMC.Name, virtBMC.Namespace, BMCServiceNameSuffix),
		}

		// Add objects to fake client
		reconciler.Client = fake.NewClientBuilder().
			WithScheme(reconciler.Scheme).
			WithObjects(virtBMC, dep, svc).
			Build()

		// Call HandleFinalizer
		_, err := reconciler.HandleFinalizer(ctx, virtBMC, log)
		Expect(err).To(BeNil())

		// Finalizer should be removed and resource deleted
		updated := &bmcv1beta1.VirtualMachineBMC{}
		err = reconciler.Client.Get(ctx, client.ObjectKeyFromObject(virtBMC), updated)
		Expect(apierrors.IsNotFound(err)).To(BeTrue())

		// Deployment should be deleted
		deletedDep := &appsv1.Deployment{}
		err = reconciler.Client.Get(ctx, client.ObjectKey{Name: dep.Name, Namespace: dep.Namespace}, deletedDep)
		Expect(apierrors.IsNotFound(err)).To(BeTrue())

		// Service should be deleted
		deletedSvc := &corev1.Service{}
		err = reconciler.Client.Get(ctx, client.ObjectKey{Name: svc.Name, Namespace: svc.Namespace}, deletedSvc)
		Expect(apierrors.IsNotFound(err)).To(BeTrue())
	})
})
