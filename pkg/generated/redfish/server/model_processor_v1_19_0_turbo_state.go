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

type ProcessorV1190TurboState string

// List of ProcessorV1190TurboState
const (
	PROCESSORV1190TURBOSTATE_ENABLED  ProcessorV1190TurboState = "Enabled"
	PROCESSORV1190TURBOSTATE_DISABLED ProcessorV1190TurboState = "Disabled"
)

// AllowedProcessorV1190TurboStateEnumValues is all the allowed values of ProcessorV1190TurboState enum
var AllowedProcessorV1190TurboStateEnumValues = []ProcessorV1190TurboState{
	"Enabled",
	"Disabled",
}

// validProcessorV1190TurboStateEnumValue provides a map of ProcessorV1190TurboStates for fast verification of use input
var validProcessorV1190TurboStateEnumValues = map[ProcessorV1190TurboState]struct{}{
	"Enabled":  {},
	"Disabled": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProcessorV1190TurboState) IsValid() bool {
	_, ok := validProcessorV1190TurboStateEnumValues[v]
	return ok
}

// NewProcessorV1190TurboStateFromValue returns a pointer to a valid ProcessorV1190TurboState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProcessorV1190TurboStateFromValue(v string) (ProcessorV1190TurboState, error) {
	ev := ProcessorV1190TurboState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ProcessorV1190TurboState: valid values are %v", v, AllowedProcessorV1190TurboStateEnumValues)
}

// AssertProcessorV1190TurboStateRequired checks if the required fields are not zero-ed
func AssertProcessorV1190TurboStateRequired(obj ProcessorV1190TurboState) error {
	return nil
}

// AssertProcessorV1190TurboStateConstraints checks if the values respects the defined constraints
func AssertProcessorV1190TurboStateConstraints(obj ProcessorV1190TurboState) error {
	return nil
}