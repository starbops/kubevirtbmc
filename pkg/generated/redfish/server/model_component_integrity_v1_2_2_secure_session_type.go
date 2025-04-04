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

type ComponentIntegrityV122SecureSessionType string

// List of ComponentIntegrityV122SecureSessionType
const (
	COMPONENTINTEGRITYV122SECURESESSIONTYPE_PLAIN                   ComponentIntegrityV122SecureSessionType = "Plain"
	COMPONENTINTEGRITYV122SECURESESSIONTYPE_ENCRYPTED_AUTHENTICATED ComponentIntegrityV122SecureSessionType = "EncryptedAuthenticated"
	COMPONENTINTEGRITYV122SECURESESSIONTYPE_AUTHENTICATED_ONLY      ComponentIntegrityV122SecureSessionType = "AuthenticatedOnly"
)

// AllowedComponentIntegrityV122SecureSessionTypeEnumValues is all the allowed values of ComponentIntegrityV122SecureSessionType enum
var AllowedComponentIntegrityV122SecureSessionTypeEnumValues = []ComponentIntegrityV122SecureSessionType{
	"Plain",
	"EncryptedAuthenticated",
	"AuthenticatedOnly",
}

// validComponentIntegrityV122SecureSessionTypeEnumValue provides a map of ComponentIntegrityV122SecureSessionTypes for fast verification of use input
var validComponentIntegrityV122SecureSessionTypeEnumValues = map[ComponentIntegrityV122SecureSessionType]struct{}{
	"Plain":                  {},
	"EncryptedAuthenticated": {},
	"AuthenticatedOnly":      {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComponentIntegrityV122SecureSessionType) IsValid() bool {
	_, ok := validComponentIntegrityV122SecureSessionTypeEnumValues[v]
	return ok
}

// NewComponentIntegrityV122SecureSessionTypeFromValue returns a pointer to a valid ComponentIntegrityV122SecureSessionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComponentIntegrityV122SecureSessionTypeFromValue(v string) (ComponentIntegrityV122SecureSessionType, error) {
	ev := ComponentIntegrityV122SecureSessionType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComponentIntegrityV122SecureSessionType: valid values are %v", v, AllowedComponentIntegrityV122SecureSessionTypeEnumValues)
}

// AssertComponentIntegrityV122SecureSessionTypeRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122SecureSessionTypeRequired(obj ComponentIntegrityV122SecureSessionType) error {
	return nil
}

// AssertComponentIntegrityV122SecureSessionTypeConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122SecureSessionTypeConstraints(obj ComponentIntegrityV122SecureSessionType) error {
	return nil
}
