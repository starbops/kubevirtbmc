// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PortV1110FibreChannelProperties - Fibre Channel-specific properties for a port.
type PortV1110FibreChannelProperties struct {

	// An array of configured World Wide Names (WWN) that are associated with this network port, including the programmed address of the lowest-numbered network device function, the configured but not active address, if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedWorldWideNames []*string `json:"AssociatedWorldWideNames,omitempty"`

	// The Fibre Channel Fabric Name provided by the switch.
	FabricName *string `json:"FabricName,omitempty"`

	// The number of ports not on the associated device that the associated device has discovered through this port.
	NumberDiscoveredRemotePorts *int64 `json:"NumberDiscoveredRemotePorts,omitempty"`

	PortConnectionType PortV1110PortConnectionType `json:"PortConnectionType,omitempty"`
}

// AssertPortV1110FibreChannelPropertiesRequired checks if the required fields are not zero-ed
func AssertPortV1110FibreChannelPropertiesRequired(obj PortV1110FibreChannelProperties) error {
	return nil
}

// AssertPortV1110FibreChannelPropertiesConstraints checks if the values respects the defined constraints
func AssertPortV1110FibreChannelPropertiesConstraints(obj PortV1110FibreChannelProperties) error {
	return nil
}