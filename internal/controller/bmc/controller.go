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
	"time"

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
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
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

// +kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=bmc.kubevirt.io,resources=virtualmachinebmcs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core,resources=pods;services;secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubevirt.io,resources=virtualmachines,verbs=get;list;watch

func (r *VirtualMachineBMCReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("VirtualMachineBMC", req.NamespacedName)

	log.Info("Reconciling")
	var virtBMC bmcv1beta1.VirtualMachineBMC
	if err := r.Get(ctx, req.NamespacedName, &virtBMC); err != nil {
		log.Error(err, "Failed to get VirtualMachineBMC")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if virtBMC.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(&virtBMC, BMCFinalizer) {
			controllerutil.AddFinalizer(&virtBMC, BMCFinalizer)
			if err := r.Update(ctx, &virtBMC); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		if controllerutil.ContainsFinalizer(&virtBMC, BMCFinalizer) {
			if err := r.deleteDeployment(ctx, &virtBMC, log); err != nil {
				return ctrl.Result{}, err
			}
			if err := r.deleteService(ctx, &virtBMC, log); err != nil {
				return ctrl.Result{}, err
			}
			controllerutil.RemoveFinalizer(&virtBMC, BMCFinalizer)
			if err := r.Update(ctx, &virtBMC); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if virtBMC.Spec.VirtualMachineRef == nil || virtBMC.Spec.VirtualMachineRef.Name == "" {
		log.Error(nil, "VirtualMachineRef is required and must specify a name")
		return ctrl.Result{Requeue: true}, nil
	}

	var vm kubevirtv1.VirtualMachine
	if err := r.Get(ctx, types.NamespacedName{
		Name:      virtBMC.Spec.VirtualMachineRef.Name,
		Namespace: virtBMC.Namespace,
	}, &vm); err != nil {
		log.Error(err, "Failed to get referenced VirtualMachine")
		return ctrl.Result{Requeue: true}, nil
	}

	if virtBMC.Spec.AuthSecretRef == nil || virtBMC.Spec.AuthSecretRef.Name == "" {
		log.Error(nil, "AuthSecretRef is required and must specify a name")
		return ctrl.Result{Requeue: true}, nil
	}

	var secret corev1.Secret
	if err := r.Get(ctx, types.NamespacedName{
		Name:      virtBMC.Spec.AuthSecretRef.Name,
		Namespace: virtBMC.Namespace,
	}, &secret); err != nil {
		log.Error(err, "Failed to get referenced Secret")
		return ctrl.Result{Requeue: true}, nil
	}

	foundDep := &appsv1.Deployment{}
	depName := types.NamespacedName{
		Name:      virtBMC.Name + BMCProxyLabelSuffix,
		Namespace: virtBMC.Namespace,
	}
	if err := r.Get(ctx, depName, foundDep); err != nil {
		if apierrors.IsNotFound(err) {
			dep := r.NewDeployment(&virtBMC)
			log.Info("Creating Deployment", "Deployment", depName)
			if err := r.Create(ctx, dep); err != nil {
				log.Error(err, "Failed to create Deployment", "Deployment", depName)
				return ctrl.Result{}, err
			}
			return ctrl.Result{Requeue: true}, nil
		}
		log.Error(err, "Failed to get Deployment", "Deployment", depName)
		return ctrl.Result{}, err
	}

	size := int32(DefaultReplicas)
	if *foundDep.Spec.Replicas != size {
		foundDep.Spec.Replicas = &size
		if err := r.Update(ctx, foundDep); err != nil {
			log.Error(err, "Failed to update Deployment replicas", "Deployment", depName)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	foundSvc := &corev1.Service{}
	svcName := types.NamespacedName{
		Name:      virtBMC.Name + BMCServiceNameSuffix,
		Namespace: virtBMC.Namespace,
	}
	if err := r.Get(ctx, svcName, foundSvc); err != nil {
		if apierrors.IsNotFound(err) {
			svc := r.NewService(&virtBMC)
			log.Info("Creating Service", "Service", svcName)
			if err := r.Create(ctx, svc); err != nil {
				log.Error(err, "Failed to create Service", "Service", svcName)
				return ctrl.Result{}, err
			}
			return ctrl.Result{Requeue: true}, nil
		}
		log.Error(err, "Failed to get Service", "Service", svcName)
		return ctrl.Result{}, err
	}

	ready := metav1.Condition{
		Type:               ConditionReady,
		Status:             metav1.ConditionTrue,
		Reason:             ReasonReady,
		Message:            MessageReady,
		LastTransitionTime: metav1.Now(),
		ObservedGeneration: virtBMC.Generation,
	}

	if err := r.updateStatusIfNeeded(ctx, log, &virtBMC, foundSvc.Spec.ClusterIP, ready); err != nil {
		log.Error(err, "Failed to update VirtualMachineBMC status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{RequeueAfter: 5 * time.Minute}, nil
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

func (r *VirtualMachineBMCReconciler) updateStatusIfNeeded(ctx context.Context, log logr.Logger, bmc *bmcv1beta1.VirtualMachineBMC, clusterIP string, condition metav1.Condition) error {
	updated := false

	if clusterIP != "" && bmc.Status.ClusterIP != clusterIP {
		bmc.Status.ClusterIP = clusterIP
		updated = true
	}

	found := false
	for i, cond := range bmc.Status.Conditions {
		if cond.Type == condition.Type {
			found = true
			if cond.Status != condition.Status || cond.Message != condition.Message {
				bmc.Status.Conditions[i] = condition
				updated = true
			}
			break
		}
	}
	if !found {
		bmc.Status.Conditions = append(bmc.Status.Conditions, condition)
		updated = true
	}

	if updated {
		log.Info("Updating VirtualMachineBMC status")
		if err := r.Status().Update(ctx, bmc); err != nil {
			log.Error(err, "Failed to update VirtualMachineBMC status")
			return err
		}
	}

	return nil
}
