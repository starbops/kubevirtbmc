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

type ReservoirV101ReservoirType string

// List of ReservoirV101ReservoirType
const (
	RESERVOIRV101RESERVOIRTYPE_RESERVE   ReservoirV101ReservoirType = "Reserve"
	RESERVOIRV101RESERVOIRTYPE_OVERFLOW  ReservoirV101ReservoirType = "Overflow"
	RESERVOIRV101RESERVOIRTYPE_INLINE    ReservoirV101ReservoirType = "Inline"
	RESERVOIRV101RESERVOIRTYPE_IMMERSION ReservoirV101ReservoirType = "Immersion"
)

// AllowedReservoirV101ReservoirTypeEnumValues is all the allowed values of ReservoirV101ReservoirType enum
var AllowedReservoirV101ReservoirTypeEnumValues = []ReservoirV101ReservoirType{
	"Reserve",
	"Overflow",
	"Inline",
	"Immersion",
}

// validReservoirV101ReservoirTypeEnumValue provides a map of ReservoirV101ReservoirTypes for fast verification of use input
var validReservoirV101ReservoirTypeEnumValues = map[ReservoirV101ReservoirType]struct{}{
	"Reserve":   {},
	"Overflow":  {},
	"Inline":    {},
	"Immersion": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReservoirV101ReservoirType) IsValid() bool {
	_, ok := validReservoirV101ReservoirTypeEnumValues[v]
	return ok
}

// NewReservoirV101ReservoirTypeFromValue returns a pointer to a valid ReservoirV101ReservoirType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReservoirV101ReservoirTypeFromValue(v string) (ReservoirV101ReservoirType, error) {
	ev := ReservoirV101ReservoirType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ReservoirV101ReservoirType: valid values are %v", v, AllowedReservoirV101ReservoirTypeEnumValues)
}

// AssertReservoirV101ReservoirTypeRequired checks if the required fields are not zero-ed
func AssertReservoirV101ReservoirTypeRequired(obj ReservoirV101ReservoirType) error {
	return nil
}

// AssertReservoirV101ReservoirTypeConstraints checks if the values respects the defined constraints
func AssertReservoirV101ReservoirTypeConstraints(obj ReservoirV101ReservoirType) error {
	return nil
}