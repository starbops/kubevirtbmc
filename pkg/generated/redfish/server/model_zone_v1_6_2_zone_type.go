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

type ZoneV162ZoneType string

// List of ZoneV162ZoneType
const (
	ZONEV162ZONETYPE_DEFAULT                 ZoneV162ZoneType = "Default"
	ZONEV162ZONETYPE_ZONE_OF_ENDPOINTS       ZoneV162ZoneType = "ZoneOfEndpoints"
	ZONEV162ZONETYPE_ZONE_OF_ZONES           ZoneV162ZoneType = "ZoneOfZones"
	ZONEV162ZONETYPE_ZONE_OF_RESOURCE_BLOCKS ZoneV162ZoneType = "ZoneOfResourceBlocks"
)

// AllowedZoneV162ZoneTypeEnumValues is all the allowed values of ZoneV162ZoneType enum
var AllowedZoneV162ZoneTypeEnumValues = []ZoneV162ZoneType{
	"Default",
	"ZoneOfEndpoints",
	"ZoneOfZones",
	"ZoneOfResourceBlocks",
}

// validZoneV162ZoneTypeEnumValue provides a map of ZoneV162ZoneTypes for fast verification of use input
var validZoneV162ZoneTypeEnumValues = map[ZoneV162ZoneType]struct{}{
	"Default":              {},
	"ZoneOfEndpoints":      {},
	"ZoneOfZones":          {},
	"ZoneOfResourceBlocks": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZoneV162ZoneType) IsValid() bool {
	_, ok := validZoneV162ZoneTypeEnumValues[v]
	return ok
}

// NewZoneV162ZoneTypeFromValue returns a pointer to a valid ZoneV162ZoneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZoneV162ZoneTypeFromValue(v string) (ZoneV162ZoneType, error) {
	ev := ZoneV162ZoneType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ZoneV162ZoneType: valid values are %v", v, AllowedZoneV162ZoneTypeEnumValues)
}

// AssertZoneV162ZoneTypeRequired checks if the required fields are not zero-ed
func AssertZoneV162ZoneTypeRequired(obj ZoneV162ZoneType) error {
	return nil
}

// AssertZoneV162ZoneTypeConstraints checks if the values respects the defined constraints
func AssertZoneV162ZoneTypeConstraints(obj ZoneV162ZoneType) error {
	return nil
}