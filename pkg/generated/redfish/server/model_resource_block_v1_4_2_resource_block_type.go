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

type ResourceBlockV142ResourceBlockType string

// List of ResourceBlockV142ResourceBlockType
const (
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_COMPUTE              ResourceBlockV142ResourceBlockType = "Compute"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_PROCESSOR            ResourceBlockV142ResourceBlockType = "Processor"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_MEMORY               ResourceBlockV142ResourceBlockType = "Memory"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_NETWORK              ResourceBlockV142ResourceBlockType = "Network"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_STORAGE              ResourceBlockV142ResourceBlockType = "Storage"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_COMPUTER_SYSTEM      ResourceBlockV142ResourceBlockType = "ComputerSystem"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_EXPANSION            ResourceBlockV142ResourceBlockType = "Expansion"
	RESOURCEBLOCKV142RESOURCEBLOCKTYPE_INDEPENDENT_RESOURCE ResourceBlockV142ResourceBlockType = "IndependentResource"
)

// AllowedResourceBlockV142ResourceBlockTypeEnumValues is all the allowed values of ResourceBlockV142ResourceBlockType enum
var AllowedResourceBlockV142ResourceBlockTypeEnumValues = []ResourceBlockV142ResourceBlockType{
	"Compute",
	"Processor",
	"Memory",
	"Network",
	"Storage",
	"ComputerSystem",
	"Expansion",
	"IndependentResource",
}

// validResourceBlockV142ResourceBlockTypeEnumValue provides a map of ResourceBlockV142ResourceBlockTypes for fast verification of use input
var validResourceBlockV142ResourceBlockTypeEnumValues = map[ResourceBlockV142ResourceBlockType]struct{}{
	"Compute":             {},
	"Processor":           {},
	"Memory":              {},
	"Network":             {},
	"Storage":             {},
	"ComputerSystem":      {},
	"Expansion":           {},
	"IndependentResource": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceBlockV142ResourceBlockType) IsValid() bool {
	_, ok := validResourceBlockV142ResourceBlockTypeEnumValues[v]
	return ok
}

// NewResourceBlockV142ResourceBlockTypeFromValue returns a pointer to a valid ResourceBlockV142ResourceBlockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceBlockV142ResourceBlockTypeFromValue(v string) (ResourceBlockV142ResourceBlockType, error) {
	ev := ResourceBlockV142ResourceBlockType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ResourceBlockV142ResourceBlockType: valid values are %v", v, AllowedResourceBlockV142ResourceBlockTypeEnumValues)
}

// AssertResourceBlockV142ResourceBlockTypeRequired checks if the required fields are not zero-ed
func AssertResourceBlockV142ResourceBlockTypeRequired(obj ResourceBlockV142ResourceBlockType) error {
	return nil
}

// AssertResourceBlockV142ResourceBlockTypeConstraints checks if the values respects the defined constraints
func AssertResourceBlockV142ResourceBlockTypeConstraints(obj ResourceBlockV142ResourceBlockType) error {
	return nil
}