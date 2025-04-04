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
	"time"
)

// SensorV181Sensor - The Sensor schema describes a sensor and its properties.
type SensorV181Sensor struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The estimated percent error of measured versus actual values.
	// Deprecated
	Accuracy *float32 `json:"Accuracy,omitempty"`

	Actions SensorV181Actions `json:"Actions,omitempty"`

	// The adjusted maximum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMaxAllowableOperatingValue *float32 `json:"AdjustedMaxAllowableOperatingValue,omitempty"`

	// The adjusted minimum allowable operating value for this equipment based on the environmental conditions.
	AdjustedMinAllowableOperatingValue *float32 `json:"AdjustedMinAllowableOperatingValue,omitempty"`

	// The product of voltage and current for an AC circuit, in volt-ampere units.
	ApparentVA *float32 `json:"ApparentVA,omitempty"`

	// Apparent energy (kVAh).
	ApparentkVAh *float32 `json:"ApparentkVAh,omitempty"`

	// The average sensor value.
	AverageReading *float32 `json:"AverageReading,omitempty"`

	// The interval over which the average sensor value is calculated.
	AveragingInterval *string `json:"AveragingInterval,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// Indicates that enough readings were collected to calculate the average sensor reading over the averaging interval time.
	AveragingIntervalAchieved *bool `json:"AveragingIntervalAchieved,omitempty"`

	// The calibration offset applied to the Reading.
	Calibration *float32 `json:"Calibration,omitempty"`

	// The date and time that the sensor was last calibrated.
	CalibrationTime *time.Time `json:"CalibrationTime,omitempty"`

	// The crest factor for this sensor.
	CrestFactor *float32 `json:"CrestFactor,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	ElectricalContext SensorElectricalContext `json:"ElectricalContext,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Implementation SensorV181ImplementationType `json:"Implementation,omitempty"`

	// The total accumulation value for this sensor.
	LifetimeReading *float32 `json:"LifetimeReading,omitempty"`

	Links SensorV181Links `json:"Links,omitempty"`

	// The power load utilization for this sensor.
	// Deprecated
	LoadPercent *float32 `json:"LoadPercent,omitempty"`

	Location ResourceLocation `json:"Location,omitempty"`

	// The lowest sensor value.
	LowestReading *float32 `json:"LowestReading,omitempty"`

	// The time when the lowest sensor value occurred.
	LowestReadingTime *time.Time `json:"LowestReadingTime,omitempty"`

	// The maximum allowable operating value for this equipment.
	MaxAllowableOperatingValue *float32 `json:"MaxAllowableOperatingValue,omitempty"`

	// The minimum allowable operating value for this equipment.
	MinAllowableOperatingValue *float32 `json:"MinAllowableOperatingValue,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The peak sensor value.
	PeakReading *float32 `json:"PeakReading,omitempty"`

	// The time when the peak sensor value occurred.
	PeakReadingTime *time.Time `json:"PeakReadingTime,omitempty"`

	// The phase angle (degrees) between the current and voltage waveforms.
	PhaseAngleDegrees *float32 `json:"PhaseAngleDegrees,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	PhysicalSubContext PhysicalContextPhysicalSubContext `json:"PhysicalSubContext,omitempty"`

	// The power factor for this sensor.
	PowerFactor *float32 `json:"PowerFactor,omitempty"`

	// The number of significant digits in the reading.
	Precision *float32 `json:"Precision,omitempty"`

	// The square root of the difference term of squared apparent VA and squared power (Reading) for a circuit, in VAR units.
	ReactiveVAR *float32 `json:"ReactiveVAR,omitempty"`

	// Reactive energy (kVARh).
	ReactivekVARh *float32 `json:"ReactivekVARh,omitempty"`

	// The sensor value.
	Reading *float32 `json:"Reading,omitempty"`

	// Accuracy (+/-) of the reading.
	ReadingAccuracy *float32 `json:"ReadingAccuracy,omitempty"`

	ReadingBasis SensorV181ReadingBasisType `json:"ReadingBasis,omitempty"`

	// The maximum possible value for this sensor.
	ReadingRangeMax *float32 `json:"ReadingRangeMax,omitempty"`

	// The minimum possible value for this sensor.
	ReadingRangeMin *float32 `json:"ReadingRangeMin,omitempty"`

	// The date and time that the reading was acquired from the sensor.
	ReadingTime *time.Time `json:"ReadingTime,omitempty"`

	ReadingType SensorV181ReadingType `json:"ReadingType,omitempty"`

	// The units of the reading and thresholds.
	ReadingUnits *string `json:"ReadingUnits,omitempty"`

	// An array of links to resources or objects that this sensor services.
	RelatedItem []OdataV4IdRef `json:"RelatedItem,omitempty"`

	// The number of items in a collection.
	RelatedItemodataCount int64 `json:"RelatedItem@odata.count,omitempty"`

	// The time interval between readings of the physical sensor.
	// Deprecated
	SensingFrequency *float32 `json:"SensingFrequency,omitempty"`

	// The time interval between readings of the sensor.
	SensingInterval *string `json:"SensingInterval,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	SensorGroup RedundancyRedundantGroup `json:"SensorGroup,omitempty"`

	// The date and time when the time-based properties were last reset.
	SensorResetTime *time.Time `json:"SensorResetTime,omitempty"`

	// The rotational speed.
	SpeedRPM *float32 `json:"SpeedRPM,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The total harmonic distortion percent (% THD).
	THDPercent *float32 `json:"THDPercent,omitempty"`

	Thresholds SensorV181Thresholds `json:"Thresholds,omitempty"`

	VoltageType SensorVoltageType `json:"VoltageType,omitempty"`
}

// AssertSensorV181SensorRequired checks if the required fields are not zero-ed
func AssertSensorV181SensorRequired(obj SensorV181Sensor) error {
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

	if err := AssertSensorV181ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertSensorV181LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationRequired(obj.Location); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertRedundancyRedundantGroupRequired(obj.SensorGroup); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertSensorV181ThresholdsRequired(obj.Thresholds); err != nil {
		return err
	}
	return nil
}

// AssertSensorV181SensorConstraints checks if the values respects the defined constraints
func AssertSensorV181SensorConstraints(obj SensorV181Sensor) error {
	if err := AssertSensorV181ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertSensorV181LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceLocationConstraints(obj.Location); err != nil {
		return err
	}
	if obj.PhaseAngleDegrees != nil && *obj.PhaseAngleDegrees < -90 {
		return &ParsingError{Param: "PhaseAngleDegrees", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PhaseAngleDegrees != nil && *obj.PhaseAngleDegrees > 90 {
		return &ParsingError{Param: "PhaseAngleDegrees", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.PowerFactor != nil && *obj.PowerFactor < -1 {
		return &ParsingError{Param: "PowerFactor", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PowerFactor != nil && *obj.PowerFactor > 1 {
		return &ParsingError{Param: "PowerFactor", Err: errors.New(errMsgMaxValueConstraint)}
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertRedundancyRedundantGroupConstraints(obj.SensorGroup); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	if obj.THDPercent != nil && *obj.THDPercent < 0 {
		return &ParsingError{Param: "THDPercent", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertSensorV181ThresholdsConstraints(obj.Thresholds); err != nil {
		return err
	}
	return nil
}
