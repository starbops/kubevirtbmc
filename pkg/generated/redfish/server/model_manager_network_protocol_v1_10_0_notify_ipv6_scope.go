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

type ManagerNetworkProtocolV1100NotifyIpv6Scope string

// List of ManagerNetworkProtocolV1100NotifyIpv6Scope
const (
	MANAGERNETWORKPROTOCOLV1100NOTIFYIPV6SCOPE_LINK         ManagerNetworkProtocolV1100NotifyIpv6Scope = "Link"
	MANAGERNETWORKPROTOCOLV1100NOTIFYIPV6SCOPE_SITE         ManagerNetworkProtocolV1100NotifyIpv6Scope = "Site"
	MANAGERNETWORKPROTOCOLV1100NOTIFYIPV6SCOPE_ORGANIZATION ManagerNetworkProtocolV1100NotifyIpv6Scope = "Organization"
)

// AllowedManagerNetworkProtocolV1100NotifyIpv6ScopeEnumValues is all the allowed values of ManagerNetworkProtocolV1100NotifyIpv6Scope enum
var AllowedManagerNetworkProtocolV1100NotifyIpv6ScopeEnumValues = []ManagerNetworkProtocolV1100NotifyIpv6Scope{
	"Link",
	"Site",
	"Organization",
}

// validManagerNetworkProtocolV1100NotifyIpv6ScopeEnumValue provides a map of ManagerNetworkProtocolV1100NotifyIpv6Scopes for fast verification of use input
var validManagerNetworkProtocolV1100NotifyIpv6ScopeEnumValues = map[ManagerNetworkProtocolV1100NotifyIpv6Scope]struct{}{
	"Link":         {},
	"Site":         {},
	"Organization": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManagerNetworkProtocolV1100NotifyIpv6Scope) IsValid() bool {
	_, ok := validManagerNetworkProtocolV1100NotifyIpv6ScopeEnumValues[v]
	return ok
}

// NewManagerNetworkProtocolV1100NotifyIpv6ScopeFromValue returns a pointer to a valid ManagerNetworkProtocolV1100NotifyIpv6Scope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManagerNetworkProtocolV1100NotifyIpv6ScopeFromValue(v string) (ManagerNetworkProtocolV1100NotifyIpv6Scope, error) {
	ev := ManagerNetworkProtocolV1100NotifyIpv6Scope(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ManagerNetworkProtocolV1100NotifyIpv6Scope: valid values are %v", v, AllowedManagerNetworkProtocolV1100NotifyIpv6ScopeEnumValues)
}

// AssertManagerNetworkProtocolV1100NotifyIpv6ScopeRequired checks if the required fields are not zero-ed
func AssertManagerNetworkProtocolV1100NotifyIpv6ScopeRequired(obj ManagerNetworkProtocolV1100NotifyIpv6Scope) error {
	return nil
}

// AssertManagerNetworkProtocolV1100NotifyIpv6ScopeConstraints checks if the values respects the defined constraints
func AssertManagerNetworkProtocolV1100NotifyIpv6ScopeConstraints(obj ManagerNetworkProtocolV1100NotifyIpv6Scope) error {
	return nil
}