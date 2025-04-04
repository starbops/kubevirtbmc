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

type PortV1110LinkState string

// List of PortV1110LinkState
const (
	PORTV1110LINKSTATE_ENABLED  PortV1110LinkState = "Enabled"
	PORTV1110LINKSTATE_DISABLED PortV1110LinkState = "Disabled"
)

// AllowedPortV1110LinkStateEnumValues is all the allowed values of PortV1110LinkState enum
var AllowedPortV1110LinkStateEnumValues = []PortV1110LinkState{
	"Enabled",
	"Disabled",
}

// validPortV1110LinkStateEnumValue provides a map of PortV1110LinkStates for fast verification of use input
var validPortV1110LinkStateEnumValues = map[PortV1110LinkState]struct{}{
	"Enabled":  {},
	"Disabled": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortV1110LinkState) IsValid() bool {
	_, ok := validPortV1110LinkStateEnumValues[v]
	return ok
}

// NewPortV1110LinkStateFromValue returns a pointer to a valid PortV1110LinkState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortV1110LinkStateFromValue(v string) (PortV1110LinkState, error) {
	ev := PortV1110LinkState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PortV1110LinkState: valid values are %v", v, AllowedPortV1110LinkStateEnumValues)
}

// AssertPortV1110LinkStateRequired checks if the required fields are not zero-ed
func AssertPortV1110LinkStateRequired(obj PortV1110LinkState) error {
	return nil
}

// AssertPortV1110LinkStateConstraints checks if the values respects the defined constraints
func AssertPortV1110LinkStateConstraints(obj PortV1110LinkState) error {
	return nil
}
