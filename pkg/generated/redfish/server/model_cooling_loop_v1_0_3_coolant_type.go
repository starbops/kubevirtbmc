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

type CoolingLoopV103CoolantType string

// List of CoolingLoopV103CoolantType
const (
	COOLINGLOOPV103COOLANTTYPE_WATER        CoolingLoopV103CoolantType = "Water"
	COOLINGLOOPV103COOLANTTYPE_HYDROCARBON  CoolingLoopV103CoolantType = "Hydrocarbon"
	COOLINGLOOPV103COOLANTTYPE_FLUOROCARBON CoolingLoopV103CoolantType = "Fluorocarbon"
	COOLINGLOOPV103COOLANTTYPE_DIELECTRIC   CoolingLoopV103CoolantType = "Dielectric"
)

// AllowedCoolingLoopV103CoolantTypeEnumValues is all the allowed values of CoolingLoopV103CoolantType enum
var AllowedCoolingLoopV103CoolantTypeEnumValues = []CoolingLoopV103CoolantType{
	"Water",
	"Hydrocarbon",
	"Fluorocarbon",
	"Dielectric",
}

// validCoolingLoopV103CoolantTypeEnumValue provides a map of CoolingLoopV103CoolantTypes for fast verification of use input
var validCoolingLoopV103CoolantTypeEnumValues = map[CoolingLoopV103CoolantType]struct{}{
	"Water":        {},
	"Hydrocarbon":  {},
	"Fluorocarbon": {},
	"Dielectric":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CoolingLoopV103CoolantType) IsValid() bool {
	_, ok := validCoolingLoopV103CoolantTypeEnumValues[v]
	return ok
}

// NewCoolingLoopV103CoolantTypeFromValue returns a pointer to a valid CoolingLoopV103CoolantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCoolingLoopV103CoolantTypeFromValue(v string) (CoolingLoopV103CoolantType, error) {
	ev := CoolingLoopV103CoolantType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CoolingLoopV103CoolantType: valid values are %v", v, AllowedCoolingLoopV103CoolantTypeEnumValues)
}

// AssertCoolingLoopV103CoolantTypeRequired checks if the required fields are not zero-ed
func AssertCoolingLoopV103CoolantTypeRequired(obj CoolingLoopV103CoolantType) error {
	return nil
}

// AssertCoolingLoopV103CoolantTypeConstraints checks if the values respects the defined constraints
func AssertCoolingLoopV103CoolantTypeConstraints(obj CoolingLoopV103CoolantType) error {
	return nil
}