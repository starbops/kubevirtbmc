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

type ComputerSystemV1230InterfaceType string

// List of ComputerSystemV1230InterfaceType
const (
	COMPUTERSYSTEMV1230INTERFACETYPE_TPM1_2 ComputerSystemV1230InterfaceType = "TPM1_2"
	COMPUTERSYSTEMV1230INTERFACETYPE_TPM2_0 ComputerSystemV1230InterfaceType = "TPM2_0"
	COMPUTERSYSTEMV1230INTERFACETYPE_TCM1_0 ComputerSystemV1230InterfaceType = "TCM1_0"
)

// AllowedComputerSystemV1230InterfaceTypeEnumValues is all the allowed values of ComputerSystemV1230InterfaceType enum
var AllowedComputerSystemV1230InterfaceTypeEnumValues = []ComputerSystemV1230InterfaceType{
	"TPM1_2",
	"TPM2_0",
	"TCM1_0",
}

// validComputerSystemV1230InterfaceTypeEnumValue provides a map of ComputerSystemV1230InterfaceTypes for fast verification of use input
var validComputerSystemV1230InterfaceTypeEnumValues = map[ComputerSystemV1230InterfaceType]struct{}{
	"TPM1_2": {},
	"TPM2_0": {},
	"TCM1_0": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1230InterfaceType) IsValid() bool {
	_, ok := validComputerSystemV1230InterfaceTypeEnumValues[v]
	return ok
}

// NewComputerSystemV1230InterfaceTypeFromValue returns a pointer to a valid ComputerSystemV1230InterfaceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1230InterfaceTypeFromValue(v string) (ComputerSystemV1230InterfaceType, error) {
	ev := ComputerSystemV1230InterfaceType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1230InterfaceType: valid values are %v", v, AllowedComputerSystemV1230InterfaceTypeEnumValues)
}

// AssertComputerSystemV1230InterfaceTypeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1230InterfaceTypeRequired(obj ComputerSystemV1230InterfaceType) error {
	return nil
}

// AssertComputerSystemV1230InterfaceTypeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1230InterfaceTypeConstraints(obj ComputerSystemV1230InterfaceType) error {
	return nil
}