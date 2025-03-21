// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PowerV172PowerSupplyReset - This action resets the targeted power supply.
type PowerV172PowerSupplyReset struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertPowerV172PowerSupplyResetRequired checks if the required fields are not zero-ed
func AssertPowerV172PowerSupplyResetRequired(obj PowerV172PowerSupplyReset) error {
	return nil
}

// AssertPowerV172PowerSupplyResetConstraints checks if the values respects the defined constraints
func AssertPowerV172PowerSupplyResetConstraints(obj PowerV172PowerSupplyReset) error {
	return nil
}
