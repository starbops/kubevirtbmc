// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// FabricAdapterV152Actions - The available actions for this resource.
type FabricAdapterV152Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertFabricAdapterV152ActionsRequired checks if the required fields are not zero-ed
func AssertFabricAdapterV152ActionsRequired(obj FabricAdapterV152Actions) error {
	return nil
}

// AssertFabricAdapterV152ActionsConstraints checks if the values respects the defined constraints
func AssertFabricAdapterV152ActionsConstraints(obj FabricAdapterV152Actions) error {
	return nil
}