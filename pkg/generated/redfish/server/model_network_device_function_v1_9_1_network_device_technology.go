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

type NetworkDeviceFunctionV191NetworkDeviceTechnology string

// List of NetworkDeviceFunctionV191NetworkDeviceTechnology
const (
	NETWORKDEVICEFUNCTIONV191NETWORKDEVICETECHNOLOGY_DISABLED                    NetworkDeviceFunctionV191NetworkDeviceTechnology = "Disabled"
	NETWORKDEVICEFUNCTIONV191NETWORKDEVICETECHNOLOGY_ETHERNET                    NetworkDeviceFunctionV191NetworkDeviceTechnology = "Ethernet"
	NETWORKDEVICEFUNCTIONV191NETWORKDEVICETECHNOLOGY_FIBRE_CHANNEL               NetworkDeviceFunctionV191NetworkDeviceTechnology = "FibreChannel"
	NETWORKDEVICEFUNCTIONV191NETWORKDEVICETECHNOLOGY_I_SCSI                      NetworkDeviceFunctionV191NetworkDeviceTechnology = "iSCSI"
	NETWORKDEVICEFUNCTIONV191NETWORKDEVICETECHNOLOGY_FIBRE_CHANNEL_OVER_ETHERNET NetworkDeviceFunctionV191NetworkDeviceTechnology = "FibreChannelOverEthernet"
	NETWORKDEVICEFUNCTIONV191NETWORKDEVICETECHNOLOGY_INFINI_BAND                 NetworkDeviceFunctionV191NetworkDeviceTechnology = "InfiniBand"
)

// AllowedNetworkDeviceFunctionV191NetworkDeviceTechnologyEnumValues is all the allowed values of NetworkDeviceFunctionV191NetworkDeviceTechnology enum
var AllowedNetworkDeviceFunctionV191NetworkDeviceTechnologyEnumValues = []NetworkDeviceFunctionV191NetworkDeviceTechnology{
	"Disabled",
	"Ethernet",
	"FibreChannel",
	"iSCSI",
	"FibreChannelOverEthernet",
	"InfiniBand",
}

// validNetworkDeviceFunctionV191NetworkDeviceTechnologyEnumValue provides a map of NetworkDeviceFunctionV191NetworkDeviceTechnologys for fast verification of use input
var validNetworkDeviceFunctionV191NetworkDeviceTechnologyEnumValues = map[NetworkDeviceFunctionV191NetworkDeviceTechnology]struct{}{
	"Disabled":                 {},
	"Ethernet":                 {},
	"FibreChannel":             {},
	"iSCSI":                    {},
	"FibreChannelOverEthernet": {},
	"InfiniBand":               {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkDeviceFunctionV191NetworkDeviceTechnology) IsValid() bool {
	_, ok := validNetworkDeviceFunctionV191NetworkDeviceTechnologyEnumValues[v]
	return ok
}

// NewNetworkDeviceFunctionV191NetworkDeviceTechnologyFromValue returns a pointer to a valid NetworkDeviceFunctionV191NetworkDeviceTechnology
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkDeviceFunctionV191NetworkDeviceTechnologyFromValue(v string) (NetworkDeviceFunctionV191NetworkDeviceTechnology, error) {
	ev := NetworkDeviceFunctionV191NetworkDeviceTechnology(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for NetworkDeviceFunctionV191NetworkDeviceTechnology: valid values are %v", v, AllowedNetworkDeviceFunctionV191NetworkDeviceTechnologyEnumValues)
}

// AssertNetworkDeviceFunctionV191NetworkDeviceTechnologyRequired checks if the required fields are not zero-ed
func AssertNetworkDeviceFunctionV191NetworkDeviceTechnologyRequired(obj NetworkDeviceFunctionV191NetworkDeviceTechnology) error {
	return nil
}

// AssertNetworkDeviceFunctionV191NetworkDeviceTechnologyConstraints checks if the values respects the defined constraints
func AssertNetworkDeviceFunctionV191NetworkDeviceTechnologyConstraints(obj NetworkDeviceFunctionV191NetworkDeviceTechnology) error {
	return nil
}