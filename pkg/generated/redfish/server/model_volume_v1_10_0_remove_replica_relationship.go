// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VolumeV1100RemoveReplicaRelationship - This action is used to disable data synchronization between a source and target volume, remove the replication relationship, and optionally delete the target volume.
type VolumeV1100RemoveReplicaRelationship struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertVolumeV1100RemoveReplicaRelationshipRequired checks if the required fields are not zero-ed
func AssertVolumeV1100RemoveReplicaRelationshipRequired(obj VolumeV1100RemoveReplicaRelationship) error {
	return nil
}

// AssertVolumeV1100RemoveReplicaRelationshipConstraints checks if the values respects the defined constraints
func AssertVolumeV1100RemoveReplicaRelationshipConstraints(obj VolumeV1100RemoveReplicaRelationship) error {
	return nil
}