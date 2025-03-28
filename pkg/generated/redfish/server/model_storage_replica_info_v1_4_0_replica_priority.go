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

// StorageReplicaInfoV140ReplicaPriority : ReplicaPriority allows the priority of background copy engine I/O to be managed relative to host I/O operations during a sequential background copy operation.
type StorageReplicaInfoV140ReplicaPriority string

// List of StorageReplicaInfoV140ReplicaPriority
const (
	STORAGEREPLICAINFOV140REPLICAPRIORITY_LOW    StorageReplicaInfoV140ReplicaPriority = "Low"
	STORAGEREPLICAINFOV140REPLICAPRIORITY_SAME   StorageReplicaInfoV140ReplicaPriority = "Same"
	STORAGEREPLICAINFOV140REPLICAPRIORITY_HIGH   StorageReplicaInfoV140ReplicaPriority = "High"
	STORAGEREPLICAINFOV140REPLICAPRIORITY_URGENT StorageReplicaInfoV140ReplicaPriority = "Urgent"
)

// AllowedStorageReplicaInfoV140ReplicaPriorityEnumValues is all the allowed values of StorageReplicaInfoV140ReplicaPriority enum
var AllowedStorageReplicaInfoV140ReplicaPriorityEnumValues = []StorageReplicaInfoV140ReplicaPriority{
	"Low",
	"Same",
	"High",
	"Urgent",
}

// validStorageReplicaInfoV140ReplicaPriorityEnumValue provides a map of StorageReplicaInfoV140ReplicaPrioritys for fast verification of use input
var validStorageReplicaInfoV140ReplicaPriorityEnumValues = map[StorageReplicaInfoV140ReplicaPriority]struct{}{
	"Low":    {},
	"Same":   {},
	"High":   {},
	"Urgent": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageReplicaInfoV140ReplicaPriority) IsValid() bool {
	_, ok := validStorageReplicaInfoV140ReplicaPriorityEnumValues[v]
	return ok
}

// NewStorageReplicaInfoV140ReplicaPriorityFromValue returns a pointer to a valid StorageReplicaInfoV140ReplicaPriority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageReplicaInfoV140ReplicaPriorityFromValue(v string) (StorageReplicaInfoV140ReplicaPriority, error) {
	ev := StorageReplicaInfoV140ReplicaPriority(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for StorageReplicaInfoV140ReplicaPriority: valid values are %v", v, AllowedStorageReplicaInfoV140ReplicaPriorityEnumValues)
}

// AssertStorageReplicaInfoV140ReplicaPriorityRequired checks if the required fields are not zero-ed
func AssertStorageReplicaInfoV140ReplicaPriorityRequired(obj StorageReplicaInfoV140ReplicaPriority) error {
	return nil
}

// AssertStorageReplicaInfoV140ReplicaPriorityConstraints checks if the values respects the defined constraints
func AssertStorageReplicaInfoV140ReplicaPriorityConstraints(obj StorageReplicaInfoV140ReplicaPriority) error {
	return nil
}
