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

type ManagerNetworkProtocolV1100SnmpEncryptionProtocols string

// List of ManagerNetworkProtocolV1100SnmpEncryptionProtocols
const (
	MANAGERNETWORKPROTOCOLV1100SNMPENCRYPTIONPROTOCOLS_NONE          ManagerNetworkProtocolV1100SnmpEncryptionProtocols = "None"
	MANAGERNETWORKPROTOCOLV1100SNMPENCRYPTIONPROTOCOLS_ACCOUNT       ManagerNetworkProtocolV1100SnmpEncryptionProtocols = "Account"
	MANAGERNETWORKPROTOCOLV1100SNMPENCRYPTIONPROTOCOLS_CBC_DES       ManagerNetworkProtocolV1100SnmpEncryptionProtocols = "CBC_DES"
	MANAGERNETWORKPROTOCOLV1100SNMPENCRYPTIONPROTOCOLS_CFB128_AES128 ManagerNetworkProtocolV1100SnmpEncryptionProtocols = "CFB128_AES128"
	MANAGERNETWORKPROTOCOLV1100SNMPENCRYPTIONPROTOCOLS_CFB128_AES192 ManagerNetworkProtocolV1100SnmpEncryptionProtocols = "CFB128_AES192"
	MANAGERNETWORKPROTOCOLV1100SNMPENCRYPTIONPROTOCOLS_CFB128_AES256 ManagerNetworkProtocolV1100SnmpEncryptionProtocols = "CFB128_AES256"
)

// AllowedManagerNetworkProtocolV1100SnmpEncryptionProtocolsEnumValues is all the allowed values of ManagerNetworkProtocolV1100SnmpEncryptionProtocols enum
var AllowedManagerNetworkProtocolV1100SnmpEncryptionProtocolsEnumValues = []ManagerNetworkProtocolV1100SnmpEncryptionProtocols{
	"None",
	"Account",
	"CBC_DES",
	"CFB128_AES128",
	"CFB128_AES192",
	"CFB128_AES256",
}

// validManagerNetworkProtocolV1100SnmpEncryptionProtocolsEnumValue provides a map of ManagerNetworkProtocolV1100SnmpEncryptionProtocolss for fast verification of use input
var validManagerNetworkProtocolV1100SnmpEncryptionProtocolsEnumValues = map[ManagerNetworkProtocolV1100SnmpEncryptionProtocols]struct{}{
	"None":          {},
	"Account":       {},
	"CBC_DES":       {},
	"CFB128_AES128": {},
	"CFB128_AES192": {},
	"CFB128_AES256": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManagerNetworkProtocolV1100SnmpEncryptionProtocols) IsValid() bool {
	_, ok := validManagerNetworkProtocolV1100SnmpEncryptionProtocolsEnumValues[v]
	return ok
}

// NewManagerNetworkProtocolV1100SnmpEncryptionProtocolsFromValue returns a pointer to a valid ManagerNetworkProtocolV1100SnmpEncryptionProtocols
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManagerNetworkProtocolV1100SnmpEncryptionProtocolsFromValue(v string) (ManagerNetworkProtocolV1100SnmpEncryptionProtocols, error) {
	ev := ManagerNetworkProtocolV1100SnmpEncryptionProtocols(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ManagerNetworkProtocolV1100SnmpEncryptionProtocols: valid values are %v", v, AllowedManagerNetworkProtocolV1100SnmpEncryptionProtocolsEnumValues)
}

// AssertManagerNetworkProtocolV1100SnmpEncryptionProtocolsRequired checks if the required fields are not zero-ed
func AssertManagerNetworkProtocolV1100SnmpEncryptionProtocolsRequired(obj ManagerNetworkProtocolV1100SnmpEncryptionProtocols) error {
	return nil
}

// AssertManagerNetworkProtocolV1100SnmpEncryptionProtocolsConstraints checks if the values respects the defined constraints
func AssertManagerNetworkProtocolV1100SnmpEncryptionProtocolsConstraints(obj ManagerNetworkProtocolV1100SnmpEncryptionProtocols) error {
	return nil
}