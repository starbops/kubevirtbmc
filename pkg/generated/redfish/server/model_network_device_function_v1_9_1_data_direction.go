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

type NetworkDeviceFunctionV191DataDirection string

// List of NetworkDeviceFunctionV191DataDirection
const (
	NETWORKDEVICEFUNCTIONV191DATADIRECTION_NONE    NetworkDeviceFunctionV191DataDirection = "None"
	NETWORKDEVICEFUNCTIONV191DATADIRECTION_INGRESS NetworkDeviceFunctionV191DataDirection = "Ingress"
	NETWORKDEVICEFUNCTIONV191DATADIRECTION_EGRESS  NetworkDeviceFunctionV191DataDirection = "Egress"
)

// AllowedNetworkDeviceFunctionV191DataDirectionEnumValues is all the allowed values of NetworkDeviceFunctionV191DataDirection enum
var AllowedNetworkDeviceFunctionV191DataDirectionEnumValues = []NetworkDeviceFunctionV191DataDirection{
	"None",
	"Ingress",
	"Egress",
}

// validNetworkDeviceFunctionV191DataDirectionEnumValue provides a map of NetworkDeviceFunctionV191DataDirections for fast verification of use input
var validNetworkDeviceFunctionV191DataDirectionEnumValues = map[NetworkDeviceFunctionV191DataDirection]struct{}{
	"None":    {},
	"Ingress": {},
	"Egress":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkDeviceFunctionV191DataDirection) IsValid() bool {
	_, ok := validNetworkDeviceFunctionV191DataDirectionEnumValues[v]
	return ok
}

// NewNetworkDeviceFunctionV191DataDirectionFromValue returns a pointer to a valid NetworkDeviceFunctionV191DataDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkDeviceFunctionV191DataDirectionFromValue(v string) (NetworkDeviceFunctionV191DataDirection, error) {
	ev := NetworkDeviceFunctionV191DataDirection(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for NetworkDeviceFunctionV191DataDirection: valid values are %v", v, AllowedNetworkDeviceFunctionV191DataDirectionEnumValues)
}

// AssertNetworkDeviceFunctionV191DataDirectionRequired checks if the required fields are not zero-ed
func AssertNetworkDeviceFunctionV191DataDirectionRequired(obj NetworkDeviceFunctionV191DataDirection) error {
	return nil
}

// AssertNetworkDeviceFunctionV191DataDirectionConstraints checks if the values respects the defined constraints
func AssertNetworkDeviceFunctionV191DataDirectionConstraints(obj NetworkDeviceFunctionV191DataDirection) error {
	return nil
}
