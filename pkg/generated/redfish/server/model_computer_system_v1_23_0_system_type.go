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

type ComputerSystemV1230SystemType string

// List of ComputerSystemV1230SystemType
const (
	COMPUTERSYSTEMV1230SYSTEMTYPE_PHYSICAL               ComputerSystemV1230SystemType = "Physical"
	COMPUTERSYSTEMV1230SYSTEMTYPE_VIRTUAL                ComputerSystemV1230SystemType = "Virtual"
	COMPUTERSYSTEMV1230SYSTEMTYPE_OS                     ComputerSystemV1230SystemType = "OS"
	COMPUTERSYSTEMV1230SYSTEMTYPE_PHYSICALLY_PARTITIONED ComputerSystemV1230SystemType = "PhysicallyPartitioned"
	COMPUTERSYSTEMV1230SYSTEMTYPE_VIRTUALLY_PARTITIONED  ComputerSystemV1230SystemType = "VirtuallyPartitioned"
	COMPUTERSYSTEMV1230SYSTEMTYPE_COMPOSED               ComputerSystemV1230SystemType = "Composed"
	COMPUTERSYSTEMV1230SYSTEMTYPE_DPU                    ComputerSystemV1230SystemType = "DPU"
)

// AllowedComputerSystemV1230SystemTypeEnumValues is all the allowed values of ComputerSystemV1230SystemType enum
var AllowedComputerSystemV1230SystemTypeEnumValues = []ComputerSystemV1230SystemType{
	"Physical",
	"Virtual",
	"OS",
	"PhysicallyPartitioned",
	"VirtuallyPartitioned",
	"Composed",
	"DPU",
}

// validComputerSystemV1230SystemTypeEnumValue provides a map of ComputerSystemV1230SystemTypes for fast verification of use input
var validComputerSystemV1230SystemTypeEnumValues = map[ComputerSystemV1230SystemType]struct{}{
	"Physical":              {},
	"Virtual":               {},
	"OS":                    {},
	"PhysicallyPartitioned": {},
	"VirtuallyPartitioned":  {},
	"Composed":              {},
	"DPU":                   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1230SystemType) IsValid() bool {
	_, ok := validComputerSystemV1230SystemTypeEnumValues[v]
	return ok
}

// NewComputerSystemV1230SystemTypeFromValue returns a pointer to a valid ComputerSystemV1230SystemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1230SystemTypeFromValue(v string) (ComputerSystemV1230SystemType, error) {
	ev := ComputerSystemV1230SystemType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1230SystemType: valid values are %v", v, AllowedComputerSystemV1230SystemTypeEnumValues)
}

// AssertComputerSystemV1230SystemTypeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1230SystemTypeRequired(obj ComputerSystemV1230SystemType) error {
	return nil
}

// AssertComputerSystemV1230SystemTypeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1230SystemTypeConstraints(obj ComputerSystemV1230SystemType) error {
	return nil
}