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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineRef defines the virtual machine to be controlled by the BMC.
type VirtualMachineRef struct {
	// Name of the virtual machine.
	Name string `json:"name"`
}

// AuthSecretRef defines the Kubernetes Secret containing BMC credentials.
type AuthSecretRef struct {
	// Name of the Secret resource.
	Name string `json:"name"`
}

// VirtualMachineBMCSpec defines the desired state of VirtualMachineBMC.
type VirtualMachineBMCSpec struct {
	// Reference to the VM to manage.
	VirtualMachine VirtualMachineRef `json:"virtualMachine"`

	// Reference to the Secret containing IPMI/Redfish credentials.
	AuthSecret AuthSecretRef `json:"authSecret"`
}

// Condition describes a status condition of the BMC (e.g., Ready).
type Condition struct {
	// Type of condition, e.g., Ready
	Type string `json:"type"`

	// Status of the condition, one of True, False, Unknown
	Status string `json:"status"`

	// Reason for last condition change
	Reason string `json:"reason,omitempty"`

	// Message describing details of transition
	Message string `json:"message,omitempty"`

	// Last time the condition transitioned
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// VirtualMachineBMCStatus defines the observed state of VirtualMachineBMC.
type VirtualMachineBMCStatus struct {
	// IP address exposed by the BMC service
	ClusterIP string `json:"clusterIP,omitempty"`

	// List of current conditions (e.g., Ready)
	Conditions []Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VirtualMachineBMC is the Schema for the virtualmachinebmcs API
type VirtualMachineBMC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineBMCSpec   `json:"spec,omitempty"`
	Status VirtualMachineBMCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineBMCList contains a list of VirtualMachineBMC
type VirtualMachineBMCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineBMC `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineBMC{}, &VirtualMachineBMCList{})
}
