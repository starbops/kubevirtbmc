// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

type ManifestManifestStanzasInner struct {

	// The OEM-defined type of stanza.
	OEMStanzaType *string `json:"OEMStanzaType,omitempty"`

	// The content of the request for the stanza.
	Request map[string]interface{} `json:"Request,omitempty"`

	// The content of the response for the stanza.
	Response map[string]interface{} `json:"Response,omitempty"`

	// The identifier of the stanza.  This is a unique identifier specified by the client and is not used by the service.
	StanzaId *string `json:"StanzaId,omitempty"`

	StanzaType ManifestV112StanzaType `json:"StanzaType,omitempty"`
}

// AssertManifestManifestStanzasInnerRequired checks if the required fields are not zero-ed
func AssertManifestManifestStanzasInnerRequired(obj ManifestManifestStanzasInner) error {
	return nil
}

// AssertManifestManifestStanzasInnerConstraints checks if the values respects the defined constraints
func AssertManifestManifestStanzasInnerConstraints(obj ManifestManifestStanzasInner) error {
	return nil
}
