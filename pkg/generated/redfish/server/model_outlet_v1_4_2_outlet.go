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

// OutletV142Outlet - The Outlet schema contains a definition for an electrical outlet.
type OutletV142Outlet struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions OutletV142Actions `json:"Actions,omitempty"`

	// Indicates whether the configuration is locked.
	ConfigurationLocked bool `json:"ConfigurationLocked,omitempty"`

	CurrentAmps SensorSensorCurrentExcerpt `json:"CurrentAmps,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of names of downstream devices that are powered by this outlet.
	ElectricalConsumerNames []*string `json:"ElectricalConsumerNames,omitempty"`

	ElectricalContext SensorElectricalContext `json:"ElectricalContext,omitempty"`

	EnergykWh SensorSensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	FrequencyHz SensorSensorExcerpt `json:"FrequencyHz,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	IndicatorLED ResourceIndicatorLed `json:"IndicatorLED,omitempty"`

	Links OutletV142Links `json:"Links,omitempty"`

	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive *bool `json:"LocationIndicatorActive,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	NominalVoltage CircuitNominalVoltageType `json:"NominalVoltage,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	OutletType OutletReceptacleType `json:"OutletType,omitempty"`

	PhaseWiringType CircuitPhaseWiringType `json:"PhaseWiringType,omitempty"`

	PolyPhaseCurrentAmps OutletV142CurrentSensors `json:"PolyPhaseCurrentAmps,omitempty"`

	PolyPhaseVoltage OutletV142VoltageSensors `json:"PolyPhaseVoltage,omitempty"`

	// Indicates whether power control requests are locked.
	PowerControlLocked bool `json:"PowerControlLocked,omitempty"`

	// The number of seconds to delay power on after a PowerControl action to cycle power.  Zero seconds indicates no delay.
	PowerCycleDelaySeconds *float32 `json:"PowerCycleDelaySeconds,omitempty"`

	// Indicates if the outlet can be powered.
	PowerEnabled *bool `json:"PowerEnabled,omitempty"`

	PowerLoadPercent SensorSensorExcerpt `json:"PowerLoadPercent,omitempty"`

	// The number of seconds to delay power off after a PowerControl action.  Zero seconds indicates no delay to power off.
	PowerOffDelaySeconds *float32 `json:"PowerOffDelaySeconds,omitempty"`

	// The number of seconds to delay power up after a power cycle or a PowerControl action.  Zero seconds indicates no delay to power up.
	PowerOnDelaySeconds *float32 `json:"PowerOnDelaySeconds,omitempty"`

	// The number of seconds to delay power on after power has been restored.  Zero seconds indicates no delay.
	PowerRestoreDelaySeconds *float32 `json:"PowerRestoreDelaySeconds,omitempty"`

	PowerRestorePolicy CircuitPowerRestorePolicyTypes `json:"PowerRestorePolicy,omitempty"`

	PowerState ResourcePowerState `json:"PowerState,omitempty"`

	// Indicates whether the power state is undergoing a delayed transition.
	PowerStateInTransition bool `json:"PowerStateInTransition,omitempty"`

	PowerWatts SensorSensorPowerExcerpt `json:"PowerWatts,omitempty"`

	// The rated maximum current allowed for this outlet.
	RatedCurrentAmps *float32 `json:"RatedCurrentAmps,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// A user-assigned label.
	UserLabel string `json:"UserLabel,omitempty"`

	Voltage SensorSensorVoltageExcerpt `json:"Voltage,omitempty"`

	VoltageType OutletV142VoltageType `json:"VoltageType,omitempty"`
}

// AssertOutletV142OutletRequired checks if the required fields are not zero-ed
func AssertOutletV142OutletRequired(obj OutletV142Outlet) error {
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

	if err := AssertOutletV142ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertSensorSensorCurrentExcerptRequired(obj.CurrentAmps); err != nil {
		return err
	}
	if err := AssertSensorSensorEnergykWhExcerptRequired(obj.EnergykWh); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.FrequencyHz); err != nil {
		return err
	}
	if err := AssertOutletV142LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertOutletV142CurrentSensorsRequired(obj.PolyPhaseCurrentAmps); err != nil {
		return err
	}
	if err := AssertOutletV142VoltageSensorsRequired(obj.PolyPhaseVoltage); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.PowerLoadPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorPowerExcerptRequired(obj.PowerWatts); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertSensorSensorVoltageExcerptRequired(obj.Voltage); err != nil {
		return err
	}
	return nil
}

// AssertOutletV142OutletConstraints checks if the values respects the defined constraints
func AssertOutletV142OutletConstraints(obj OutletV142Outlet) error {
	if err := AssertOutletV142ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertSensorSensorCurrentExcerptConstraints(obj.CurrentAmps); err != nil {
		return err
	}
	if err := AssertSensorSensorEnergykWhExcerptConstraints(obj.EnergykWh); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.FrequencyHz); err != nil {
		return err
	}
	if err := AssertOutletV142LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertOutletV142CurrentSensorsConstraints(obj.PolyPhaseCurrentAmps); err != nil {
		return err
	}
	if err := AssertOutletV142VoltageSensorsConstraints(obj.PolyPhaseVoltage); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.PowerLoadPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorPowerExcerptConstraints(obj.PowerWatts); err != nil {
		return err
	}
	if obj.RatedCurrentAmps != nil && *obj.RatedCurrentAmps < 0 {
		return &ParsingError{Param: "RatedCurrentAmps", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	if err := AssertSensorSensorVoltageExcerptConstraints(obj.Voltage); err != nil {
		return err
	}
	return nil
}