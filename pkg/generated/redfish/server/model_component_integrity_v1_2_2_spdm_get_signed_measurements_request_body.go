// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ComponentIntegrityV122SpdmGetSignedMeasurementsRequestBody - This action generates an SPDM cryptographic signed statement over the given nonce and measurements of the SPDM Responder.
type ComponentIntegrityV122SpdmGetSignedMeasurementsRequestBody struct {

	// An array of indices that identify the measurement blocks to sign.
	MeasurementIndices []int64 `json:"MeasurementIndices,omitempty"`

	// A 32-byte hex-encoded string that is signed with the measurements.  The value should be unique.
	Nonce string `json:"Nonce,omitempty" validate:"regexp=^[0-9a-fA-F]{64}$"`

	// The slot identifier for the certificate containing the private key to generate the signature over the measurements.
	SlotId int64 `json:"SlotId,omitempty"`
}

// AssertComponentIntegrityV122SpdmGetSignedMeasurementsRequestBodyRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122SpdmGetSignedMeasurementsRequestBodyRequired(obj ComponentIntegrityV122SpdmGetSignedMeasurementsRequestBody) error {
	return nil
}

// AssertComponentIntegrityV122SpdmGetSignedMeasurementsRequestBodyConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122SpdmGetSignedMeasurementsRequestBodyConstraints(obj ComponentIntegrityV122SpdmGetSignedMeasurementsRequestBody) error {
	return nil
}
