// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// LogEntryV1170Links - The links to other resources that are related to this resource.
type LogEntryV1170Links struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	OriginOfCondition OdataV4IdRef `json:"OriginOfCondition,omitempty"`

	// An array of links to resources associated with this log entry.
	RelatedItem []OdataV4IdRef `json:"RelatedItem,omitempty"`

	// The number of items in a collection.
	RelatedItemodataCount int64 `json:"RelatedItem@odata.count,omitempty"`

	// An array of links to other log entries that are related to this log entry.
	RelatedLogEntries []OdataV4IdRef `json:"RelatedLogEntries,omitempty"`

	// The number of items in a collection.
	RelatedLogEntriesodataCount int64 `json:"RelatedLogEntries@odata.count,omitempty"`
}

// AssertLogEntryV1170LinksRequired checks if the required fields are not zero-ed
func AssertLogEntryV1170LinksRequired(obj LogEntryV1170Links) error {
	if err := AssertOdataV4IdRefRequired(obj.OriginOfCondition); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RelatedLogEntries {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertLogEntryV1170LinksConstraints checks if the values respects the defined constraints
func AssertLogEntryV1170LinksConstraints(obj LogEntryV1170Links) error {
	if err := AssertOdataV4IdRefConstraints(obj.OriginOfCondition); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RelatedLogEntries {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}