// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PortV1110LinkConfiguration - Properties of the link for which this port is configured.
type PortV1110LinkConfiguration struct {

	// An indication of whether the port is capable of autonegotiating speed.
	AutoSpeedNegotiationCapable *bool `json:"AutoSpeedNegotiationCapable,omitempty"`

	// Controls whether this port is configured to enable autonegotiating speed.
	AutoSpeedNegotiationEnabled *bool `json:"AutoSpeedNegotiationEnabled,omitempty"`

	// The set of link speed capabilities of this port.
	CapableLinkSpeedGbps []*float32 `json:"CapableLinkSpeedGbps,omitempty"`

	// The set of link speed and width pairs this port is configured to use for autonegotiation.
	ConfiguredNetworkLinks []PortV1110ConfiguredNetworkLink `json:"ConfiguredNetworkLinks,omitempty"`
}

// AssertPortV1110LinkConfigurationRequired checks if the required fields are not zero-ed
func AssertPortV1110LinkConfigurationRequired(obj PortV1110LinkConfiguration) error {
	for _, el := range obj.ConfiguredNetworkLinks {
		if err := AssertPortV1110ConfiguredNetworkLinkRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPortV1110LinkConfigurationConstraints checks if the values respects the defined constraints
func AssertPortV1110LinkConfigurationConstraints(obj PortV1110LinkConfiguration) error {
	for _, el := range obj.ConfiguredNetworkLinks {
		if err := AssertPortV1110ConfiguredNetworkLinkConstraints(el); err != nil {
			return err
		}
	}
	return nil
}