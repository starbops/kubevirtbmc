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

type SessionV171SessionTypes string

// List of SessionV171SessionTypes
const (
	SESSIONV171SESSIONTYPES_HOST_CONSOLE        SessionV171SessionTypes = "HostConsole"
	SESSIONV171SESSIONTYPES_MANAGER_CONSOLE     SessionV171SessionTypes = "ManagerConsole"
	SESSIONV171SESSIONTYPES_IPMI                SessionV171SessionTypes = "IPMI"
	SESSIONV171SESSIONTYPES_KVMIP               SessionV171SessionTypes = "KVMIP"
	SESSIONV171SESSIONTYPES_OEM                 SessionV171SessionTypes = "OEM"
	SESSIONV171SESSIONTYPES_REDFISH             SessionV171SessionTypes = "Redfish"
	SESSIONV171SESSIONTYPES_VIRTUAL_MEDIA       SessionV171SessionTypes = "VirtualMedia"
	SESSIONV171SESSIONTYPES_WEB_UI              SessionV171SessionTypes = "WebUI"
	SESSIONV171SESSIONTYPES_OUTBOUND_CONNECTION SessionV171SessionTypes = "OutboundConnection"
)

// AllowedSessionV171SessionTypesEnumValues is all the allowed values of SessionV171SessionTypes enum
var AllowedSessionV171SessionTypesEnumValues = []SessionV171SessionTypes{
	"HostConsole",
	"ManagerConsole",
	"IPMI",
	"KVMIP",
	"OEM",
	"Redfish",
	"VirtualMedia",
	"WebUI",
	"OutboundConnection",
}

// validSessionV171SessionTypesEnumValue provides a map of SessionV171SessionTypess for fast verification of use input
var validSessionV171SessionTypesEnumValues = map[SessionV171SessionTypes]struct{}{
	"HostConsole":        {},
	"ManagerConsole":     {},
	"IPMI":               {},
	"KVMIP":              {},
	"OEM":                {},
	"Redfish":            {},
	"VirtualMedia":       {},
	"WebUI":              {},
	"OutboundConnection": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionV171SessionTypes) IsValid() bool {
	_, ok := validSessionV171SessionTypesEnumValues[v]
	return ok
}

// NewSessionV171SessionTypesFromValue returns a pointer to a valid SessionV171SessionTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionV171SessionTypesFromValue(v string) (SessionV171SessionTypes, error) {
	ev := SessionV171SessionTypes(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for SessionV171SessionTypes: valid values are %v", v, AllowedSessionV171SessionTypesEnumValues)
}

// AssertSessionV171SessionTypesRequired checks if the required fields are not zero-ed
func AssertSessionV171SessionTypesRequired(obj SessionV171SessionTypes) error {
	return nil
}

// AssertSessionV171SessionTypesConstraints checks if the values respects the defined constraints
func AssertSessionV171SessionTypesConstraints(obj SessionV171SessionTypes) error {
	return nil
}