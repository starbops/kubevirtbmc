// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// LogEntryV1170Cper - Details for a CPER section or record associated with a log entry.
type LogEntryV1170Cper struct {

	// The CPER Notification Type for a CPER record.
	NotificationType *string `json:"NotificationType,omitempty" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The CPER Section Type.
	SectionType *string `json:"SectionType,omitempty" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
}

// AssertLogEntryV1170CperRequired checks if the required fields are not zero-ed
func AssertLogEntryV1170CperRequired(obj LogEntryV1170Cper) error {
	return nil
}

// AssertLogEntryV1170CperConstraints checks if the values respects the defined constraints
func AssertLogEntryV1170CperConstraints(obj LogEntryV1170Cper) error {
	return nil
}