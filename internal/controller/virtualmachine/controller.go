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

package virtualmachine

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kubevirtv1 "kubevirt.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	virtualmachinev1 "zespre.com/kubebmc/api/v1"
	ctlkubebmc "zespre.com/kubebmc/internal/controller/kubebmc"
)

const (
	finalizerName = "kubebmc-virtualmachine-controller"
)

// VirtualMachineReconciler reconciles a VirtualMachine object
type VirtualMachineReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (v *VirtualMachineReconciler) handleFinalizer(ctx context.Context, vm *kubevirtv1.VirtualMachine) error {
	log := log.FromContext(ctx)
	if vm.ObjectMeta.DeletionTimestamp.IsZero() {
		// add finalizer
		if !controllerutil.ContainsFinalizer(vm, finalizerName) {
			ok := controllerutil.AddFinalizer(vm, finalizerName)
			log.V(1).Info(fmt.Sprintf("add finalizer %s : %t", finalizerName, ok))
			return v.Update(ctx, vm)
		}
	} else {
		// remove finalizer
		if controllerutil.ContainsFinalizer(vm, finalizerName) {
			ok := controllerutil.RemoveFinalizer(vm, finalizerName)
			log.V(1).Info(fmt.Sprintf("remove finalizer %s : %t", finalizerName, ok))
			return v.Update(ctx, vm)
		}
	}
	return nil
}

func (v *VirtualMachineReconciler) constructKubeBMCForVirtualMachine(vm *kubevirtv1.VirtualMachine) (*virtualmachinev1.KubeBMC, error) {
	name := fmt.Sprintf("%s-%s", vm.Namespace, vm.Name)

	kubeBMC := &virtualmachinev1.KubeBMC{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels: map[string]string{
				ctlkubebmc.ManagedByLabel: "kubebmc-controller",
			},
			Name:      name,
			Namespace: ctlkubebmc.KBMCNamespace,
		},
		Spec: virtualmachinev1.KubeBMCSpec{
			Username:                ctlkubebmc.DefaultUsername,
			Password:                ctlkubebmc.DefaultPassword,
			VirtualMachineNamespace: vm.Namespace,
			VirtualMachineName:      vm.Name,
		},
	}

	return kubeBMC, nil
}

//+kubebuilder:rbac:groups=kubevirt.io,resources=virtualmachines,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=virtualmachine.zespre.com,resources=kubebmcs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=virtualmachine.zespre.com,resources=kubebmcs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=virtualmachine.zespre.com,resources=kubebmcs/finalizers,verbs=update

func (v *VirtualMachineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var vm kubevirtv1.VirtualMachine
	if err := v.Get(ctx, req.NamespacedName, &vm); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch VirtualMachine")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err := v.handleFinalizer(ctx, &vm); err != nil {
		log.Error(err, "failed to update finalizer")
		return ctrl.Result{}, err
	}

	// Prepare the KubeBMC
	kubeBMC, err := v.constructKubeBMCForVirtualMachine(&vm)
	if err != nil {
		log.Error(err, "unable to construct kubebmc from template")
		return ctrl.Result{}, err
	}

	// Create the KubeBMC on the cluster
	if err := v.Create(ctx, kubeBMC); err != nil && !apierrors.IsAlreadyExists(err) {
		log.Error(err, "unable to create KubeBMC for VirtualMachine", "kubeBMC", kubeBMC)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created KubeBMC for VirtualMachine", "kubeBMC", kubeBMC)

	knn := types.NamespacedName{
		Namespace: ctlkubebmc.KBMCNamespace,
		Name:      fmt.Sprintf("%s-%s", vm.Namespace, vm.Name),
	}
	var kbmc virtualmachinev1.KubeBMC
	if err := v.Get(ctx, knn, &kbmc); err != nil {
		log.Error(err, "unable to fetch KubeBMC")
		return ctrl.Result{}, err
	}

	if !vm.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := v.Delete(ctx, &kbmc); err != nil {
			log.Error(err, "unable to delete KubeBMC for VirtualMachine", "kbmc", kbmc)
			return ctrl.Result{}, err
		}
		log.V(1).Info("removed KubeBMC for VirtualMachine", "kbmc", kbmc)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (v *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubevirtv1.VirtualMachine{}).
		Complete(v)
}
