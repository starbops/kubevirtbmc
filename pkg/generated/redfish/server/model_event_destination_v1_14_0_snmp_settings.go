// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EventDestinationV1140SnmpSettings - Settings for an SNMP event destination.
type EventDestinationV1140SnmpSettings struct {

	// The secret authentication key for SNMPv3.
	AuthenticationKey *string `json:"AuthenticationKey,omitempty" validate:"regexp=(^[ !#-~]+$)|(^Passphrase:[ ^[ !#-~]+$)|(^Hex:[0-9A-Fa-f]{24,96})|(^\\\\*+$)"`

	// Indicates if the AuthenticationKey property is set.
	AuthenticationKeySet bool `json:"AuthenticationKeySet,omitempty"`

	AuthenticationProtocol EventDestinationV1140SnmpAuthenticationProtocols `json:"AuthenticationProtocol,omitempty"`

	// The secret authentication key for SNMPv3.
	EncryptionKey *string `json:"EncryptionKey,omitempty" validate:"regexp=(^[A-Za-z0-9]+$)|(^\\\\*+$)"`

	// Indicates if the EncryptionKey property is set.
	EncryptionKeySet bool `json:"EncryptionKeySet,omitempty"`

	EncryptionProtocol EventDestinationV1140SnmpEncryptionProtocols `json:"EncryptionProtocol,omitempty"`

	// The SNMP trap community string.
	TrapCommunity *string `json:"TrapCommunity,omitempty"`
}

// AssertEventDestinationV1140SnmpSettingsRequired checks if the required fields are not zero-ed
func AssertEventDestinationV1140SnmpSettingsRequired(obj EventDestinationV1140SnmpSettings) error {
	return nil
}

// AssertEventDestinationV1140SnmpSettingsConstraints checks if the values respects the defined constraints
func AssertEventDestinationV1140SnmpSettingsConstraints(obj EventDestinationV1140SnmpSettings) error {
	return nil
}
