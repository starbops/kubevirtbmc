// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// ManagerV1192ModifyRedundancySet - The `ModifyRedundancySet` operation adds members to or removes members from a redundant group of managers.
type ManagerV1192ModifyRedundancySet struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertManagerV1192ModifyRedundancySetRequired checks if the required fields are not zero-ed
func AssertManagerV1192ModifyRedundancySetRequired(obj ManagerV1192ModifyRedundancySet) error {
	return nil
}

// AssertManagerV1192ModifyRedundancySetConstraints checks if the values respects the defined constraints
func AssertManagerV1192ModifyRedundancySetConstraints(obj ManagerV1192ModifyRedundancySet) error {
	return nil
}