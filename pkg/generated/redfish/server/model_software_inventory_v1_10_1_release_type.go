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

type SoftwareInventoryV1101ReleaseType string

// List of SoftwareInventoryV1101ReleaseType
const (
	SOFTWAREINVENTORYV1101RELEASETYPE_PRODUCTION SoftwareInventoryV1101ReleaseType = "Production"
	SOFTWAREINVENTORYV1101RELEASETYPE_PROTOTYPE  SoftwareInventoryV1101ReleaseType = "Prototype"
	SOFTWAREINVENTORYV1101RELEASETYPE_OTHER      SoftwareInventoryV1101ReleaseType = "Other"
)

// AllowedSoftwareInventoryV1101ReleaseTypeEnumValues is all the allowed values of SoftwareInventoryV1101ReleaseType enum
var AllowedSoftwareInventoryV1101ReleaseTypeEnumValues = []SoftwareInventoryV1101ReleaseType{
	"Production",
	"Prototype",
	"Other",
}

// validSoftwareInventoryV1101ReleaseTypeEnumValue provides a map of SoftwareInventoryV1101ReleaseTypes for fast verification of use input
var validSoftwareInventoryV1101ReleaseTypeEnumValues = map[SoftwareInventoryV1101ReleaseType]struct{}{
	"Production": {},
	"Prototype":  {},
	"Other":      {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SoftwareInventoryV1101ReleaseType) IsValid() bool {
	_, ok := validSoftwareInventoryV1101ReleaseTypeEnumValues[v]
	return ok
}

// NewSoftwareInventoryV1101ReleaseTypeFromValue returns a pointer to a valid SoftwareInventoryV1101ReleaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSoftwareInventoryV1101ReleaseTypeFromValue(v string) (SoftwareInventoryV1101ReleaseType, error) {
	ev := SoftwareInventoryV1101ReleaseType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for SoftwareInventoryV1101ReleaseType: valid values are %v", v, AllowedSoftwareInventoryV1101ReleaseTypeEnumValues)
}

// AssertSoftwareInventoryV1101ReleaseTypeRequired checks if the required fields are not zero-ed
func AssertSoftwareInventoryV1101ReleaseTypeRequired(obj SoftwareInventoryV1101ReleaseType) error {
	return nil
}

// AssertSoftwareInventoryV1101ReleaseTypeConstraints checks if the values respects the defined constraints
func AssertSoftwareInventoryV1101ReleaseTypeConstraints(obj SoftwareInventoryV1101ReleaseType) error {
	return nil
}
