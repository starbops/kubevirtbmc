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
	"errors"
)

// PortV1110FunctionMaxBandwidth - A maximum bandwidth allocation percentage for a network device functions associated with a port.
type PortV1110FunctionMaxBandwidth struct {

	// The maximum bandwidth allocation percentage allocated to the corresponding network device function instance.
	AllocationPercent *int64 `json:"AllocationPercent,omitempty"`

	NetworkDeviceFunction OdataV4IdRef `json:"NetworkDeviceFunction,omitempty"`
}

// AssertPortV1110FunctionMaxBandwidthRequired checks if the required fields are not zero-ed
func AssertPortV1110FunctionMaxBandwidthRequired(obj PortV1110FunctionMaxBandwidth) error {
	if err := AssertOdataV4IdRefRequired(obj.NetworkDeviceFunction); err != nil {
		return err
	}
	return nil
}

// AssertPortV1110FunctionMaxBandwidthConstraints checks if the values respects the defined constraints
func AssertPortV1110FunctionMaxBandwidthConstraints(obj PortV1110FunctionMaxBandwidth) error {
	if obj.AllocationPercent != nil && *obj.AllocationPercent < 0 {
		return &ParsingError{Param: "AllocationPercent", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AllocationPercent != nil && *obj.AllocationPercent > 100 {
		return &ParsingError{Param: "AllocationPercent", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertOdataV4IdRefConstraints(obj.NetworkDeviceFunction); err != nil {
		return err
	}
	return nil
}