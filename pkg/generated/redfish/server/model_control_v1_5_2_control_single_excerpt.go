// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ControlV152ControlSingleExcerpt - The `Control` schema describes a control point and its properties.
type ControlV152ControlSingleExcerpt struct {

	// The maximum possible setting for this control.
	AllowableMax *float32 `json:"AllowableMax,omitempty"`

	// The minimum possible setting for this control.
	AllowableMin *float32 `json:"AllowableMin,omitempty"`

	ControlMode *ControlV152ControlMode `json:"ControlMode,omitempty"`

	// The link to the resource that provides the data for this control.
	DataSourceUri *string `json:"DataSourceUri,omitempty"`

	// The default set point of the control.
	DefaultSetPoint *float32 `json:"DefaultSetPoint,omitempty"`

	// The reading of the sensor associated with this control.
	Reading *float32 `json:"Reading,omitempty"`

	// The units of the sensor reading associated with this control.
	ReadingUnits *string `json:"ReadingUnits,omitempty"`

	// The desired set point of the control.
	SetPoint *float32 `json:"SetPoint,omitempty"`
}

// AssertControlV152ControlSingleExcerptRequired checks if the required fields are not zero-ed
func AssertControlV152ControlSingleExcerptRequired(obj ControlV152ControlSingleExcerpt) error {
	return nil
}

// AssertControlV152ControlSingleExcerptConstraints checks if the values respects the defined constraints
func AssertControlV152ControlSingleExcerptConstraints(obj ControlV152ControlSingleExcerpt) error {
	return nil
}