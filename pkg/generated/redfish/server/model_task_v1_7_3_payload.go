// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// TaskV173Payload - The HTTP and JSON payload details for this Task.
type TaskV173Payload struct {

	// An array of HTTP headers that this task includes.
	HttpHeaders []string `json:"HttpHeaders,omitempty"`

	// The HTTP operation to perform to execute this task.
	HttpOperation string `json:"HttpOperation,omitempty"`

	// The JSON payload to use in the execution of this task.
	JsonBody string `json:"JsonBody,omitempty"`

	// The URI of the target for this task.
	TargetUri string `json:"TargetUri,omitempty"`
}

// AssertTaskV173PayloadRequired checks if the required fields are not zero-ed
func AssertTaskV173PayloadRequired(obj TaskV173Payload) error {
	return nil
}

// AssertTaskV173PayloadConstraints checks if the values respects the defined constraints
func AssertTaskV173PayloadConstraints(obj TaskV173Payload) error {
	return nil
}