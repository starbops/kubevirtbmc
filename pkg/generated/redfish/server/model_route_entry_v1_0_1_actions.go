// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// RouteEntryV101Actions - The available actions for this Resource.
type RouteEntryV101Actions struct {

	// The available OEM-specific actions for this Resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertRouteEntryV101ActionsRequired checks if the required fields are not zero-ed
func AssertRouteEntryV101ActionsRequired(obj RouteEntryV101Actions) error {
	return nil
}

// AssertRouteEntryV101ActionsConstraints checks if the values respects the defined constraints
func AssertRouteEntryV101ActionsConstraints(obj RouteEntryV101Actions) error {
	return nil
}
