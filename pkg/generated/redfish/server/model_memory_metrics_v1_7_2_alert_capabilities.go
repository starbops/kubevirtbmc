// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryMetricsV172AlertCapabilities - The conditions that would generate an alert to the CXL Fabric Manager or host.
type MemoryMetricsV172AlertCapabilities struct {

	// Indicates whether correctable ECC errors generate an alert to the CXL Fabric Manager or host.
	CorrectableECCError *bool `json:"CorrectableECCError,omitempty"`

	// Indicates whether spare block conditions generate an alert to the CXL Fabric Manager or host.
	SpareBlock *bool `json:"SpareBlock,omitempty"`

	// Indicates whether temperature conditions generate an alert to the CXL Fabric Manager or host.
	Temperature *bool `json:"Temperature,omitempty"`

	// Indicates whether uncorrectable ECC errors generate an alert to the CXL Fabric Manager or host.
	UncorrectableECCError *bool `json:"UncorrectableECCError,omitempty"`
}

// AssertMemoryMetricsV172AlertCapabilitiesRequired checks if the required fields are not zero-ed
func AssertMemoryMetricsV172AlertCapabilitiesRequired(obj MemoryMetricsV172AlertCapabilities) error {
	return nil
}

// AssertMemoryMetricsV172AlertCapabilitiesConstraints checks if the values respects the defined constraints
func AssertMemoryMetricsV172AlertCapabilitiesConstraints(obj MemoryMetricsV172AlertCapabilities) error {
	return nil
}
