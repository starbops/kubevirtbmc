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

// MetricReportDefinitionV145MetricReportDefinition - The MetricReportDefinition schema describes set of metrics that are collected into a metric report.
type MetricReportDefinitionV145MetricReportDefinition struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions MetricReportDefinitionV145Actions `json:"Actions,omitempty"`

	// The maximum number of entries that can be appended to a metric report.  When the metric report reaches its limit, its behavior is dictated by the ReportUpdates property.
	AppendLimit int64 `json:"AppendLimit,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Links MetricReportDefinitionV145Links `json:"Links,omitempty"`

	// The list of URIs with wildcards and property identifiers to include in the metric report.  If a URI has wildcards, the wildcards are substituted as specified in the Wildcards property.
	MetricProperties []*string `json:"MetricProperties,omitempty"`

	MetricReport OdataV4IdRef `json:"MetricReport,omitempty"`

	// An indication of whether the generation of new metric reports is enabled.
	MetricReportDefinitionEnabled *bool `json:"MetricReportDefinitionEnabled,omitempty"`

	MetricReportDefinitionType MetricReportDefinitionV145MetricReportDefinitionType `json:"MetricReportDefinitionType,omitempty"`

	// The interval at which to send the complete metric report because the Redfish client wants refreshed metric data even when the data has not changed.  This property value is always greater than the recurrence interval of a metric report, and it only applies when the SuppressRepeatedMetricValue property is `true`.
	MetricReportHeartbeatInterval *string `json:"MetricReportHeartbeatInterval,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The list of metrics to include in the metric report.  The metrics might include calculations to apply to metric properties.
	Metrics []MetricReportDefinitionV145Metric `json:"Metrics,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The set of actions to perform when a metric report is generated.
	ReportActions []MetricReportDefinitionV145ReportActionsEnum `json:"ReportActions,omitempty"`

	// The maximum timespan that a metric report can cover.
	ReportTimespan *string `json:"ReportTimespan,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	ReportUpdates MetricReportDefinitionV145ReportUpdatesEnum `json:"ReportUpdates,omitempty"`

	Schedule ScheduleSchedule `json:"Schedule,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// An indication of whether any metrics are suppressed from the generated metric report.  If `true`, any metric that equals the same value in the previously generated metric report is suppressed from the current report.  Also, duplicate metrics are suppressed.  If `false`, no metrics are suppressed from the current report.  The current report might contain no metrics if all metrics equal the values in the previously generated metric report.
	SuppressRepeatedMetricValue *bool `json:"SuppressRepeatedMetricValue,omitempty"`

	// The set of wildcards and their substitution values for the entries in the MetricProperties property.
	Wildcards []MetricReportDefinitionV145Wildcard `json:"Wildcards,omitempty"`
}

// AssertMetricReportDefinitionV145MetricReportDefinitionRequired checks if the required fields are not zero-ed
func AssertMetricReportDefinitionV145MetricReportDefinitionRequired(obj MetricReportDefinitionV145MetricReportDefinition) error {
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

	if err := AssertMetricReportDefinitionV145ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertMetricReportDefinitionV145LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.MetricReport); err != nil {
		return err
	}
	for _, el := range obj.Metrics {
		if err := AssertMetricReportDefinitionV145MetricRequired(el); err != nil {
			return err
		}
	}
	if err := AssertScheduleScheduleRequired(obj.Schedule); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	for _, el := range obj.Wildcards {
		if err := AssertMetricReportDefinitionV145WildcardRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMetricReportDefinitionV145MetricReportDefinitionConstraints checks if the values respects the defined constraints
func AssertMetricReportDefinitionV145MetricReportDefinitionConstraints(obj MetricReportDefinitionV145MetricReportDefinition) error {
	if err := AssertMetricReportDefinitionV145ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if obj.AppendLimit < 0 {
		return &ParsingError{Param: "AppendLimit", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertMetricReportDefinitionV145LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.MetricReport); err != nil {
		return err
	}
	for _, el := range obj.Metrics {
		if err := AssertMetricReportDefinitionV145MetricConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertScheduleScheduleConstraints(obj.Schedule); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	for _, el := range obj.Wildcards {
		if err := AssertMetricReportDefinitionV145WildcardConstraints(el); err != nil {
			return err
		}
	}
	return nil
}