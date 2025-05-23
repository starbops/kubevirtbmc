// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ChassisV1250ResetRequestBody - This action resets the chassis.  Additionally, it could reset systems or other contained resources depending on the ResetType used to invoke this action.
type ChassisV1250ResetRequestBody struct {
	ResetType ResourceResetType `json:"ResetType,omitempty"`
}

// AssertChassisV1250ResetRequestBodyRequired checks if the required fields are not zero-ed
func AssertChassisV1250ResetRequestBodyRequired(obj ChassisV1250ResetRequestBody) error {
	return nil
}

// AssertChassisV1250ResetRequestBodyConstraints checks if the values respects the defined constraints
func AssertChassisV1250ResetRequestBodyConstraints(obj ChassisV1250ResetRequestBody) error {
	return nil
}
