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
	virtualmachinev1 "kubevirt.org/kubevirtbmc/api/v1"
	ctlvirtualmachinebmc "kubevirt.org/kubevirtbmc/internal/controller/virtualmachinebmc"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	finalizerName = "virtualmachinebmc-virtualmachine-controller"
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

func (v *VirtualMachineReconciler) constructVirtualMachineBMCForVirtualMachine(vm *kubevirtv1.VirtualMachine) (*virtualmachinev1.VirtualMachineBMC, error) {
	name := fmt.Sprintf("%s-%s", vm.Namespace, vm.Name)

	kubeBMC := &virtualmachinev1.VirtualMachineBMC{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: make(map[string]string),
			Labels: map[string]string{
				ctlvirtualmachinebmc.ManagedByLabel: "virt-bmc-controller",
			},
			Name:      name,
			Namespace: ctlvirtualmachinebmc.KBMCNamespace,
		},
		Spec: virtualmachinev1.VirtualMachineBMCSpec{
			Username:                ctlvirtualmachinebmc.DefaultUsername,
			Password:                ctlvirtualmachinebmc.DefaultPassword,
			VirtualMachineNamespace: vm.Namespace,
			VirtualMachineName:      vm.Name,
		},
	}

	return kubeBMC, nil
}

//+kubebuilder:rbac:groups=kubevirt.io,resources=virtualmachines,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=virtualmachine.kubevirt.org,resources=virtualmachinebmcs/finalizers,verbs=update

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

	// Prepare the VirtualMachineBMC
	kubeBMC, err := v.constructVirtualMachineBMCForVirtualMachine(&vm)
	if err != nil {
		log.Error(err, "unable to construct virtualmachinebmc from template")
		return ctrl.Result{}, err
	}

	// Create the VirtualMachineBMC on the cluster
	if err := v.Create(ctx, kubeBMC); err != nil && !apierrors.IsAlreadyExists(err) {
		log.Error(err, "unable to create VirtualMachineBMC for VirtualMachine", "kubeBMC", kubeBMC)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created VirtualMachineBMC for VirtualMachine", "kubeBMC", kubeBMC)

	knn := types.NamespacedName{
		Namespace: ctlvirtualmachinebmc.KBMCNamespace,
		Name:      fmt.Sprintf("%s-%s", vm.Namespace, vm.Name),
	}
	var virt-bmc virtualmachinev1.VirtualMachineBMC
	if err := v.Get(ctx, knn, &virt-bmc); err != nil {
		log.Error(err, "unable to fetch VirtualMachineBMC")
		return ctrl.Result{}, err
	}

	if !vm.ObjectMeta.DeletionTimestamp.IsZero() {
		if err := v.Delete(ctx, &virt-bmc); err != nil {
			log.Error(err, "unable to delete VirtualMachineBMC for VirtualMachine", "virt-bmc", virt-bmc)
			return ctrl.Result{}, err
		}
		log.V(1).Info("removed VirtualMachineBMC for VirtualMachine", "virt-bmc", virt-bmc)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (v *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubevirtv1.VirtualMachine{}).
		Complete(v)
}
