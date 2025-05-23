// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// JobServiceV105JobServiceCapabilities - The supported capabilities of this job service implementation.
type JobServiceV105JobServiceCapabilities struct {

	// The maximum number of jobs supported.
	MaxJobs *int64 `json:"MaxJobs,omitempty"`

	// The maximum number of job steps supported.
	MaxSteps *int64 `json:"MaxSteps,omitempty"`

	// An indication of whether scheduling of jobs is supported.
	Scheduling *bool `json:"Scheduling,omitempty"`
}

// AssertJobServiceV105JobServiceCapabilitiesRequired checks if the required fields are not zero-ed
func AssertJobServiceV105JobServiceCapabilitiesRequired(obj JobServiceV105JobServiceCapabilities) error {
	return nil
}

// AssertJobServiceV105JobServiceCapabilitiesConstraints checks if the values respects the defined constraints
func AssertJobServiceV105JobServiceCapabilitiesConstraints(obj JobServiceV105JobServiceCapabilities) error {
	return nil
}
