// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VolumeV1100ForceEnable - Request system force the volume to an enabled state regardless of data loss.
type VolumeV1100ForceEnable struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertVolumeV1100ForceEnableRequired checks if the required fields are not zero-ed
func AssertVolumeV1100ForceEnableRequired(obj VolumeV1100ForceEnable) error {
	return nil
}

// AssertVolumeV1100ForceEnableConstraints checks if the values respects the defined constraints
func AssertVolumeV1100ForceEnableConstraints(obj VolumeV1100ForceEnable) error {
	return nil
}
