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

// ProcessorMetricsV163ProcessorMetrics - The ProcessorMetrics schema contains usage and health statistics for a processor.
type ProcessorMetricsV163ProcessorMetrics struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ProcessorMetricsV163Actions `json:"Actions,omitempty"`

	// The average frequency of the processor.
	// Deprecated
	AverageFrequencyMHz *float32 `json:"AverageFrequencyMHz,omitempty"`

	// The bandwidth usage of this processor as a percentage.
	BandwidthPercent *float32 `json:"BandwidthPercent,omitempty"`

	// The processor cache metrics.
	Cache []ProcessorMetricsV163CacheMetrics `json:"Cache,omitempty"`

	CacheMetricsTotal ProcessorMetricsV163CacheMetricsTotal `json:"CacheMetricsTotal,omitempty"`

	// The power, in watt units, that the processor has consumed.
	// Deprecated
	ConsumedPowerWatt *float32 `json:"ConsumedPowerWatt,omitempty"`

	// The processor core metrics.
	CoreMetrics []ProcessorMetricsV163CoreMetrics `json:"CoreMetrics,omitempty"`

	CoreVoltage SensorSensorVoltageExcerpt `json:"CoreVoltage,omitempty"`

	// The number of correctable core errors.
	CorrectableCoreErrorCount *int64 `json:"CorrectableCoreErrorCount,omitempty"`

	// The number of correctable errors of all other components.
	CorrectableOtherErrorCount *int64 `json:"CorrectableOtherErrorCount,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The frequency relative to the nominal processor frequency ratio.
	FrequencyRatio *float32 `json:"FrequencyRatio,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The percentage of time spent in kernel mode.
	KernelPercent *float32 `json:"KernelPercent,omitempty"`

	// The local memory bandwidth usage in bytes.
	LocalMemoryBandwidthBytes *int64 `json:"LocalMemoryBandwidthBytes,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// Operating speed of the processor in MHz.
	OperatingSpeedMHz *int64 `json:"OperatingSpeedMHz,omitempty"`

	PCIeErrors PcieDevicePcieErrors `json:"PCIeErrors,omitempty"`

	// The total duration of throttling caused by a power limit of the processor since reset.
	PowerLimitThrottleDuration *string `json:"PowerLimitThrottleDuration,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The remote memory bandwidth usage in bytes.
	RemoteMemoryBandwidthBytes *int64 `json:"RemoteMemoryBandwidthBytes,omitempty"`

	// The temperature of the processor.
	// Deprecated
	TemperatureCelsius *float32 `json:"TemperatureCelsius,omitempty"`

	// The total duration of throttling caused by a thermal limit of the processor since reset.
	ThermalLimitThrottleDuration *string `json:"ThermalLimitThrottleDuration,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The CPU margin to throttle (temperature offset in degree Celsius units).
	ThrottlingCelsius *float32 `json:"ThrottlingCelsius,omitempty"`

	// The number of uncorrectable core errors.
	UncorrectableCoreErrorCount *int64 `json:"UncorrectableCoreErrorCount,omitempty"`

	// The number of uncorrectable errors of all other components.
	UncorrectableOtherErrorCount *int64 `json:"UncorrectableOtherErrorCount,omitempty"`

	// The percentage of time spent in user mode.
	UserPercent *float32 `json:"UserPercent,omitempty"`
}

// AssertProcessorMetricsV163ProcessorMetricsRequired checks if the required fields are not zero-ed
func AssertProcessorMetricsV163ProcessorMetricsRequired(obj ProcessorMetricsV163ProcessorMetrics) error {
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

	if err := AssertProcessorMetricsV163ActionsRequired(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.Cache {
		if err := AssertProcessorMetricsV163CacheMetricsRequired(el); err != nil {
			return err
		}
	}
	if err := AssertProcessorMetricsV163CacheMetricsTotalRequired(obj.CacheMetricsTotal); err != nil {
		return err
	}
	for _, el := range obj.CoreMetrics {
		if err := AssertProcessorMetricsV163CoreMetricsRequired(el); err != nil {
			return err
		}
	}
	if err := AssertSensorSensorVoltageExcerptRequired(obj.CoreVoltage); err != nil {
		return err
	}
	if err := AssertPcieDevicePcieErrorsRequired(obj.PCIeErrors); err != nil {
		return err
	}
	return nil
}

// AssertProcessorMetricsV163ProcessorMetricsConstraints checks if the values respects the defined constraints
func AssertProcessorMetricsV163ProcessorMetricsConstraints(obj ProcessorMetricsV163ProcessorMetrics) error {
	if err := AssertProcessorMetricsV163ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if obj.BandwidthPercent != nil && *obj.BandwidthPercent < 0 {
		return &ParsingError{Param: "BandwidthPercent", Err: errors.New(errMsgMinValueConstraint)}
	}
	for _, el := range obj.Cache {
		if err := AssertProcessorMetricsV163CacheMetricsConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertProcessorMetricsV163CacheMetricsTotalConstraints(obj.CacheMetricsTotal); err != nil {
		return err
	}
	for _, el := range obj.CoreMetrics {
		if err := AssertProcessorMetricsV163CoreMetricsConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertSensorSensorVoltageExcerptConstraints(obj.CoreVoltage); err != nil {
		return err
	}
	if obj.KernelPercent != nil && *obj.KernelPercent < 0 {
		return &ParsingError{Param: "KernelPercent", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertPcieDevicePcieErrorsConstraints(obj.PCIeErrors); err != nil {
		return err
	}
	if obj.UserPercent != nil && *obj.UserPercent < 0 {
		return &ParsingError{Param: "UserPercent", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
