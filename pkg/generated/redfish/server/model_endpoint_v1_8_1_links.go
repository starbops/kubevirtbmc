// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EndpointV181Links - The links to other resources that are related to this resource.
type EndpointV181Links struct {

	// An array of links to the address pools associated with this endpoint.
	AddressPools []OdataV4IdRef `json:"AddressPools,omitempty"`

	// The number of items in a collection.
	AddressPoolsodataCount int64 `json:"AddressPools@odata.count,omitempty"`

	// An array of links to the switch ports or remote device ports at the other end of the link.
	ConnectedPorts []OdataV4IdRef `json:"ConnectedPorts,omitempty"`

	// The number of items in a collection.
	ConnectedPortsodataCount int64 `json:"ConnectedPorts@odata.count,omitempty"`

	// The connections to which this endpoint belongs.
	Connections []OdataV4IdRef `json:"Connections,omitempty"`

	// The number of items in a collection.
	ConnectionsodataCount int64 `json:"Connections@odata.count,omitempty"`

	// An array of links to the device ports that this endpoint represents.
	LocalPorts []OdataV4IdRef `json:"LocalPorts,omitempty"`

	// The number of items in a collection.
	LocalPortsodataCount int64 `json:"LocalPorts@odata.count,omitempty"`

	// An array of links to the endpoints that cannot be used in zones if this endpoint is in a zone.
	MutuallyExclusiveEndpoints []OdataV4IdRef `json:"MutuallyExclusiveEndpoints,omitempty"`

	// The number of items in a collection.
	MutuallyExclusiveEndpointsodataCount int64 `json:"MutuallyExclusiveEndpoints@odata.count,omitempty"`

	// When NetworkDeviceFunction resources are present, this array contains links to the network device functions that connect to this endpoint.
	NetworkDeviceFunction []OdataV4IdRef `json:"NetworkDeviceFunction,omitempty"`

	// The number of items in a collection.
	NetworkDeviceFunctionodataCount int64 `json:"NetworkDeviceFunction@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the physical ports associated with this endpoint.
	// Deprecated
	Ports []OdataV4IdRef `json:"Ports,omitempty"`

	// The number of items in a collection.
	PortsodataCount int64 `json:"Ports@odata.count,omitempty"`

	// The zones to which this endpoint belongs.
	Zones []OdataV4IdRef `json:"Zones,omitempty"`

	// The number of items in a collection.
	ZonesodataCount int64 `json:"Zones@odata.count,omitempty"`
}

// AssertEndpointV181LinksRequired checks if the required fields are not zero-ed
func AssertEndpointV181LinksRequired(obj EndpointV181Links) error {
	for _, el := range obj.AddressPools {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnectedPorts {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Connections {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.LocalPorts {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MutuallyExclusiveEndpoints {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NetworkDeviceFunction {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ports {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Zones {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEndpointV181LinksConstraints checks if the values respects the defined constraints
func AssertEndpointV181LinksConstraints(obj EndpointV181Links) error {
	for _, el := range obj.AddressPools {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConnectedPorts {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Connections {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.LocalPorts {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.MutuallyExclusiveEndpoints {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NetworkDeviceFunction {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ports {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Zones {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
