// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

import (
	"fmt"
)

type ComputerSystemV1230GraphicalConnectTypesSupported string

// List of ComputerSystemV1230GraphicalConnectTypesSupported
const (
	COMPUTERSYSTEMV1230GRAPHICALCONNECTTYPESSUPPORTED_KVMIP ComputerSystemV1230GraphicalConnectTypesSupported = "KVMIP"
	COMPUTERSYSTEMV1230GRAPHICALCONNECTTYPESSUPPORTED_OEM   ComputerSystemV1230GraphicalConnectTypesSupported = "OEM"
)

// AllowedComputerSystemV1230GraphicalConnectTypesSupportedEnumValues is all the allowed values of ComputerSystemV1230GraphicalConnectTypesSupported enum
var AllowedComputerSystemV1230GraphicalConnectTypesSupportedEnumValues = []ComputerSystemV1230GraphicalConnectTypesSupported{
	"KVMIP",
	"OEM",
}

// validComputerSystemV1230GraphicalConnectTypesSupportedEnumValue provides a map of ComputerSystemV1230GraphicalConnectTypesSupporteds for fast verification of use input
var validComputerSystemV1230GraphicalConnectTypesSupportedEnumValues = map[ComputerSystemV1230GraphicalConnectTypesSupported]struct{}{
	"KVMIP": {},
	"OEM":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1230GraphicalConnectTypesSupported) IsValid() bool {
	_, ok := validComputerSystemV1230GraphicalConnectTypesSupportedEnumValues[v]
	return ok
}

// NewComputerSystemV1230GraphicalConnectTypesSupportedFromValue returns a pointer to a valid ComputerSystemV1230GraphicalConnectTypesSupported
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1230GraphicalConnectTypesSupportedFromValue(v string) (ComputerSystemV1230GraphicalConnectTypesSupported, error) {
	ev := ComputerSystemV1230GraphicalConnectTypesSupported(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1230GraphicalConnectTypesSupported: valid values are %v", v, AllowedComputerSystemV1230GraphicalConnectTypesSupportedEnumValues)
}

// AssertComputerSystemV1230GraphicalConnectTypesSupportedRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1230GraphicalConnectTypesSupportedRequired(obj ComputerSystemV1230GraphicalConnectTypesSupported) error {
	return nil
}

// AssertComputerSystemV1230GraphicalConnectTypesSupportedConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1230GraphicalConnectTypesSupportedConstraints(obj ComputerSystemV1230GraphicalConnectTypesSupported) error {
	return nil
}