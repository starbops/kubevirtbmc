// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// TelemetryServiceV133ClearMetricReports - The action to clear the metric reports for this telemetry service.
type TelemetryServiceV133ClearMetricReports struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertTelemetryServiceV133ClearMetricReportsRequired checks if the required fields are not zero-ed
func AssertTelemetryServiceV133ClearMetricReportsRequired(obj TelemetryServiceV133ClearMetricReports) error {
	return nil
}

// AssertTelemetryServiceV133ClearMetricReportsConstraints checks if the values respects the defined constraints
func AssertTelemetryServiceV133ClearMetricReportsConstraints(obj TelemetryServiceV133ClearMetricReports) error {
	return nil
}
