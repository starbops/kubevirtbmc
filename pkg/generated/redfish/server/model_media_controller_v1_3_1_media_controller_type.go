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

type MediaControllerV131MediaControllerType string

// List of MediaControllerV131MediaControllerType
const (
	MEDIACONTROLLERV131MEDIACONTROLLERTYPE_MEMORY MediaControllerV131MediaControllerType = "Memory"
)

// AllowedMediaControllerV131MediaControllerTypeEnumValues is all the allowed values of MediaControllerV131MediaControllerType enum
var AllowedMediaControllerV131MediaControllerTypeEnumValues = []MediaControllerV131MediaControllerType{
	"Memory",
}

// validMediaControllerV131MediaControllerTypeEnumValue provides a map of MediaControllerV131MediaControllerTypes for fast verification of use input
var validMediaControllerV131MediaControllerTypeEnumValues = map[MediaControllerV131MediaControllerType]struct{}{
	"Memory": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaControllerV131MediaControllerType) IsValid() bool {
	_, ok := validMediaControllerV131MediaControllerTypeEnumValues[v]
	return ok
}

// NewMediaControllerV131MediaControllerTypeFromValue returns a pointer to a valid MediaControllerV131MediaControllerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaControllerV131MediaControllerTypeFromValue(v string) (MediaControllerV131MediaControllerType, error) {
	ev := MediaControllerV131MediaControllerType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for MediaControllerV131MediaControllerType: valid values are %v", v, AllowedMediaControllerV131MediaControllerTypeEnumValues)
}

// AssertMediaControllerV131MediaControllerTypeRequired checks if the required fields are not zero-ed
func AssertMediaControllerV131MediaControllerTypeRequired(obj MediaControllerV131MediaControllerType) error {
	return nil
}

// AssertMediaControllerV131MediaControllerTypeConstraints checks if the values respects the defined constraints
func AssertMediaControllerV131MediaControllerTypeConstraints(obj MediaControllerV131MediaControllerType) error {
	return nil
}