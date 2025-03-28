// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// StorageV1151SetEncryptionKey - This action sets the local encryption key for the storage subsystem.
type StorageV1151SetEncryptionKey struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertStorageV1151SetEncryptionKeyRequired checks if the required fields are not zero-ed
func AssertStorageV1151SetEncryptionKeyRequired(obj StorageV1151SetEncryptionKey) error {
	return nil
}

// AssertStorageV1151SetEncryptionKeyConstraints checks if the values respects the defined constraints
func AssertStorageV1151SetEncryptionKeyConstraints(obj StorageV1151SetEncryptionKey) error {
	return nil
}
