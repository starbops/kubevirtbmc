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

type SignatureSignatureTypeRegistry string

// List of SignatureSignatureTypeRegistry
const (
	SIGNATURESIGNATURETYPEREGISTRY_UEFI SignatureSignatureTypeRegistry = "UEFI"
)

// AllowedSignatureSignatureTypeRegistryEnumValues is all the allowed values of SignatureSignatureTypeRegistry enum
var AllowedSignatureSignatureTypeRegistryEnumValues = []SignatureSignatureTypeRegistry{
	"UEFI",
}

// validSignatureSignatureTypeRegistryEnumValue provides a map of SignatureSignatureTypeRegistrys for fast verification of use input
var validSignatureSignatureTypeRegistryEnumValues = map[SignatureSignatureTypeRegistry]struct{}{
	"UEFI": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignatureSignatureTypeRegistry) IsValid() bool {
	_, ok := validSignatureSignatureTypeRegistryEnumValues[v]
	return ok
}

// NewSignatureSignatureTypeRegistryFromValue returns a pointer to a valid SignatureSignatureTypeRegistry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignatureSignatureTypeRegistryFromValue(v string) (SignatureSignatureTypeRegistry, error) {
	ev := SignatureSignatureTypeRegistry(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for SignatureSignatureTypeRegistry: valid values are %v", v, AllowedSignatureSignatureTypeRegistryEnumValues)
}

// AssertSignatureSignatureTypeRegistryRequired checks if the required fields are not zero-ed
func AssertSignatureSignatureTypeRegistryRequired(obj SignatureSignatureTypeRegistry) error {
	return nil
}

// AssertSignatureSignatureTypeRegistryConstraints checks if the values respects the defined constraints
func AssertSignatureSignatureTypeRegistryConstraints(obj SignatureSignatureTypeRegistry) error {
	return nil
}