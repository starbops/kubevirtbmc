package bmc

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

// handleFinalizer ensures the finalizer is added or removed appropriately.
func (r *VirtualMachineBMCReconciler) handleFinalizer(ctx context.Context, bmc *bmcv1beta1.VirtualMachineBMC, log logr.Logger) (*ctrl.Result, error) {
	if bmc.ObjectMeta.DeletionTimestamp.IsZero() {
		// Add finalizer if not present
		if !controllerutil.ContainsFinalizer(bmc, FinalizerName) {
			log.Info("Adding finalizer")
			controllerutil.AddFinalizer(bmc, FinalizerName)
			if err := r.Update(ctx, bmc); err != nil {
				log.Error(err, "Failed to add finalizer")
				return &ctrl.Result{}, err
			}
		}
	} else {
		// Handle deletion
		if controllerutil.ContainsFinalizer(bmc, FinalizerName) {
			log.Info("Cleaning up owned resources before deletion")
			if err := r.cleanupOwnedResources(ctx, bmc, log); err != nil {
				return &ctrl.Result{}, err
			}
			controllerutil.RemoveFinalizer(bmc, FinalizerName)
			if err := r.Update(ctx, bmc); err != nil {
				log.Error(err, "Failed to remove finalizer")
				return &ctrl.Result{}, err
			}
		}
		return &ctrl.Result{}, nil
	}
	return nil, nil
}

// cleanupOwnedResources deletes the Deployment and Service owned by the BMC
func (r *VirtualMachineBMCReconciler) cleanupOwnedResources(ctx context.Context, bmc *bmcv1beta1.VirtualMachineBMC, log logr.Logger) error {
	deploy := &appsv1.Deployment{}
	if err := r.Get(ctx, types.NamespacedName{Name: bmc.Name + "-bmc-proxy", Namespace: bmc.Namespace}, deploy); err == nil {
		log.Info("Deleting Deployment")
		if err := r.Delete(ctx, deploy); err != nil && !apierrors.IsNotFound(err) {
			return fmt.Errorf("failed to delete Deployment: %w", err)
		}
	}

	svc := &corev1.Service{}
	if err := r.Get(ctx, types.NamespacedName{Name: bmc.Name + "-bmc-service", Namespace: bmc.Namespace}, svc); err == nil {
		log.Info("Deleting Service")
		if err := r.Delete(ctx, svc); err != nil && !apierrors.IsNotFound(err) {
			return fmt.Errorf("failed to delete Service: %w", err)
		}
	}

	return nil
}
