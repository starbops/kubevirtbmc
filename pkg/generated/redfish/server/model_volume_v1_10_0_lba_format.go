// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

type VolumeV1100LbaFormat struct {

	// The LBA data size in bytes.
	LBADataSizeBytes *int64 `json:"LBADataSizeBytes,omitempty"`

	LBAFormatType *VolumeLbaFormatType `json:"LBAFormatType,omitempty"`

	// The LBA metadata size in bytes.
	LBAMetadataSizeBytes *int64 `json:"LBAMetadataSizeBytes,omitempty"`

	RelativePerformance *VolumeLbaRelativePerformanceType `json:"RelativePerformance,omitempty"`
}

// AssertVolumeV1100LbaFormatRequired checks if the required fields are not zero-ed
func AssertVolumeV1100LbaFormatRequired(obj VolumeV1100LbaFormat) error {
	return nil
}

// AssertVolumeV1100LbaFormatConstraints checks if the values respects the defined constraints
func AssertVolumeV1100LbaFormatConstraints(obj VolumeV1100LbaFormat) error {
	return nil
}