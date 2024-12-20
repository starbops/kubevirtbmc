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

type ProcessorV1190ThrottleCause string

// List of ProcessorV1190ThrottleCause
const (
	PROCESSORV1190THROTTLECAUSE_POWER_LIMIT               ProcessorV1190ThrottleCause = "PowerLimit"
	PROCESSORV1190THROTTLECAUSE_THERMAL_LIMIT             ProcessorV1190ThrottleCause = "ThermalLimit"
	PROCESSORV1190THROTTLECAUSE_CLOCK_LIMIT               ProcessorV1190ThrottleCause = "ClockLimit"
	PROCESSORV1190THROTTLECAUSE_MANAGEMENT_DETECTED_FAULT ProcessorV1190ThrottleCause = "ManagementDetectedFault"
	PROCESSORV1190THROTTLECAUSE_UNKNOWN                   ProcessorV1190ThrottleCause = "Unknown"
	PROCESSORV1190THROTTLECAUSE_OEM                       ProcessorV1190ThrottleCause = "OEM"
)

// AllowedProcessorV1190ThrottleCauseEnumValues is all the allowed values of ProcessorV1190ThrottleCause enum
var AllowedProcessorV1190ThrottleCauseEnumValues = []ProcessorV1190ThrottleCause{
	"PowerLimit",
	"ThermalLimit",
	"ClockLimit",
	"ManagementDetectedFault",
	"Unknown",
	"OEM",
}

// validProcessorV1190ThrottleCauseEnumValue provides a map of ProcessorV1190ThrottleCauses for fast verification of use input
var validProcessorV1190ThrottleCauseEnumValues = map[ProcessorV1190ThrottleCause]struct{}{
	"PowerLimit":              {},
	"ThermalLimit":            {},
	"ClockLimit":              {},
	"ManagementDetectedFault": {},
	"Unknown":                 {},
	"OEM":                     {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProcessorV1190ThrottleCause) IsValid() bool {
	_, ok := validProcessorV1190ThrottleCauseEnumValues[v]
	return ok
}

// NewProcessorV1190ThrottleCauseFromValue returns a pointer to a valid ProcessorV1190ThrottleCause
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProcessorV1190ThrottleCauseFromValue(v string) (ProcessorV1190ThrottleCause, error) {
	ev := ProcessorV1190ThrottleCause(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ProcessorV1190ThrottleCause: valid values are %v", v, AllowedProcessorV1190ThrottleCauseEnumValues)
}

// AssertProcessorV1190ThrottleCauseRequired checks if the required fields are not zero-ed
func AssertProcessorV1190ThrottleCauseRequired(obj ProcessorV1190ThrottleCause) error {
	return nil
}

// AssertProcessorV1190ThrottleCauseConstraints checks if the values respects the defined constraints
func AssertProcessorV1190ThrottleCauseConstraints(obj ProcessorV1190ThrottleCause) error {
	return nil
}
