// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// OutboundConnectionV101Actions - The available actions for this resource.
type OutboundConnectionV101Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertOutboundConnectionV101ActionsRequired checks if the required fields are not zero-ed
func AssertOutboundConnectionV101ActionsRequired(obj OutboundConnectionV101Actions) error {
	return nil
}

// AssertOutboundConnectionV101ActionsConstraints checks if the values respects the defined constraints
func AssertOutboundConnectionV101ActionsConstraints(obj OutboundConnectionV101Actions) error {
	return nil
}
