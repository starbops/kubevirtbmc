package bmc

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

type VirtualMachineBMCReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	AgentImageName string
	AgentImageTag  string
}

func (r *VirtualMachineBMCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("vmBMC", req.NamespacedName)

	// 1. Fetch the resource
	vmBMC := &bmcv1beta1.VirtualMachineBMC{}
	if err := r.Get(ctx, req.NamespacedName, vmBMC); err != nil {
		if apierrors.IsNotFound(err) {
			// Object deleted, nothing to do
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to fetch VirtualMachineBMC")
		return ctrl.Result{}, err
	}

	// 2. Handle finalizer/deletion
	if result, err := r.handleFinalizer(ctx, vmBMC, logger); result != nil || err != nil {
		return *result, err
	}

	// 3. Validate references to VM and Secret
	vm, secret, result, err := r.validateReferences(ctx, vmBMC, logger)
	if result != nil || err != nil {
		return *result, err
	}

	// 4. Reconcile Deployment and Service
	svc := r.serviceForVirtualMachineBMC(vmBMC)
	deploy := r.deploymentForVirtualMachineBMC(vmBMC, vm, secret)

	foundSvc := &corev1.Service{}
	foundDeploy := &appsv1.Deployment{}

	// Reconcile Service
	if err := r.reconcileService(ctx, vmBMC, svc, foundSvc, logger); err != nil {
		return ctrl.Result{}, err
	}

	// Reconcile Deployment
	if err := r.reconcileDeployment(ctx, vmBMC, deploy, foundDeploy, logger); err != nil {
		return ctrl.Result{}, err
	}

	// 5. Update Status
	if _, err := r.updateStatus(ctx, vmBMC, foundSvc, foundDeploy, logger); err != nil {
		return ctrl.Result{}, err
	}

	// 6. Requeue every 5 minutes
	return ctrl.Result{RequeueAfter: 5 * time.Minute}, nil
}
func (r *VirtualMachineBMCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bmcv1beta1.VirtualMachineBMC{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Watches(
			&kubevirtv1.VirtualMachine{},
			handler.EnqueueRequestsFromMapFunc(mapVMToBMCs(r.Client)),
		).
		Watches(
			&corev1.Secret{},
			handler.EnqueueRequestsFromMapFunc(mapSecretToBMCs(r.Client)),
		).
		Complete(r)
}

func (r *VirtualMachineBMCReconciler) reconcileService(ctx context.Context, _ *bmcv1beta1.VirtualMachineBMC, desired *corev1.Service, found *corev1.Service, logger logr.Logger) error {
	err := r.Get(ctx, types.NamespacedName{Name: desired.Name, Namespace: desired.Namespace}, found)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("Creating BMC Service", "name", desired.Name)
			return r.Create(ctx, desired)
		}
		logger.Error(err, "Failed to get Service")
		return err
	}

	// Compare mutable fields only
	needUpdate := false
	updated := found.DeepCopy()

	if !equalPorts(found.Spec.Ports, desired.Spec.Ports) {
		updated.Spec.Ports = desired.Spec.Ports
		needUpdate = true
	}

	if !equalMap(found.Spec.Selector, desired.Spec.Selector) {
		updated.Spec.Selector = desired.Spec.Selector
		needUpdate = true
	}

	if needUpdate {
		logger.Info("Updating BMC Service", "name", desired.Name)
		return r.Update(ctx, updated)
	}

	return nil
}

func (r *VirtualMachineBMCReconciler) reconcileDeployment(ctx context.Context, _ *bmcv1beta1.VirtualMachineBMC, desired *appsv1.Deployment, found *appsv1.Deployment, logger logr.Logger) error {
	err := r.Get(ctx, types.NamespacedName{Name: desired.Name, Namespace: desired.Namespace}, found)
	if err != nil {
		if apierrors.IsNotFound(err) {
			logger.Info("Creating BMC Deployment", "name", desired.Name)
			return r.Create(ctx, desired)
		}
		logger.Error(err, "Failed to get Deployment")
		return err
	}

	// Compare key mutable fields
	updated := found.DeepCopy()
	needUpdate := false

	if !equalMap(found.Spec.Template.Labels, desired.Spec.Template.Labels) {
		updated.Spec.Template.Labels = desired.Spec.Template.Labels
		needUpdate = true
	}

	if len(found.Spec.Template.Spec.Containers) > 0 && len(desired.Spec.Template.Spec.Containers) > 0 {
		foundC := found.Spec.Template.Spec.Containers[0]
		desiredC := desired.Spec.Template.Spec.Containers[0]

		if foundC.Image != desiredC.Image || !equalEnv(foundC.Env, desiredC.Env) || !equalMounts(foundC.VolumeMounts, desiredC.VolumeMounts) {
			updated.Spec.Template.Spec.Containers[0] = desiredC
			needUpdate = true
		}
	}

	if needUpdate {
		logger.Info("Updating BMC Deployment", "name", desired.Name)
		return r.Update(ctx, updated)
	}

	return nil
}
