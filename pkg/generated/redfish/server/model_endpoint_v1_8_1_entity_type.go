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

type EndpointV181EntityType string

// List of EndpointV181EntityType
const (
	ENDPOINTV181ENTITYTYPE_STORAGE_INITIATOR     EndpointV181EntityType = "StorageInitiator"
	ENDPOINTV181ENTITYTYPE_ROOT_COMPLEX          EndpointV181EntityType = "RootComplex"
	ENDPOINTV181ENTITYTYPE_NETWORK_CONTROLLER    EndpointV181EntityType = "NetworkController"
	ENDPOINTV181ENTITYTYPE_DRIVE                 EndpointV181EntityType = "Drive"
	ENDPOINTV181ENTITYTYPE_STORAGE_EXPANDER      EndpointV181EntityType = "StorageExpander"
	ENDPOINTV181ENTITYTYPE_DISPLAY_CONTROLLER    EndpointV181EntityType = "DisplayController"
	ENDPOINTV181ENTITYTYPE_BRIDGE                EndpointV181EntityType = "Bridge"
	ENDPOINTV181ENTITYTYPE_PROCESSOR             EndpointV181EntityType = "Processor"
	ENDPOINTV181ENTITYTYPE_VOLUME                EndpointV181EntityType = "Volume"
	ENDPOINTV181ENTITYTYPE_ACCELERATION_FUNCTION EndpointV181EntityType = "AccelerationFunction"
	ENDPOINTV181ENTITYTYPE_MEDIA_CONTROLLER      EndpointV181EntityType = "MediaController"
	ENDPOINTV181ENTITYTYPE_MEMORY_CHUNK          EndpointV181EntityType = "MemoryChunk"
	ENDPOINTV181ENTITYTYPE_SWITCH                EndpointV181EntityType = "Switch"
	ENDPOINTV181ENTITYTYPE_FABRIC_BRIDGE         EndpointV181EntityType = "FabricBridge"
	ENDPOINTV181ENTITYTYPE_MANAGER               EndpointV181EntityType = "Manager"
	ENDPOINTV181ENTITYTYPE_STORAGE_SUBSYSTEM     EndpointV181EntityType = "StorageSubsystem"
	ENDPOINTV181ENTITYTYPE_MEMORY                EndpointV181EntityType = "Memory"
	ENDPOINTV181ENTITYTYPE_CXL_DEVICE            EndpointV181EntityType = "CXLDevice"
)

// AllowedEndpointV181EntityTypeEnumValues is all the allowed values of EndpointV181EntityType enum
var AllowedEndpointV181EntityTypeEnumValues = []EndpointV181EntityType{
	"StorageInitiator",
	"RootComplex",
	"NetworkController",
	"Drive",
	"StorageExpander",
	"DisplayController",
	"Bridge",
	"Processor",
	"Volume",
	"AccelerationFunction",
	"MediaController",
	"MemoryChunk",
	"Switch",
	"FabricBridge",
	"Manager",
	"StorageSubsystem",
	"Memory",
	"CXLDevice",
}

// validEndpointV181EntityTypeEnumValue provides a map of EndpointV181EntityTypes for fast verification of use input
var validEndpointV181EntityTypeEnumValues = map[EndpointV181EntityType]struct{}{
	"StorageInitiator":     {},
	"RootComplex":          {},
	"NetworkController":    {},
	"Drive":                {},
	"StorageExpander":      {},
	"DisplayController":    {},
	"Bridge":               {},
	"Processor":            {},
	"Volume":               {},
	"AccelerationFunction": {},
	"MediaController":      {},
	"MemoryChunk":          {},
	"Switch":               {},
	"FabricBridge":         {},
	"Manager":              {},
	"StorageSubsystem":     {},
	"Memory":               {},
	"CXLDevice":            {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EndpointV181EntityType) IsValid() bool {
	_, ok := validEndpointV181EntityTypeEnumValues[v]
	return ok
}

// NewEndpointV181EntityTypeFromValue returns a pointer to a valid EndpointV181EntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndpointV181EntityTypeFromValue(v string) (EndpointV181EntityType, error) {
	ev := EndpointV181EntityType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for EndpointV181EntityType: valid values are %v", v, AllowedEndpointV181EntityTypeEnumValues)
}

// AssertEndpointV181EntityTypeRequired checks if the required fields are not zero-ed
func AssertEndpointV181EntityTypeRequired(obj EndpointV181EntityType) error {
	return nil
}

// AssertEndpointV181EntityTypeConstraints checks if the values respects the defined constraints
func AssertEndpointV181EntityTypeConstraints(obj EndpointV181EntityType) error {
	return nil
}