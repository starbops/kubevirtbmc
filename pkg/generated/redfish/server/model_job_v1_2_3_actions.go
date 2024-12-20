// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// JobV123Actions - The available actions for this resource.
type JobV123Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertJobV123ActionsRequired checks if the required fields are not zero-ed
func AssertJobV123ActionsRequired(obj JobV123Actions) error {
	return nil
}

// AssertJobV123ActionsConstraints checks if the values respects the defined constraints
func AssertJobV123ActionsConstraints(obj JobV123Actions) error {
	return nil
}
