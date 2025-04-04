// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VolumeV1100CreateReplicaTargetRequestBody - This action is used to create a new volume resource to provide expanded data protection through a replica relationship with the specified source volume.
type VolumeV1100CreateReplicaTargetRequestBody struct {
	ReplicaType StorageReplicaInfoReplicaType `json:"ReplicaType"`

	ReplicaUpdateMode StorageReplicaInfoReplicaUpdateMode `json:"ReplicaUpdateMode"`

	// The Uri to the existing target Storage Pool.
	TargetStoragePool string `json:"TargetStoragePool"`

	// The Name for the new target volume.
	VolumeName string `json:"VolumeName,omitempty"`
}

// AssertVolumeV1100CreateReplicaTargetRequestBodyRequired checks if the required fields are not zero-ed
func AssertVolumeV1100CreateReplicaTargetRequestBodyRequired(obj VolumeV1100CreateReplicaTargetRequestBody) error {
	elements := map[string]interface{}{
		"ReplicaType":       obj.ReplicaType,
		"ReplicaUpdateMode": obj.ReplicaUpdateMode,
		"TargetStoragePool": obj.TargetStoragePool,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertVolumeV1100CreateReplicaTargetRequestBodyConstraints checks if the values respects the defined constraints
func AssertVolumeV1100CreateReplicaTargetRequestBodyConstraints(obj VolumeV1100CreateReplicaTargetRequestBody) error {
	return nil
}
