// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AccountServiceV1160MfaBypass - Multi-factor authentication bypass settings.
type AccountServiceV1160MfaBypass struct {

	// The types of multi-factor authentication this account or role mapping is allowed to bypass.
	BypassTypes []AccountServiceMfaBypassBypassTypesInner `json:"BypassTypes,omitempty"`
}

// AssertAccountServiceV1160MfaBypassRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1160MfaBypassRequired(obj AccountServiceV1160MfaBypass) error {
	for _, el := range obj.BypassTypes {
		if err := AssertAccountServiceMfaBypassBypassTypesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAccountServiceV1160MfaBypassConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1160MfaBypassConstraints(obj AccountServiceV1160MfaBypass) error {
	for _, el := range obj.BypassTypes {
		if err := AssertAccountServiceMfaBypassBypassTypesInnerConstraints(el); err != nil {
			return err
		}
	}
	return nil
}