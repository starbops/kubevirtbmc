// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ProcessorMetricsV163CurrentPeriod - The cache memory metrics since the last system reset or ClearCurrentPeriod action for a processor.
type ProcessorMetricsV163CurrentPeriod struct {

	// The number of correctable errors of cache memory since reset or ClearCurrentPeriod action for this processor.
	CorrectableECCErrorCount *int64 `json:"CorrectableECCErrorCount,omitempty"`

	// The number of uncorrectable errors of cache memory since reset or ClearCurrentPeriod action for this processor.
	UncorrectableECCErrorCount *int64 `json:"UncorrectableECCErrorCount,omitempty"`
}

// AssertProcessorMetricsV163CurrentPeriodRequired checks if the required fields are not zero-ed
func AssertProcessorMetricsV163CurrentPeriodRequired(obj ProcessorMetricsV163CurrentPeriod) error {
	return nil
}

// AssertProcessorMetricsV163CurrentPeriodConstraints checks if the values respects the defined constraints
func AssertProcessorMetricsV163CurrentPeriodConstraints(obj ProcessorMetricsV163CurrentPeriod) error {
	return nil
}