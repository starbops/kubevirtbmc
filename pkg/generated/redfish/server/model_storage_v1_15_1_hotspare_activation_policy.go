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

type StorageV1151HotspareActivationPolicy string

// List of StorageV1151HotspareActivationPolicy
const (
	STORAGEV1151HOTSPAREACTIVATIONPOLICY_ON_DRIVE_FAILURE           StorageV1151HotspareActivationPolicy = "OnDriveFailure"
	STORAGEV1151HOTSPAREACTIVATIONPOLICY_ON_DRIVE_PREDICTED_FAILURE StorageV1151HotspareActivationPolicy = "OnDrivePredictedFailure"
	STORAGEV1151HOTSPAREACTIVATIONPOLICY_OEM                        StorageV1151HotspareActivationPolicy = "OEM"
)

// AllowedStorageV1151HotspareActivationPolicyEnumValues is all the allowed values of StorageV1151HotspareActivationPolicy enum
var AllowedStorageV1151HotspareActivationPolicyEnumValues = []StorageV1151HotspareActivationPolicy{
	"OnDriveFailure",
	"OnDrivePredictedFailure",
	"OEM",
}

// validStorageV1151HotspareActivationPolicyEnumValue provides a map of StorageV1151HotspareActivationPolicys for fast verification of use input
var validStorageV1151HotspareActivationPolicyEnumValues = map[StorageV1151HotspareActivationPolicy]struct{}{
	"OnDriveFailure":          {},
	"OnDrivePredictedFailure": {},
	"OEM":                     {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageV1151HotspareActivationPolicy) IsValid() bool {
	_, ok := validStorageV1151HotspareActivationPolicyEnumValues[v]
	return ok
}

// NewStorageV1151HotspareActivationPolicyFromValue returns a pointer to a valid StorageV1151HotspareActivationPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageV1151HotspareActivationPolicyFromValue(v string) (StorageV1151HotspareActivationPolicy, error) {
	ev := StorageV1151HotspareActivationPolicy(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for StorageV1151HotspareActivationPolicy: valid values are %v", v, AllowedStorageV1151HotspareActivationPolicyEnumValues)
}

// AssertStorageV1151HotspareActivationPolicyRequired checks if the required fields are not zero-ed
func AssertStorageV1151HotspareActivationPolicyRequired(obj StorageV1151HotspareActivationPolicy) error {
	return nil
}

// AssertStorageV1151HotspareActivationPolicyConstraints checks if the values respects the defined constraints
func AssertStorageV1151HotspareActivationPolicyConstraints(obj StorageV1151HotspareActivationPolicy) error {
	return nil
}