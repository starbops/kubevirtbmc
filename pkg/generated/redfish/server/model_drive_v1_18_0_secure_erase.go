// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// DriveV1180SecureErase - This action securely erases the contents of the drive.
type DriveV1180SecureErase struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertDriveV1180SecureEraseRequired checks if the required fields are not zero-ed
func AssertDriveV1180SecureEraseRequired(obj DriveV1180SecureErase) error {
	return nil
}

// AssertDriveV1180SecureEraseConstraints checks if the values respects the defined constraints
func AssertDriveV1180SecureEraseConstraints(obj DriveV1180SecureErase) error {
	return nil
}