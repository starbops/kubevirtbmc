// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

import (
	"fmt"
)

// StorageReplicaInfoV140ConsistencyState : The values of ConsistencyState indicate the consistency type used by the source and its associated target group.
type StorageReplicaInfoV140ConsistencyState string

// List of StorageReplicaInfoV140ConsistencyState
const (
	STORAGEREPLICAINFOV140CONSISTENCYSTATE_CONSISTENT   StorageReplicaInfoV140ConsistencyState = "Consistent"
	STORAGEREPLICAINFOV140CONSISTENCYSTATE_INCONSISTENT StorageReplicaInfoV140ConsistencyState = "Inconsistent"
)

// AllowedStorageReplicaInfoV140ConsistencyStateEnumValues is all the allowed values of StorageReplicaInfoV140ConsistencyState enum
var AllowedStorageReplicaInfoV140ConsistencyStateEnumValues = []StorageReplicaInfoV140ConsistencyState{
	"Consistent",
	"Inconsistent",
}

// validStorageReplicaInfoV140ConsistencyStateEnumValue provides a map of StorageReplicaInfoV140ConsistencyStates for fast verification of use input
var validStorageReplicaInfoV140ConsistencyStateEnumValues = map[StorageReplicaInfoV140ConsistencyState]struct{}{
	"Consistent":   {},
	"Inconsistent": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageReplicaInfoV140ConsistencyState) IsValid() bool {
	_, ok := validStorageReplicaInfoV140ConsistencyStateEnumValues[v]
	return ok
}

// NewStorageReplicaInfoV140ConsistencyStateFromValue returns a pointer to a valid StorageReplicaInfoV140ConsistencyState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageReplicaInfoV140ConsistencyStateFromValue(v string) (StorageReplicaInfoV140ConsistencyState, error) {
	ev := StorageReplicaInfoV140ConsistencyState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for StorageReplicaInfoV140ConsistencyState: valid values are %v", v, AllowedStorageReplicaInfoV140ConsistencyStateEnumValues)
}

// AssertStorageReplicaInfoV140ConsistencyStateRequired checks if the required fields are not zero-ed
func AssertStorageReplicaInfoV140ConsistencyStateRequired(obj StorageReplicaInfoV140ConsistencyState) error {
	return nil
}

// AssertStorageReplicaInfoV140ConsistencyStateConstraints checks if the values respects the defined constraints
func AssertStorageReplicaInfoV140ConsistencyStateConstraints(obj StorageReplicaInfoV140ConsistencyState) error {
	return nil
}