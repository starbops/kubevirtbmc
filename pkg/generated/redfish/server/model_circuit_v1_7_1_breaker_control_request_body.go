// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CircuitV171BreakerControlRequestBody - This action attempts to reset the circuit breaker.
type CircuitV171BreakerControlRequestBody struct {
	PowerState CircuitPowerState `json:"PowerState,omitempty"`
}

// AssertCircuitV171BreakerControlRequestBodyRequired checks if the required fields are not zero-ed
func AssertCircuitV171BreakerControlRequestBodyRequired(obj CircuitV171BreakerControlRequestBody) error {
	return nil
}

// AssertCircuitV171BreakerControlRequestBodyConstraints checks if the values respects the defined constraints
func AssertCircuitV171BreakerControlRequestBodyConstraints(obj CircuitV171BreakerControlRequestBody) error {
	return nil
}