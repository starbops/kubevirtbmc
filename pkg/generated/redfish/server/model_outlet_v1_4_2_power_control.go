// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// OutletV142PowerControl - This action turns the outlet on or off.
type OutletV142PowerControl struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertOutletV142PowerControlRequired checks if the required fields are not zero-ed
func AssertOutletV142PowerControlRequired(obj OutletV142PowerControl) error {
	return nil
}

// AssertOutletV142PowerControlConstraints checks if the values respects the defined constraints
func AssertOutletV142PowerControlConstraints(obj OutletV142PowerControl) error {
	return nil
}