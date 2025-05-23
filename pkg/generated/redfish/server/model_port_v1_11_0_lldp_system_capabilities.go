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

type PortV1110LldpSystemCapabilities string

// List of PortV1110LldpSystemCapabilities
const (
	PORTV1110LLDPSYSTEMCAPABILITIES_NONE                PortV1110LldpSystemCapabilities = "None"
	PORTV1110LLDPSYSTEMCAPABILITIES_BRIDGE              PortV1110LldpSystemCapabilities = "Bridge"
	PORTV1110LLDPSYSTEMCAPABILITIES_DOCSIS_CABLE_DEVICE PortV1110LldpSystemCapabilities = "DOCSISCableDevice"
	PORTV1110LLDPSYSTEMCAPABILITIES_OTHER               PortV1110LldpSystemCapabilities = "Other"
	PORTV1110LLDPSYSTEMCAPABILITIES_REPEATER            PortV1110LldpSystemCapabilities = "Repeater"
	PORTV1110LLDPSYSTEMCAPABILITIES_ROUTER              PortV1110LldpSystemCapabilities = "Router"
	PORTV1110LLDPSYSTEMCAPABILITIES_STATION             PortV1110LldpSystemCapabilities = "Station"
	PORTV1110LLDPSYSTEMCAPABILITIES_TELEPHONE           PortV1110LldpSystemCapabilities = "Telephone"
	PORTV1110LLDPSYSTEMCAPABILITIES_WLAN_ACCESS_POINT   PortV1110LldpSystemCapabilities = "WLANAccessPoint"
)

// AllowedPortV1110LldpSystemCapabilitiesEnumValues is all the allowed values of PortV1110LldpSystemCapabilities enum
var AllowedPortV1110LldpSystemCapabilitiesEnumValues = []PortV1110LldpSystemCapabilities{
	"None",
	"Bridge",
	"DOCSISCableDevice",
	"Other",
	"Repeater",
	"Router",
	"Station",
	"Telephone",
	"WLANAccessPoint",
}

// validPortV1110LldpSystemCapabilitiesEnumValue provides a map of PortV1110LldpSystemCapabilitiess for fast verification of use input
var validPortV1110LldpSystemCapabilitiesEnumValues = map[PortV1110LldpSystemCapabilities]struct{}{
	"None":              {},
	"Bridge":            {},
	"DOCSISCableDevice": {},
	"Other":             {},
	"Repeater":          {},
	"Router":            {},
	"Station":           {},
	"Telephone":         {},
	"WLANAccessPoint":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortV1110LldpSystemCapabilities) IsValid() bool {
	_, ok := validPortV1110LldpSystemCapabilitiesEnumValues[v]
	return ok
}

// NewPortV1110LldpSystemCapabilitiesFromValue returns a pointer to a valid PortV1110LldpSystemCapabilities
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortV1110LldpSystemCapabilitiesFromValue(v string) (PortV1110LldpSystemCapabilities, error) {
	ev := PortV1110LldpSystemCapabilities(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PortV1110LldpSystemCapabilities: valid values are %v", v, AllowedPortV1110LldpSystemCapabilitiesEnumValues)
}

// AssertPortV1110LldpSystemCapabilitiesRequired checks if the required fields are not zero-ed
func AssertPortV1110LldpSystemCapabilitiesRequired(obj PortV1110LldpSystemCapabilities) error {
	return nil
}

// AssertPortV1110LldpSystemCapabilitiesConstraints checks if the values respects the defined constraints
func AssertPortV1110LldpSystemCapabilitiesConstraints(obj PortV1110LldpSystemCapabilities) error {
	return nil
}
