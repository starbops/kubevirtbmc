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

// EventServiceV1101SubmitTestEventRequestBody - This action generates a test event.
type EventServiceV1101SubmitTestEventRequestBody struct {

	// The group identifier for the event.
	EventGroupId int64 `json:"EventGroupId,omitempty"`

	// The ID for the event to add.
	EventId string `json:"EventId,omitempty"`

	// The date and time for the event to add.
	EventTimestamp time.Time `json:"EventTimestamp,omitempty"`

	// Deprecated
	EventType EventEventType `json:"EventType,omitempty"`

	// The human-readable message for the event to add.
	Message string `json:"Message,omitempty"`

	// An array of message arguments for the event to add.
	MessageArgs []string `json:"MessageArgs,omitempty"`

	// The MessageId for the event to add.
	MessageId string `json:"MessageId" validate:"regexp=^\\\\w+\\\\.\\\\d+\\\\.\\\\d+\\\\.\\\\w+$"`

	MessageSeverity ResourceHealth `json:"MessageSeverity,omitempty"`

	// The URL in the OriginOfCondition property of the event to add.  It is not a reference object.
	OriginOfCondition string `json:"OriginOfCondition,omitempty"`

	// The severity for the event to add.
	Severity string `json:"Severity,omitempty"`
}

// AssertEventServiceV1101SubmitTestEventRequestBodyRequired checks if the required fields are not zero-ed
func AssertEventServiceV1101SubmitTestEventRequestBodyRequired(obj EventServiceV1101SubmitTestEventRequestBody) error {
	elements := map[string]interface{}{
		"MessageId": obj.MessageId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertEventServiceV1101SubmitTestEventRequestBodyConstraints checks if the values respects the defined constraints
func AssertEventServiceV1101SubmitTestEventRequestBodyConstraints(obj EventServiceV1101SubmitTestEventRequestBody) error {
	return nil
}
