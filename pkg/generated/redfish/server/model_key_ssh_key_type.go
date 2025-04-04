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

type KeySshKeyType string

// List of KeySshKeyType
const (
	KEYSSHKEYTYPE_RSA     KeySshKeyType = "RSA"
	KEYSSHKEYTYPE_DSA     KeySshKeyType = "DSA"
	KEYSSHKEYTYPE_ECDSA   KeySshKeyType = "ECDSA"
	KEYSSHKEYTYPE_ED25519 KeySshKeyType = "Ed25519"
)

// AllowedKeySshKeyTypeEnumValues is all the allowed values of KeySshKeyType enum
var AllowedKeySshKeyTypeEnumValues = []KeySshKeyType{
	"RSA",
	"DSA",
	"ECDSA",
	"Ed25519",
}

// validKeySshKeyTypeEnumValue provides a map of KeySshKeyTypes for fast verification of use input
var validKeySshKeyTypeEnumValues = map[KeySshKeyType]struct{}{
	"RSA":     {},
	"DSA":     {},
	"ECDSA":   {},
	"Ed25519": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeySshKeyType) IsValid() bool {
	_, ok := validKeySshKeyTypeEnumValues[v]
	return ok
}

// NewKeySshKeyTypeFromValue returns a pointer to a valid KeySshKeyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeySshKeyTypeFromValue(v string) (KeySshKeyType, error) {
	ev := KeySshKeyType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for KeySshKeyType: valid values are %v", v, AllowedKeySshKeyTypeEnumValues)
}

// AssertKeySshKeyTypeRequired checks if the required fields are not zero-ed
func AssertKeySshKeyTypeRequired(obj KeySshKeyType) error {
	return nil
}

// AssertKeySshKeyTypeConstraints checks if the values respects the defined constraints
func AssertKeySshKeyTypeConstraints(obj KeySshKeyType) error {
	return nil
}
