// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SecurityPolicyV101SpdmAlgorithmSet - The SPDM algorithm settings.
type SecurityPolicyV101SpdmAlgorithmSet struct {

	// The AEAD algorithms.
	AEAD []*string `json:"AEAD,omitempty"`

	// The asymmetric signature algorithms.
	BaseAsym []*string `json:"BaseAsym,omitempty"`

	// The hash algorithms.
	BaseHash []*string `json:"BaseHash,omitempty"`
}

// AssertSecurityPolicyV101SpdmAlgorithmSetRequired checks if the required fields are not zero-ed
func AssertSecurityPolicyV101SpdmAlgorithmSetRequired(obj SecurityPolicyV101SpdmAlgorithmSet) error {
	return nil
}

// AssertSecurityPolicyV101SpdmAlgorithmSetConstraints checks if the values respects the defined constraints
func AssertSecurityPolicyV101SpdmAlgorithmSetConstraints(obj SecurityPolicyV101SpdmAlgorithmSet) error {
	return nil
}
