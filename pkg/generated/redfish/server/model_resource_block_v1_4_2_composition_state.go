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

type ResourceBlockV142CompositionState string

// List of ResourceBlockV142CompositionState
const (
	RESOURCEBLOCKV142COMPOSITIONSTATE_COMPOSING              ResourceBlockV142CompositionState = "Composing"
	RESOURCEBLOCKV142COMPOSITIONSTATE_COMPOSED_AND_AVAILABLE ResourceBlockV142CompositionState = "ComposedAndAvailable"
	RESOURCEBLOCKV142COMPOSITIONSTATE_COMPOSED               ResourceBlockV142CompositionState = "Composed"
	RESOURCEBLOCKV142COMPOSITIONSTATE_UNUSED                 ResourceBlockV142CompositionState = "Unused"
	RESOURCEBLOCKV142COMPOSITIONSTATE_FAILED                 ResourceBlockV142CompositionState = "Failed"
	RESOURCEBLOCKV142COMPOSITIONSTATE_UNAVAILABLE            ResourceBlockV142CompositionState = "Unavailable"
)

// AllowedResourceBlockV142CompositionStateEnumValues is all the allowed values of ResourceBlockV142CompositionState enum
var AllowedResourceBlockV142CompositionStateEnumValues = []ResourceBlockV142CompositionState{
	"Composing",
	"ComposedAndAvailable",
	"Composed",
	"Unused",
	"Failed",
	"Unavailable",
}

// validResourceBlockV142CompositionStateEnumValue provides a map of ResourceBlockV142CompositionStates for fast verification of use input
var validResourceBlockV142CompositionStateEnumValues = map[ResourceBlockV142CompositionState]struct{}{
	"Composing":            {},
	"ComposedAndAvailable": {},
	"Composed":             {},
	"Unused":               {},
	"Failed":               {},
	"Unavailable":          {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceBlockV142CompositionState) IsValid() bool {
	_, ok := validResourceBlockV142CompositionStateEnumValues[v]
	return ok
}

// NewResourceBlockV142CompositionStateFromValue returns a pointer to a valid ResourceBlockV142CompositionState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceBlockV142CompositionStateFromValue(v string) (ResourceBlockV142CompositionState, error) {
	ev := ResourceBlockV142CompositionState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ResourceBlockV142CompositionState: valid values are %v", v, AllowedResourceBlockV142CompositionStateEnumValues)
}

// AssertResourceBlockV142CompositionStateRequired checks if the required fields are not zero-ed
func AssertResourceBlockV142CompositionStateRequired(obj ResourceBlockV142CompositionState) error {
	return nil
}

// AssertResourceBlockV142CompositionStateConstraints checks if the values respects the defined constraints
func AssertResourceBlockV142CompositionStateConstraints(obj ResourceBlockV142CompositionState) error {
	return nil
}