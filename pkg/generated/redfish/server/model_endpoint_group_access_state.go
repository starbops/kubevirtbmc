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

// EndpointGroupAccessState : Describes the options for the access characteristics of an endpoint.
type EndpointGroupAccessState string

// List of EndpointGroupAccessState
const (
	ENDPOINTGROUPACCESSSTATE_OPTIMIZED     EndpointGroupAccessState = "Optimized"
	ENDPOINTGROUPACCESSSTATE_NON_OPTIMIZED EndpointGroupAccessState = "NonOptimized"
	ENDPOINTGROUPACCESSSTATE_STANDBY       EndpointGroupAccessState = "Standby"
	ENDPOINTGROUPACCESSSTATE_UNAVAILABLE   EndpointGroupAccessState = "Unavailable"
	ENDPOINTGROUPACCESSSTATE_TRANSITIONING EndpointGroupAccessState = "Transitioning"
)

// AllowedEndpointGroupAccessStateEnumValues is all the allowed values of EndpointGroupAccessState enum
var AllowedEndpointGroupAccessStateEnumValues = []EndpointGroupAccessState{
	"Optimized",
	"NonOptimized",
	"Standby",
	"Unavailable",
	"Transitioning",
}

// validEndpointGroupAccessStateEnumValue provides a map of EndpointGroupAccessStates for fast verification of use input
var validEndpointGroupAccessStateEnumValues = map[EndpointGroupAccessState]struct{}{
	"Optimized":     {},
	"NonOptimized":  {},
	"Standby":       {},
	"Unavailable":   {},
	"Transitioning": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EndpointGroupAccessState) IsValid() bool {
	_, ok := validEndpointGroupAccessStateEnumValues[v]
	return ok
}

// NewEndpointGroupAccessStateFromValue returns a pointer to a valid EndpointGroupAccessState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndpointGroupAccessStateFromValue(v string) (EndpointGroupAccessState, error) {
	ev := EndpointGroupAccessState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for EndpointGroupAccessState: valid values are %v", v, AllowedEndpointGroupAccessStateEnumValues)
}

// AssertEndpointGroupAccessStateRequired checks if the required fields are not zero-ed
func AssertEndpointGroupAccessStateRequired(obj EndpointGroupAccessState) error {
	return nil
}

// AssertEndpointGroupAccessStateConstraints checks if the values respects the defined constraints
func AssertEndpointGroupAccessStateConstraints(obj EndpointGroupAccessState) error {
	return nil
}
