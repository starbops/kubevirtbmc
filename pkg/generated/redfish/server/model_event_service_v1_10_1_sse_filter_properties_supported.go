// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EventServiceV1101SseFilterPropertiesSupported - The set of properties that are supported in the `$filter` query parameter for the ServerSentEventUri.
type EventServiceV1101SseFilterPropertiesSupported struct {

	// An indication of whether the service supports filtering by the EventFormatType property.
	EventFormatType bool `json:"EventFormatType,omitempty"`

	// An indication of whether the service supports filtering by the EventTypes property.
	// Deprecated
	EventType bool `json:"EventType,omitempty"`

	// An indication of whether the service supports filtering by the MessageIds property.
	MessageId bool `json:"MessageId,omitempty"`

	// An indication of whether the service supports filtering by the MetricReportDefinitions property.
	MetricReportDefinition bool `json:"MetricReportDefinition,omitempty"`

	// An indication of whether the service supports filtering by the OriginResources property.
	OriginResource bool `json:"OriginResource,omitempty"`

	// An indication of whether the service supports filtering by the RegistryPrefixes property.
	RegistryPrefix bool `json:"RegistryPrefix,omitempty"`

	// An indication of whether the service supports filtering by the ResourceTypes property.
	ResourceType bool `json:"ResourceType,omitempty"`

	// An indication of whether the service supports filtering by the SubordinateResources property.
	SubordinateResources bool `json:"SubordinateResources,omitempty"`
}

// AssertEventServiceV1101SseFilterPropertiesSupportedRequired checks if the required fields are not zero-ed
func AssertEventServiceV1101SseFilterPropertiesSupportedRequired(obj EventServiceV1101SseFilterPropertiesSupported) error {
	return nil
}

// AssertEventServiceV1101SseFilterPropertiesSupportedConstraints checks if the values respects the defined constraints
func AssertEventServiceV1101SseFilterPropertiesSupportedConstraints(obj EventServiceV1101SseFilterPropertiesSupported) error {
	return nil
}