// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VolumeV1100AssignReplicaTarget - This action is used to establish a replication relationship by assigning an existing volume to serve as a target replica for an existing source volume.
type VolumeV1100AssignReplicaTarget struct {

	// Link to invoke action
	Target string `json:"target,omitempty"`

	// Friendly action name
	Title string `json:"title,omitempty"`
}

// AssertVolumeV1100AssignReplicaTargetRequired checks if the required fields are not zero-ed
func AssertVolumeV1100AssignReplicaTargetRequired(obj VolumeV1100AssignReplicaTarget) error {
	return nil
}

// AssertVolumeV1100AssignReplicaTargetConstraints checks if the values respects the defined constraints
func AssertVolumeV1100AssignReplicaTargetConstraints(obj VolumeV1100AssignReplicaTarget) error {
	return nil
}