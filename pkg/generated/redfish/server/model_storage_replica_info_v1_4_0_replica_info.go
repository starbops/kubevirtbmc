// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// StorageReplicaInfoV140ReplicaInfo - Defines the characteristics of a replica.
type StorageReplicaInfoV140ReplicaInfo struct {

	// True if consistency is enabled.
	ConsistencyEnabled *bool `json:"ConsistencyEnabled,omitempty"`

	ConsistencyState *StorageReplicaInfoV140ConsistencyState `json:"ConsistencyState,omitempty"`

	ConsistencyStatus *StorageReplicaInfoV140ConsistencyStatus `json:"ConsistencyStatus,omitempty"`

	ConsistencyType *StorageReplicaInfoV140ConsistencyType `json:"ConsistencyType,omitempty"`

	DataProtectionLineOfService OdataV4IdRef `json:"DataProtectionLineOfService,omitempty"`

	// If true, the storage array tells host to stop sending data to source element if copying to a remote element fails.
	FailedCopyStopsHostIO *bool `json:"FailedCopyStopsHostIO,omitempty"`

	// Specifies the percent of the work completed to reach synchronization.
	PercentSynced *int64 `json:"PercentSynced,omitempty"`

	// ReplicaFaultDomain describes the fault domain (local or remote) of the replica relationship.
	RemoteSourceReplica *string `json:"RemoteSourceReplica,omitempty"`

	Replica OdataV4IdRef `json:"Replica,omitempty"`

	ReplicaFaultDomain *StorageReplicaInfoReplicaFaultDomain `json:"ReplicaFaultDomain,omitempty"`

	ReplicaPriority *StorageReplicaInfoV140ReplicaPriority `json:"ReplicaPriority,omitempty"`

	ReplicaProgressStatus *StorageReplicaInfoV140ReplicaProgressStatus `json:"ReplicaProgressStatus,omitempty"`

	ReplicaReadOnlyAccess *StorageReplicaInfoV140ReplicaReadOnlyAccess `json:"ReplicaReadOnlyAccess,omitempty"`

	ReplicaRecoveryMode *StorageReplicaInfoV140ReplicaRecoveryMode `json:"ReplicaRecoveryMode,omitempty"`

	// Deprecated
	ReplicaRole *StorageReplicaInfoV140ReplicaRole `json:"ReplicaRole,omitempty"`

	// Applies to Adaptive mode and it describes maximum number of bytes the SyncedElement (target) can be out of sync.
	ReplicaSkewBytes *int64 `json:"ReplicaSkewBytes,omitempty"`

	ReplicaState *StorageReplicaInfoV140ReplicaState `json:"ReplicaState,omitempty"`

	ReplicaType *StorageReplicaInfoReplicaType `json:"ReplicaType,omitempty"`

	ReplicaUpdateMode *StorageReplicaInfoReplicaUpdateMode `json:"ReplicaUpdateMode,omitempty"`

	RequestedReplicaState *StorageReplicaInfoV140ReplicaState `json:"RequestedReplicaState,omitempty"`

	SourceReplica OdataV4IdRef `json:"SourceReplica,omitempty"`

	// Synchronization is maintained.
	SyncMaintained *bool `json:"SyncMaintained,omitempty"`

	UndiscoveredElement *StorageReplicaInfoV140UndiscoveredElement `json:"UndiscoveredElement,omitempty"`

	// Specifies when point-in-time copy was taken or when the replication relationship is activated, reactivated, resumed or re-established.
	WhenActivated *string `json:"WhenActivated,omitempty"`

	// Specifies when the replication relationship is deactivated.
	WhenDeactivated *string `json:"WhenDeactivated,omitempty"`

	// Specifies when the replication relationship is established.
	WhenEstablished *string `json:"WhenEstablished,omitempty"`

	// Specifies when the replication relationship is suspended.
	WhenSuspended *string `json:"WhenSuspended,omitempty"`

	// The point in time that the Elements were synchronized.
	WhenSynced *string `json:"WhenSynced,omitempty"`

	// Specifies when the replication relationship is synchronized.
	WhenSynchronized *string `json:"WhenSynchronized,omitempty"`
}

// AssertStorageReplicaInfoV140ReplicaInfoRequired checks if the required fields are not zero-ed
func AssertStorageReplicaInfoV140ReplicaInfoRequired(obj StorageReplicaInfoV140ReplicaInfo) error {
	if err := AssertOdataV4IdRefRequired(obj.DataProtectionLineOfService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Replica); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.SourceReplica); err != nil {
		return err
	}
	return nil
}

// AssertStorageReplicaInfoV140ReplicaInfoConstraints checks if the values respects the defined constraints
func AssertStorageReplicaInfoV140ReplicaInfoConstraints(obj StorageReplicaInfoV140ReplicaInfo) error {
	if err := AssertOdataV4IdRefConstraints(obj.DataProtectionLineOfService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Replica); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.SourceReplica); err != nil {
		return err
	}
	return nil
}