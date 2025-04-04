// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SignatureV102Actions - The available actions for this resource.
type SignatureV102Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertSignatureV102ActionsRequired checks if the required fields are not zero-ed
func AssertSignatureV102ActionsRequired(obj SignatureV102Actions) error {
	return nil
}

// AssertSignatureV102ActionsConstraints checks if the values respects the defined constraints
func AssertSignatureV102ActionsConstraints(obj SignatureV102Actions) error {
	return nil
}
