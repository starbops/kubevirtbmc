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

type ComponentIntegrityV122SpdMmeasurementSummaryType string

// List of ComponentIntegrityV122SpdMmeasurementSummaryType
const (
	COMPONENTINTEGRITYV122SPDMMEASUREMENTSUMMARYTYPE_TCB ComponentIntegrityV122SpdMmeasurementSummaryType = "TCB"
	COMPONENTINTEGRITYV122SPDMMEASUREMENTSUMMARYTYPE_ALL ComponentIntegrityV122SpdMmeasurementSummaryType = "All"
)

// AllowedComponentIntegrityV122SpdMmeasurementSummaryTypeEnumValues is all the allowed values of ComponentIntegrityV122SpdMmeasurementSummaryType enum
var AllowedComponentIntegrityV122SpdMmeasurementSummaryTypeEnumValues = []ComponentIntegrityV122SpdMmeasurementSummaryType{
	"TCB",
	"All",
}

// validComponentIntegrityV122SpdMmeasurementSummaryTypeEnumValue provides a map of ComponentIntegrityV122SpdMmeasurementSummaryTypes for fast verification of use input
var validComponentIntegrityV122SpdMmeasurementSummaryTypeEnumValues = map[ComponentIntegrityV122SpdMmeasurementSummaryType]struct{}{
	"TCB": {},
	"All": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComponentIntegrityV122SpdMmeasurementSummaryType) IsValid() bool {
	_, ok := validComponentIntegrityV122SpdMmeasurementSummaryTypeEnumValues[v]
	return ok
}

// NewComponentIntegrityV122SpdMmeasurementSummaryTypeFromValue returns a pointer to a valid ComponentIntegrityV122SpdMmeasurementSummaryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComponentIntegrityV122SpdMmeasurementSummaryTypeFromValue(v string) (ComponentIntegrityV122SpdMmeasurementSummaryType, error) {
	ev := ComponentIntegrityV122SpdMmeasurementSummaryType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComponentIntegrityV122SpdMmeasurementSummaryType: valid values are %v", v, AllowedComponentIntegrityV122SpdMmeasurementSummaryTypeEnumValues)
}

// AssertComponentIntegrityV122SpdMmeasurementSummaryTypeRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122SpdMmeasurementSummaryTypeRequired(obj ComponentIntegrityV122SpdMmeasurementSummaryType) error {
	return nil
}

// AssertComponentIntegrityV122SpdMmeasurementSummaryTypeConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122SpdMmeasurementSummaryTypeConstraints(obj ComponentIntegrityV122SpdMmeasurementSummaryType) error {
	return nil
}