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

// ChassisV1250Chassis - The Chassis schema represents the physical components of a system.  This resource represents the sheet-metal confined spaces and logical zones such as racks, enclosures, chassis and all other containers.  Subsystems, such as sensors, that operate outside of a system's data plane are linked either directly or indirectly through this resource.  A subsystem that operates outside of a system's data plane are not accessible to software that runs on the system.  It also describes the location, such as a slot, socket, or bay, where a unit can be installed, by populating a resource instance with an absent state if a unit is not present.
type ChassisV1250Chassis struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ChassisV1250Actions `json:"Actions,omitempty"`

	Assembly OdataV4IdRef `json:"Assembly,omitempty"`

	// The user-assigned asset tag of this chassis.
	AssetTag *string `json:"AssetTag,omitempty"`

	Certificates OdataV4IdRef `json:"Certificates,omitempty"`

	ChassisType ChassisV1250ChassisType `json:"ChassisType"`

	Controls OdataV4IdRef `json:"Controls,omitempty"`

	// The depth of the chassis.
	DepthMm *float32 `json:"DepthMm,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	Doors ChassisV1250Doors `json:"Doors,omitempty"`

	Drives OdataV4IdRef `json:"Drives,omitempty"`

	// The URIs of the management interfaces for the external electrical source connections for this chassis.
	ElectricalSourceManagerURIs []*string `json:"ElectricalSourceManagerURIs,omitempty"`

	// The names of the external electrical sources, such as circuits or outlets, connected to this chassis.
	ElectricalSourceNames []*string `json:"ElectricalSourceNames,omitempty"`

	EnvironmentMetrics OdataV4IdRef `json:"EnvironmentMetrics,omitempty"`

	EnvironmentalClass ChassisV1250EnvironmentalClass `json:"EnvironmentalClass,omitempty"`

	FabricAdapters OdataV4IdRef `json:"FabricAdapters,omitempty"`

	// The names of the external heating or cooling equipment, such as coolant distribution units, connected to this chassis.
	HeatingCoolingEquipmentNames []*string `json:"HeatingCoolingEquipmentNames,omitempty"`

	// The URIs of the management interfaces for the external heating or cooling equipment for this chassis.
	HeatingCoolingManagerURIs []*string `json:"HeatingCoolingManagerURIs,omitempty"`

	// The height of the chassis.
	HeightMm *float32 `json:"HeightMm,omitempty"`

	// An indication of whether this component can be inserted or removed while the equipment is in operation.
	HotPluggable *bool `json:"HotPluggable,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	IndicatorLED ChassisV1250IndicatorLed `json:"IndicatorLED,omitempty"`

	Links ChassisV1250Links `json:"Links,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	LogServices OdataV4IdRef `json:"LogServices,omitempty"`

	// The manufacturer of this chassis.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	// The upper bound of the total power consumed by the chassis.
	MaxPowerWatts *float32 `json:"MaxPowerWatts,omitempty"`

	// An array of DSP0274-defined measurement blocks.
	// Deprecated
	Measurements []SoftwareInventoryMeasurementBlock `json:"Measurements,omitempty"`

	MediaControllers OdataV4IdRef `json:"MediaControllers,omitempty"`

	Memory OdataV4IdRef `json:"Memory,omitempty"`

	MemoryDomains OdataV4IdRef `json:"MemoryDomains,omitempty"`

	// The lower bound of the total power consumed by the chassis.
	MinPowerWatts *float32 `json:"MinPowerWatts,omitempty"`

	// The model number of the chassis.
	Model *string `json:"Model,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	NetworkAdapters OdataV4IdRef `json:"NetworkAdapters,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeDevices OdataV4IdRef `json:"PCIeDevices,omitempty"`

	PCIeSlots OdataV4IdRef `json:"PCIeSlots,omitempty"`

	// The part number of the chassis.
	PartNumber *string `json:"PartNumber,omitempty"`

	PhysicalSecurity ChassisV1250PhysicalSecurity `json:"PhysicalSecurity,omitempty"`

	Power OdataV4IdRef `json:"Power,omitempty"`

	PowerState ResourcePowerState `json:"PowerState,omitempty"`

	PowerSubsystem OdataV4IdRef `json:"PowerSubsystem,omitempty"`

	// Indicates that the chassis receives power from the containing chassis.
	PoweredByParent *bool `json:"PoweredByParent,omitempty"`

	Processors OdataV4IdRef `json:"Processors,omitempty"`

	// An indication of whether this component can be independently replaced as allowed by the vendor's replacement policy.
	Replaceable *bool `json:"Replaceable,omitempty"`

	// The SKU of the chassis.
	SKU *string `json:"SKU,omitempty"`

	Sensors OdataV4IdRef `json:"Sensors,omitempty"`

	// The serial number of the chassis.
	SerialNumber *string `json:"SerialNumber,omitempty"`

	// The spare part number of the chassis.
	SparePartNumber *string `json:"SparePartNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	Thermal OdataV4IdRef `json:"Thermal,omitempty"`

	ThermalDirection ChassisV1250ThermalDirection `json:"ThermalDirection,omitempty"`

	// Indicates that the chassis is thermally managed by the parent chassis.
	ThermalManagedByParent *bool `json:"ThermalManagedByParent,omitempty"`

	ThermalSubsystem OdataV4IdRef `json:"ThermalSubsystem,omitempty"`

	TrustedComponents OdataV4IdRef `json:"TrustedComponents,omitempty"`

	UUID string `json:"UUID,omitempty" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`

	// The hardware version of this chassis.
	Version *string `json:"Version,omitempty"`

	// The weight of the chassis.
	WeightKg *float32 `json:"WeightKg,omitempty"`

	// The width of the chassis.
	WidthMm *float32 `json:"WidthMm,omitempty"`
}

// AssertChassisV1250ChassisRequired checks if the required fields are not zero-ed
func AssertChassisV1250ChassisRequired(obj ChassisV1250Chassis) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"ChassisType": obj.ChassisType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertChassisV1250ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Assembly); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Certificates); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Controls); err != nil {
		return err
	}
	if err := AssertChassisV1250DoorsRequired(obj.Doors); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Drives); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.EnvironmentMetrics); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.FabricAdapters); err != nil {
		return err
	}
	if err := AssertChassisV1250LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.LogServices); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.MediaControllers); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Memory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.MemoryDomains); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.NetworkAdapters); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.PCIeDevices); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.PCIeSlots); err != nil {
		return err
	}
	if err := AssertChassisV1250PhysicalSecurityRequired(obj.PhysicalSecurity); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Power); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.PowerSubsystem); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Processors); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Sensors); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Thermal); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ThermalSubsystem); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.TrustedComponents); err != nil {
		return err
	}
	return nil
}

// AssertChassisV1250ChassisConstraints checks if the values respects the defined constraints
func AssertChassisV1250ChassisConstraints(obj ChassisV1250Chassis) error {
	if err := AssertChassisV1250ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Assembly); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Certificates); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Controls); err != nil {
		return err
	}
	if obj.DepthMm != nil && *obj.DepthMm < 0 {
		return &ParsingError{Param: "DepthMm", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertChassisV1250DoorsConstraints(obj.Doors); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Drives); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.EnvironmentMetrics); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.FabricAdapters); err != nil {
		return err
	}
	if obj.HeightMm != nil && *obj.HeightMm < 0 {
		return &ParsingError{Param: "HeightMm", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertChassisV1250LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationConstraints(obj.Location); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.LogServices); err != nil {
		return err
	}
	for _, el := range obj.Measurements {
		if err := AssertSoftwareInventoryMeasurementBlockConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefConstraints(obj.MediaControllers); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Memory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.MemoryDomains); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.NetworkAdapters); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.PCIeDevices); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.PCIeSlots); err != nil {
		return err
	}
	if err := AssertChassisV1250PhysicalSecurityConstraints(obj.PhysicalSecurity); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Power); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.PowerSubsystem); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Processors); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Sensors); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Thermal); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ThermalSubsystem); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.TrustedComponents); err != nil {
		return err
	}
	if obj.WeightKg != nil && *obj.WeightKg < 0 {
		return &ParsingError{Param: "WeightKg", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.WidthMm != nil && *obj.WidthMm < 0 {
		return &ParsingError{Param: "WidthMm", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}