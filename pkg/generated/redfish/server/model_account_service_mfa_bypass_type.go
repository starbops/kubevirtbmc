// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

import (
	"fmt"
)

type AccountServiceMfaBypassType string

// List of AccountServiceMfaBypassType
const (
	ACCOUNTSERVICEMFABYPASSTYPE_ALL                          AccountServiceMfaBypassType = "All"
	ACCOUNTSERVICEMFABYPASSTYPE_SECUR_ID                     AccountServiceMfaBypassType = "SecurID"
	ACCOUNTSERVICEMFABYPASSTYPE_GOOGLE_AUTHENTICATOR         AccountServiceMfaBypassType = "GoogleAuthenticator"
	ACCOUNTSERVICEMFABYPASSTYPE_MICROSOFT_AUTHENTICATOR      AccountServiceMfaBypassType = "MicrosoftAuthenticator"
	ACCOUNTSERVICEMFABYPASSTYPE_CLIENT_CERTIFICATE           AccountServiceMfaBypassType = "ClientCertificate"
	ACCOUNTSERVICEMFABYPASSTYPE_ONE_TIME_PASSCODE            AccountServiceMfaBypassType = "OneTimePasscode"
	ACCOUNTSERVICEMFABYPASSTYPE_TIME_BASED_ONE_TIME_PASSWORD AccountServiceMfaBypassType = "TimeBasedOneTimePassword"
	ACCOUNTSERVICEMFABYPASSTYPE_OEM                          AccountServiceMfaBypassType = "OEM"
)

// AllowedAccountServiceMfaBypassTypeEnumValues is all the allowed values of AccountServiceMfaBypassType enum
var AllowedAccountServiceMfaBypassTypeEnumValues = []AccountServiceMfaBypassType{
	"All",
	"SecurID",
	"GoogleAuthenticator",
	"MicrosoftAuthenticator",
	"ClientCertificate",
	"OneTimePasscode",
	"TimeBasedOneTimePassword",
	"OEM",
}

// validAccountServiceMfaBypassTypeEnumValue provides a map of AccountServiceMfaBypassTypes for fast verification of use input
var validAccountServiceMfaBypassTypeEnumValues = map[AccountServiceMfaBypassType]struct{}{
	"All":                      {},
	"SecurID":                  {},
	"GoogleAuthenticator":      {},
	"MicrosoftAuthenticator":   {},
	"ClientCertificate":        {},
	"OneTimePasscode":          {},
	"TimeBasedOneTimePassword": {},
	"OEM":                      {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountServiceMfaBypassType) IsValid() bool {
	_, ok := validAccountServiceMfaBypassTypeEnumValues[v]
	return ok
}

// NewAccountServiceMfaBypassTypeFromValue returns a pointer to a valid AccountServiceMfaBypassType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountServiceMfaBypassTypeFromValue(v string) (AccountServiceMfaBypassType, error) {
	ev := AccountServiceMfaBypassType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for AccountServiceMfaBypassType: valid values are %v", v, AllowedAccountServiceMfaBypassTypeEnumValues)
}

// AssertAccountServiceMfaBypassTypeRequired checks if the required fields are not zero-ed
func AssertAccountServiceMfaBypassTypeRequired(obj AccountServiceMfaBypassType) error {
	return nil
}

// AssertAccountServiceMfaBypassTypeConstraints checks if the values respects the defined constraints
func AssertAccountServiceMfaBypassTypeConstraints(obj AccountServiceMfaBypassType) error {
	return nil
}
