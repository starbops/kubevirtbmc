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

type ComputerSystemV1220TrustedModuleRequiredToBoot string

// List of ComputerSystemV1220TrustedModuleRequiredToBoot
const (
	COMPUTERSYSTEMV1220TRUSTEDMODULEREQUIREDTOBOOT_DISABLED ComputerSystemV1220TrustedModuleRequiredToBoot = "Disabled"
	COMPUTERSYSTEMV1220TRUSTEDMODULEREQUIREDTOBOOT_REQUIRED ComputerSystemV1220TrustedModuleRequiredToBoot = "Required"
)

// AllowedComputerSystemV1220TrustedModuleRequiredToBootEnumValues is all the allowed values of ComputerSystemV1220TrustedModuleRequiredToBoot enum
var AllowedComputerSystemV1220TrustedModuleRequiredToBootEnumValues = []ComputerSystemV1220TrustedModuleRequiredToBoot{
	"Disabled",
	"Required",
}

// validComputerSystemV1220TrustedModuleRequiredToBootEnumValue provides a map of ComputerSystemV1220TrustedModuleRequiredToBoots for fast verification of use input
var validComputerSystemV1220TrustedModuleRequiredToBootEnumValues = map[ComputerSystemV1220TrustedModuleRequiredToBoot]struct{}{
	"Disabled": {},
	"Required": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputerSystemV1220TrustedModuleRequiredToBoot) IsValid() bool {
	_, ok := validComputerSystemV1220TrustedModuleRequiredToBootEnumValues[v]
	return ok
}

// NewComputerSystemV1220TrustedModuleRequiredToBootFromValue returns a pointer to a valid ComputerSystemV1220TrustedModuleRequiredToBoot
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputerSystemV1220TrustedModuleRequiredToBootFromValue(v string) (ComputerSystemV1220TrustedModuleRequiredToBoot, error) {
	ev := ComputerSystemV1220TrustedModuleRequiredToBoot(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComputerSystemV1220TrustedModuleRequiredToBoot: valid values are %v", v, AllowedComputerSystemV1220TrustedModuleRequiredToBootEnumValues)
}

// AssertComputerSystemV1220TrustedModuleRequiredToBootRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1220TrustedModuleRequiredToBootRequired(obj ComputerSystemV1220TrustedModuleRequiredToBoot) error {
	return nil
}

// AssertComputerSystemV1220TrustedModuleRequiredToBootConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1220TrustedModuleRequiredToBootConstraints(obj ComputerSystemV1220TrustedModuleRequiredToBoot) error {
	return nil
}