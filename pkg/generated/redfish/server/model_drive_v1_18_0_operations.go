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
	"errors"
)

// DriveV1180Operations - An operation currently running on this resource.
type DriveV1180Operations struct {
	AssociatedTask OdataV4IdRef `json:"AssociatedTask,omitempty"`

	Operation VolumeOperationType `json:"Operation,omitempty"`

	// The name of the operation.
	// Deprecated
	OperationName *string `json:"OperationName,omitempty"`

	// The percentage of the operation that has been completed.
	PercentageComplete *int64 `json:"PercentageComplete,omitempty"`
}

// AssertDriveV1180OperationsRequired checks if the required fields are not zero-ed
func AssertDriveV1180OperationsRequired(obj DriveV1180Operations) error {
	if err := AssertOdataV4IdRefRequired(obj.AssociatedTask); err != nil {
		return err
	}
	return nil
}

// AssertDriveV1180OperationsConstraints checks if the values respects the defined constraints
func AssertDriveV1180OperationsConstraints(obj DriveV1180Operations) error {
	if err := AssertOdataV4IdRefConstraints(obj.AssociatedTask); err != nil {
		return err
	}
	if obj.PercentageComplete != nil && *obj.PercentageComplete < 0 {
		return &ParsingError{Param: "PercentageComplete", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PercentageComplete != nil && *obj.PercentageComplete > 100 {
		return &ParsingError{Param: "PercentageComplete", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
