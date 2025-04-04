// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CapacityV121CapacityInfo - The capacity of specific data type in a data store.
type CapacityV121CapacityInfo struct {

	// The number of bytes currently allocated by the storage system in this data store for this data type.
	AllocatedBytes *int64 `json:"AllocatedBytes,omitempty"`

	// The number of bytes consumed in this data store for this data type.
	ConsumedBytes *int64 `json:"ConsumedBytes,omitempty"`

	// The number of bytes the storage system guarantees can be allocated in this data store for this data type.
	GuaranteedBytes *int64 `json:"GuaranteedBytes,omitempty"`

	// The maximum number of bytes that can be allocated in this data store for this data type.
	ProvisionedBytes *int64 `json:"ProvisionedBytes,omitempty"`
}

// AssertCapacityV121CapacityInfoRequired checks if the required fields are not zero-ed
func AssertCapacityV121CapacityInfoRequired(obj CapacityV121CapacityInfo) error {
	return nil
}

// AssertCapacityV121CapacityInfoConstraints checks if the values respects the defined constraints
func AssertCapacityV121CapacityInfoConstraints(obj CapacityV121CapacityInfo) error {
	return nil
}
