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

type ComputerSystemV1230BootSourceOverrideEnabled string

// List of ComputerSystemV1230BootSourceOverrideEnabled
const (
	COMPUTERSYSTEMV1230BOOTSOURCEOVERRIDEENABLED_DISABLED   ComputerSystemV1230BootSourceOverrideEnabled = "Disabled"
	COMPUTERSYSTEMV1230BOOTSOURCEOVERRIDEENABLED_ONCE       ComputerSystemV1230BootSourceOverrideEnabled = "Once"
	COMPUTERSYSTEMV1230BOOTSOURCEOVERRIDEENABLED_CONTINUOUS ComputerSystemV1230BootSourceOverrideEnabled = "Continuous"
)

// AllowedComputerSystemV1230BootSourceOverrideEnabledEnumValues is all the allowed values of ComputerSystemV1230BootSourceOverrideEnabled enum
var AllowedComputerSystemV1230BootSourceOverrideEnabledEnumValues = []ComputerSystemV1230BootSourceOverrideEnabled{
	"Disabled",
	"Once",
	"Continuous",
}

// validComputerSystemV1230BootSourceOverrideEnabledEnumValue provides a map of ComputerSystemV1230BootSourceOverrideEnableds for fast verification of use input
var validComputerSystemV1230BootSourceOverrideEnabledEnumValues = map[ComputerSystemV1230BootSourceOverrideEnabled]struct{}{
	"Disabled":   {},
	"Once":       {},
	"Continuous": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1230BootSourceOverrideEnabled) IsValid() bool {
	_, ok := validComputerSystemV1230BootSourceOverrideEnabledEnumValues[v]
	return ok
}

// NewComputerSystemV1230BootSourceOverrideEnabledFromValue returns a pointer to a valid ComputerSystemV1230BootSourceOverrideEnabled
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1230BootSourceOverrideEnabledFromValue(v string) (ComputerSystemV1230BootSourceOverrideEnabled, error) {
	ev := ComputerSystemV1230BootSourceOverrideEnabled(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1230BootSourceOverrideEnabled: valid values are %v", v, AllowedComputerSystemV1230BootSourceOverrideEnabledEnumValues)
}

// AssertComputerSystemV1230BootSourceOverrideEnabledRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1230BootSourceOverrideEnabledRequired(obj ComputerSystemV1230BootSourceOverrideEnabled) error {
	return nil
}

// AssertComputerSystemV1230BootSourceOverrideEnabledConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1230BootSourceOverrideEnabledConstraints(obj ComputerSystemV1230BootSourceOverrideEnabled) error {
	return nil
}