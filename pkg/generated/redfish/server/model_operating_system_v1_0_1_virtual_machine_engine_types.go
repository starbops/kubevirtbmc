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

type OperatingSystemV101VirtualMachineEngineTypes string

// List of OperatingSystemV101VirtualMachineEngineTypes
const (
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_V_MWARE_ESX OperatingSystemV101VirtualMachineEngineTypes = "VMwareESX"
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_HYPER_V     OperatingSystemV101VirtualMachineEngineTypes = "HyperV"
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_XEN         OperatingSystemV101VirtualMachineEngineTypes = "Xen"
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_KVM         OperatingSystemV101VirtualMachineEngineTypes = "KVM"
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_QEMU        OperatingSystemV101VirtualMachineEngineTypes = "QEMU"
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_VIRTUAL_BOX OperatingSystemV101VirtualMachineEngineTypes = "VirtualBox"
	OPERATINGSYSTEMV101VIRTUALMACHINEENGINETYPES_POWER_VM    OperatingSystemV101VirtualMachineEngineTypes = "PowerVM"
)

// AllowedOperatingSystemV101VirtualMachineEngineTypesEnumValues is all the allowed values of OperatingSystemV101VirtualMachineEngineTypes enum
var AllowedOperatingSystemV101VirtualMachineEngineTypesEnumValues = []OperatingSystemV101VirtualMachineEngineTypes{
	"VMwareESX",
	"HyperV",
	"Xen",
	"KVM",
	"QEMU",
	"VirtualBox",
	"PowerVM",
}

// validOperatingSystemV101VirtualMachineEngineTypesEnumValue provides a map of OperatingSystemV101VirtualMachineEngineTypess for fast verification of use input
var validOperatingSystemV101VirtualMachineEngineTypesEnumValues = map[OperatingSystemV101VirtualMachineEngineTypes]struct{}{
	"VMwareESX":  {},
	"HyperV":     {},
	"Xen":        {},
	"KVM":        {},
	"QEMU":       {},
	"VirtualBox": {},
	"PowerVM":    {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperatingSystemV101VirtualMachineEngineTypes) IsValid() bool {
	_, ok := validOperatingSystemV101VirtualMachineEngineTypesEnumValues[v]
	return ok
}

// NewOperatingSystemV101VirtualMachineEngineTypesFromValue returns a pointer to a valid OperatingSystemV101VirtualMachineEngineTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatingSystemV101VirtualMachineEngineTypesFromValue(v string) (OperatingSystemV101VirtualMachineEngineTypes, error) {
	ev := OperatingSystemV101VirtualMachineEngineTypes(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for OperatingSystemV101VirtualMachineEngineTypes: valid values are %v", v, AllowedOperatingSystemV101VirtualMachineEngineTypesEnumValues)
}

// AssertOperatingSystemV101VirtualMachineEngineTypesRequired checks if the required fields are not zero-ed
func AssertOperatingSystemV101VirtualMachineEngineTypesRequired(obj OperatingSystemV101VirtualMachineEngineTypes) error {
	return nil
}

// AssertOperatingSystemV101VirtualMachineEngineTypesConstraints checks if the values respects the defined constraints
func AssertOperatingSystemV101VirtualMachineEngineTypesConstraints(obj OperatingSystemV101VirtualMachineEngineTypes) error {
	return nil
}