// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ExternalAccountProviderV171OAuth2Service - Various settings to parse an OAuth 2.0 service.
type ExternalAccountProviderV171OAuth2Service struct {

	// The allowable audience strings of the Redfish service.
	Audience []string `json:"Audience,omitempty"`

	// The issuer string of the OAuth 2.0 service.  Clients should configure this property if Mode contains `Offline`.
	Issuer *string `json:"Issuer,omitempty"`

	Mode ExternalAccountProviderV171OAuth2Mode `json:"Mode,omitempty"`

	// The Base64-encoded signing keys of the issuer of the OAuth 2.0 service.  Clients should configure this property if Mode contains `Offline`.
	OAuthServiceSigningKeys *string `json:"OAuthServiceSigningKeys,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertExternalAccountProviderV171OAuth2ServiceRequired checks if the required fields are not zero-ed
func AssertExternalAccountProviderV171OAuth2ServiceRequired(obj ExternalAccountProviderV171OAuth2Service) error {
	return nil
}

// AssertExternalAccountProviderV171OAuth2ServiceConstraints checks if the values respects the defined constraints
func AssertExternalAccountProviderV171OAuth2ServiceConstraints(obj ExternalAccountProviderV171OAuth2Service) error {
	return nil
}