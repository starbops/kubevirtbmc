// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AggregateV102RemoveElements - This action is used to remove one or more resources from the aggregate.
type AggregateV102RemoveElements struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertAggregateV102RemoveElementsRequired checks if the required fields are not zero-ed
func AssertAggregateV102RemoveElementsRequired(obj AggregateV102RemoveElements) error {
	return nil
}

// AssertAggregateV102RemoveElementsConstraints checks if the values respects the defined constraints
func AssertAggregateV102RemoveElementsConstraints(obj AggregateV102RemoveElements) error {
	return nil
}