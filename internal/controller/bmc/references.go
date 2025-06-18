package bmc

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// validateReferences fetches and validates the referenced VM and Secret.
// Returns nil for both if not found, with an early Requeue result.
func (r *VirtualMachineBMCReconciler) validateReferences(ctx context.Context, vmBMC *bmcv1beta1.VirtualMachineBMC, logger logr.Logger) (*kubevirtv1.VirtualMachine, *corev1.Secret, *ctrl.Result, error) {
	if vmBMC.Spec.VirtualMachineRef == nil || vmBMC.Spec.VirtualMachineRef.Name == "" {
		r.setCondition(ctx, vmBMC, bmcv1beta1.ConditionReady, metav1.ConditionFalse, "VirtualMachineRefMissing", "VirtualMachineRef is required and must specify a name.")
		return nil, nil, &ctrl.Result{RequeueAfter: time.Minute}, nil
	}
	var vm *kubevirtv1.VirtualMachine
	if err := r.Get(ctx, types.NamespacedName{Namespace: vmBMC.Namespace, Name: vmBMC.Spec.VirtualMachineRef.Name}, vm); err != nil {
		if errors.IsNotFound(err) {
			r.setCondition(ctx, vmBMC, bmcv1beta1.ConditionReady, metav1.ConditionFalse, "VirtualMachineNotFound", fmt.Sprintf("Referenced VirtualMachine '%s' not found.", vmBMC.Spec.VirtualMachineRef.Name))
			return nil, nil, &ctrl.Result{RequeueAfter: time.Minute}, nil
		}
		logger.Error(err, "Failed to get VirtualMachine")
		return nil, nil, &ctrl.Result{}, err
	}
	if vmBMC.Spec.AuthSecretRef == nil || vmBMC.Spec.AuthSecretRef.Name == "" {
		r.setCondition(ctx, vmBMC, bmcv1beta1.ConditionReady, metav1.ConditionFalse, "AuthSecretRefMissing", "AuthSecretRef is required and must specify a name.")
		return nil, nil, &ctrl.Result{RequeueAfter: time.Minute}, nil
	}
	secret := &corev1.Secret{}
	if err := r.Get(ctx, types.NamespacedName{Namespace: vmBMC.Namespace, Name: vmBMC.Spec.AuthSecretRef.Name}, secret); err != nil {
		if errors.IsNotFound(err) {
			r.setCondition(ctx, vmBMC, bmcv1beta1.ConditionReady, metav1.ConditionFalse, "AuthSecretNotFound", fmt.Sprintf("Referenced Secret '%s' not found.", vmBMC.Spec.AuthSecretRef.Name))
			return nil, nil, &ctrl.Result{RequeueAfter: time.Minute}, nil
		}
		logger.Error(err, "Failed to get Secret")
		return nil, nil, &ctrl.Result{}, err
	}
	r.setCondition(ctx, vmBMC, bmcv1beta1.ConditionReady, metav1.ConditionUnknown, "ReferencesValidated", "All required references are found.")
	return vm, secret, nil, nil
}
