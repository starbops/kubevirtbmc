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

// PortV1110LldpTransmit - Link Layer Data Protocol (LLDP) data being transmitted on this link.
type PortV1110LldpTransmit struct {

	// Link Layer Data Protocol (LLDP) chassis ID.
	ChassisId *string `json:"ChassisId,omitempty"`

	ChassisIdSubtype PortV1110Ieee802IdSubtype `json:"ChassisIdSubtype,omitempty"`

	// The IPv4 management address to be transmitted from this endpoint.
	ManagementAddressIPv4 *string `json:"ManagementAddressIPv4,omitempty" validate:"regexp=(^(?:[0-9]{1,3}\\\\.){3}[0-9]{1,3}$)|(^$)"`

	// The IPv6 management address to be transmitted from this endpoint.
	ManagementAddressIPv6 *string `json:"ManagementAddressIPv6,omitempty"`

	// The management MAC address to be transmitted from this endpoint.
	ManagementAddressMAC *string `json:"ManagementAddressMAC,omitempty" validate:"regexp=(^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$)|(^$)"`

	// The management VLAN ID to be transmitted from this endpoint.
	ManagementVlanId *int64 `json:"ManagementVlanId,omitempty"`

	// A colon-delimited string of hexadecimal octets identifying a port to be transmitted from this endpoint.
	PortId *string `json:"PortId,omitempty" validate:"regexp=(^([0-9A-F]{2})([:]([0-9A-F]){2}){0,63}$)|(^$)"`

	PortIdSubtype PortV1110Ieee802IdSubtype `json:"PortIdSubtype,omitempty"`

	// The system capabilities to be transmitted from this endpoint.
	SystemCapabilities []PortV1110LldpSystemCapabilities `json:"SystemCapabilities,omitempty"`

	// The system description to be transmitted from this endpoint.
	SystemDescription *string `json:"SystemDescription,omitempty"`

	// The system name to be transmitted from this endpoint.
	SystemName *string `json:"SystemName,omitempty"`
}

// AssertPortV1110LldpTransmitRequired checks if the required fields are not zero-ed
func AssertPortV1110LldpTransmitRequired(obj PortV1110LldpTransmit) error {
	return nil
}

// AssertPortV1110LldpTransmitConstraints checks if the values respects the defined constraints
func AssertPortV1110LldpTransmitConstraints(obj PortV1110LldpTransmit) error {
	if obj.ManagementVlanId != nil && *obj.ManagementVlanId < 0 {
		return &ParsingError{Param: "ManagementVlanId", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.ManagementVlanId != nil && *obj.ManagementVlanId > 4095 {
		return &ParsingError{Param: "ManagementVlanId", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}