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

// CoolingLoopV102CoolingLoop - This schema describes a cooling loop.  A cooling loop might be any coolant-carrying vessel, such as facility-level pipework, an immersion cooling tank, or a manifold.  A loop might describe its connectors and instrumentation but does not generally include active cooling components or subsystems.
type CoolingLoopV102CoolingLoop struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions CoolingLoopV102Actions `json:"Actions,omitempty"`

	// An array of names of downstream devices that receive coolant from this cooling loop.
	ConsumingEquipmentNames []*string `json:"ConsumingEquipmentNames,omitempty"`

	Coolant CoolingLoopV102Coolant `json:"Coolant,omitempty"`

	CoolantLevelPercent SensorSensorExcerpt `json:"CoolantLevelPercent,omitempty"`

	CoolantLevelStatus ResourceHealth `json:"CoolantLevelStatus,omitempty"`

	CoolantQuality ResourceHealth `json:"CoolantQuality,omitempty"`

	// The link to the application that manages the cooling loop.
	CoolingManagerURI *string `json:"CoolingManagerURI,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Links CoolingLoopV102Links `json:"Links,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PrimaryCoolantConnectors OdataV4IdRef `json:"PrimaryCoolantConnectors,omitempty"`

	// The rated liquid flow (L/min) for this cooling loop.
	RatedFlowLitersPerMinute *float32 `json:"RatedFlowLitersPerMinute,omitempty"`

	// The rated pressure (kPa) for this cooling loop.
	RatedPressurekPa *float32 `json:"RatedPressurekPa,omitempty"`

	SecondaryCoolantConnectors OdataV4IdRef `json:"SecondaryCoolantConnectors,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// An array of names of upstream devices that supply coolant to this cooling loop.
	SupplyEquipmentNames []*string `json:"SupplyEquipmentNames,omitempty"`

	// A user-assigned label.
	UserLabel string `json:"UserLabel,omitempty"`
}

// AssertCoolingLoopV102CoolingLoopRequired checks if the required fields are not zero-ed
func AssertCoolingLoopV102CoolingLoopRequired(obj CoolingLoopV102CoolingLoop) error {
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

	if err := AssertCoolingLoopV102ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertCoolingLoopV102CoolantRequired(obj.Coolant); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.CoolantLevelPercent); err != nil {
		return err
	}
	if err := AssertCoolingLoopV102LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.PrimaryCoolantConnectors); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.SecondaryCoolantConnectors); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertCoolingLoopV102CoolingLoopConstraints checks if the values respects the defined constraints
func AssertCoolingLoopV102CoolingLoopConstraints(obj CoolingLoopV102CoolingLoop) error {
	if err := AssertCoolingLoopV102ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertCoolingLoopV102CoolantConstraints(obj.Coolant); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.CoolantLevelPercent); err != nil {
		return err
	}
	if err := AssertCoolingLoopV102LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.PrimaryCoolantConnectors); err != nil {
		return err
	}
	if obj.RatedFlowLitersPerMinute != nil && *obj.RatedFlowLitersPerMinute < 0 {
		return &ParsingError{Param: "RatedFlowLitersPerMinute", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.RatedPressurekPa != nil && *obj.RatedPressurekPa < 0 {
		return &ParsingError{Param: "RatedPressurekPa", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertOdataV4IdRefConstraints(obj.SecondaryCoolantConnectors); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}