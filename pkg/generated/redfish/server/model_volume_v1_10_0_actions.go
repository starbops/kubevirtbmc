// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

type VolumeV1100Actions struct {
	VolumeAssignReplicaTarget VolumeV1100AssignReplicaTarget `json:"#Volume.AssignReplicaTarget,omitempty"`

	VolumeChangeRAIDLayout VolumeV1100ChangeRaidLayout `json:"#Volume.ChangeRAIDLayout,omitempty"`

	VolumeCheckConsistency VolumeV1100CheckConsistency `json:"#Volume.CheckConsistency,omitempty"`

	VolumeCreateReplicaTarget VolumeV1100CreateReplicaTarget `json:"#Volume.CreateReplicaTarget,omitempty"`

	VolumeForceEnable VolumeV1100ForceEnable `json:"#Volume.ForceEnable,omitempty"`

	VolumeInitialize VolumeV1100Initialize `json:"#Volume.Initialize,omitempty"`

	VolumeRemoveReplicaRelationship VolumeV1100RemoveReplicaRelationship `json:"#Volume.RemoveReplicaRelationship,omitempty"`

	VolumeResumeReplication VolumeV1100ResumeReplication `json:"#Volume.ResumeReplication,omitempty"`

	VolumeReverseReplicationRelationship VolumeV1100ReverseReplicationRelationship `json:"#Volume.ReverseReplicationRelationship,omitempty"`

	VolumeSplitReplication VolumeV1100SplitReplication `json:"#Volume.SplitReplication,omitempty"`

	VolumeSuspendReplication VolumeV1100SuspendReplication `json:"#Volume.SuspendReplication,omitempty"`

	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertVolumeV1100ActionsRequired checks if the required fields are not zero-ed
func AssertVolumeV1100ActionsRequired(obj VolumeV1100Actions) error {
	if err := AssertVolumeV1100AssignReplicaTargetRequired(obj.VolumeAssignReplicaTarget); err != nil {
		return err
	}
	if err := AssertVolumeV1100ChangeRaidLayoutRequired(obj.VolumeChangeRAIDLayout); err != nil {
		return err
	}
	if err := AssertVolumeV1100CheckConsistencyRequired(obj.VolumeCheckConsistency); err != nil {
		return err
	}
	if err := AssertVolumeV1100CreateReplicaTargetRequired(obj.VolumeCreateReplicaTarget); err != nil {
		return err
	}
	if err := AssertVolumeV1100ForceEnableRequired(obj.VolumeForceEnable); err != nil {
		return err
	}
	if err := AssertVolumeV1100InitializeRequired(obj.VolumeInitialize); err != nil {
		return err
	}
	if err := AssertVolumeV1100RemoveReplicaRelationshipRequired(obj.VolumeRemoveReplicaRelationship); err != nil {
		return err
	}
	if err := AssertVolumeV1100ResumeReplicationRequired(obj.VolumeResumeReplication); err != nil {
		return err
	}
	if err := AssertVolumeV1100ReverseReplicationRelationshipRequired(obj.VolumeReverseReplicationRelationship); err != nil {
		return err
	}
	if err := AssertVolumeV1100SplitReplicationRequired(obj.VolumeSplitReplication); err != nil {
		return err
	}
	if err := AssertVolumeV1100SuspendReplicationRequired(obj.VolumeSuspendReplication); err != nil {
		return err
	}
	return nil
}

// AssertVolumeV1100ActionsConstraints checks if the values respects the defined constraints
func AssertVolumeV1100ActionsConstraints(obj VolumeV1100Actions) error {
	if err := AssertVolumeV1100AssignReplicaTargetConstraints(obj.VolumeAssignReplicaTarget); err != nil {
		return err
	}
	if err := AssertVolumeV1100ChangeRaidLayoutConstraints(obj.VolumeChangeRAIDLayout); err != nil {
		return err
	}
	if err := AssertVolumeV1100CheckConsistencyConstraints(obj.VolumeCheckConsistency); err != nil {
		return err
	}
	if err := AssertVolumeV1100CreateReplicaTargetConstraints(obj.VolumeCreateReplicaTarget); err != nil {
		return err
	}
	if err := AssertVolumeV1100ForceEnableConstraints(obj.VolumeForceEnable); err != nil {
		return err
	}
	if err := AssertVolumeV1100InitializeConstraints(obj.VolumeInitialize); err != nil {
		return err
	}
	if err := AssertVolumeV1100RemoveReplicaRelationshipConstraints(obj.VolumeRemoveReplicaRelationship); err != nil {
		return err
	}
	if err := AssertVolumeV1100ResumeReplicationConstraints(obj.VolumeResumeReplication); err != nil {
		return err
	}
	if err := AssertVolumeV1100ReverseReplicationRelationshipConstraints(obj.VolumeReverseReplicationRelationship); err != nil {
		return err
	}
	if err := AssertVolumeV1100SplitReplicationConstraints(obj.VolumeSplitReplication); err != nil {
		return err
	}
	if err := AssertVolumeV1100SuspendReplicationConstraints(obj.VolumeSuspendReplication); err != nil {
		return err
	}
	return nil
}