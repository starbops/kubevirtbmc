// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AggregateV102Reset - This action is used to reset a collection of resources.  For example, this could be an aggregate or a list of computer systems.
type AggregateV102Reset struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertAggregateV102ResetRequired checks if the required fields are not zero-ed
func AssertAggregateV102ResetRequired(obj AggregateV102Reset) error {
	return nil
}

// AssertAggregateV102ResetConstraints checks if the values respects the defined constraints
func AssertAggregateV102ResetConstraints(obj AggregateV102Reset) error {
	return nil
}