// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EthernetInterfaceV1120Links - The links to other resources that are related to this resource.
type EthernetInterfaceV1120Links struct {

	// The links to the Ethernet interfaces that are affiliated with this interface, such as a VLAN or a team that uses this interface.
	AffiliatedInterfaces []OdataV4IdRef `json:"AffiliatedInterfaces,omitempty"`

	// The number of items in a collection.
	AffiliatedInterfacesodataCount int64 `json:"AffiliatedInterfaces@odata.count,omitempty"`

	Chassis OdataV4IdRef `json:"Chassis,omitempty"`

	// An array of links to the endpoints that connect to this Ethernet interface.
	Endpoints []OdataV4IdRef `json:"Endpoints,omitempty"`

	// The number of items in a collection.
	EndpointsodataCount int64 `json:"Endpoints@odata.count,omitempty"`

	HostInterface OdataV4IdRef `json:"HostInterface,omitempty"`

	NetworkDeviceFunction OdataV4IdRef `json:"NetworkDeviceFunction,omitempty"`

	// The link to the network device functions that constitute this Ethernet interface.
	NetworkDeviceFunctions []OdataV4IdRef `json:"NetworkDeviceFunctions,omitempty"`

	// The number of items in a collection.
	NetworkDeviceFunctionsodataCount int64 `json:"NetworkDeviceFunctions@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The links to the ports providing this Ethernet interface.
	Ports []OdataV4IdRef `json:"Ports,omitempty"`

	// The number of items in a collection.
	PortsodataCount int64 `json:"Ports@odata.count,omitempty"`

	// The links to the Ethernet interfaces that constitute this Ethernet interface.
	RelatedInterfaces []OdataV4IdRef `json:"RelatedInterfaces,omitempty"`

	// The number of items in a collection.
	RelatedInterfacesodataCount int64 `json:"RelatedInterfaces@odata.count,omitempty"`
}

// AssertEthernetInterfaceV1120LinksRequired checks if the required fields are not zero-ed
func AssertEthernetInterfaceV1120LinksRequired(obj EthernetInterfaceV1120Links) error {
	for _, el := range obj.AffiliatedInterfaces {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.Chassis); err != nil {
		return err
	}
	for _, el := range obj.Endpoints {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.HostInterface); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.NetworkDeviceFunction); err != nil {
		return err
	}
	for _, el := range obj.NetworkDeviceFunctions {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ports {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RelatedInterfaces {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEthernetInterfaceV1120LinksConstraints checks if the values respects the defined constraints
func AssertEthernetInterfaceV1120LinksConstraints(obj EthernetInterfaceV1120Links) error {
	for _, el := range obj.AffiliatedInterfaces {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefConstraints(obj.Chassis); err != nil {
		return err
	}
	for _, el := range obj.Endpoints {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefConstraints(obj.HostInterface); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.NetworkDeviceFunction); err != nil {
		return err
	}
	for _, el := range obj.NetworkDeviceFunctions {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ports {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RelatedInterfaces {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}