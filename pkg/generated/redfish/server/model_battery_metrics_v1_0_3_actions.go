// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// BatteryMetricsV103Actions - The available actions for this resource.
type BatteryMetricsV103Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertBatteryMetricsV103ActionsRequired checks if the required fields are not zero-ed
func AssertBatteryMetricsV103ActionsRequired(obj BatteryMetricsV103Actions) error {
	return nil
}

// AssertBatteryMetricsV103ActionsConstraints checks if the values respects the defined constraints
func AssertBatteryMetricsV103ActionsConstraints(obj BatteryMetricsV103Actions) error {
	return nil
}