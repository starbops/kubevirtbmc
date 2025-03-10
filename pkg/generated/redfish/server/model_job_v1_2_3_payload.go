// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// JobV123Payload - The HTTP and JSON payload details for this job.
type JobV123Payload struct {

	// An array of HTTP headers in this job.
	HttpHeaders []string `json:"HttpHeaders,omitempty"`

	// The HTTP operation that executes this job.
	HttpOperation string `json:"HttpOperation,omitempty"`

	// The JSON payload to use in the execution of this job.
	JsonBody string `json:"JsonBody,omitempty"`

	// The link to the target for this job.
	TargetUri string `json:"TargetUri,omitempty"`
}

// AssertJobV123PayloadRequired checks if the required fields are not zero-ed
func AssertJobV123PayloadRequired(obj JobV123Payload) error {
	return nil
}

// AssertJobV123PayloadConstraints checks if the values respects the defined constraints
func AssertJobV123PayloadConstraints(obj JobV123Payload) error {
	return nil
}
