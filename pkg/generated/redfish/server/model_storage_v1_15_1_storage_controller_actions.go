// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// StorageV1151StorageControllerActions - The available actions for this resource.
type StorageV1151StorageControllerActions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertStorageV1151StorageControllerActionsRequired checks if the required fields are not zero-ed
func AssertStorageV1151StorageControllerActionsRequired(obj StorageV1151StorageControllerActions) error {
	return nil
}

// AssertStorageV1151StorageControllerActionsConstraints checks if the values respects the defined constraints
func AssertStorageV1151StorageControllerActionsConstraints(obj StorageV1151StorageControllerActions) error {
	return nil
}