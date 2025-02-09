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

type MemoryV1190OperatingMemoryModes string

// List of MemoryV1190OperatingMemoryModes
const (
	MEMORYV1190OPERATINGMEMORYMODES_VOLATILE MemoryV1190OperatingMemoryModes = "Volatile"
	MEMORYV1190OPERATINGMEMORYMODES_PMEM     MemoryV1190OperatingMemoryModes = "PMEM"
	MEMORYV1190OPERATINGMEMORYMODES_BLOCK    MemoryV1190OperatingMemoryModes = "Block"
)

// AllowedMemoryV1190OperatingMemoryModesEnumValues is all the allowed values of MemoryV1190OperatingMemoryModes enum
var AllowedMemoryV1190OperatingMemoryModesEnumValues = []MemoryV1190OperatingMemoryModes{
	"Volatile",
	"PMEM",
	"Block",
}

// validMemoryV1190OperatingMemoryModesEnumValue provides a map of MemoryV1190OperatingMemoryModess for fast verification of use input
var validMemoryV1190OperatingMemoryModesEnumValues = map[MemoryV1190OperatingMemoryModes]struct{}{
	"Volatile": {},
	"PMEM":     {},
	"Block":    {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MemoryV1190OperatingMemoryModes) IsValid() bool {
	_, ok := validMemoryV1190OperatingMemoryModesEnumValues[v]
	return ok
}

// NewMemoryV1190OperatingMemoryModesFromValue returns a pointer to a valid MemoryV1190OperatingMemoryModes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMemoryV1190OperatingMemoryModesFromValue(v string) (MemoryV1190OperatingMemoryModes, error) {
	ev := MemoryV1190OperatingMemoryModes(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for MemoryV1190OperatingMemoryModes: valid values are %v", v, AllowedMemoryV1190OperatingMemoryModesEnumValues)
}

// AssertMemoryV1190OperatingMemoryModesRequired checks if the required fields are not zero-ed
func AssertMemoryV1190OperatingMemoryModesRequired(obj MemoryV1190OperatingMemoryModes) error {
	return nil
}

// AssertMemoryV1190OperatingMemoryModesConstraints checks if the values respects the defined constraints
func AssertMemoryV1190OperatingMemoryModesConstraints(obj MemoryV1190OperatingMemoryModes) error {
	return nil
}
