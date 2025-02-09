// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// StorageControllerMetricsV102Actions - The available actions for this resource.
type StorageControllerMetricsV102Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertStorageControllerMetricsV102ActionsRequired checks if the required fields are not zero-ed
func AssertStorageControllerMetricsV102ActionsRequired(obj StorageControllerMetricsV102Actions) error {
	return nil
}

// AssertStorageControllerMetricsV102ActionsConstraints checks if the values respects the defined constraints
func AssertStorageControllerMetricsV102ActionsConstraints(obj StorageControllerMetricsV102Actions) error {
	return nil
}
