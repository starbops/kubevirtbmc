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

type CircuitPhaseWiringType string

// List of CircuitPhaseWiringType
const (
	CIRCUITPHASEWIRINGTYPE_ONE_PHASE3_WIRE        CircuitPhaseWiringType = "OnePhase3Wire"
	CIRCUITPHASEWIRINGTYPE_TWO_PHASE3_WIRE        CircuitPhaseWiringType = "TwoPhase3Wire"
	CIRCUITPHASEWIRINGTYPE_ONE_OR_TWO_PHASE3_WIRE CircuitPhaseWiringType = "OneOrTwoPhase3Wire"
	CIRCUITPHASEWIRINGTYPE_TWO_PHASE4_WIRE        CircuitPhaseWiringType = "TwoPhase4Wire"
	CIRCUITPHASEWIRINGTYPE_THREE_PHASE4_WIRE      CircuitPhaseWiringType = "ThreePhase4Wire"
	CIRCUITPHASEWIRINGTYPE_THREE_PHASE5_WIRE      CircuitPhaseWiringType = "ThreePhase5Wire"
)

// AllowedCircuitPhaseWiringTypeEnumValues is all the allowed values of CircuitPhaseWiringType enum
var AllowedCircuitPhaseWiringTypeEnumValues = []CircuitPhaseWiringType{
	"OnePhase3Wire",
	"TwoPhase3Wire",
	"OneOrTwoPhase3Wire",
	"TwoPhase4Wire",
	"ThreePhase4Wire",
	"ThreePhase5Wire",
}

// validCircuitPhaseWiringTypeEnumValue provides a map of CircuitPhaseWiringTypes for fast verification of use input
var validCircuitPhaseWiringTypeEnumValues = map[CircuitPhaseWiringType]struct{}{
	"OnePhase3Wire":      {},
	"TwoPhase3Wire":      {},
	"OneOrTwoPhase3Wire": {},
	"TwoPhase4Wire":      {},
	"ThreePhase4Wire":    {},
	"ThreePhase5Wire":    {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CircuitPhaseWiringType) IsValid() bool {
	_, ok := validCircuitPhaseWiringTypeEnumValues[v]
	return ok
}

// NewCircuitPhaseWiringTypeFromValue returns a pointer to a valid CircuitPhaseWiringType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCircuitPhaseWiringTypeFromValue(v string) (CircuitPhaseWiringType, error) {
	ev := CircuitPhaseWiringType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CircuitPhaseWiringType: valid values are %v", v, AllowedCircuitPhaseWiringTypeEnumValues)
}

// AssertCircuitPhaseWiringTypeRequired checks if the required fields are not zero-ed
func AssertCircuitPhaseWiringTypeRequired(obj CircuitPhaseWiringType) error {
	return nil
}

// AssertCircuitPhaseWiringTypeConstraints checks if the values respects the defined constraints
func AssertCircuitPhaseWiringTypeConstraints(obj CircuitPhaseWiringType) error {
	return nil
}