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

type AccountServiceV1150LocalAccountAuth string

// List of AccountServiceV1150LocalAccountAuth
const (
	ACCOUNTSERVICEV1150LOCALACCOUNTAUTH_ENABLED     AccountServiceV1150LocalAccountAuth = "Enabled"
	ACCOUNTSERVICEV1150LOCALACCOUNTAUTH_DISABLED    AccountServiceV1150LocalAccountAuth = "Disabled"
	ACCOUNTSERVICEV1150LOCALACCOUNTAUTH_FALLBACK    AccountServiceV1150LocalAccountAuth = "Fallback"
	ACCOUNTSERVICEV1150LOCALACCOUNTAUTH_LOCAL_FIRST AccountServiceV1150LocalAccountAuth = "LocalFirst"
)

// AllowedAccountServiceV1150LocalAccountAuthEnumValues is all the allowed values of AccountServiceV1150LocalAccountAuth enum
var AllowedAccountServiceV1150LocalAccountAuthEnumValues = []AccountServiceV1150LocalAccountAuth{
	"Enabled",
	"Disabled",
	"Fallback",
	"LocalFirst",
}

// validAccountServiceV1150LocalAccountAuthEnumValue provides a map of AccountServiceV1150LocalAccountAuths for fast verification of use input
var validAccountServiceV1150LocalAccountAuthEnumValues = map[AccountServiceV1150LocalAccountAuth]struct{}{
	"Enabled":    {},
	"Disabled":   {},
	"Fallback":   {},
	"LocalFirst": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountServiceV1150LocalAccountAuth) IsValid() bool {
	_, ok := validAccountServiceV1150LocalAccountAuthEnumValues[v]
	return ok
}

// NewAccountServiceV1150LocalAccountAuthFromValue returns a pointer to a valid AccountServiceV1150LocalAccountAuth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountServiceV1150LocalAccountAuthFromValue(v string) (AccountServiceV1150LocalAccountAuth, error) {
	ev := AccountServiceV1150LocalAccountAuth(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for AccountServiceV1150LocalAccountAuth: valid values are %v", v, AllowedAccountServiceV1150LocalAccountAuthEnumValues)
}

// AssertAccountServiceV1150LocalAccountAuthRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150LocalAccountAuthRequired(obj AccountServiceV1150LocalAccountAuth) error {
	return nil
}

// AssertAccountServiceV1150LocalAccountAuthConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150LocalAccountAuthConstraints(obj AccountServiceV1150LocalAccountAuth) error {
	return nil
}