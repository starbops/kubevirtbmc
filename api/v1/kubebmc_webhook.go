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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var kubebmclog = logf.Log.WithName("kubebmc-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *KubeBMC) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-virtualmachine-zespre-com-v1-kubebmc,mutating=true,failurePolicy=fail,sideEffects=None,groups=virtualmachine.zespre.com,resources=kubebmcs,verbs=create;update,versions=v1,name=mkubebmc.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &KubeBMC{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *KubeBMC) Default() {
	kubebmclog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	if r.Spec.Username == "" {
		r.Spec.Username = "admin"
	}
	if r.Spec.Password == "" {
		r.Spec.Password = "password"
	}
	if r.Spec.VirtualMachineNamespace == "" {
		r.Spec.VirtualMachineNamespace = "default"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-virtualmachine-zespre-com-v1-kubebmc,mutating=false,failurePolicy=fail,sideEffects=None,groups=virtualmachine.zespre.com,resources=kubebmcs,verbs=create;update,versions=v1,name=vkubebmc.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &KubeBMC{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *KubeBMC) ValidateCreate() (admission.Warnings, error) {
	kubebmclog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *KubeBMC) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	kubebmclog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *KubeBMC) ValidateDelete() (admission.Warnings, error) {
	kubebmclog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
