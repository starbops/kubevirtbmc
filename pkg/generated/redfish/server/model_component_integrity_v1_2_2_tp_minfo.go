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
	"errors"
)

// ComponentIntegrityV122TpMinfo - Integrity information about a Trusted Platform Module (TPM).
type ComponentIntegrityV122TpMinfo struct {
	ComponentCommunication ComponentIntegrityV122TpMcommunication `json:"ComponentCommunication,omitempty"`

	IdentityAuthentication ComponentIntegrityV122TpMauth `json:"IdentityAuthentication,omitempty"`

	MeasurementSet ComponentIntegrityV122TpMmeasurementSet `json:"MeasurementSet,omitempty"`

	// The maximum number of bytes that can be specified in the Nonce parameter of the TPMGetSignedMeasurements action.
	NonceSizeBytesMaximum *int64 `json:"NonceSizeBytesMaximum,omitempty"`
}

// AssertComponentIntegrityV122TpMinfoRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122TpMinfoRequired(obj ComponentIntegrityV122TpMinfo) error {
	if err := AssertComponentIntegrityV122TpMcommunicationRequired(obj.ComponentCommunication); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpMauthRequired(obj.IdentityAuthentication); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpMmeasurementSetRequired(obj.MeasurementSet); err != nil {
		return err
	}
	return nil
}

// AssertComponentIntegrityV122TpMinfoConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122TpMinfoConstraints(obj ComponentIntegrityV122TpMinfo) error {
	if err := AssertComponentIntegrityV122TpMcommunicationConstraints(obj.ComponentCommunication); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpMauthConstraints(obj.IdentityAuthentication); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpMmeasurementSetConstraints(obj.MeasurementSet); err != nil {
		return err
	}
	if obj.NonceSizeBytesMaximum != nil && *obj.NonceSizeBytesMaximum < 0 {
		return &ParsingError{Param: "NonceSizeBytesMaximum", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}