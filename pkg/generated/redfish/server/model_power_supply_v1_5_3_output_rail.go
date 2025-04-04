// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PowerSupplyV153OutputRail - This type describes an output power rail of a power supply.
type PowerSupplyV153OutputRail struct {

	// The nominal voltage of this output power rail.
	NominalVoltage *float32 `json:"NominalVoltage,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`
}

// AssertPowerSupplyV153OutputRailRequired checks if the required fields are not zero-ed
func AssertPowerSupplyV153OutputRailRequired(obj PowerSupplyV153OutputRail) error {
	return nil
}

// AssertPowerSupplyV153OutputRailConstraints checks if the values respects the defined constraints
func AssertPowerSupplyV153OutputRailConstraints(obj PowerSupplyV153OutputRail) error {
	return nil
}
