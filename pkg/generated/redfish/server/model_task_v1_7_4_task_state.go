// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

import (
	"fmt"
)

type TaskV174TaskState string

// List of TaskV174TaskState
const (
	TASKV174TASKSTATE_NEW         TaskV174TaskState = "New"
	TASKV174TASKSTATE_STARTING    TaskV174TaskState = "Starting"
	TASKV174TASKSTATE_RUNNING     TaskV174TaskState = "Running"
	TASKV174TASKSTATE_SUSPENDED   TaskV174TaskState = "Suspended"
	TASKV174TASKSTATE_INTERRUPTED TaskV174TaskState = "Interrupted"
	TASKV174TASKSTATE_PENDING     TaskV174TaskState = "Pending"
	TASKV174TASKSTATE_STOPPING    TaskV174TaskState = "Stopping"
	TASKV174TASKSTATE_COMPLETED   TaskV174TaskState = "Completed"
	TASKV174TASKSTATE_KILLED      TaskV174TaskState = "Killed"
	TASKV174TASKSTATE_EXCEPTION   TaskV174TaskState = "Exception"
	TASKV174TASKSTATE_SERVICE     TaskV174TaskState = "Service"
	TASKV174TASKSTATE_CANCELLING  TaskV174TaskState = "Cancelling"
	TASKV174TASKSTATE_CANCELLED   TaskV174TaskState = "Cancelled"
)

// AllowedTaskV174TaskStateEnumValues is all the allowed values of TaskV174TaskState enum
var AllowedTaskV174TaskStateEnumValues = []TaskV174TaskState{
	"New",
	"Starting",
	"Running",
	"Suspended",
	"Interrupted",
	"Pending",
	"Stopping",
	"Completed",
	"Killed",
	"Exception",
	"Service",
	"Cancelling",
	"Cancelled",
}

// validTaskV174TaskStateEnumValue provides a map of TaskV174TaskStates for fast verification of use input
var validTaskV174TaskStateEnumValues = map[TaskV174TaskState]struct{}{
	"New":         {},
	"Starting":    {},
	"Running":     {},
	"Suspended":   {},
	"Interrupted": {},
	"Pending":     {},
	"Stopping":    {},
	"Completed":   {},
	"Killed":      {},
	"Exception":   {},
	"Service":     {},
	"Cancelling":  {},
	"Cancelled":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskV174TaskState) IsValid() bool {
	_, ok := validTaskV174TaskStateEnumValues[v]
	return ok
}

// NewTaskV174TaskStateFromValue returns a pointer to a valid TaskV174TaskState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskV174TaskStateFromValue(v string) (TaskV174TaskState, error) {
	ev := TaskV174TaskState(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TaskV174TaskState: valid values are %v", v, AllowedTaskV174TaskStateEnumValues)
}

// AssertTaskV174TaskStateRequired checks if the required fields are not zero-ed
func AssertTaskV174TaskStateRequired(obj TaskV174TaskState) error {
	return nil
}

// AssertTaskV174TaskStateConstraints checks if the values respects the defined constraints
func AssertTaskV174TaskStateConstraints(obj TaskV174TaskState) error {
	return nil
}