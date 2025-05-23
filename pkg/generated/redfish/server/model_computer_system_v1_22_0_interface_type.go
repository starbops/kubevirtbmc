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

type ComputerSystemV1220InterfaceType string

// List of ComputerSystemV1220InterfaceType
const (
	COMPUTERSYSTEMV1220INTERFACETYPE_TPM1_2 ComputerSystemV1220InterfaceType = "TPM1_2"
	COMPUTERSYSTEMV1220INTERFACETYPE_TPM2_0 ComputerSystemV1220InterfaceType = "TPM2_0"
	COMPUTERSYSTEMV1220INTERFACETYPE_TCM1_0 ComputerSystemV1220InterfaceType = "TCM1_0"
)

// AllowedComputerSystemV1220InterfaceTypeEnumValues is all the allowed values of ComputerSystemV1220InterfaceType enum
var AllowedComputerSystemV1220InterfaceTypeEnumValues = []ComputerSystemV1220InterfaceType{
	"TPM1_2",
	"TPM2_0",
	"TCM1_0",
}

// validComputerSystemV1220InterfaceTypeEnumValue provides a map of ComputerSystemV1220InterfaceTypes for fast verification of use input
var validComputerSystemV1220InterfaceTypeEnumValues = map[ComputerSystemV1220InterfaceType]struct{}{
	"TPM1_2": {},
	"TPM2_0": {},
	"TCM1_0": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1220InterfaceType) IsValid() bool {
	_, ok := validComputerSystemV1220InterfaceTypeEnumValues[v]
	return ok
}

// NewComputerSystemV1220InterfaceTypeFromValue returns a pointer to a valid ComputerSystemV1220InterfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1220InterfaceTypeFromValue(v string) (ComputerSystemV1220InterfaceType, error) {
	ev := ComputerSystemV1220InterfaceType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1220InterfaceType: valid values are %v", v, AllowedComputerSystemV1220InterfaceTypeEnumValues)
}

// AssertComputerSystemV1220InterfaceTypeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220InterfaceTypeRequired(obj ComputerSystemV1220InterfaceType) error {
	return nil
}

// AssertComputerSystemV1220InterfaceTypeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220InterfaceTypeConstraints(obj ComputerSystemV1220InterfaceType) error {
	return nil
}
