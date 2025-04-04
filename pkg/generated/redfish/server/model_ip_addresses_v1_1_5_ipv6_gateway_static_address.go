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

// IpAddressesV115Ipv6GatewayStaticAddress - This type represents a single IPv6 static address to be assigned on a network interface.
type IpAddressesV115Ipv6GatewayStaticAddress struct {

	// A valid IPv6 address.
	Address *string `json:"Address"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PrefixLength int64 `json:"PrefixLength,omitempty"`
}

// AssertIpAddressesV115Ipv6GatewayStaticAddressRequired checks if the required fields are not zero-ed
func AssertIpAddressesV115Ipv6GatewayStaticAddressRequired(obj IpAddressesV115Ipv6GatewayStaticAddress) error {
	elements := map[string]interface{}{
		"Address": obj.Address,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertIpAddressesV115Ipv6GatewayStaticAddressConstraints checks if the values respects the defined constraints
func AssertIpAddressesV115Ipv6GatewayStaticAddressConstraints(obj IpAddressesV115Ipv6GatewayStaticAddress) error {
	if obj.PrefixLength < 0 {
		return &ParsingError{Param: "PrefixLength", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PrefixLength > 128 {
		return &ParsingError{Param: "PrefixLength", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
