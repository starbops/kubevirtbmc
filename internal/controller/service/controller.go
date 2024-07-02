/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	virtualmachinev1 "kubevirt.org/kubevirtbmc/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ServiceReconciler reconciles a Service object
type ServiceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs,verbs=get;list;watch
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs/status,verbs=get;update;patch

func (s *ServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var svc corev1.Service
	if err := s.Get(ctx, req.NamespacedName, &svc); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch Service")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if svc.Labels == nil {
		return ctrl.Result{}, nil
	}
	virt - bmcName, ok := svc.Labels[ctlvirtualmachinebmc.virt-bmcNameLabel]
	if !ok {
		return ctrl.Result{}, nil
	}
	knn := types.NamespacedName{
		Namespace: req.Namespace,
		Name:      virt - bmcName,
	}

	var kubeBMC virtualmachinev1.VirtualMachineBMC
	if err := s.Get(ctx, knn, &kubeBMC); err != nil {
		log.Error(err, "unable to fetch VirtualMachineBMC")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Update VirtualMachineBMC status
	if svc.Spec.ClusterIP == "" {
		return ctrl.Result{RequeueAfter: time.Duration(time.Second * 10)}, fmt.Errorf("clusterIP is not ready yet")
	}
	kubeBMC.Status.Ready = true
	kubeBMC.Status.ServiceIP = svc.Spec.ClusterIP
	if err := s.Status().Update(ctx, &kubeBMC); err != nil {
		log.Error(err, "unable to update VirtualMachineBMC status")
		return ctrl.Result{}, err
	}

	log.V(1).Info("updated VirtualMachineBMC status for Service", "kubeBMC", kubeBMC)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (s *ServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Service{}).
		Complete(s)
}
