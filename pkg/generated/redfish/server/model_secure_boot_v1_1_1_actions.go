// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SecureBootV111Actions - The available actions for this resource.
type SecureBootV111Actions struct {
	SecureBootResetKeys SecureBootV111ResetKeys `json:"#SecureBoot.ResetKeys,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertSecureBootV111ActionsRequired checks if the required fields are not zero-ed
func AssertSecureBootV111ActionsRequired(obj SecureBootV111Actions) error {
	if err := AssertSecureBootV111ResetKeysRequired(obj.SecureBootResetKeys); err != nil {
		return err
	}
	return nil
}

// AssertSecureBootV111ActionsConstraints checks if the values respects the defined constraints
func AssertSecureBootV111ActionsConstraints(obj SecureBootV111Actions) error {
	if err := AssertSecureBootV111ResetKeysConstraints(obj.SecureBootResetKeys); err != nil {
		return err
	}
	return nil
}