// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SwitchV192Actions - The available actions for this resource.
type SwitchV192Actions struct {
	SwitchReset SwitchV192Reset `json:"#Switch.Reset,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertSwitchV192ActionsRequired checks if the required fields are not zero-ed
func AssertSwitchV192ActionsRequired(obj SwitchV192Actions) error {
	if err := AssertSwitchV192ResetRequired(obj.SwitchReset); err != nil {
		return err
	}
	return nil
}

// AssertSwitchV192ActionsConstraints checks if the values respects the defined constraints
func AssertSwitchV192ActionsConstraints(obj SwitchV192Actions) error {
	if err := AssertSwitchV192ResetConstraints(obj.SwitchReset); err != nil {
		return err
	}
	return nil
}