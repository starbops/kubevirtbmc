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

// FanV151Fan - The Fan schema describes a cooling fan unit for a computer system or similar devices contained within a chassis.  It also describes the location, such as a slot, socket, or bay, where a unit can be installed, by populating a resource instance with an absent state if a unit is not present.
type FanV151Fan struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions FanV151Actions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The diameter of the fan assembly in millimeter units.
	FanDiameterMm *int64 `json:"FanDiameterMm,omitempty"`

	// An indication of whether this device can be inserted or removed while the equipment is in operation.
	HotPluggable *bool `json:"HotPluggable,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Links FanV151Links `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	// The manufacturer of this fan.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// The model number for this fan.
	Model *string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this fan.
	PartNumber *string `json:"PartNumber,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	PowerWatts SensorSensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// An indication of whether this component can be independently replaced as allowed by the vendor's replacement policy.
	Replaceable *bool `json:"Replaceable,omitempty"`

	SecondarySpeedPercent SensorSensorFanExcerpt `json:"SecondarySpeedPercent,omitempty"`

	// The serial number for this fan.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// The spare part number for this fan.
	SparePartNumber *string `json:"SparePartNumber,omitempty"`

	SpeedPercent SensorSensorFanExcerpt `json:"SpeedPercent,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`
}

// AssertFanV151FanRequired checks if the required fields are not zero-ed
func AssertFanV151FanRequired(obj FanV151Fan) error {
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

	if err := AssertFanV151ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertFanV151LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	if err := AssertSensorSensorPowerExcerptRequired(obj.PowerWatts); err != nil {
		return err
	}
	if err := AssertSensorSensorFanExcerptRequired(obj.SecondarySpeedPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorFanExcerptRequired(obj.SpeedPercent); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertFanV151FanConstraints checks if the values respects the defined constraints
func AssertFanV151FanConstraints(obj FanV151Fan) error {
	if err := AssertFanV151ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Assembly); err != nil {
		return err
	}
	if obj.FanDiameterMm != nil && *obj.FanDiameterMm < 0 {
		return &ParsingError{Param: "FanDiameterMm", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertFanV151LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationConstraints(obj.Location); err != nil {
		return err
	}
	if err := AssertSensorSensorPowerExcerptConstraints(obj.PowerWatts); err != nil {
		return err
	}
	if err := AssertSensorSensorFanExcerptConstraints(obj.SecondarySpeedPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorFanExcerptConstraints(obj.SpeedPercent); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}
