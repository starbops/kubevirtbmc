// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ReservoirV101Reservoir - The Reservoir schema describes a reservoir unit for a cooling system or similar device.
type ReservoirV101Reservoir struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ReservoirV101Actions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The capacity of the reservoir (L).
	CapacityLiters *float32 `json:"CapacityLiters,omitempty"`

	Coolant CoolingLoopCoolant `json:"Coolant,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	Filters OdataV4IdRef `json:"Filters,omitempty"`

	FluidLevelPercent SensorSensorExcerpt `json:"FluidLevelPercent,omitempty"`

	FluidLevelStatus ResourceHealth `json:"FluidLevelStatus,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	InternalPressurekPa SensorSensorExcerpt `json:"InternalPressurekPa,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	// The manufacturer of this reservoir.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// The model number for this reservoir.
	Model *string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The part number for this reservoir.
	PartNumber *string `json:"PartNumber,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	ReservoirType ReservoirV101ReservoirType `json:"ReservoirType,omitempty"`

	// The serial number for this reservoir.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// The spare part number for this reservoir.
	SparePartNumber *string `json:"SparePartNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// A user-assigned label.
	UserLabel string `json:"UserLabel,omitempty"`
}

// AssertReservoirV101ReservoirRequired checks if the required fields are not zero-ed
func AssertReservoirV101ReservoirRequired(obj ReservoirV101Reservoir) error {
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

	if err := AssertReservoirV101ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertCoolingLoopCoolantRequired(obj.Coolant); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Filters); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.FluidLevelPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.InternalPressurekPa); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertReservoirV101ReservoirConstraints checks if the values respects the defined constraints
func AssertReservoirV101ReservoirConstraints(obj ReservoirV101Reservoir) error {
	if err := AssertReservoirV101ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Assembly); err != nil {
		return err
	}
	if err := AssertCoolingLoopCoolantConstraints(obj.Coolant); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Filters); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.FluidLevelPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.InternalPressurekPa); err != nil {
		return err
	}
	if err := AssertResourceLocationConstraints(obj.Location); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}