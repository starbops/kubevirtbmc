// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SwitchMetricsV101Actions - The available actions for this resource.
type SwitchMetricsV101Actions struct {
	SwitchMetricsClearCurrentPeriod SwitchMetricsV101ClearCurrentPeriod `json:"#SwitchMetrics.ClearCurrentPeriod,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertSwitchMetricsV101ActionsRequired checks if the required fields are not zero-ed
func AssertSwitchMetricsV101ActionsRequired(obj SwitchMetricsV101Actions) error {
	if err := AssertSwitchMetricsV101ClearCurrentPeriodRequired(obj.SwitchMetricsClearCurrentPeriod); err != nil {
		return err
	}
	return nil
}

// AssertSwitchMetricsV101ActionsConstraints checks if the values respects the defined constraints
func AssertSwitchMetricsV101ActionsConstraints(obj SwitchMetricsV101Actions) error {
	if err := AssertSwitchMetricsV101ClearCurrentPeriodConstraints(obj.SwitchMetricsClearCurrentPeriod); err != nil {
		return err
	}
	return nil
}
