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

type PortV1110LinkNetworkTechnology string

// List of PortV1110LinkNetworkTechnology
const (
	PORTV1110LINKNETWORKTECHNOLOGY_ETHERNET      PortV1110LinkNetworkTechnology = "Ethernet"
	PORTV1110LINKNETWORKTECHNOLOGY_INFINI_BAND   PortV1110LinkNetworkTechnology = "InfiniBand"
	PORTV1110LINKNETWORKTECHNOLOGY_FIBRE_CHANNEL PortV1110LinkNetworkTechnology = "FibreChannel"
	PORTV1110LINKNETWORKTECHNOLOGY_GEN_Z         PortV1110LinkNetworkTechnology = "GenZ"
	PORTV1110LINKNETWORKTECHNOLOGY_PCIE          PortV1110LinkNetworkTechnology = "PCIe"
)

// AllowedPortV1110LinkNetworkTechnologyEnumValues is all the allowed values of PortV1110LinkNetworkTechnology enum
var AllowedPortV1110LinkNetworkTechnologyEnumValues = []PortV1110LinkNetworkTechnology{
	"Ethernet",
	"InfiniBand",
	"FibreChannel",
	"GenZ",
	"PCIe",
}

// validPortV1110LinkNetworkTechnologyEnumValue provides a map of PortV1110LinkNetworkTechnologys for fast verification of use input
var validPortV1110LinkNetworkTechnologyEnumValues = map[PortV1110LinkNetworkTechnology]struct{}{
	"Ethernet":     {},
	"InfiniBand":   {},
	"FibreChannel": {},
	"GenZ":         {},
	"PCIe":         {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortV1110LinkNetworkTechnology) IsValid() bool {
	_, ok := validPortV1110LinkNetworkTechnologyEnumValues[v]
	return ok
}

// NewPortV1110LinkNetworkTechnologyFromValue returns a pointer to a valid PortV1110LinkNetworkTechnology
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortV1110LinkNetworkTechnologyFromValue(v string) (PortV1110LinkNetworkTechnology, error) {
	ev := PortV1110LinkNetworkTechnology(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PortV1110LinkNetworkTechnology: valid values are %v", v, AllowedPortV1110LinkNetworkTechnologyEnumValues)
}

// AssertPortV1110LinkNetworkTechnologyRequired checks if the required fields are not zero-ed
func AssertPortV1110LinkNetworkTechnologyRequired(obj PortV1110LinkNetworkTechnology) error {
	return nil
}

// AssertPortV1110LinkNetworkTechnologyConstraints checks if the values respects the defined constraints
func AssertPortV1110LinkNetworkTechnologyConstraints(obj PortV1110LinkNetworkTechnology) error {
	return nil
}