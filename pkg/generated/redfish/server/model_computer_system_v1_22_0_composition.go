// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ComputerSystemV1220Composition - Information about the composition capabilities and state of a computer system.
type ComputerSystemV1220Composition struct {

	// The composition use cases in which this computer system can participate.
	UseCases []ComputerSystemV1220CompositionUseCase `json:"UseCases,omitempty"`
}

// AssertComputerSystemV1220CompositionRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220CompositionRequired(obj ComputerSystemV1220Composition) error {
	return nil
}

// AssertComputerSystemV1220CompositionConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220CompositionConstraints(obj ComputerSystemV1220Composition) error {
	return nil
}