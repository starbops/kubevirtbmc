// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// SessionV171Links - The links to other resources that are related to this resource.
type SessionV171Links struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	OutboundConnection OdataV4IdRef `json:"OutboundConnection,omitempty"`
}

// AssertSessionV171LinksRequired checks if the required fields are not zero-ed
func AssertSessionV171LinksRequired(obj SessionV171Links) error {
	if err := AssertOdataV4IdRefRequired(obj.OutboundConnection); err != nil {
		return err
	}
	return nil
}

// AssertSessionV171LinksConstraints checks if the values respects the defined constraints
func AssertSessionV171LinksConstraints(obj SessionV171Links) error {
	if err := AssertOdataV4IdRefConstraints(obj.OutboundConnection); err != nil {
		return err
	}
	return nil
}