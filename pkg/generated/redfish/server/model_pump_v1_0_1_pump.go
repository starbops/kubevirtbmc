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
	"time"
)

// PumpV101Pump - The Pump schema describes a pump unit for a cooling system or similar device.
type PumpV101Pump struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions PumpV101Actions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The user-assigned asset tag for this equipment.
	AssetTag *string `json:"AssetTag,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	Filters OdataV4IdRef `json:"Filters,omitempty"`

	// The firmware version of this equipment.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Location ResourceLocation `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	// The manufacturer of this pump.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// The model number for this pump.
	Model *string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this pump.
	PartNumber *string `json:"PartNumber,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	// The production or manufacturing date of this equipment.
	ProductionDate *time.Time `json:"ProductionDate,omitempty"`

	PumpSpeedPercent SensorSensorPumpExcerpt `json:"PumpSpeedPercent,omitempty"`

	PumpType PumpV101PumpType `json:"PumpType,omitempty"`

	// The serial number for this pump.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// The hours of service this pump has provided.
	ServiceHours *float32 `json:"ServiceHours,omitempty"`

	// The spare part number for this pump.
	SparePartNumber *string `json:"SparePartNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// A user-assigned label.
	UserLabel string `json:"UserLabel,omitempty"`

	// The hardware version of this equipment.
	Version *string `json:"Version,omitempty"`
}

// AssertPumpV101PumpRequired checks if the required fields are not zero-ed
func AssertPumpV101PumpRequired(obj PumpV101Pump) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPumpV101ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Filters); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	if err := AssertSensorSensorPumpExcerptRequired(obj.PumpSpeedPercent); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertPumpV101PumpConstraints checks if the values respects the defined constraints
func AssertPumpV101PumpConstraints(obj PumpV101Pump) error {
	if err := AssertPumpV101ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Assembly); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Filters); err != nil {
		return err
	}
	if err := AssertResourceLocationConstraints(obj.Location); err != nil {
		return err
	}
	if err := AssertSensorSensorPumpExcerptConstraints(obj.PumpSpeedPercent); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}