// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// ContainerV101ResetRequestBody - This action resets the container.
type ContainerV101ResetRequestBody struct {
	ResetType ResourceResetType `json:"ResetType,omitempty"`
}

// AssertContainerV101ResetRequestBodyRequired checks if the required fields are not zero-ed
func AssertContainerV101ResetRequestBodyRequired(obj ContainerV101ResetRequestBody) error {
	return nil
}

// AssertContainerV101ResetRequestBodyConstraints checks if the values respects the defined constraints
func AssertContainerV101ResetRequestBodyConstraints(obj ContainerV101ResetRequestBody) error {
	return nil
}