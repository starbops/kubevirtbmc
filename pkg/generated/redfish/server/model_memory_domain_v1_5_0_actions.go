// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryDomainV150Actions - The available actions for this Resource.
type MemoryDomainV150Actions struct {

	// The available OEM-specific actions for this Resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertMemoryDomainV150ActionsRequired checks if the required fields are not zero-ed
func AssertMemoryDomainV150ActionsRequired(obj MemoryDomainV150Actions) error {
	return nil
}

// AssertMemoryDomainV150ActionsConstraints checks if the values respects the defined constraints
func AssertMemoryDomainV150ActionsConstraints(obj MemoryDomainV150Actions) error {
	return nil
}