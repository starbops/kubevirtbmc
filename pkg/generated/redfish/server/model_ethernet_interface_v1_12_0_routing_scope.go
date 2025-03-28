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

type EthernetInterfaceV1120RoutingScope string

// List of EthernetInterfaceV1120RoutingScope
const (
	ETHERNETINTERFACEV1120ROUTINGSCOPE_EXTERNAL  EthernetInterfaceV1120RoutingScope = "External"
	ETHERNETINTERFACEV1120ROUTINGSCOPE_HOST_ONLY EthernetInterfaceV1120RoutingScope = "HostOnly"
	ETHERNETINTERFACEV1120ROUTINGSCOPE_INTERNAL  EthernetInterfaceV1120RoutingScope = "Internal"
	ETHERNETINTERFACEV1120ROUTINGSCOPE_LIMITED   EthernetInterfaceV1120RoutingScope = "Limited"
)

// AllowedEthernetInterfaceV1120RoutingScopeEnumValues is all the allowed values of EthernetInterfaceV1120RoutingScope enum
var AllowedEthernetInterfaceV1120RoutingScopeEnumValues = []EthernetInterfaceV1120RoutingScope{
	"External",
	"HostOnly",
	"Internal",
	"Limited",
}

// validEthernetInterfaceV1120RoutingScopeEnumValue provides a map of EthernetInterfaceV1120RoutingScopes for fast verification of use input
var validEthernetInterfaceV1120RoutingScopeEnumValues = map[EthernetInterfaceV1120RoutingScope]struct{}{
	"External": {},
	"HostOnly": {},
	"Internal": {},
	"Limited":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EthernetInterfaceV1120RoutingScope) IsValid() bool {
	_, ok := validEthernetInterfaceV1120RoutingScopeEnumValues[v]
	return ok
}

// NewEthernetInterfaceV1120RoutingScopeFromValue returns a pointer to a valid EthernetInterfaceV1120RoutingScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEthernetInterfaceV1120RoutingScopeFromValue(v string) (EthernetInterfaceV1120RoutingScope, error) {
	ev := EthernetInterfaceV1120RoutingScope(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for EthernetInterfaceV1120RoutingScope: valid values are %v", v, AllowedEthernetInterfaceV1120RoutingScopeEnumValues)
}

// AssertEthernetInterfaceV1120RoutingScopeRequired checks if the required fields are not zero-ed
func AssertEthernetInterfaceV1120RoutingScopeRequired(obj EthernetInterfaceV1120RoutingScope) error {
	return nil
}

// AssertEthernetInterfaceV1120RoutingScopeConstraints checks if the values respects the defined constraints
func AssertEthernetInterfaceV1120RoutingScopeConstraints(obj EthernetInterfaceV1120RoutingScope) error {
	return nil
}
