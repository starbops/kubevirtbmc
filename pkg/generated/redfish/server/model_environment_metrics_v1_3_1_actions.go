// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EnvironmentMetricsV131Actions - The available actions for this resource.
type EnvironmentMetricsV131Actions struct {
	EnvironmentMetricsResetMetrics EnvironmentMetricsV131ResetMetrics `json:"#EnvironmentMetrics.ResetMetrics,omitempty"`

	EnvironmentMetricsResetToDefaults EnvironmentMetricsV131ResetToDefaults `json:"#EnvironmentMetrics.ResetToDefaults,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertEnvironmentMetricsV131ActionsRequired checks if the required fields are not zero-ed
func AssertEnvironmentMetricsV131ActionsRequired(obj EnvironmentMetricsV131Actions) error {
	if err := AssertEnvironmentMetricsV131ResetMetricsRequired(obj.EnvironmentMetricsResetMetrics); err != nil {
		return err
	}
	if err := AssertEnvironmentMetricsV131ResetToDefaultsRequired(obj.EnvironmentMetricsResetToDefaults); err != nil {
		return err
	}
	return nil
}

// AssertEnvironmentMetricsV131ActionsConstraints checks if the values respects the defined constraints
func AssertEnvironmentMetricsV131ActionsConstraints(obj EnvironmentMetricsV131Actions) error {
	if err := AssertEnvironmentMetricsV131ResetMetricsConstraints(obj.EnvironmentMetricsResetMetrics); err != nil {
		return err
	}
	if err := AssertEnvironmentMetricsV131ResetToDefaultsConstraints(obj.EnvironmentMetricsResetToDefaults); err != nil {
		return err
	}
	return nil
}