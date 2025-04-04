// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AggregationSourceV140GenerateSshIdentityKeyPairRequestBody - This action generates a new SSH identity key-pair to be used with this aggregation source.  The generated public key is stored in the Key resource referenced by the PublicIdentityKey property in SSHSettings.  Any existing key-pair is deleted and replaced by the new key-pair.
type AggregationSourceV140GenerateSshIdentityKeyPairRequestBody struct {
	Curve KeyEcdsaCurveType `json:"Curve,omitempty"`

	// The length of the SSH key, in bits, if the KeyType parameter contains `RSA`.
	KeyLength int64 `json:"KeyLength,omitempty"`

	KeyType KeySshKeyType `json:"KeyType"`
}

// AssertAggregationSourceV140GenerateSshIdentityKeyPairRequestBodyRequired checks if the required fields are not zero-ed
func AssertAggregationSourceV140GenerateSshIdentityKeyPairRequestBodyRequired(obj AggregationSourceV140GenerateSshIdentityKeyPairRequestBody) error {
	elements := map[string]interface{}{
		"KeyType": obj.KeyType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAggregationSourceV140GenerateSshIdentityKeyPairRequestBodyConstraints checks if the values respects the defined constraints
func AssertAggregationSourceV140GenerateSshIdentityKeyPairRequestBodyConstraints(obj AggregationSourceV140GenerateSshIdentityKeyPairRequestBody) error {
	return nil
}
