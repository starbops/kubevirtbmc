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
	"time"
)

// MetricReportV150MetricReport - The MetricReport schema represents a set of collected metrics.
type MetricReportV150MetricReport struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions MetricReportV150Actions `json:"Actions,omitempty"`

	// A context can be supplied at subscription time.  This property is the context value supplied by the subscriber.
	Context string `json:"Context,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	MetricReportDefinition OdataV4IdRef `json:"MetricReportDefinition,omitempty"`

	// An array of metric values for the metered items of this metric report.
	MetricValues []MetricReportV150MetricValue `json:"MetricValues,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The current sequence identifier for this metric report.
	// Deprecated
	ReportSequence string `json:"ReportSequence,omitempty"`

	// The time associated with the metric report in its entirety.  The time of the metric report can be relevant when the time of individual metrics are minimally different.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
}

// AssertMetricReportV150MetricReportRequired checks if the required fields are not zero-ed
func AssertMetricReportV150MetricReportRequired(obj MetricReportV150MetricReport) error {
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

	if err := AssertMetricReportV150ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.MetricReportDefinition); err != nil {
		return err
	}
	for _, el := range obj.MetricValues {
		if err := AssertMetricReportV150MetricValueRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMetricReportV150MetricReportConstraints checks if the values respects the defined constraints
func AssertMetricReportV150MetricReportConstraints(obj MetricReportV150MetricReport) error {
	if err := AssertMetricReportV150ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.MetricReportDefinition); err != nil {
		return err
	}
	for _, el := range obj.MetricValues {
		if err := AssertMetricReportV150MetricValueConstraints(el); err != nil {
			return err
		}
	}
	return nil
}