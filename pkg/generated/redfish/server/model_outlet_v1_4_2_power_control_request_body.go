// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// OutletV142PowerControlRequestBody - This action turns the outlet on or off.
type OutletV142PowerControlRequestBody struct {
	PowerState CircuitPowerState `json:"PowerState,omitempty"`
}

// AssertOutletV142PowerControlRequestBodyRequired checks if the required fields are not zero-ed
func AssertOutletV142PowerControlRequestBodyRequired(obj OutletV142PowerControlRequestBody) error {
	return nil
}

// AssertOutletV142PowerControlRequestBodyConstraints checks if the values respects the defined constraints
func AssertOutletV142PowerControlRequestBodyConstraints(obj OutletV142PowerControlRequestBody) error {
	return nil
}