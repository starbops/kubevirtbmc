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

type RegisteredClientV111ClientType string

// List of RegisteredClientV111ClientType
const (
	REGISTEREDCLIENTV111CLIENTTYPE_MONITOR   RegisteredClientV111ClientType = "Monitor"
	REGISTEREDCLIENTV111CLIENTTYPE_CONFIGURE RegisteredClientV111ClientType = "Configure"
)

// AllowedRegisteredClientV111ClientTypeEnumValues is all the allowed values of RegisteredClientV111ClientType enum
var AllowedRegisteredClientV111ClientTypeEnumValues = []RegisteredClientV111ClientType{
	"Monitor",
	"Configure",
}

// validRegisteredClientV111ClientTypeEnumValue provides a map of RegisteredClientV111ClientTypes for fast verification of use input
var validRegisteredClientV111ClientTypeEnumValues = map[RegisteredClientV111ClientType]struct{}{
	"Monitor":   {},
	"Configure": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegisteredClientV111ClientType) IsValid() bool {
	_, ok := validRegisteredClientV111ClientTypeEnumValues[v]
	return ok
}

// NewRegisteredClientV111ClientTypeFromValue returns a pointer to a valid RegisteredClientV111ClientType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegisteredClientV111ClientTypeFromValue(v string) (RegisteredClientV111ClientType, error) {
	ev := RegisteredClientV111ClientType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for RegisteredClientV111ClientType: valid values are %v", v, AllowedRegisteredClientV111ClientTypeEnumValues)
}

// AssertRegisteredClientV111ClientTypeRequired checks if the required fields are not zero-ed
func AssertRegisteredClientV111ClientTypeRequired(obj RegisteredClientV111ClientType) error {
	return nil
}

// AssertRegisteredClientV111ClientTypeConstraints checks if the values respects the defined constraints
func AssertRegisteredClientV111ClientTypeConstraints(obj RegisteredClientV111ClientType) error {
	return nil
}