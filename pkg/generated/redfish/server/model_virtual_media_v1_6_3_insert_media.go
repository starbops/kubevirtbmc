// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VirtualMediaV163InsertMedia - This action attaches remote media to virtual media.
type VirtualMediaV163InsertMedia struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertVirtualMediaV163InsertMediaRequired checks if the required fields are not zero-ed
func AssertVirtualMediaV163InsertMediaRequired(obj VirtualMediaV163InsertMedia) error {
	return nil
}

// AssertVirtualMediaV163InsertMediaConstraints checks if the values respects the defined constraints
func AssertVirtualMediaV163InsertMediaConstraints(obj VirtualMediaV163InsertMedia) error {
	return nil
}