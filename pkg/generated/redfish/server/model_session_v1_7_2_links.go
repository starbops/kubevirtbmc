// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// SessionV172Links - The links to other resources that are related to this resource.
type SessionV172Links struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	OutboundConnection OdataV4IdRef `json:"OutboundConnection,omitempty"`
}

// AssertSessionV172LinksRequired checks if the required fields are not zero-ed
func AssertSessionV172LinksRequired(obj SessionV172Links) error {
	if err := AssertOdataV4IdRefRequired(obj.OutboundConnection); err != nil {
		return err
	}
	return nil
}

// AssertSessionV172LinksConstraints checks if the values respects the defined constraints
func AssertSessionV172LinksConstraints(obj SessionV172Links) error {
	if err := AssertOdataV4IdRefConstraints(obj.OutboundConnection); err != nil {
		return err
	}
	return nil
}