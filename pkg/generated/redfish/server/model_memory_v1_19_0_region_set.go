// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryV1190RegionSet - Memory region information within a memory device.
type MemoryV1190RegionSet struct {

	// An indication of whether the master passphrase is enabled for this region.
	MasterPassphraseEnabled *bool `json:"MasterPassphraseEnabled,omitempty"`

	MemoryClassification MemoryV1190MemoryClassification `json:"MemoryClassification,omitempty"`

	// Offset within the memory that corresponds to the start of this memory region in mebibytes (MiB).
	OffsetMiB *int64 `json:"OffsetMiB,omitempty"`

	// An indication of whether the passphrase is enabled for this region.
	PassphraseEnabled *bool `json:"PassphraseEnabled,omitempty"`

	// An indication of whether the state of the passphrase for this region is enabled.
	// Deprecated
	PassphraseState *bool `json:"PassphraseState,omitempty"`

	// Unique region ID representing a specific region within the memory device.
	RegionId *string `json:"RegionId,omitempty"`

	// Size of this memory region in mebibytes (MiB).
	SizeMiB *int64 `json:"SizeMiB,omitempty"`
}

// AssertMemoryV1190RegionSetRequired checks if the required fields are not zero-ed
func AssertMemoryV1190RegionSetRequired(obj MemoryV1190RegionSet) error {
	return nil
}

// AssertMemoryV1190RegionSetConstraints checks if the values respects the defined constraints
func AssertMemoryV1190RegionSetConstraints(obj MemoryV1190RegionSet) error {
	return nil
}