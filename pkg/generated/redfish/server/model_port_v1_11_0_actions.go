// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PortV1110Actions - The available actions for this resource.
type PortV1110Actions struct {
	PortReset PortV1110Reset `json:"#Port.Reset,omitempty"`

	PortResetPPB PortV1110ResetPpb `json:"#Port.ResetPPB,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertPortV1110ActionsRequired checks if the required fields are not zero-ed
func AssertPortV1110ActionsRequired(obj PortV1110Actions) error {
	if err := AssertPortV1110ResetRequired(obj.PortReset); err != nil {
		return err
	}
	if err := AssertPortV1110ResetPpbRequired(obj.PortResetPPB); err != nil {
		return err
	}
	return nil
}

// AssertPortV1110ActionsConstraints checks if the values respects the defined constraints
func AssertPortV1110ActionsConstraints(obj PortV1110Actions) error {
	if err := AssertPortV1110ResetConstraints(obj.PortReset); err != nil {
		return err
	}
	if err := AssertPortV1110ResetPpbConstraints(obj.PortResetPPB); err != nil {
		return err
	}
	return nil
}