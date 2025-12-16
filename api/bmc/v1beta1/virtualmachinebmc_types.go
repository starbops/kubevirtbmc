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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Condition type constant.
const (
	ConditionReady                   = "ServiceReady"
	ConditionNotReady                = "ServiceNotReady"
	ConditionVirtualMachineAvailable = "VirtualMachineAvailable"
	ConditionSecretAvailable         = "SecretAvailable"
)

// VirtualMachineBMCSpec defines the desired state of VirtualMachineBMC.
type VirtualMachineBMCSpec struct {
	// BMC Service configuration
	// +optional
	Service *BMCServiceSpec `json:"service,omitempty"`

	// Reference to the Secret containing IPMI/Redfish credentials
	// +Required
	AuthSecretRef *corev1.LocalObjectReference `json:"authSecretRef"`

	// Reference to the VM to manage
	// +Required
	VirtualMachineRef *corev1.LocalObjectReference `json:"virtualMachineRef"`
}

// Service configuration for the BMC service.
type BMCServiceSpec struct {
	// Type of service
	// +kubebuilder:default=ClusterIP
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	Type corev1.ServiceType `json:"type,omitempty"`

	// Additional labels to apply to the service
	Labels map[string]string `json:"labels,omitempty"`
}

// VirtualMachineBMCStatus defines the observed state of VirtualMachineBMC.
type VirtualMachineBMCStatus struct {
	// IP address assigned to the LoadBalancer Type BMC service
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`

	// IP address assigned to the ClusterIP Type BMC service
	ClusterIP string `json:"clusterIP,omitempty"`

	// List of current conditions (e.g., Ready)
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=vmbmc;vmbmcs
// +kubebuilder:printcolumn:name="VIRTUALMACHINE",type="string",JSONPath=`.spec.virtualMachineRef.name`
// +kubebuilder:printcolumn:name="SECRET",type="string",JSONPath=`.spec.authSecretRef.name`
// +kubebuilder:printcolumn:name="CLUSTERIP",type="string",JSONPath=`.status.clusterIP`
// +kubebuilder:printcolumn:name="SERVICEREADY",type="string",JSONPath=`.status.conditions[?(@.type=='ServiceReady')].status`

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
