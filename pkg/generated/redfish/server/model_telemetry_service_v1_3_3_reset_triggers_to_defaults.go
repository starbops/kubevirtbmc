// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// TelemetryServiceV133ResetTriggersToDefaults - The action to reset the triggers to factory defaults.
type TelemetryServiceV133ResetTriggersToDefaults struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertTelemetryServiceV133ResetTriggersToDefaultsRequired checks if the required fields are not zero-ed
func AssertTelemetryServiceV133ResetTriggersToDefaultsRequired(obj TelemetryServiceV133ResetTriggersToDefaults) error {
	return nil
}

// AssertTelemetryServiceV133ResetTriggersToDefaultsConstraints checks if the values respects the defined constraints
func AssertTelemetryServiceV133ResetTriggersToDefaultsConstraints(obj TelemetryServiceV133ResetTriggersToDefaults) error {
	return nil
}