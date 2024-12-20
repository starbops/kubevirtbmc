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

// StorageReplicaInfoV140ReplicaProgressStatus : Values of ReplicaProgressStatus describe the status of the session with respect to Replication activity.
type StorageReplicaInfoV140ReplicaProgressStatus string

// List of StorageReplicaInfoV140ReplicaProgressStatus
const (
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_COMPLETED         StorageReplicaInfoV140ReplicaProgressStatus = "Completed"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_DORMANT           StorageReplicaInfoV140ReplicaProgressStatus = "Dormant"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_INITIALIZING      StorageReplicaInfoV140ReplicaProgressStatus = "Initializing"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_PREPARING         StorageReplicaInfoV140ReplicaProgressStatus = "Preparing"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_SYNCHRONIZING     StorageReplicaInfoV140ReplicaProgressStatus = "Synchronizing"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_RESYNCING         StorageReplicaInfoV140ReplicaProgressStatus = "Resyncing"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_RESTORING         StorageReplicaInfoV140ReplicaProgressStatus = "Restoring"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_FRACTURING        StorageReplicaInfoV140ReplicaProgressStatus = "Fracturing"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_SPLITTING         StorageReplicaInfoV140ReplicaProgressStatus = "Splitting"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_FAILING_OVER      StorageReplicaInfoV140ReplicaProgressStatus = "FailingOver"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_FAILING_BACK      StorageReplicaInfoV140ReplicaProgressStatus = "FailingBack"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_DETACHING         StorageReplicaInfoV140ReplicaProgressStatus = "Detaching"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_ABORTING          StorageReplicaInfoV140ReplicaProgressStatus = "Aborting"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_MIXED             StorageReplicaInfoV140ReplicaProgressStatus = "Mixed"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_SUSPENDING        StorageReplicaInfoV140ReplicaProgressStatus = "Suspending"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_REQUIRES_FRACTURE StorageReplicaInfoV140ReplicaProgressStatus = "RequiresFracture"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_REQUIRES_RESYNC   StorageReplicaInfoV140ReplicaProgressStatus = "RequiresResync"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_REQUIRES_ACTIVATE StorageReplicaInfoV140ReplicaProgressStatus = "RequiresActivate"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_PENDING           StorageReplicaInfoV140ReplicaProgressStatus = "Pending"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_REQUIRES_DETACH   StorageReplicaInfoV140ReplicaProgressStatus = "RequiresDetach"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_TERMINATING       StorageReplicaInfoV140ReplicaProgressStatus = "Terminating"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_REQUIRES_SPLIT    StorageReplicaInfoV140ReplicaProgressStatus = "RequiresSplit"
	STORAGEREPLICAINFOV140REPLICAPROGRESSSTATUS_REQUIRES_RESUME   StorageReplicaInfoV140ReplicaProgressStatus = "RequiresResume"
)

// AllowedStorageReplicaInfoV140ReplicaProgressStatusEnumValues is all the allowed values of StorageReplicaInfoV140ReplicaProgressStatus enum
var AllowedStorageReplicaInfoV140ReplicaProgressStatusEnumValues = []StorageReplicaInfoV140ReplicaProgressStatus{
	"Completed",
	"Dormant",
	"Initializing",
	"Preparing",
	"Synchronizing",
	"Resyncing",
	"Restoring",
	"Fracturing",
	"Splitting",
	"FailingOver",
	"FailingBack",
	"Detaching",
	"Aborting",
	"Mixed",
	"Suspending",
	"RequiresFracture",
	"RequiresResync",
	"RequiresActivate",
	"Pending",
	"RequiresDetach",
	"Terminating",
	"RequiresSplit",
	"RequiresResume",
}

// validStorageReplicaInfoV140ReplicaProgressStatusEnumValue provides a map of StorageReplicaInfoV140ReplicaProgressStatuss for fast verification of use input
var validStorageReplicaInfoV140ReplicaProgressStatusEnumValues = map[StorageReplicaInfoV140ReplicaProgressStatus]struct{}{
	"Completed":        {},
	"Dormant":          {},
	"Initializing":     {},
	"Preparing":        {},
	"Synchronizing":    {},
	"Resyncing":        {},
	"Restoring":        {},
	"Fracturing":       {},
	"Splitting":        {},
	"FailingOver":      {},
	"FailingBack":      {},
	"Detaching":        {},
	"Aborting":         {},
	"Mixed":            {},
	"Suspending":       {},
	"RequiresFracture": {},
	"RequiresResync":   {},
	"RequiresActivate": {},
	"Pending":          {},
	"RequiresDetach":   {},
	"Terminating":      {},
	"RequiresSplit":    {},
	"RequiresResume":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageReplicaInfoV140ReplicaProgressStatus) IsValid() bool {
	_, ok := validStorageReplicaInfoV140ReplicaProgressStatusEnumValues[v]
	return ok
}

// NewStorageReplicaInfoV140ReplicaProgressStatusFromValue returns a pointer to a valid StorageReplicaInfoV140ReplicaProgressStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageReplicaInfoV140ReplicaProgressStatusFromValue(v string) (StorageReplicaInfoV140ReplicaProgressStatus, error) {
	ev := StorageReplicaInfoV140ReplicaProgressStatus(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for StorageReplicaInfoV140ReplicaProgressStatus: valid values are %v", v, AllowedStorageReplicaInfoV140ReplicaProgressStatusEnumValues)
}

// AssertStorageReplicaInfoV140ReplicaProgressStatusRequired checks if the required fields are not zero-ed
func AssertStorageReplicaInfoV140ReplicaProgressStatusRequired(obj StorageReplicaInfoV140ReplicaProgressStatus) error {
	return nil
}

// AssertStorageReplicaInfoV140ReplicaProgressStatusConstraints checks if the values respects the defined constraints
func AssertStorageReplicaInfoV140ReplicaProgressStatusConstraints(obj StorageReplicaInfoV140ReplicaProgressStatus) error {
	return nil
}
