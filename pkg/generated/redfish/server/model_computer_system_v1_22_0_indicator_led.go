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

type ComputerSystemV1220IndicatorLed string

// List of ComputerSystemV1220IndicatorLed
const (
	COMPUTERSYSTEMV1220INDICATORLED_UNKNOWN  ComputerSystemV1220IndicatorLed = "Unknown"
	COMPUTERSYSTEMV1220INDICATORLED_LIT      ComputerSystemV1220IndicatorLed = "Lit"
	COMPUTERSYSTEMV1220INDICATORLED_BLINKING ComputerSystemV1220IndicatorLed = "Blinking"
	COMPUTERSYSTEMV1220INDICATORLED_OFF      ComputerSystemV1220IndicatorLed = "Off"
)

// AllowedComputerSystemV1220IndicatorLedEnumValues is all the allowed values of ComputerSystemV1220IndicatorLed enum
var AllowedComputerSystemV1220IndicatorLedEnumValues = []ComputerSystemV1220IndicatorLed{
	"Unknown",
	"Lit",
	"Blinking",
	"Off",
}

// validComputerSystemV1220IndicatorLedEnumValue provides a map of ComputerSystemV1220IndicatorLeds for fast verification of use input
var validComputerSystemV1220IndicatorLedEnumValues = map[ComputerSystemV1220IndicatorLed]struct{}{
	"Unknown":  {},
	"Lit":      {},
	"Blinking": {},
	"Off":      {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1220IndicatorLed) IsValid() bool {
	_, ok := validComputerSystemV1220IndicatorLedEnumValues[v]
	return ok
}

// NewComputerSystemV1220IndicatorLedFromValue returns a pointer to a valid ComputerSystemV1220IndicatorLed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1220IndicatorLedFromValue(v string) (ComputerSystemV1220IndicatorLed, error) {
	ev := ComputerSystemV1220IndicatorLed(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1220IndicatorLed: valid values are %v", v, AllowedComputerSystemV1220IndicatorLedEnumValues)
}

// AssertComputerSystemV1220IndicatorLedRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220IndicatorLedRequired(obj ComputerSystemV1220IndicatorLed) error {
	return nil
}

// AssertComputerSystemV1220IndicatorLedConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220IndicatorLedConstraints(obj ComputerSystemV1220IndicatorLed) error {
	return nil
}