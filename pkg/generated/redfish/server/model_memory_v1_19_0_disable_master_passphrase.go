// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryV1190DisableMasterPassphrase - Disables the master passphrase for the given region.
type MemoryV1190DisableMasterPassphrase struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertMemoryV1190DisableMasterPassphraseRequired checks if the required fields are not zero-ed
func AssertMemoryV1190DisableMasterPassphraseRequired(obj MemoryV1190DisableMasterPassphrase) error {
	return nil
}

// AssertMemoryV1190DisableMasterPassphraseConstraints checks if the values respects the defined constraints
func AssertMemoryV1190DisableMasterPassphraseConstraints(obj MemoryV1190DisableMasterPassphrase) error {
	return nil
}