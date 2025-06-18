package bmc

import (
	"context"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

// setCondition updates or inserts a status condition on the VirtualMachineBMC.
func (r *VirtualMachineBMCReconciler) setCondition(ctx context.Context, bmc *bmcv1beta1.VirtualMachineBMC, conditionType string, status metav1.ConditionStatus, reason, message string) {
	newCond := metav1.Condition{
		Type:               conditionType,
		Status:             status,
		Reason:             reason,
		Message:            message,
		LastTransitionTime: metav1.Now(),
		ObservedGeneration: bmc.Generation,
	}

	// Check for an existing condition to update
	for i, cond := range bmc.Status.Conditions {
		if cond.Type == conditionType {
			if cond.Status != status || cond.Reason != reason || cond.Message != message {
				bmc.Status.Conditions[i] = newCond
			}
			return
		}
	}

	// Condition not found, add it
	bmc.Status.Conditions = append(bmc.Status.Conditions, newCond)
}

// updateStatus updates .status.ClusterIP and .status.conditions for the BMC
func (r *VirtualMachineBMCReconciler) updateStatus(ctx context.Context, bmc *bmcv1beta1.VirtualMachineBMC, svc *corev1.Service, deploy *appsv1.Deployment, logger logr.Logger) (*ctrl.Result, error) {
	needUpdate := false

	// Update ClusterIP if changed
	if svc.Spec.ClusterIP != "" && bmc.Status.ClusterIP != svc.Spec.ClusterIP {
		bmc.Status.ClusterIP = svc.Spec.ClusterIP
		needUpdate = true
	}

	// Determine desired condition
	newStatus := metav1.ConditionFalse
	reason := ReasonBMCNotReady
	message := MsgBMCNotReady

	if deploy != nil && deploy.Spec.Replicas != nil && deploy.Status.ReadyReplicas == *deploy.Spec.Replicas {
		newStatus = metav1.ConditionTrue
		reason = ReasonBMCReady
		message = MsgBMCReady
	}

	// Check if current condition is different
	current := metav1.ConditionUnknown
	for _, cond := range bmc.Status.Conditions {
		if cond.Type == ConditionReady {
			current = cond.Status
			break
		}
	}

	if current != newStatus {
		r.setCondition(ctx, bmc, ConditionReady, newStatus, reason, message)
		needUpdate = true
	}

	if needUpdate {
		if err := r.Status().Update(ctx, bmc); err != nil {
			logger.Error(err, "failed to update VirtualMachineBMC status")
			return &ctrl.Result{}, err
		}
	}

	return nil, nil
}
