// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryRegionV101Actions - The available actions for this resource.
type MemoryRegionV101Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertMemoryRegionV101ActionsRequired checks if the required fields are not zero-ed
func AssertMemoryRegionV101ActionsRequired(obj MemoryRegionV101Actions) error {
	return nil
}

// AssertMemoryRegionV101ActionsConstraints checks if the values respects the defined constraints
func AssertMemoryRegionV101ActionsConstraints(obj MemoryRegionV101Actions) error {
	return nil
}
