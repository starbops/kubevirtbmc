// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AddressPoolV124RouteDistinguisherRange - The Route Distinguisher (RD) number range for an Ethernet fabric.
type AddressPoolV124RouteDistinguisherRange struct {

	// Lower Route Distinguisher (RD) number.
	Lower int64 `json:"Lower,omitempty"`

	// Upper Route Distinguisher (RD) number.
	Upper int64 `json:"Upper,omitempty"`
}

// AssertAddressPoolV124RouteDistinguisherRangeRequired checks if the required fields are not zero-ed
func AssertAddressPoolV124RouteDistinguisherRangeRequired(obj AddressPoolV124RouteDistinguisherRange) error {
	return nil
}

// AssertAddressPoolV124RouteDistinguisherRangeConstraints checks if the values respects the defined constraints
func AssertAddressPoolV124RouteDistinguisherRangeConstraints(obj AddressPoolV124RouteDistinguisherRange) error {
	return nil
}