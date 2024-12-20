// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ConnectionV131GenZConnectionKey - The Gen-Z-specific permission key information for a connection.
type ConnectionV131GenZConnectionKey struct {

	// The Access Key for this connection.
	AccessKey string `json:"AccessKey,omitempty" validate:"regexp=^0[xX]([a-fA-F]|[0-9]){2}$"`

	// Indicates whether Region Key domain checking is enabled for this connection.
	RKeyDomainCheckingEnabled bool `json:"RKeyDomainCheckingEnabled,omitempty"`

	// The read-only Region Key for this connection.
	RKeyReadOnlyKey string `json:"RKeyReadOnlyKey,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){4}$"`

	// The read-write Region Key for this connection.
	RKeyReadWriteKey string `json:"RKeyReadWriteKey,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){4}$"`
}

// AssertConnectionV131GenZConnectionKeyRequired checks if the required fields are not zero-ed
func AssertConnectionV131GenZConnectionKeyRequired(obj ConnectionV131GenZConnectionKey) error {
	return nil
}

// AssertConnectionV131GenZConnectionKeyConstraints checks if the values respects the defined constraints
func AssertConnectionV131GenZConnectionKeyConstraints(obj ConnectionV131GenZConnectionKey) error {
	return nil
}
