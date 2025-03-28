// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MessageRegistryFileV114Actions - The available actions for this resource.
type MessageRegistryFileV114Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertMessageRegistryFileV114ActionsRequired checks if the required fields are not zero-ed
func AssertMessageRegistryFileV114ActionsRequired(obj MessageRegistryFileV114Actions) error {
	return nil
}

// AssertMessageRegistryFileV114ActionsConstraints checks if the values respects the defined constraints
func AssertMessageRegistryFileV114ActionsConstraints(obj MessageRegistryFileV114Actions) error {
	return nil
}
