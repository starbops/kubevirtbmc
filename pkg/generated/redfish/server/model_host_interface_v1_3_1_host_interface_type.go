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

type HostInterfaceV131HostInterfaceType string

// List of HostInterfaceV131HostInterfaceType
const (
	HOSTINTERFACEV131HOSTINTERFACETYPE_NETWORK_HOST_INTERFACE HostInterfaceV131HostInterfaceType = "NetworkHostInterface"
)

// AllowedHostInterfaceV131HostInterfaceTypeEnumValues is all the allowed values of HostInterfaceV131HostInterfaceType enum
var AllowedHostInterfaceV131HostInterfaceTypeEnumValues = []HostInterfaceV131HostInterfaceType{
	"NetworkHostInterface",
}

// validHostInterfaceV131HostInterfaceTypeEnumValue provides a map of HostInterfaceV131HostInterfaceTypes for fast verification of use input
var validHostInterfaceV131HostInterfaceTypeEnumValues = map[HostInterfaceV131HostInterfaceType]struct{}{
	"NetworkHostInterface": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostInterfaceV131HostInterfaceType) IsValid() bool {
	_, ok := validHostInterfaceV131HostInterfaceTypeEnumValues[v]
	return ok
}

// NewHostInterfaceV131HostInterfaceTypeFromValue returns a pointer to a valid HostInterfaceV131HostInterfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostInterfaceV131HostInterfaceTypeFromValue(v string) (HostInterfaceV131HostInterfaceType, error) {
	ev := HostInterfaceV131HostInterfaceType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for HostInterfaceV131HostInterfaceType: valid values are %v", v, AllowedHostInterfaceV131HostInterfaceTypeEnumValues)
}

// AssertHostInterfaceV131HostInterfaceTypeRequired checks if the required fields are not zero-ed
func AssertHostInterfaceV131HostInterfaceTypeRequired(obj HostInterfaceV131HostInterfaceType) error {
	return nil
}

// AssertHostInterfaceV131HostInterfaceTypeConstraints checks if the values respects the defined constraints
func AssertHostInterfaceV131HostInterfaceTypeConstraints(obj HostInterfaceV131HostInterfaceType) error {
	return nil
}