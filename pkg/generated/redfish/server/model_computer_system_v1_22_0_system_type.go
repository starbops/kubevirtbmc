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

type ComputerSystemV1220SystemType string

// List of ComputerSystemV1220SystemType
const (
	COMPUTERSYSTEMV1220SYSTEMTYPE_PHYSICAL               ComputerSystemV1220SystemType = "Physical"
	COMPUTERSYSTEMV1220SYSTEMTYPE_VIRTUAL                ComputerSystemV1220SystemType = "Virtual"
	COMPUTERSYSTEMV1220SYSTEMTYPE_OS                     ComputerSystemV1220SystemType = "OS"
	COMPUTERSYSTEMV1220SYSTEMTYPE_PHYSICALLY_PARTITIONED ComputerSystemV1220SystemType = "PhysicallyPartitioned"
	COMPUTERSYSTEMV1220SYSTEMTYPE_VIRTUALLY_PARTITIONED  ComputerSystemV1220SystemType = "VirtuallyPartitioned"
	COMPUTERSYSTEMV1220SYSTEMTYPE_COMPOSED               ComputerSystemV1220SystemType = "Composed"
	COMPUTERSYSTEMV1220SYSTEMTYPE_DPU                    ComputerSystemV1220SystemType = "DPU"
)

// AllowedComputerSystemV1220SystemTypeEnumValues is all the allowed values of ComputerSystemV1220SystemType enum
var AllowedComputerSystemV1220SystemTypeEnumValues = []ComputerSystemV1220SystemType{
	"Physical",
	"Virtual",
	"OS",
	"PhysicallyPartitioned",
	"VirtuallyPartitioned",
	"Composed",
	"DPU",
}

// validComputerSystemV1220SystemTypeEnumValue provides a map of ComputerSystemV1220SystemTypes for fast verification of use input
var validComputerSystemV1220SystemTypeEnumValues = map[ComputerSystemV1220SystemType]struct{}{
	"Physical":              {},
	"Virtual":               {},
	"OS":                    {},
	"PhysicallyPartitioned": {},
	"VirtuallyPartitioned":  {},
	"Composed":              {},
	"DPU":                   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1220SystemType) IsValid() bool {
	_, ok := validComputerSystemV1220SystemTypeEnumValues[v]
	return ok
}

// NewComputerSystemV1220SystemTypeFromValue returns a pointer to a valid ComputerSystemV1220SystemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1220SystemTypeFromValue(v string) (ComputerSystemV1220SystemType, error) {
	ev := ComputerSystemV1220SystemType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1220SystemType: valid values are %v", v, AllowedComputerSystemV1220SystemTypeEnumValues)
}

// AssertComputerSystemV1220SystemTypeRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220SystemTypeRequired(obj ComputerSystemV1220SystemType) error {
	return nil
}

// AssertComputerSystemV1220SystemTypeConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220SystemTypeConstraints(obj ComputerSystemV1220SystemType) error {
	return nil
}