// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EthernetInterfaceV1120EthernetInterface - The EthernetInterface schema represents a single, logical Ethernet interface or network interface controller (NIC).
type EthernetInterfaceV1120EthernetInterface struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions EthernetInterfaceV1120Actions `json:"Actions,omitempty"`

	// An indication of whether the speed and duplex are automatically negotiated and configured on this interface.
	AutoNeg *bool `json:"AutoNeg,omitempty"`

	DHCPv4 EthernetInterfaceV1120Dhcpv4Configuration `json:"DHCPv4,omitempty"`

	DHCPv6 EthernetInterfaceV1120Dhcpv6Configuration `json:"DHCPv6,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	EthernetInterfaceType EthernetInterfaceV1120EthernetDeviceType `json:"EthernetInterfaceType,omitempty"`

	// The complete, fully qualified domain name that DNS obtains for this interface.
	FQDN *string `json:"FQDN,omitempty"`

	// An indication of whether full-duplex mode is enabled on the Ethernet connection for this interface.
	FullDuplex *bool `json:"FullDuplex,omitempty"`

	// The DNS host name, without any domain information.
	HostName *string `json:"HostName,omitempty"`

	// The IPv4 addresses currently in use by this interface.
	IPv4Addresses []IpAddressesIpv4Address `json:"IPv4Addresses,omitempty"`

	// The IPv4 static addresses assigned to this interface.  See IPv4Addresses for the addresses in use by this interface.
	IPv4StaticAddresses []IpAddressesIpv4Address `json:"IPv4StaticAddresses,omitempty"`

	// An array that represents the RFC6724-defined address selection policy table.
	IPv6AddressPolicyTable []EthernetInterfaceV1120Ipv6AddressPolicyEntry `json:"IPv6AddressPolicyTable,omitempty"`

	// The IPv6 addresses currently in use by this interface.
	IPv6Addresses []IpAddressesIpv6Address `json:"IPv6Addresses,omitempty"`

	// The IPv6 default gateway address in use on this interface.
	IPv6DefaultGateway *string `json:"IPv6DefaultGateway,omitempty"`

	// An indication of whether IPv6 is enabled on this interface.
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty"`

	// The IPv6 static addresses assigned to this interface.  See IPv6Addresses for the addresses in use by this interface.
	IPv6StaticAddresses []IpAddressesIpv6StaticAddress `json:"IPv6StaticAddresses,omitempty"`

	// The IPv6 static default gateways for this interface.
	IPv6StaticDefaultGateways []IpAddressesIpv6GatewayStaticAddress `json:"IPv6StaticDefaultGateways,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// An indication of whether this interface is enabled.
	InterfaceEnabled *bool `json:"InterfaceEnabled,omitempty"`

	LinkStatus EthernetInterfaceV1120LinkStatus `json:"LinkStatus,omitempty"`

	Links EthernetInterfaceV1120Links `json:"Links,omitempty"`

	MACAddress string `json:"MACAddress,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`

	// The currently configured maximum transmission unit (MTU), in bytes, on this interface.
	MTUSize *int64 `json:"MTUSize,omitempty"`

	// The maximum number of static IPv6 addresses that can be configured on this interface.
	MaxIPv6StaticAddresses *int64 `json:"MaxIPv6StaticAddresses,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The DNS servers in use on this interface.
	NameServers []string `json:"NameServers,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PermanentMACAddress string `json:"PermanentMACAddress,omitempty" validate:"regexp=^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`

	RoutingScope EthernetInterfaceV1120RoutingScope `json:"RoutingScope,omitempty"`

	// The current speed, in Mbit/s, of this interface.
	SpeedMbps *int64 `json:"SpeedMbps,omitempty"`

	StatelessAddressAutoConfig EthernetInterfaceV1120StatelessAddressAutoConfiguration `json:"StatelessAddressAutoConfig,omitempty"`

	// The statically-defined set of DNS server IPv4 and IPv6 addresses.
	StaticNameServers []*string `json:"StaticNameServers,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	TeamMode EthernetInterfaceV1120TeamMode `json:"TeamMode,omitempty"`

	// The UEFI device path for this interface.
	UefiDevicePath *string `json:"UefiDevicePath,omitempty"`

	VLAN VLanNetworkInterfaceVlan `json:"VLAN,omitempty"`

	VLANs OdataV4IdRef `json:"VLANs,omitempty"`
}

// AssertEthernetInterfaceV1120EthernetInterfaceRequired checks if the required fields are not zero-ed
func AssertEthernetInterfaceV1120EthernetInterfaceRequired(obj EthernetInterfaceV1120EthernetInterface) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertEthernetInterfaceV1120ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertEthernetInterfaceV1120Dhcpv4ConfigurationRequired(obj.DHCPv4); err != nil {
		return err
	}
	if err := AssertEthernetInterfaceV1120Dhcpv6ConfigurationRequired(obj.DHCPv6); err != nil {
		return err
	}
	for _, el := range obj.IPv4Addresses {
		if err := AssertIpAddressesIpv4AddressRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv4StaticAddresses {
		if err := AssertIpAddressesIpv4AddressRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6AddressPolicyTable {
		if err := AssertEthernetInterfaceV1120Ipv6AddressPolicyEntryRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6Addresses {
		if err := AssertIpAddressesIpv6AddressRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6StaticAddresses {
		if err := AssertIpAddressesIpv6StaticAddressRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6StaticDefaultGateways {
		if err := AssertIpAddressesIpv6GatewayStaticAddressRequired(el); err != nil {
			return err
		}
	}
	if err := AssertEthernetInterfaceV1120LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertEthernetInterfaceV1120StatelessAddressAutoConfigurationRequired(obj.StatelessAddressAutoConfig); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertVLanNetworkInterfaceVlanRequired(obj.VLAN); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.VLANs); err != nil {
		return err
	}
	return nil
}

// AssertEthernetInterfaceV1120EthernetInterfaceConstraints checks if the values respects the defined constraints
func AssertEthernetInterfaceV1120EthernetInterfaceConstraints(obj EthernetInterfaceV1120EthernetInterface) error {
	if err := AssertEthernetInterfaceV1120ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertEthernetInterfaceV1120Dhcpv4ConfigurationConstraints(obj.DHCPv4); err != nil {
		return err
	}
	if err := AssertEthernetInterfaceV1120Dhcpv6ConfigurationConstraints(obj.DHCPv6); err != nil {
		return err
	}
	for _, el := range obj.IPv4Addresses {
		if err := AssertIpAddressesIpv4AddressConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv4StaticAddresses {
		if err := AssertIpAddressesIpv4AddressConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6AddressPolicyTable {
		if err := AssertEthernetInterfaceV1120Ipv6AddressPolicyEntryConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6Addresses {
		if err := AssertIpAddressesIpv6AddressConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6StaticAddresses {
		if err := AssertIpAddressesIpv6StaticAddressConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.IPv6StaticDefaultGateways {
		if err := AssertIpAddressesIpv6GatewayStaticAddressConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertEthernetInterfaceV1120LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertEthernetInterfaceV1120StatelessAddressAutoConfigurationConstraints(obj.StatelessAddressAutoConfig); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	if err := AssertVLanNetworkInterfaceVlanConstraints(obj.VLAN); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.VLANs); err != nil {
		return err
	}
	return nil
}