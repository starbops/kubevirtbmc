// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ContainerV100Actions - The available actions for this resource.
type ContainerV100Actions struct {
	ContainerReset ContainerV100Reset `json:"#Container.Reset,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertContainerV100ActionsRequired checks if the required fields are not zero-ed
func AssertContainerV100ActionsRequired(obj ContainerV100Actions) error {
	if err := AssertContainerV100ResetRequired(obj.ContainerReset); err != nil {
		return err
	}
	return nil
}

// AssertContainerV100ActionsConstraints checks if the values respects the defined constraints
func AssertContainerV100ActionsConstraints(obj ContainerV100Actions) error {
	if err := AssertContainerV100ResetConstraints(obj.ContainerReset); err != nil {
		return err
	}
	return nil
}