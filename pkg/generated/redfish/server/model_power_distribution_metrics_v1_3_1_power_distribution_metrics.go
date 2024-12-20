// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PowerDistributionMetricsV131PowerDistributionMetrics - This is the schema definition for the metrics of a power distribution component or unit, such as a floor power distribution unit (PDU) or switchgear.
type PowerDistributionMetricsV131PowerDistributionMetrics struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	AbsoluteHumidity SensorSensorExcerpt `json:"AbsoluteHumidity,omitempty"`

	Actions PowerDistributionMetricsV131Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	EnergykWh SensorSensorEnergykWhExcerpt `json:"EnergykWh,omitempty"`

	HumidityPercent SensorSensorExcerpt `json:"HumidityPercent,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PowerLoadPercent SensorSensorExcerpt `json:"PowerLoadPercent,omitempty"`

	PowerWatts SensorSensorPowerExcerpt `json:"PowerWatts,omitempty"`

	TemperatureCelsius SensorSensorExcerpt `json:"TemperatureCelsius,omitempty"`
}

// AssertPowerDistributionMetricsV131PowerDistributionMetricsRequired checks if the required fields are not zero-ed
func AssertPowerDistributionMetricsV131PowerDistributionMetricsRequired(obj PowerDistributionMetricsV131PowerDistributionMetrics) error {
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

	if err := AssertSensorSensorExcerptRequired(obj.AbsoluteHumidity); err != nil {
		return err
	}
	if err := AssertPowerDistributionMetricsV131ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertSensorSensorEnergykWhExcerptRequired(obj.EnergykWh); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.HumidityPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.PowerLoadPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorPowerExcerptRequired(obj.PowerWatts); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptRequired(obj.TemperatureCelsius); err != nil {
		return err
	}
	return nil
}

// AssertPowerDistributionMetricsV131PowerDistributionMetricsConstraints checks if the values respects the defined constraints
func AssertPowerDistributionMetricsV131PowerDistributionMetricsConstraints(obj PowerDistributionMetricsV131PowerDistributionMetrics) error {
	if err := AssertSensorSensorExcerptConstraints(obj.AbsoluteHumidity); err != nil {
		return err
	}
	if err := AssertPowerDistributionMetricsV131ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertSensorSensorEnergykWhExcerptConstraints(obj.EnergykWh); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.HumidityPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.PowerLoadPercent); err != nil {
		return err
	}
	if err := AssertSensorSensorPowerExcerptConstraints(obj.PowerWatts); err != nil {
		return err
	}
	if err := AssertSensorSensorExcerptConstraints(obj.TemperatureCelsius); err != nil {
		return err
	}
	return nil
}
