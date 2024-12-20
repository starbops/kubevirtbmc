// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MetricDefinitionV133MetricDefinition - The MetricDefinition schema describes the metadata information for a metric.
type MetricDefinitionV133MetricDefinition struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The estimated percent error of measured versus actual values.
	Accuracy *float32 `json:"Accuracy,omitempty"`

	Actions MetricDefinitionV133Actions `json:"Actions,omitempty"`

	Calculable MetricDefinitionV133Calculable `json:"Calculable,omitempty"`

	CalculationAlgorithm MetricDefinitionV133CalculationAlgorithmEnum `json:"CalculationAlgorithm,omitempty"`

	// The metric properties that are part of a calculation that this metric definition defines.
	CalculationParameters []MetricDefinitionV133CalculationParamsType `json:"CalculationParameters,omitempty"`

	// The time interval over which the metric calculation is performed.
	CalculationTimeInterval *string `json:"CalculationTimeInterval,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The calibration offset added to the metric reading.
	Calibration *float32 `json:"Calibration,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// This array property specifies possible values of a discrete metric.
	DiscreteValues []*string `json:"DiscreteValues,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Implementation MetricDefinitionV133ImplementationType `json:"Implementation,omitempty"`

	// An indication of whether the metric values are linear versus non-linear.
	IsLinear *bool `json:"IsLinear,omitempty"`

	// The logical contexts related to the metric.
	LogicalContexts []PhysicalContextLogicalContext `json:"LogicalContexts,omitempty"`

	// Maximum value for metric reading.
	MaxReadingRange *float32 `json:"MaxReadingRange,omitempty"`

	MetricDataType MetricDefinitionV133MetricDataType `json:"MetricDataType,omitempty"`

	// The list of URIs with wildcards and property identifiers that this metric definition defines.  If a URI has wildcards, the wildcards are substituted as specified in the Wildcards property.
	MetricProperties []*string `json:"MetricProperties,omitempty"`

	MetricType MetricDefinitionV133MetricType `json:"MetricType,omitempty"`

	// Minimum value for metric reading.
	MinReadingRange *float32 `json:"MinReadingRange,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM-defined calculation that is performed on a source metric to obtain the metric being defined.
	OEMCalculationAlgorithm *string `json:"OEMCalculationAlgorithm,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PhysicalContext PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	// Number of significant digits in the metric reading.
	Precision *int64 `json:"Precision,omitempty"`

	// The time interval between when a metric is updated.
	SensingInterval *string `json:"SensingInterval,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The accuracy of the timestamp.
	TimestampAccuracy *string `json:"TimestampAccuracy,omitempty" validate:"regexp=^P(\\\\d+D)?(T(\\\\d+H)?(\\\\d+M)?(\\\\d+(.\\\\d+)?S)?)?$"`

	// The units of measure for this metric.
	Units *string `json:"Units,omitempty"`

	// The wildcards and their substitution values for the entries in the MetricProperties array property.
	Wildcards []MetricDefinitionV133Wildcard `json:"Wildcards,omitempty"`
}

// AssertMetricDefinitionV133MetricDefinitionRequired checks if the required fields are not zero-ed
func AssertMetricDefinitionV133MetricDefinitionRequired(obj MetricDefinitionV133MetricDefinition) error {
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

	if err := AssertMetricDefinitionV133ActionsRequired(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.CalculationParameters {
		if err := AssertMetricDefinitionV133CalculationParamsTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Wildcards {
		if err := AssertMetricDefinitionV133WildcardRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMetricDefinitionV133MetricDefinitionConstraints checks if the values respects the defined constraints
func AssertMetricDefinitionV133MetricDefinitionConstraints(obj MetricDefinitionV133MetricDefinition) error {
	if err := AssertMetricDefinitionV133ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.CalculationParameters {
		if err := AssertMetricDefinitionV133CalculationParamsTypeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Wildcards {
		if err := AssertMetricDefinitionV133WildcardConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
