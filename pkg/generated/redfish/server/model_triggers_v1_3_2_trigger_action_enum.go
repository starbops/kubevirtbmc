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

// TriggersV132TriggerActionEnum : The actions to perform when a trigger condition is met.
type TriggersV132TriggerActionEnum string

// List of TriggersV132TriggerActionEnum
const (
	TRIGGERSV132TRIGGERACTIONENUM_LOG_TO_LOG_SERVICE    TriggersV132TriggerActionEnum = "LogToLogService"
	TRIGGERSV132TRIGGERACTIONENUM_REDFISH_EVENT         TriggersV132TriggerActionEnum = "RedfishEvent"
	TRIGGERSV132TRIGGERACTIONENUM_REDFISH_METRIC_REPORT TriggersV132TriggerActionEnum = "RedfishMetricReport"
)

// AllowedTriggersV132TriggerActionEnumEnumValues is all the allowed values of TriggersV132TriggerActionEnum enum
var AllowedTriggersV132TriggerActionEnumEnumValues = []TriggersV132TriggerActionEnum{
	"LogToLogService",
	"RedfishEvent",
	"RedfishMetricReport",
}

// validTriggersV132TriggerActionEnumEnumValue provides a map of TriggersV132TriggerActionEnums for fast verification of use input
var validTriggersV132TriggerActionEnumEnumValues = map[TriggersV132TriggerActionEnum]struct{}{
	"LogToLogService":     {},
	"RedfishEvent":        {},
	"RedfishMetricReport": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggersV132TriggerActionEnum) IsValid() bool {
	_, ok := validTriggersV132TriggerActionEnumEnumValues[v]
	return ok
}

// NewTriggersV132TriggerActionEnumFromValue returns a pointer to a valid TriggersV132TriggerActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggersV132TriggerActionEnumFromValue(v string) (TriggersV132TriggerActionEnum, error) {
	ev := TriggersV132TriggerActionEnum(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TriggersV132TriggerActionEnum: valid values are %v", v, AllowedTriggersV132TriggerActionEnumEnumValues)
}

// AssertTriggersV132TriggerActionEnumRequired checks if the required fields are not zero-ed
func AssertTriggersV132TriggerActionEnumRequired(obj TriggersV132TriggerActionEnum) error {
	return nil
}

// AssertTriggersV132TriggerActionEnumConstraints checks if the values respects the defined constraints
func AssertTriggersV132TriggerActionEnumConstraints(obj TriggersV132TriggerActionEnum) error {
	return nil
}