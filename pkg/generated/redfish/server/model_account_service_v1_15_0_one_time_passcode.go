// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AccountServiceV1150OneTimePasscode - Various settings for one-time passcode (OTP) multi-factor authentication.
type AccountServiceV1150OneTimePasscode struct {

	// An indication of whether multi-factor authentication using a one-time passcode is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
}

// AssertAccountServiceV1150OneTimePasscodeRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150OneTimePasscodeRequired(obj AccountServiceV1150OneTimePasscode) error {
	return nil
}

// AssertAccountServiceV1150OneTimePasscodeConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150OneTimePasscodeConstraints(obj AccountServiceV1150OneTimePasscode) error {
	return nil
}