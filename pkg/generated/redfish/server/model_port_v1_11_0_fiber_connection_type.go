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

type PortV1110FiberConnectionType string

// List of PortV1110FiberConnectionType
const (
	PORTV1110FIBERCONNECTIONTYPE_SINGLE_MODE PortV1110FiberConnectionType = "SingleMode"
	PORTV1110FIBERCONNECTIONTYPE_MULTI_MODE  PortV1110FiberConnectionType = "MultiMode"
)

// AllowedPortV1110FiberConnectionTypeEnumValues is all the allowed values of PortV1110FiberConnectionType enum
var AllowedPortV1110FiberConnectionTypeEnumValues = []PortV1110FiberConnectionType{
	"SingleMode",
	"MultiMode",
}

// validPortV1110FiberConnectionTypeEnumValue provides a map of PortV1110FiberConnectionTypes for fast verification of use input
var validPortV1110FiberConnectionTypeEnumValues = map[PortV1110FiberConnectionType]struct{}{
	"SingleMode": {},
	"MultiMode":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortV1110FiberConnectionType) IsValid() bool {
	_, ok := validPortV1110FiberConnectionTypeEnumValues[v]
	return ok
}

// NewPortV1110FiberConnectionTypeFromValue returns a pointer to a valid PortV1110FiberConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortV1110FiberConnectionTypeFromValue(v string) (PortV1110FiberConnectionType, error) {
	ev := PortV1110FiberConnectionType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PortV1110FiberConnectionType: valid values are %v", v, AllowedPortV1110FiberConnectionTypeEnumValues)
}

// AssertPortV1110FiberConnectionTypeRequired checks if the required fields are not zero-ed
func AssertPortV1110FiberConnectionTypeRequired(obj PortV1110FiberConnectionType) error {
	return nil
}

// AssertPortV1110FiberConnectionTypeConstraints checks if the values respects the defined constraints
func AssertPortV1110FiberConnectionTypeConstraints(obj PortV1110FiberConnectionType) error {
	return nil
}
