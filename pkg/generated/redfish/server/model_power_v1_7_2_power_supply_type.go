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

type PowerV172PowerSupplyType string

// List of PowerV172PowerSupplyType
const (
	POWERV172POWERSUPPLYTYPE_UNKNOWN  PowerV172PowerSupplyType = "Unknown"
	POWERV172POWERSUPPLYTYPE_AC       PowerV172PowerSupplyType = "AC"
	POWERV172POWERSUPPLYTYPE_DC       PowerV172PowerSupplyType = "DC"
	POWERV172POWERSUPPLYTYPE_A_COR_DC PowerV172PowerSupplyType = "ACorDC"
)

// AllowedPowerV172PowerSupplyTypeEnumValues is all the allowed values of PowerV172PowerSupplyType enum
var AllowedPowerV172PowerSupplyTypeEnumValues = []PowerV172PowerSupplyType{
	"Unknown",
	"AC",
	"DC",
	"ACorDC",
}

// validPowerV172PowerSupplyTypeEnumValue provides a map of PowerV172PowerSupplyTypes for fast verification of use input
var validPowerV172PowerSupplyTypeEnumValues = map[PowerV172PowerSupplyType]struct{}{
	"Unknown": {},
	"AC":      {},
	"DC":      {},
	"ACorDC":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerV172PowerSupplyType) IsValid() bool {
	_, ok := validPowerV172PowerSupplyTypeEnumValues[v]
	return ok
}

// NewPowerV172PowerSupplyTypeFromValue returns a pointer to a valid PowerV172PowerSupplyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerV172PowerSupplyTypeFromValue(v string) (PowerV172PowerSupplyType, error) {
	ev := PowerV172PowerSupplyType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PowerV172PowerSupplyType: valid values are %v", v, AllowedPowerV172PowerSupplyTypeEnumValues)
}

// AssertPowerV172PowerSupplyTypeRequired checks if the required fields are not zero-ed
func AssertPowerV172PowerSupplyTypeRequired(obj PowerV172PowerSupplyType) error {
	return nil
}

// AssertPowerV172PowerSupplyTypeConstraints checks if the values respects the defined constraints
func AssertPowerV172PowerSupplyTypeConstraints(obj PowerV172PowerSupplyType) error {
	return nil
}