// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryV1190SetPassphrase - Set passphrase for the given regions.
type MemoryV1190SetPassphrase struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertMemoryV1190SetPassphraseRequired checks if the required fields are not zero-ed
func AssertMemoryV1190SetPassphraseRequired(obj MemoryV1190SetPassphrase) error {
	return nil
}

// AssertMemoryV1190SetPassphraseConstraints checks if the values respects the defined constraints
func AssertMemoryV1190SetPassphraseConstraints(obj MemoryV1190SetPassphrase) error {
	return nil
}
