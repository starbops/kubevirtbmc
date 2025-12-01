/*
Copyright 2025.

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

	virtualmachinev1alpha1 "kubevirt.io/kubevirtbmc/api/bmc/v1beta1"
)

// nolint:unused
// log is for logging in this package.
var virtualmachinebmclog = logf.Log.WithName("virtualmachinebmc-resource")

// SetupVirtualMachineBMCWebhookWithManager registers the webhook for VirtualMachineBMC in the manager.
func SetupVirtualMachineBMCWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&virtualmachinev1alpha1.VirtualMachineBMC{}).
		WithValidator(&VirtualMachineBMCCustomValidator{}).
		WithDefaulter(&VirtualMachineBMCCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-virtualmachine-kubevirt-io-v1alpha1-virtualmachinebmc,mutating=true,failurePolicy=fail,sideEffects=None,groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs,verbs=create;update,versions=v1alpha1,name=mvirtualmachinebmc-v1alpha1.kb.io,admissionReviewVersions=v1

// VirtualMachineBMCCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind VirtualMachineBMC when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type VirtualMachineBMCCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &VirtualMachineBMCCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind VirtualMachineBMC.
func (d *VirtualMachineBMCCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	virtualmachinebmc, ok := obj.(*virtualmachinev1alpha1.VirtualMachineBMC)

	if !ok {
		return fmt.Errorf("expected an VirtualMachineBMC object but got %T", obj)
	}
	virtualmachinebmclog.Info("Defaulting for VirtualMachineBMC", "name", virtualmachinebmc.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-virtualmachine-kubevirt-io-v1alpha1-virtualmachinebmc,mutating=false,failurePolicy=fail,sideEffects=None,groups=virtualmachine.kubevirt.io,resources=virtualmachinebmcs,verbs=create;update,versions=v1alpha1,name=vvirtualmachinebmc-v1alpha1.kb.io,admissionReviewVersions=v1

// VirtualMachineBMCCustomValidator struct is responsible for validating the VirtualMachineBMC resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type VirtualMachineBMCCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &VirtualMachineBMCCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type VirtualMachineBMC.
func (v *VirtualMachineBMCCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	virtualmachinebmc, ok := obj.(*virtualmachinev1alpha1.VirtualMachineBMC)
	if !ok {
		return nil, fmt.Errorf("expected a VirtualMachineBMC object but got %T", obj)
	}
	virtualmachinebmclog.Info("Validation for VirtualMachineBMC upon creation", "name", virtualmachinebmc.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type VirtualMachineBMC.
func (v *VirtualMachineBMCCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	virtualmachinebmc, ok := newObj.(*virtualmachinev1alpha1.VirtualMachineBMC)
	if !ok {
		return nil, fmt.Errorf("expected a VirtualMachineBMC object for the newObj but got %T", newObj)
	}
	virtualmachinebmclog.Info("Validation for VirtualMachineBMC upon update", "name", virtualmachinebmc.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type VirtualMachineBMC.
func (v *VirtualMachineBMCCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	virtualmachinebmc, ok := obj.(*virtualmachinev1alpha1.VirtualMachineBMC)
	if !ok {
		return nil, fmt.Errorf("expected a VirtualMachineBMC object but got %T", obj)
	}
	virtualmachinebmclog.Info("Validation for VirtualMachineBMC upon deletion", "name", virtualmachinebmc.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
