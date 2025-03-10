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

type LogServiceV160LogDiagnosticDataTypes string

// List of LogServiceV160LogDiagnosticDataTypes
const (
	LOGSERVICEV160LOGDIAGNOSTICDATATYPES_MANAGER LogServiceV160LogDiagnosticDataTypes = "Manager"
	LOGSERVICEV160LOGDIAGNOSTICDATATYPES_PRE_OS  LogServiceV160LogDiagnosticDataTypes = "PreOS"
	LOGSERVICEV160LOGDIAGNOSTICDATATYPES_OS      LogServiceV160LogDiagnosticDataTypes = "OS"
	LOGSERVICEV160LOGDIAGNOSTICDATATYPES_OEM     LogServiceV160LogDiagnosticDataTypes = "OEM"
)

// AllowedLogServiceV160LogDiagnosticDataTypesEnumValues is all the allowed values of LogServiceV160LogDiagnosticDataTypes enum
var AllowedLogServiceV160LogDiagnosticDataTypesEnumValues = []LogServiceV160LogDiagnosticDataTypes{
	"Manager",
	"PreOS",
	"OS",
	"OEM",
}

// validLogServiceV160LogDiagnosticDataTypesEnumValue provides a map of LogServiceV160LogDiagnosticDataTypess for fast verification of use input
var validLogServiceV160LogDiagnosticDataTypesEnumValues = map[LogServiceV160LogDiagnosticDataTypes]struct{}{
	"Manager": {},
	"PreOS":   {},
	"OS":      {},
	"OEM":     {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogServiceV160LogDiagnosticDataTypes) IsValid() bool {
	_, ok := validLogServiceV160LogDiagnosticDataTypesEnumValues[v]
	return ok
}

// NewLogServiceV160LogDiagnosticDataTypesFromValue returns a pointer to a valid LogServiceV160LogDiagnosticDataTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogServiceV160LogDiagnosticDataTypesFromValue(v string) (LogServiceV160LogDiagnosticDataTypes, error) {
	ev := LogServiceV160LogDiagnosticDataTypes(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for LogServiceV160LogDiagnosticDataTypes: valid values are %v", v, AllowedLogServiceV160LogDiagnosticDataTypesEnumValues)
}

// AssertLogServiceV160LogDiagnosticDataTypesRequired checks if the required fields are not zero-ed
func AssertLogServiceV160LogDiagnosticDataTypesRequired(obj LogServiceV160LogDiagnosticDataTypes) error {
	return nil
}

// AssertLogServiceV160LogDiagnosticDataTypesConstraints checks if the values respects the defined constraints
func AssertLogServiceV160LogDiagnosticDataTypesConstraints(obj LogServiceV160LogDiagnosticDataTypes) error {
	return nil
}
