// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ProcessorMetricsV163ClearCurrentPeriod - This action sets the CurrentPeriod property's values to 0.
type ProcessorMetricsV163ClearCurrentPeriod struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertProcessorMetricsV163ClearCurrentPeriodRequired checks if the required fields are not zero-ed
func AssertProcessorMetricsV163ClearCurrentPeriodRequired(obj ProcessorMetricsV163ClearCurrentPeriod) error {
	return nil
}

// AssertProcessorMetricsV163ClearCurrentPeriodConstraints checks if the values respects the defined constraints
func AssertProcessorMetricsV163ClearCurrentPeriodConstraints(obj ProcessorMetricsV163ClearCurrentPeriod) error {
	return nil
}