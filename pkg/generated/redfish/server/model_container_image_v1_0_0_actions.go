// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ContainerImageV100Actions - The available actions for this resource.
type ContainerImageV100Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertContainerImageV100ActionsRequired checks if the required fields are not zero-ed
func AssertContainerImageV100ActionsRequired(obj ContainerImageV100Actions) error {
	return nil
}

// AssertContainerImageV100ActionsConstraints checks if the values respects the defined constraints
func AssertContainerImageV100ActionsConstraints(obj ContainerImageV100Actions) error {
	return nil
}