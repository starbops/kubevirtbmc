// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PowerDistributionMetricsV131ResetMetrics - This action resets the summary metrics related to this equipment.
type PowerDistributionMetricsV131ResetMetrics struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertPowerDistributionMetricsV131ResetMetricsRequired checks if the required fields are not zero-ed
func AssertPowerDistributionMetricsV131ResetMetricsRequired(obj PowerDistributionMetricsV131ResetMetrics) error {
	return nil
}

// AssertPowerDistributionMetricsV131ResetMetricsConstraints checks if the values respects the defined constraints
func AssertPowerDistributionMetricsV131ResetMetricsConstraints(obj PowerDistributionMetricsV131ResetMetrics) error {
	return nil
}
