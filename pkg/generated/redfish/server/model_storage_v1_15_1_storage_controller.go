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

// StorageV1151StorageController - The StorageController schema describes a storage controller and its properties.  A storage controller represents a physical or virtual storage device that produces volumes.
type StorageV1151StorageController struct {

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	Actions StorageV1151StorageControllerActions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this storage controller.
	AssetTag *string `json:"AssetTag,omitempty"`

	CacheSummary StorageV1151CacheSummary `json:"CacheSummary,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	ControllerRates StorageV1151Rates `json:"ControllerRates,omitempty"`

	// The firmware version of this storage controller.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`

	// The durable names for the storage controller.
	Identifiers []ResourceIdentifier `json:"Identifiers,omitempty"`

	Links StorageV1151StorageControllerLinks `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// The manufacturer of this storage controller.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	// Deprecated
	Measurements []SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`

	// The unique identifier for the member within an array.
	MemberId string `json:"MemberId"`

	// The model number for the storage controller.
	Model *string `json:"Model,omitempty"`

	// The name of the storage controller.
	Name *string `json:"Name,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeInterface PcieDevicePcieInterface `json:"PCIeInterface,omitempty"`

	// The part number for this storage controller.
	PartNumber *string `json:"PartNumber,omitempty"`

	Ports OdataV4IdRef `json:"Ports,omitempty"`

	// The SKU for this storage controller.
	SKU *string `json:"SKU,omitempty"`

	// The serial number for this storage controller.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// The maximum speed of the storage controller's device interface.
	SpeedGbps *float32 `json:"SpeedGbps,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The supported set of protocols for communicating with this storage controller.
	SupportedControllerProtocols []ProtocolProtocol `json:"SupportedControllerProtocols,omitempty"`

	// The protocols that the storage controller can use to communicate with attached devices.
	SupportedDeviceProtocols []ProtocolProtocol `json:"SupportedDeviceProtocols,omitempty"`

	// The set of RAID types supported by the storage controller.
	SupportedRAIDTypes []VolumeRaidType `json:"SupportedRAIDTypes,omitempty"`
}

// AssertStorageV1151StorageControllerRequired checks if the required fields are not zero-ed
func AssertStorageV1151StorageControllerRequired(obj StorageV1151StorageController) error {
	elements := map[string]interface{}{
		"@odata.id": obj.OdataId,
		"MemberId":  obj.MemberId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertStorageV1151StorageControllerActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertStorageV1151CacheSummaryRequired(obj.CacheSummary); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Certificates); err != nil {
		return err
	}
	if err := AssertStorageV1151RatesRequired(obj.ControllerRates); err != nil {
		return err
	}
	for _, el := range obj.Identifiers {
		if err := AssertResourceIdentifierRequired(el); err != nil {
			return err
		}
	}
	if err := AssertStorageV1151StorageControllerLinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockRequired(el); err != nil {
			return err
		}
	}
	if err := AssertPcieDevicePcieInterfaceRequired(obj.PCIeInterface); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Ports); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertStorageV1151StorageControllerConstraints checks if the values respects the defined constraints
func AssertStorageV1151StorageControllerConstraints(obj StorageV1151StorageController) error {
	if err := AssertStorageV1151StorageControllerActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Assembly); err != nil {
		return err
	}
	if err := AssertStorageV1151CacheSummaryConstraints(obj.CacheSummary); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Certificates); err != nil {
		return err
	}
	if err := AssertStorageV1151RatesConstraints(obj.ControllerRates); err != nil {
		return err
	}
	for _, el := range obj.Identifiers {
		if err := AssertResourceIdentifierConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertStorageV1151StorageControllerLinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationConstraints(obj.Location); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertPcieDevicePcieInterfaceConstraints(obj.PCIeInterface); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Ports); err != nil {
		return err
	}
	if obj.SpeedGbps != nil && *obj.SpeedGbps < 0 {
		return &ParsingError{Param: "SpeedGbps", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}