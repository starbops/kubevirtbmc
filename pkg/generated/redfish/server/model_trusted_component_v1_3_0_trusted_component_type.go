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

type TrustedComponentV130TrustedComponentType string

// List of TrustedComponentV130TrustedComponentType
const (
	TRUSTEDCOMPONENTV130TRUSTEDCOMPONENTTYPE_DISCRETE   TrustedComponentV130TrustedComponentType = "Discrete"
	TRUSTEDCOMPONENTV130TRUSTEDCOMPONENTTYPE_INTEGRATED TrustedComponentV130TrustedComponentType = "Integrated"
)

// AllowedTrustedComponentV130TrustedComponentTypeEnumValues is all the allowed values of TrustedComponentV130TrustedComponentType enum
var AllowedTrustedComponentV130TrustedComponentTypeEnumValues = []TrustedComponentV130TrustedComponentType{
	"Discrete",
	"Integrated",
}

// validTrustedComponentV130TrustedComponentTypeEnumValue provides a map of TrustedComponentV130TrustedComponentTypes for fast verification of use input
var validTrustedComponentV130TrustedComponentTypeEnumValues = map[TrustedComponentV130TrustedComponentType]struct{}{
	"Discrete":   {},
	"Integrated": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrustedComponentV130TrustedComponentType) IsValid() bool {
	_, ok := validTrustedComponentV130TrustedComponentTypeEnumValues[v]
	return ok
}

// NewTrustedComponentV130TrustedComponentTypeFromValue returns a pointer to a valid TrustedComponentV130TrustedComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrustedComponentV130TrustedComponentTypeFromValue(v string) (TrustedComponentV130TrustedComponentType, error) {
	ev := TrustedComponentV130TrustedComponentType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TrustedComponentV130TrustedComponentType: valid values are %v", v, AllowedTrustedComponentV130TrustedComponentTypeEnumValues)
}

// AssertTrustedComponentV130TrustedComponentTypeRequired checks if the required fields are not zero-ed
func AssertTrustedComponentV130TrustedComponentTypeRequired(obj TrustedComponentV130TrustedComponentType) error {
	return nil
}

// AssertTrustedComponentV130TrustedComponentTypeConstraints checks if the values respects the defined constraints
func AssertTrustedComponentV130TrustedComponentTypeConstraints(obj TrustedComponentV130TrustedComponentType) error {
	return nil
}
