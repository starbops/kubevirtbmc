// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// NetworkAdapterV1100Sriov - Single-root input/output virtualization (SR-IOV) capabilities.
type NetworkAdapterV1100Sriov struct {

	// An indication of whether this controller supports single root input/output virtualization (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA) mode.
	SRIOVVEPACapable *bool `json:"SRIOVVEPACapable,omitempty"`
}

// AssertNetworkAdapterV1100SriovRequired checks if the required fields are not zero-ed
func AssertNetworkAdapterV1100SriovRequired(obj NetworkAdapterV1100Sriov) error {
	return nil
}

// AssertNetworkAdapterV1100SriovConstraints checks if the values respects the defined constraints
func AssertNetworkAdapterV1100SriovConstraints(obj NetworkAdapterV1100Sriov) error {
	return nil
}
