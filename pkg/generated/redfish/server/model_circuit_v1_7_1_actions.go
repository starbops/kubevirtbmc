// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CircuitV171Actions - The available actions for this resource.
type CircuitV171Actions struct {
	CircuitBreakerControl CircuitV171BreakerControl `json:"#Circuit.BreakerControl,omitempty"`

	CircuitPowerControl CircuitV171PowerControl `json:"#Circuit.PowerControl,omitempty"`

	CircuitResetMetrics CircuitV171ResetMetrics `json:"#Circuit.ResetMetrics,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCircuitV171ActionsRequired checks if the required fields are not zero-ed
func AssertCircuitV171ActionsRequired(obj CircuitV171Actions) error {
	if err := AssertCircuitV171BreakerControlRequired(obj.CircuitBreakerControl); err != nil {
		return err
	}
	if err := AssertCircuitV171PowerControlRequired(obj.CircuitPowerControl); err != nil {
		return err
	}
	if err := AssertCircuitV171ResetMetricsRequired(obj.CircuitResetMetrics); err != nil {
		return err
	}
	return nil
}

// AssertCircuitV171ActionsConstraints checks if the values respects the defined constraints
func AssertCircuitV171ActionsConstraints(obj CircuitV171Actions) error {
	if err := AssertCircuitV171BreakerControlConstraints(obj.CircuitBreakerControl); err != nil {
		return err
	}
	if err := AssertCircuitV171PowerControlConstraints(obj.CircuitPowerControl); err != nil {
		return err
	}
	if err := AssertCircuitV171ResetMetricsConstraints(obj.CircuitResetMetrics); err != nil {
		return err
	}
	return nil
}