// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// TriggersV132Threshold - A threshold definition for a sensor.
type TriggersV132Threshold struct {
	Activation TriggersV132ThresholdActivation `json:"Activation,omitempty"`

	// The duration the sensor value must violate the threshold before the threshold is activated.
	DwellTime *string `json:"DwellTime,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The threshold value.
	Reading *float32 `json:"Reading,omitempty"`
}

// AssertTriggersV132ThresholdRequired checks if the required fields are not zero-ed
func AssertTriggersV132ThresholdRequired(obj TriggersV132Threshold) error {
	return nil
}

// AssertTriggersV132ThresholdConstraints checks if the values respects the defined constraints
func AssertTriggersV132ThresholdConstraints(obj TriggersV132Threshold) error {
	return nil
}