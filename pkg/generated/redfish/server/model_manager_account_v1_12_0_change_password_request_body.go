// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ManagerAccountV1120ChangePasswordRequestBody - This action changes the account password.
type ManagerAccountV1120ChangePasswordRequestBody struct {

	// The new account password.
	NewPassword string `json:"NewPassword"`

	// The password of the account tied to the current session.
	SessionAccountPassword string `json:"SessionAccountPassword"`
}

// AssertManagerAccountV1120ChangePasswordRequestBodyRequired checks if the required fields are not zero-ed
func AssertManagerAccountV1120ChangePasswordRequestBodyRequired(obj ManagerAccountV1120ChangePasswordRequestBody) error {
	elements := map[string]interface{}{
		"NewPassword":            obj.NewPassword,
		"SessionAccountPassword": obj.SessionAccountPassword,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertManagerAccountV1120ChangePasswordRequestBodyConstraints checks if the values respects the defined constraints
func AssertManagerAccountV1120ChangePasswordRequestBodyConstraints(obj ManagerAccountV1120ChangePasswordRequestBody) error {
	return nil
}