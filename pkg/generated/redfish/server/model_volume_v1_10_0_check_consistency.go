// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VolumeV1100CheckConsistency - This action is used to force a check of the Volume's parity or redundant data to ensure it matches calculated values.
type VolumeV1100CheckConsistency struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertVolumeV1100CheckConsistencyRequired checks if the required fields are not zero-ed
func AssertVolumeV1100CheckConsistencyRequired(obj VolumeV1100CheckConsistency) error {
	return nil
}

// AssertVolumeV1100CheckConsistencyConstraints checks if the values respects the defined constraints
func AssertVolumeV1100CheckConsistencyConstraints(obj VolumeV1100CheckConsistency) error {
	return nil
}
