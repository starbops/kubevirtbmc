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

type SerialInterfaceV120StopBits string

// List of SerialInterfaceV120StopBits
const (
	SERIALINTERFACEV120STOPBITS__1 SerialInterfaceV120StopBits = "1"
	SERIALINTERFACEV120STOPBITS__2 SerialInterfaceV120StopBits = "2"
)

// AllowedSerialInterfaceV120StopBitsEnumValues is all the allowed values of SerialInterfaceV120StopBits enum
var AllowedSerialInterfaceV120StopBitsEnumValues = []SerialInterfaceV120StopBits{
	"1",
	"2",
}

// validSerialInterfaceV120StopBitsEnumValue provides a map of SerialInterfaceV120StopBitss for fast verification of use input
var validSerialInterfaceV120StopBitsEnumValues = map[SerialInterfaceV120StopBits]struct{}{
	"1": {},
	"2": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SerialInterfaceV120StopBits) IsValid() bool {
	_, ok := validSerialInterfaceV120StopBitsEnumValues[v]
	return ok
}

// NewSerialInterfaceV120StopBitsFromValue returns a pointer to a valid SerialInterfaceV120StopBits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSerialInterfaceV120StopBitsFromValue(v string) (SerialInterfaceV120StopBits, error) {
	ev := SerialInterfaceV120StopBits(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for SerialInterfaceV120StopBits: valid values are %v", v, AllowedSerialInterfaceV120StopBitsEnumValues)
}

// AssertSerialInterfaceV120StopBitsRequired checks if the required fields are not zero-ed
func AssertSerialInterfaceV120StopBitsRequired(obj SerialInterfaceV120StopBits) error {
	return nil
}

// AssertSerialInterfaceV120StopBitsConstraints checks if the values respects the defined constraints
func AssertSerialInterfaceV120StopBitsConstraints(obj SerialInterfaceV120StopBits) error {
	return nil
}