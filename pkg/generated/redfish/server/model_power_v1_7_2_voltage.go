// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

type PowerV172Voltage struct {

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	Actions PowerV172VoltageActions `json:"Actions,omitempty"`

	// The value at which the reading is below normal range but not yet fatal.
	LowerThresholdCritical *float32 `json:"LowerThresholdCritical,omitempty"`

	// The value at which the reading is below normal range and fatal.
	LowerThresholdFatal *float32 `json:"LowerThresholdFatal,omitempty"`

	// The value at which the reading is below normal range.
	LowerThresholdNonCritical *float32 `json:"LowerThresholdNonCritical,omitempty"`

	// Maximum value for this sensor.
	MaxReadingRange *float32 `json:"MaxReadingRange,omitempty"`

	// The unique identifier for the member within an array.
	MemberId string `json:"MemberId"`

	// Minimum value for this sensor.
	MinReadingRange *float32 `json:"MinReadingRange,omitempty"`

	// Voltage sensor name.
	Name *string `json:"Name,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	// The reading of the voltage sensor.
	ReadingVolts *float32 `json:"ReadingVolts,omitempty"`

	// An array of links to resources or objects to which this voltage measurement applies.
	RelatedItem []OdataV4IdRef `json:"RelatedItem,omitempty"`

	// The number of items in a collection.
	RelatedItemodataCount int64 `json:"RelatedItem@odata.count,omitempty"`

	// A numerical identifier to represent the voltage sensor.
	SensorNumber *int64 `json:"SensorNumber,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The value at which the reading is above normal range but not yet fatal.
	UpperThresholdCritical *float32 `json:"UpperThresholdCritical,omitempty"`

	// The value at which the reading is above normal range and fatal.
	UpperThresholdFatal *float32 `json:"UpperThresholdFatal,omitempty"`

	// The value at which the reading is above normal range.
	UpperThresholdNonCritical *float32 `json:"UpperThresholdNonCritical,omitempty"`
}

// AssertPowerV172VoltageRequired checks if the required fields are not zero-ed
func AssertPowerV172VoltageRequired(obj PowerV172Voltage) error {
	elements := map[string]interface{}{
		"@odata.id": obj.OdataId,
		"MemberId":  obj.MemberId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPowerV172VoltageActionsRequired(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertPowerV172VoltageConstraints checks if the values respects the defined constraints
func AssertPowerV172VoltageConstraints(obj PowerV172Voltage) error {
	if err := AssertPowerV172VoltageActionsConstraints(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}
