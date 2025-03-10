// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SecureBootDatabaseV102ResetKeys - This action is used to reset the UEFI Secure Boot keys of this database.
type SecureBootDatabaseV102ResetKeys struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertSecureBootDatabaseV102ResetKeysRequired checks if the required fields are not zero-ed
func AssertSecureBootDatabaseV102ResetKeysRequired(obj SecureBootDatabaseV102ResetKeys) error {
	return nil
}

// AssertSecureBootDatabaseV102ResetKeysConstraints checks if the values respects the defined constraints
func AssertSecureBootDatabaseV102ResetKeysConstraints(obj SecureBootDatabaseV102ResetKeys) error {
	return nil
}
