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
	"fmt"
)

// MetricDefinitionV133MetricType : The type of metric.  Provides information to the client on how the metric can be handled.
type MetricDefinitionV133MetricType string

// List of MetricDefinitionV133MetricType
const (
	METRICDEFINITIONV133METRICTYPE_NUMERIC   MetricDefinitionV133MetricType = "Numeric"
	METRICDEFINITIONV133METRICTYPE_DISCRETE  MetricDefinitionV133MetricType = "Discrete"
	METRICDEFINITIONV133METRICTYPE_GAUGE     MetricDefinitionV133MetricType = "Gauge"
	METRICDEFINITIONV133METRICTYPE_COUNTER   MetricDefinitionV133MetricType = "Counter"
	METRICDEFINITIONV133METRICTYPE_COUNTDOWN MetricDefinitionV133MetricType = "Countdown"
	METRICDEFINITIONV133METRICTYPE_STRING    MetricDefinitionV133MetricType = "String"
)

// AllowedMetricDefinitionV133MetricTypeEnumValues is all the allowed values of MetricDefinitionV133MetricType enum
var AllowedMetricDefinitionV133MetricTypeEnumValues = []MetricDefinitionV133MetricType{
	"Numeric",
	"Discrete",
	"Gauge",
	"Counter",
	"Countdown",
	"String",
}

// validMetricDefinitionV133MetricTypeEnumValue provides a map of MetricDefinitionV133MetricTypes for fast verification of use input
var validMetricDefinitionV133MetricTypeEnumValues = map[MetricDefinitionV133MetricType]struct{}{
	"Numeric":   {},
	"Discrete":  {},
	"Gauge":     {},
	"Counter":   {},
	"Countdown": {},
	"String":    {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricDefinitionV133MetricType) IsValid() bool {
	_, ok := validMetricDefinitionV133MetricTypeEnumValues[v]
	return ok
}

// NewMetricDefinitionV133MetricTypeFromValue returns a pointer to a valid MetricDefinitionV133MetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricDefinitionV133MetricTypeFromValue(v string) (MetricDefinitionV133MetricType, error) {
	ev := MetricDefinitionV133MetricType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for MetricDefinitionV133MetricType: valid values are %v", v, AllowedMetricDefinitionV133MetricTypeEnumValues)
}

// AssertMetricDefinitionV133MetricTypeRequired checks if the required fields are not zero-ed
func AssertMetricDefinitionV133MetricTypeRequired(obj MetricDefinitionV133MetricType) error {
	return nil
}

// AssertMetricDefinitionV133MetricTypeConstraints checks if the values respects the defined constraints
func AssertMetricDefinitionV133MetricTypeConstraints(obj MetricDefinitionV133MetricType) error {
	return nil
}
