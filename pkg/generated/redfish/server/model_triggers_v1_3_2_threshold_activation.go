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

type TriggersV132ThresholdActivation string

// List of TriggersV132ThresholdActivation
const (
	TRIGGERSV132THRESHOLDACTIVATION_INCREASING TriggersV132ThresholdActivation = "Increasing"
	TRIGGERSV132THRESHOLDACTIVATION_DECREASING TriggersV132ThresholdActivation = "Decreasing"
	TRIGGERSV132THRESHOLDACTIVATION_EITHER     TriggersV132ThresholdActivation = "Either"
	TRIGGERSV132THRESHOLDACTIVATION_DISABLED   TriggersV132ThresholdActivation = "Disabled"
)

// AllowedTriggersV132ThresholdActivationEnumValues is all the allowed values of TriggersV132ThresholdActivation enum
var AllowedTriggersV132ThresholdActivationEnumValues = []TriggersV132ThresholdActivation{
	"Increasing",
	"Decreasing",
	"Either",
	"Disabled",
}

// validTriggersV132ThresholdActivationEnumValue provides a map of TriggersV132ThresholdActivations for fast verification of use input
var validTriggersV132ThresholdActivationEnumValues = map[TriggersV132ThresholdActivation]struct{}{
	"Increasing": {},
	"Decreasing": {},
	"Either":     {},
	"Disabled":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggersV132ThresholdActivation) IsValid() bool {
	_, ok := validTriggersV132ThresholdActivationEnumValues[v]
	return ok
}

// NewTriggersV132ThresholdActivationFromValue returns a pointer to a valid TriggersV132ThresholdActivation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggersV132ThresholdActivationFromValue(v string) (TriggersV132ThresholdActivation, error) {
	ev := TriggersV132ThresholdActivation(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TriggersV132ThresholdActivation: valid values are %v", v, AllowedTriggersV132ThresholdActivationEnumValues)
}

// AssertTriggersV132ThresholdActivationRequired checks if the required fields are not zero-ed
func AssertTriggersV132ThresholdActivationRequired(obj TriggersV132ThresholdActivation) error {
	return nil
}

// AssertTriggersV132ThresholdActivationConstraints checks if the values respects the defined constraints
func AssertTriggersV132ThresholdActivationConstraints(obj TriggersV132ThresholdActivation) error {
	return nil
}