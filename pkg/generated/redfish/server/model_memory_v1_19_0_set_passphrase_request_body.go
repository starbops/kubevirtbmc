// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// MemoryV1190SetPassphraseRequestBody - Set passphrase for the given regions.
type MemoryV1190SetPassphraseRequestBody struct {

	// Passphrase for doing the operation.
	Passphrase string `json:"Passphrase"`

	// The memory region ID to which to apply this action.
	RegionId string `json:"RegionId"`
}

// AssertMemoryV1190SetPassphraseRequestBodyRequired checks if the required fields are not zero-ed
func AssertMemoryV1190SetPassphraseRequestBodyRequired(obj MemoryV1190SetPassphraseRequestBody) error {
	elements := map[string]interface{}{
		"Passphrase": obj.Passphrase,
		"RegionId":   obj.RegionId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertMemoryV1190SetPassphraseRequestBodyConstraints checks if the values respects the defined constraints
func AssertMemoryV1190SetPassphraseRequestBodyConstraints(obj MemoryV1190SetPassphraseRequestBody) error {
	return nil
}