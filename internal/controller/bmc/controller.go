/*
 * This file is part of the KubeVirt/KubeVirtBMC project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors.
 *
 */
package bmc

import (
	"context"
	"errors"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"

	bmcv1beta1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

type VirtualMachineBMCReconciler struct {
	client.Client
	Scheme         *runtime.Scheme
	AgentImageName string
	AgentImageTag  string
	Log            logr.Logger
}

// +kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubevirt.io,resources=virtualmachines,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch

const (
	ConditionReady    = "Ready"
	ConditionNotReady = "NotReady"

	ReasonReady  = "AllComponentsReady"
	MessageReady = "VirtualMachineBMC is ready"

	VirtualMachineNotFound        = "VirtualMachineNotFound"
	MessageVirtualMachineNotFound = "Referenced VirtualMachine does not exist"

	AuthSecretNotFound        = "AuthSecretNotFound"
	MessageAuthSecretNotFound = "Referenced Secret does not exist"
)

var (
	ErrVirtualMachineBMCNil   = errors.New("virtualMachineBMC is nil")
	ErrVirtualMachineRefNil   = errors.New("virtualMachineRef is nil or empty")
	ErrAuthSecretRefNil       = errors.New("authSecretRef is nil or empty")
	ErrVirtualMachineNotFound = errors.New("referenced virtualMachine not found")
	ErrAuthSecretNotFound     = errors.New("referenced secret not found")
)

func (r *VirtualMachineBMCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log.WithValues("VirtualMachineBMC", req.NamespacedName).Info("Reconciling VirtualMachineBMC")

	var virtBMC bmcv1beta1.VirtualMachineBMC
	if err := r.Get(ctx, req.NamespacedName, &virtBMC); err != nil {
		r.Log.Error(err, "Failed to get VirtualMachineBMC")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	res, err := r.ValidateRefs(ctx, &virtBMC)
	if err != nil || res.Requeue {
		if errors.Is(err, ErrVirtualMachineNotFound) || errors.Is(err, ErrAuthSecretNotFound) {
			var (
				errorReason  = AuthSecretNotFound
				errorMessage = MessageAuthSecretNotFound
			)

			if errors.Is(err, ErrVirtualMachineNotFound) {
				errorReason = VirtualMachineNotFound
				errorMessage = MessageVirtualMachineNotFound
			}
			r.Log.Error(err, "Referenced resource not found, setting NotReady condition", "Reason", errorReason, "Message", errorMessage)
			cond := metav1.Condition{
				Type:               ConditionNotReady,
				Status:             metav1.ConditionFalse,
				Reason:             errorReason,
				Message:            errorMessage,
				ObservedGeneration: virtBMC.Generation,
			}
			setCondition(&virtBMC, cond)
			if updateErr := r.Status().Update(ctx, &virtBMC); updateErr != nil {
				r.Log.Error(updateErr, "Failed to update VirtualMachineBMC status")
				return ctrl.Result{}, updateErr
			}
			r.Log.Info("Cleaning up resources due to missing references")
			if err := r.DeleteResources(ctx, &virtBMC); err != nil {
				return ctrl.Result{}, err
			}

			r.Log.Info("Resources cleaned up due to missing references")
			return ctrl.Result{Requeue: true}, nil
		}
		return res, err
	}

	if res, err := r.ReconcileDeployment(ctx, &virtBMC); err != nil || res.Requeue {
		return res, err
	}

	svc, res, err := r.ReconcileService(ctx, &virtBMC)
	if err != nil || res.Requeue {
		return res, err
	}

	if err := r.CheckDeploymentReady(&virtBMC); err != nil {
		r.Log.Error(err, "Deployment is not ready, requeuing VirtualMachineBMC")
		return ctrl.Result{Requeue: true}, nil
	}

	ready := metav1.Condition{
		Type:               ConditionReady,
		Status:             metav1.ConditionTrue,
		Reason:             ReasonReady,
		Message:            MessageReady,
		LastTransitionTime: metav1.Now(),
		ObservedGeneration: virtBMC.Generation,
	}

	if err := r.UpdateStatusIfNeeded(ctx, &virtBMC, svc.Spec.ClusterIP, ready); err != nil {
		r.Log.Error(err, "Failed to update VirtualMachineBMC status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *VirtualMachineBMCReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &bmcv1beta1.VirtualMachineBMC{}, "spec.virtualMachineRef.name", func(rawObj client.Object) []string {
		bmc := rawObj.(*bmcv1beta1.VirtualMachineBMC)
		if bmc.Spec.VirtualMachineRef == nil {
			return nil
		}
		return []string{bmc.Spec.VirtualMachineRef.Name}
	}); err != nil {
		return err
	}

	if err := mgr.GetFieldIndexer().IndexField(context.TODO(), &bmcv1beta1.VirtualMachineBMC{}, "spec.authSecretRef.name", func(rawObj client.Object) []string {
		bmc := rawObj.(*bmcv1beta1.VirtualMachineBMC)
		if bmc.Spec.AuthSecretRef == nil {
			return nil
		}
		return []string{bmc.Spec.AuthSecretRef.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&bmcv1beta1.VirtualMachineBMC{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Watches(
			&kubevirtv1.VirtualMachine{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []ctrl.Request {
				var list bmcv1beta1.VirtualMachineBMCList
				if err := r.List(ctx, &list, client.MatchingFields{"spec.virtualMachineRef.name": obj.GetName()}); err != nil {
					return nil
				}
				var reqs []ctrl.Request
				for _, item := range list.Items {
					reqs = append(reqs, ctrl.Request{NamespacedName: types.NamespacedName{
						Name: item.Name, Namespace: item.Namespace,
					}})
				}
				return reqs
			}),
		).
		Watches(
			&corev1.Secret{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []ctrl.Request {
				var list bmcv1beta1.VirtualMachineBMCList
				if err := r.List(ctx, &list, client.MatchingFields{"spec.authSecretRef.name": obj.GetName()}); err != nil {
					return nil
				}
				var reqs []ctrl.Request
				for _, item := range list.Items {
					reqs = append(reqs, ctrl.Request{NamespacedName: types.NamespacedName{
						Name: item.Name, Namespace: item.Namespace,
					}})
				}
				return reqs
			}),
		).
		Complete(r)
}

func (r *VirtualMachineBMCReconciler) UpdateStatusIfNeeded(ctx context.Context, virtBMC *bmcv1beta1.VirtualMachineBMC, clusterIP string, condition metav1.Condition) error {
	updated := false

	if clusterIP != "" && virtBMC.Status.ClusterIP != clusterIP {
		virtBMC.Status.ClusterIP = clusterIP
		updated = true
	}

	found := false
	for i, cond := range virtBMC.Status.Conditions {
		if cond.Type == condition.Type {
			found = true
			if cond.Status != condition.Status || cond.Message != condition.Message {
				virtBMC.Status.Conditions[i] = condition
				updated = true
			}
			break
		}
	}
	if !found {
		virtBMC.Status.Conditions = append(virtBMC.Status.Conditions, condition)
		updated = true
	}

	if updated {
		r.Log.Info("Updating VirtualMachineBMC status")
		if err := r.Status().Update(ctx, virtBMC); err != nil {
			r.Log.Error(err, "Failed to update VirtualMachineBMC status")
			return err
		}
	}

	return nil
}

func (r *VirtualMachineBMCReconciler) ValidateRefs(ctx context.Context, virtBMC *bmcv1beta1.VirtualMachineBMC) (ctrl.Result, error) {
	if virtBMC == nil {
		return ctrl.Result{}, ErrVirtualMachineBMCNil
	}

	log := r.Log.WithValues(
		"VirtualMachineBMC", virtBMC.Name,
		"Namespace", virtBMC.Namespace,
	)

	if virtBMC.Spec.VirtualMachineRef == nil || virtBMC.Spec.VirtualMachineRef.Name == "" {
		log.Error(ErrVirtualMachineRefNil, ErrVirtualMachineRefNil.Error())
		return ctrl.Result{}, ErrVirtualMachineRefNil
	}

	var vm kubevirtv1.VirtualMachine
	if err := r.Get(ctx, types.NamespacedName{
		Name:      virtBMC.Spec.VirtualMachineRef.Name,
		Namespace: virtBMC.Namespace,
	}, &vm); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{Requeue: true}, ErrVirtualMachineNotFound
		}
		return ctrl.Result{}, err
	}

	if virtBMC.Spec.AuthSecretRef == nil || virtBMC.Spec.AuthSecretRef.Name == "" {
		log.Error(ErrAuthSecretRefNil, ErrAuthSecretRefNil.Error())
		return ctrl.Result{}, ErrAuthSecretRefNil
	}

	var secret corev1.Secret
	if err := r.Get(ctx, types.NamespacedName{
		Name:      virtBMC.Spec.AuthSecretRef.Name,
		Namespace: virtBMC.Namespace,
	}, &secret); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{Requeue: true}, ErrAuthSecretNotFound
		}
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func LabelsForBMC(bmcName string, vmName string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":      "kubevirtbmc",
		"app.kubernetes.io/instance":  bmcName,
		"app.kubernetes.io/component": "proxy",
		"bmc.kubevirt.io/vm-name":     vmName,
	}
}

func setCondition(virtBMC *bmcv1beta1.VirtualMachineBMC, condition metav1.Condition) {
	found := false
	for i, cond := range virtBMC.Status.Conditions {
		if cond.Type == condition.Type {
			virtBMC.Status.Conditions[i] = condition
			found = true
			break
		}
	}

	if !found {
		virtBMC.Status.Conditions = append(virtBMC.Status.Conditions, condition)
	}
}

func (r *VirtualMachineBMCReconciler) DeleteResources(ctx context.Context, virtBMC *bmcv1beta1.VirtualMachineBMC) error {
	deploymentErr := client.IgnoreNotFound(r.Delete(ctx, &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      virtBMC.Name + BMCDeploymentNameSuffix,
			Namespace: virtBMC.Namespace,
		},
	}))

	serviceErr := client.IgnoreNotFound(r.Delete(ctx, &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      virtBMC.Name + BMCServiceNameSuffix,
			Namespace: virtBMC.Namespace,
		},
	}))
	if deploymentErr != nil {
		return deploymentErr
	}
	if serviceErr != nil {
		return serviceErr
	}
	return nil
}
