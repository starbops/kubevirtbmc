// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AddressPoolV124Ipv4AddressRange - IPv4-related address range for an Ethernet fabric.
type AddressPoolV124Ipv4AddressRange struct {

	// Lower IPv4 network address.
	Lower *string `json:"Lower,omitempty" validate:"regexp=^(?:[0-9]{1,3}\\\\.){3}[0-9]{1,3}$"`

	// Upper IPv4 network address.
	Upper *string `json:"Upper,omitempty" validate:"regexp=^(?:[0-9]{1,3}\\\\.){3}[0-9]{1,3}$"`
}

// AssertAddressPoolV124Ipv4AddressRangeRequired checks if the required fields are not zero-ed
func AssertAddressPoolV124Ipv4AddressRangeRequired(obj AddressPoolV124Ipv4AddressRange) error {
	return nil
}

// AssertAddressPoolV124Ipv4AddressRangeConstraints checks if the values respects the defined constraints
func AssertAddressPoolV124Ipv4AddressRangeConstraints(obj AddressPoolV124Ipv4AddressRange) error {
	return nil
}