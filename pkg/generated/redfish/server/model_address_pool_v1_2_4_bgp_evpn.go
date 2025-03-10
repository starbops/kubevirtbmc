// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AddressPoolV124BgpEvpn - BGP Ethernet Virtual Private Network (BGP EVPN) related properties for an Ethernet fabric.
type AddressPoolV124BgpEvpn struct {

	// Address Resolution Protocol (ARP) proxy status.
	ARPProxyEnabled *bool `json:"ARPProxyEnabled,omitempty"`

	// Address Resolution Protocol (ARP) suppression status.
	ARPSupressionEnabled *bool `json:"ARPSupressionEnabled,omitempty"`

	// The anycast gateway IPv4 address.
	AnycastGatewayIPAddress *string `json:"AnycastGatewayIPAddress,omitempty" validate:"regexp=^(?:[0-9]{1,3}\\\\.){3}[0-9]{1,3}$"`

	// The anycast gateway MAC address.
	AnycastGatewayMACAddress *string `json:"AnycastGatewayMACAddress,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`

	ESINumberRange AddressPoolV124EsiNumberRange `json:"ESINumberRange,omitempty"`

	EVINumberRange AddressPoolV124EviNumberRange `json:"EVINumberRange,omitempty"`

	// The gateway IPv4 address.
	GatewayIPAddress *string `json:"GatewayIPAddress,omitempty" validate:"regexp=^(?:[0-9]{1,3}\\\\.){3}[0-9]{1,3}$"`

	GatewayIPAddressRange AddressPoolV124GatewayIpAddressRange `json:"GatewayIPAddressRange,omitempty"`

	// Network Discovery Protocol (NDP) proxy status.
	NDPProxyEnabled *bool `json:"NDPProxyEnabled,omitempty"`

	// Network Discovery Protocol (NDP) suppression status.
	NDPSupressionEnabled *bool `json:"NDPSupressionEnabled,omitempty"`

	// The Route Distinguisher (RD) Administrator subfield.
	RouteDistinguisherAdministratorSubfield *string `json:"RouteDistinguisherAdministratorSubfield,omitempty"`

	RouteDistinguisherRange AddressPoolV124RouteDistinguisherRange `json:"RouteDistinguisherRange,omitempty"`

	// The Route Target (RT) Administrator Subfield.
	RouteTargetAdministratorSubfield *string `json:"RouteTargetAdministratorSubfield,omitempty"`

	RouteTargetRange AddressPoolV124RouteTargetRange `json:"RouteTargetRange,omitempty"`

	// Underlay multicast status.
	UnderlayMulticastEnabled *bool `json:"UnderlayMulticastEnabled,omitempty"`

	// Suppression of unknown unicast packets.
	UnknownUnicastSuppressionEnabled *bool `json:"UnknownUnicastSuppressionEnabled,omitempty"`

	VLANIdentifierAddressRange AddressPoolV124VlanIdentifierAddressRange `json:"VLANIdentifierAddressRange,omitempty"`
}

// AssertAddressPoolV124BgpEvpnRequired checks if the required fields are not zero-ed
func AssertAddressPoolV124BgpEvpnRequired(obj AddressPoolV124BgpEvpn) error {
	if err := AssertAddressPoolV124EsiNumberRangeRequired(obj.ESINumberRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124EviNumberRangeRequired(obj.EVINumberRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124GatewayIpAddressRangeRequired(obj.GatewayIPAddressRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124RouteDistinguisherRangeRequired(obj.RouteDistinguisherRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124RouteTargetRangeRequired(obj.RouteTargetRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124VlanIdentifierAddressRangeRequired(obj.VLANIdentifierAddressRange); err != nil {
		return err
	}
	return nil
}

// AssertAddressPoolV124BgpEvpnConstraints checks if the values respects the defined constraints
func AssertAddressPoolV124BgpEvpnConstraints(obj AddressPoolV124BgpEvpn) error {
	if err := AssertAddressPoolV124EsiNumberRangeConstraints(obj.ESINumberRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124EviNumberRangeConstraints(obj.EVINumberRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124GatewayIpAddressRangeConstraints(obj.GatewayIPAddressRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124RouteDistinguisherRangeConstraints(obj.RouteDistinguisherRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124RouteTargetRangeConstraints(obj.RouteTargetRange); err != nil {
		return err
	}
	if err := AssertAddressPoolV124VlanIdentifierAddressRangeConstraints(obj.VLANIdentifierAddressRange); err != nil {
		return err
	}
	return nil
}
