// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ManagerAccountV1120SnmpUserInfo - The SNMP settings for an account.
type ManagerAccountV1120SnmpUserInfo struct {

	// The secret authentication key for SNMPv3.
	AuthenticationKey *string `json:"AuthenticationKey,omitempty" validate:"regexp=(^[ !#-~]+$)|(^Passphrase:[ ^[ !#-~]+$)|(^Hex:[0-9A-Fa-f]{24,96})|(^\\\\*+$)"`

	// Indicates if the AuthenticationKey property is set.
	AuthenticationKeySet bool `json:"AuthenticationKeySet,omitempty"`

	AuthenticationProtocol ManagerAccountV1120SnmpAuthenticationProtocols `json:"AuthenticationProtocol,omitempty"`

	// The secret encryption key used in SNMPv3.
	EncryptionKey *string `json:"EncryptionKey,omitempty" validate:"regexp=(^[ !#-~]+$)|(^Passphrase:[ ^[ !#-~]+$)|(^Hex:[0-9A-Fa-f]{32})|(^\\\\*+$)"`

	// Indicates if the EncryptionKey property is set.
	EncryptionKeySet bool `json:"EncryptionKeySet,omitempty"`

	EncryptionProtocol ManagerAccountV1120SnmpEncryptionProtocols `json:"EncryptionProtocol,omitempty"`
}

// AssertManagerAccountV1120SnmpUserInfoRequired checks if the required fields are not zero-ed
func AssertManagerAccountV1120SnmpUserInfoRequired(obj ManagerAccountV1120SnmpUserInfo) error {
	return nil
}

// AssertManagerAccountV1120SnmpUserInfoConstraints checks if the values respects the defined constraints
func AssertManagerAccountV1120SnmpUserInfoConstraints(obj ManagerAccountV1120SnmpUserInfo) error {
	return nil
}
