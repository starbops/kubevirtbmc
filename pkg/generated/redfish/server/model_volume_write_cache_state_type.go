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

type VolumeWriteCacheStateType string

// List of VolumeWriteCacheStateType
const (
	VOLUMEWRITECACHESTATETYPE_UNPROTECTED VolumeWriteCacheStateType = "Unprotected"
	VOLUMEWRITECACHESTATETYPE_PROTECTED   VolumeWriteCacheStateType = "Protected"
	VOLUMEWRITECACHESTATETYPE_DEGRADED    VolumeWriteCacheStateType = "Degraded"
)

// AllowedVolumeWriteCacheStateTypeEnumValues is all the allowed values of VolumeWriteCacheStateType enum
var AllowedVolumeWriteCacheStateTypeEnumValues = []VolumeWriteCacheStateType{
	"Unprotected",
	"Protected",
	"Degraded",
}

// validVolumeWriteCacheStateTypeEnumValue provides a map of VolumeWriteCacheStateTypes for fast verification of use input
var validVolumeWriteCacheStateTypeEnumValues = map[VolumeWriteCacheStateType]struct{}{
	"Unprotected": {},
	"Protected":   {},
	"Degraded":    {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VolumeWriteCacheStateType) IsValid() bool {
	_, ok := validVolumeWriteCacheStateTypeEnumValues[v]
	return ok
}

// NewVolumeWriteCacheStateTypeFromValue returns a pointer to a valid VolumeWriteCacheStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVolumeWriteCacheStateTypeFromValue(v string) (VolumeWriteCacheStateType, error) {
	ev := VolumeWriteCacheStateType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for VolumeWriteCacheStateType: valid values are %v", v, AllowedVolumeWriteCacheStateTypeEnumValues)
}

// AssertVolumeWriteCacheStateTypeRequired checks if the required fields are not zero-ed
func AssertVolumeWriteCacheStateTypeRequired(obj VolumeWriteCacheStateType) error {
	return nil
}

// AssertVolumeWriteCacheStateTypeConstraints checks if the values respects the defined constraints
func AssertVolumeWriteCacheStateTypeConstraints(obj VolumeWriteCacheStateType) error {
	return nil
}