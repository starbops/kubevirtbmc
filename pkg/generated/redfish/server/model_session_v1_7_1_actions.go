// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SessionV171Actions - The available actions for this resource.
type SessionV171Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertSessionV171ActionsRequired checks if the required fields are not zero-ed
func AssertSessionV171ActionsRequired(obj SessionV171Actions) error {
	return nil
}

// AssertSessionV171ActionsConstraints checks if the values respects the defined constraints
func AssertSessionV171ActionsConstraints(obj SessionV171Actions) error {
	return nil
}