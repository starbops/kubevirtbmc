// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// NetworkAdapterMetricsV101Actions - The available actions for this resource.
type NetworkAdapterMetricsV101Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertNetworkAdapterMetricsV101ActionsRequired checks if the required fields are not zero-ed
func AssertNetworkAdapterMetricsV101ActionsRequired(obj NetworkAdapterMetricsV101Actions) error {
	return nil
}

// AssertNetworkAdapterMetricsV101ActionsConstraints checks if the values respects the defined constraints
func AssertNetworkAdapterMetricsV101ActionsConstraints(obj NetworkAdapterMetricsV101Actions) error {
	return nil
}
