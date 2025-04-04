// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// NetworkAdapterV1100ResetSettingsToDefault - This action is to clear the settings back to factory defaults.
type NetworkAdapterV1100ResetSettingsToDefault struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertNetworkAdapterV1100ResetSettingsToDefaultRequired checks if the required fields are not zero-ed
func AssertNetworkAdapterV1100ResetSettingsToDefaultRequired(obj NetworkAdapterV1100ResetSettingsToDefault) error {
	return nil
}

// AssertNetworkAdapterV1100ResetSettingsToDefaultConstraints checks if the values respects the defined constraints
func AssertNetworkAdapterV1100ResetSettingsToDefaultConstraints(obj NetworkAdapterV1100ResetSettingsToDefault) error {
	return nil
}
