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
	"errors"
)

// ManagerNetworkProtocolV1100SnmpProtocol - The settings for a network protocol associated with a manager.
type ManagerNetworkProtocolV1100SnmpProtocol struct {
	AuthenticationProtocol ManagerNetworkProtocolV1100SnmpAuthenticationProtocols `json:"AuthenticationProtocol,omitempty"`

	CommunityAccessMode ManagerNetworkProtocolV1100SnmpCommunityAccessMode `json:"CommunityAccessMode,omitempty"`

	// The SNMP community strings.
	CommunityStrings []ManagerNetworkProtocolV1100SnmpCommunity `json:"CommunityStrings,omitempty"`

	// Indicates if access via SNMPv1 is enabled.
	EnableSNMPv1 *bool `json:"EnableSNMPv1,omitempty"`

	// Indicates if access via SNMPv2c is enabled.
	EnableSNMPv2c *bool `json:"EnableSNMPv2c,omitempty"`

	// Indicates if access via SNMPv3 is enabled.
	EnableSNMPv3 *bool `json:"EnableSNMPv3,omitempty"`

	EncryptionProtocol ManagerNetworkProtocolV1100SnmpEncryptionProtocols `json:"EncryptionProtocol,omitempty"`

	EngineId ManagerNetworkProtocolV1100EngineId `json:"EngineId,omitempty"`

	// Indicates if the community strings should be hidden.
	HideCommunityStrings *bool `json:"HideCommunityStrings,omitempty"`

	// The protocol port.
	Port *int64 `json:"Port,omitempty"`

	// An indication of whether the protocol is enabled.
	ProtocolEnabled *bool `json:"ProtocolEnabled,omitempty"`

	// The SNMP trap port.
	TrapPort *int64 `json:"TrapPort,omitempty"`
}

// AssertManagerNetworkProtocolV1100SnmpProtocolRequired checks if the required fields are not zero-ed
func AssertManagerNetworkProtocolV1100SnmpProtocolRequired(obj ManagerNetworkProtocolV1100SnmpProtocol) error {
	for _, el := range obj.CommunityStrings {
		if err := AssertManagerNetworkProtocolV1100SnmpCommunityRequired(el); err != nil {
			return err
		}
	}
	if err := AssertManagerNetworkProtocolV1100EngineIdRequired(obj.EngineId); err != nil {
		return err
	}
	return nil
}

// AssertManagerNetworkProtocolV1100SnmpProtocolConstraints checks if the values respects the defined constraints
func AssertManagerNetworkProtocolV1100SnmpProtocolConstraints(obj ManagerNetworkProtocolV1100SnmpProtocol) error {
	for _, el := range obj.CommunityStrings {
		if err := AssertManagerNetworkProtocolV1100SnmpCommunityConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertManagerNetworkProtocolV1100EngineIdConstraints(obj.EngineId); err != nil {
		return err
	}
	if obj.Port != nil && *obj.Port < 0 {
		return &ParsingError{Param: "Port", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.TrapPort != nil && *obj.TrapPort < 0 {
		return &ParsingError{Param: "TrapPort", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}