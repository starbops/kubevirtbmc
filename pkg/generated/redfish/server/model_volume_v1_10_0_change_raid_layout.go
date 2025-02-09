// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VolumeV1100ChangeRaidLayout - Request system change the RAID layout of the volume.
type VolumeV1100ChangeRaidLayout struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertVolumeV1100ChangeRaidLayoutRequired checks if the required fields are not zero-ed
func AssertVolumeV1100ChangeRaidLayoutRequired(obj VolumeV1100ChangeRaidLayout) error {
	return nil
}

// AssertVolumeV1100ChangeRaidLayoutConstraints checks if the values respects the defined constraints
func AssertVolumeV1100ChangeRaidLayoutConstraints(obj VolumeV1100ChangeRaidLayout) error {
	return nil
}
