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

type NetworkPortV142SupportedEthernetCapabilities string

// List of NetworkPortV142SupportedEthernetCapabilities
const (
	NETWORKPORTV142SUPPORTEDETHERNETCAPABILITIES_WAKE_ON_LAN NetworkPortV142SupportedEthernetCapabilities = "WakeOnLAN"
	NETWORKPORTV142SUPPORTEDETHERNETCAPABILITIES_EEE         NetworkPortV142SupportedEthernetCapabilities = "EEE"
)

// AllowedNetworkPortV142SupportedEthernetCapabilitiesEnumValues is all the allowed values of NetworkPortV142SupportedEthernetCapabilities enum
var AllowedNetworkPortV142SupportedEthernetCapabilitiesEnumValues = []NetworkPortV142SupportedEthernetCapabilities{
	"WakeOnLAN",
	"EEE",
}

// validNetworkPortV142SupportedEthernetCapabilitiesEnumValue provides a map of NetworkPortV142SupportedEthernetCapabilitiess for fast verification of use input
var validNetworkPortV142SupportedEthernetCapabilitiesEnumValues = map[NetworkPortV142SupportedEthernetCapabilities]struct{}{
	"WakeOnLAN": {},
	"EEE":       {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkPortV142SupportedEthernetCapabilities) IsValid() bool {
	_, ok := validNetworkPortV142SupportedEthernetCapabilitiesEnumValues[v]
	return ok
}

// NewNetworkPortV142SupportedEthernetCapabilitiesFromValue returns a pointer to a valid NetworkPortV142SupportedEthernetCapabilities
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkPortV142SupportedEthernetCapabilitiesFromValue(v string) (NetworkPortV142SupportedEthernetCapabilities, error) {
	ev := NetworkPortV142SupportedEthernetCapabilities(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for NetworkPortV142SupportedEthernetCapabilities: valid values are %v", v, AllowedNetworkPortV142SupportedEthernetCapabilitiesEnumValues)
}

// AssertNetworkPortV142SupportedEthernetCapabilitiesRequired checks if the required fields are not zero-ed
func AssertNetworkPortV142SupportedEthernetCapabilitiesRequired(obj NetworkPortV142SupportedEthernetCapabilities) error {
	return nil
}

// AssertNetworkPortV142SupportedEthernetCapabilitiesConstraints checks if the values respects the defined constraints
func AssertNetworkPortV142SupportedEthernetCapabilitiesConstraints(obj NetworkPortV142SupportedEthernetCapabilities) error {
	return nil
}
