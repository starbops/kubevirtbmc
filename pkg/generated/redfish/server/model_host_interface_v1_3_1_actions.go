// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// HostInterfaceV131Actions - The available actions for this resource.
type HostInterfaceV131Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertHostInterfaceV131ActionsRequired checks if the required fields are not zero-ed
func AssertHostInterfaceV131ActionsRequired(obj HostInterfaceV131Actions) error {
	return nil
}

// AssertHostInterfaceV131ActionsConstraints checks if the values respects the defined constraints
func AssertHostInterfaceV131ActionsConstraints(obj HostInterfaceV131Actions) error {
	return nil
}
