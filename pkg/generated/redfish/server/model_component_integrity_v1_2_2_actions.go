// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ComponentIntegrityV122Actions - The available actions for this resource.
type ComponentIntegrityV122Actions struct {
	ComponentIntegritySPDMGetSignedMeasurements ComponentIntegrityV122SpdmGetSignedMeasurements `json:"#ComponentIntegrity.SPDMGetSignedMeasurements,omitempty"`

	ComponentIntegrityTPMGetSignedMeasurements ComponentIntegrityV122TpmGetSignedMeasurements `json:"#ComponentIntegrity.TPMGetSignedMeasurements,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertComponentIntegrityV122ActionsRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122ActionsRequired(obj ComponentIntegrityV122Actions) error {
	if err := AssertComponentIntegrityV122SpdmGetSignedMeasurementsRequired(obj.ComponentIntegritySPDMGetSignedMeasurements); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpmGetSignedMeasurementsRequired(obj.ComponentIntegrityTPMGetSignedMeasurements); err != nil {
		return err
	}
	return nil
}

// AssertComponentIntegrityV122ActionsConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122ActionsConstraints(obj ComponentIntegrityV122Actions) error {
	if err := AssertComponentIntegrityV122SpdmGetSignedMeasurementsConstraints(obj.ComponentIntegritySPDMGetSignedMeasurements); err != nil {
		return err
	}
	if err := AssertComponentIntegrityV122TpmGetSignedMeasurementsConstraints(obj.ComponentIntegrityTPMGetSignedMeasurements); err != nil {
		return err
	}
	return nil
}