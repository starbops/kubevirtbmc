// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// TaskV174Actions - The available actions for this resource.
type TaskV174Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertTaskV174ActionsRequired checks if the required fields are not zero-ed
func AssertTaskV174ActionsRequired(obj TaskV174Actions) error {
	return nil
}

// AssertTaskV174ActionsConstraints checks if the values respects the defined constraints
func AssertTaskV174ActionsConstraints(obj TaskV174Actions) error {
	return nil
}