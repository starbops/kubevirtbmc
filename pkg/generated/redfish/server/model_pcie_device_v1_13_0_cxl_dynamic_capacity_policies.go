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

type PcieDeviceV1130CxlDynamicCapacityPolicies string

// List of PcieDeviceV1130CxlDynamicCapacityPolicies
const (
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_FREE         PcieDeviceV1130CxlDynamicCapacityPolicies = "Free"
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_CONTIGUOUS   PcieDeviceV1130CxlDynamicCapacityPolicies = "Contiguous"
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_PRESCRIPTIVE PcieDeviceV1130CxlDynamicCapacityPolicies = "Prescriptive"
	PCIEDEVICEV1130CXLDYNAMICCAPACITYPOLICIES_TAG_BASED    PcieDeviceV1130CxlDynamicCapacityPolicies = "TagBased"
)

// AllowedPcieDeviceV1130CxlDynamicCapacityPoliciesEnumValues is all the allowed values of PcieDeviceV1130CxlDynamicCapacityPolicies enum
var AllowedPcieDeviceV1130CxlDynamicCapacityPoliciesEnumValues = []PcieDeviceV1130CxlDynamicCapacityPolicies{
	"Free",
	"Contiguous",
	"Prescriptive",
	"TagBased",
}

// validPcieDeviceV1130CxlDynamicCapacityPoliciesEnumValue provides a map of PcieDeviceV1130CxlDynamicCapacityPoliciess for fast verification of use input
var validPcieDeviceV1130CxlDynamicCapacityPoliciesEnumValues = map[PcieDeviceV1130CxlDynamicCapacityPolicies]struct{}{
	"Free":         {},
	"Contiguous":   {},
	"Prescriptive": {},
	"TagBased":     {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PcieDeviceV1130CxlDynamicCapacityPolicies) IsValid() bool {
	_, ok := validPcieDeviceV1130CxlDynamicCapacityPoliciesEnumValues[v]
	return ok
}

// NewPcieDeviceV1130CxlDynamicCapacityPoliciesFromValue returns a pointer to a valid PcieDeviceV1130CxlDynamicCapacityPolicies
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPcieDeviceV1130CxlDynamicCapacityPoliciesFromValue(v string) (PcieDeviceV1130CxlDynamicCapacityPolicies, error) {
	ev := PcieDeviceV1130CxlDynamicCapacityPolicies(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for PcieDeviceV1130CxlDynamicCapacityPolicies: valid values are %v", v, AllowedPcieDeviceV1130CxlDynamicCapacityPoliciesEnumValues)
}

// AssertPcieDeviceV1130CxlDynamicCapacityPoliciesRequired checks if the required fields are not zero-ed
func AssertPcieDeviceV1130CxlDynamicCapacityPoliciesRequired(obj PcieDeviceV1130CxlDynamicCapacityPolicies) error {
	return nil
}

// AssertPcieDeviceV1130CxlDynamicCapacityPoliciesConstraints checks if the values respects the defined constraints
func AssertPcieDeviceV1130CxlDynamicCapacityPoliciesConstraints(obj PcieDeviceV1130CxlDynamicCapacityPolicies) error {
	return nil
}