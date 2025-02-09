// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CapacityV121Actions - The available actions for this resource.
type CapacityV121Actions struct {

	// The available OEM specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCapacityV121ActionsRequired checks if the required fields are not zero-ed
func AssertCapacityV121ActionsRequired(obj CapacityV121Actions) error {
	return nil
}

// AssertCapacityV121ActionsConstraints checks if the values respects the defined constraints
func AssertCapacityV121ActionsConstraints(obj CapacityV121Actions) error {
	return nil
}
