/*
Copyright 2024.

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

package v1alpha1

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var virtualmachinebmclog = logf.Log.WithName("virtualmachinebmc-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *VirtualMachineBMC) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-virtualmachine-kubevirt-io-v1alpha1-virtualmachinebmc,mutating=true,failurePolicy=fail,sideEffects=None,groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs,verbs=create;update,versions=v1alpha1,name=mvirtualmachinebmc.kb.io,admissionReviewVersions=v1

var _ webhook.CustomDefaulter = &VirtualMachineBMC{}

// Default implements webhook.CustomDefaulter
func (r *VirtualMachineBMC) Default(ctx context.Context, obj runtime.Object) error {
	virtualmachinebmclog.Info("default", "name", r.Name)

	// Cast the runtime.Object to your type
	vm, ok := obj.(*VirtualMachineBMC)
	if !ok {
		return fmt.Errorf("expected a VirtualMachineBMC object but got %T", obj)
	}

	// Defaulting logic
	if vm.Spec.Username == "" {
		vm.Spec.Username = "admin"
	}
	if vm.Spec.Password == "" {
		vm.Spec.Password = "password"
	}
	if vm.Spec.VirtualMachineNamespace == "" {
		vm.Spec.VirtualMachineNamespace = "default"
	}

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-virtualmachine-kubevirt-io-v1alpha1-virtualmachinebmc,mutating=false,failurePolicy=fail,sideEffects=None,groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs,verbs=create;update,versions=v1alpha1,name=vvirtualmachinebmc.kb.io,admissionReviewVersions=v1

var _ webhook.CustomValidator = &VirtualMachineBMC{}

// ValidateCreate implements webhook.CustomValidator
func (r *VirtualMachineBMC) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	virtualmachinebmclog.Info("validate create", "name", r.Name)
	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator
func (r *VirtualMachineBMC) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	virtualmachinebmclog.Info("validate update", "name", r.Name)
	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator
func (r *VirtualMachineBMC) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	virtualmachinebmclog.Info("validate delete", "name", r.Name)
	return nil, nil
}
