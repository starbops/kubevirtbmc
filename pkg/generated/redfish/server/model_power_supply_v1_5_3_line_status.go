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

type PowerSupplyV153LineStatus string

// List of PowerSupplyV153LineStatus
const (
	POWERSUPPLYV153LINESTATUS_NORMAL        PowerSupplyV153LineStatus = "Normal"
	POWERSUPPLYV153LINESTATUS_LOSS_OF_INPUT PowerSupplyV153LineStatus = "LossOfInput"
	POWERSUPPLYV153LINESTATUS_OUT_OF_RANGE  PowerSupplyV153LineStatus = "OutOfRange"
)

// AllowedPowerSupplyV153LineStatusEnumValues is all the allowed values of PowerSupplyV153LineStatus enum
var AllowedPowerSupplyV153LineStatusEnumValues = []PowerSupplyV153LineStatus{
	"Normal",
	"LossOfInput",
	"OutOfRange",
}

// validPowerSupplyV153LineStatusEnumValue provides a map of PowerSupplyV153LineStatuss for fast verification of use input
var validPowerSupplyV153LineStatusEnumValues = map[PowerSupplyV153LineStatus]struct{}{
	"Normal":      {},
	"LossOfInput": {},
	"OutOfRange":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerSupplyV153LineStatus) IsValid() bool {
	_, ok := validPowerSupplyV153LineStatusEnumValues[v]
	return ok
}

// NewPowerSupplyV153LineStatusFromValue returns a pointer to a valid PowerSupplyV153LineStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerSupplyV153LineStatusFromValue(v string) (PowerSupplyV153LineStatus, error) {
	ev := PowerSupplyV153LineStatus(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PowerSupplyV153LineStatus: valid values are %v", v, AllowedPowerSupplyV153LineStatusEnumValues)
}

// AssertPowerSupplyV153LineStatusRequired checks if the required fields are not zero-ed
func AssertPowerSupplyV153LineStatusRequired(obj PowerSupplyV153LineStatus) error {
	return nil
}

// AssertPowerSupplyV153LineStatusConstraints checks if the values respects the defined constraints
func AssertPowerSupplyV153LineStatusConstraints(obj PowerSupplyV153LineStatus) error {
	return nil
}