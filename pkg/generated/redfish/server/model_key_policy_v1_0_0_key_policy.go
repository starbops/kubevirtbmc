// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// KeyPolicyV100KeyPolicy - The KeyPolicy schema describes settings for how keys are allowed to be used for accessing devices or services.
type KeyPolicyV100KeyPolicy struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions KeyPolicyV100Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// Indicates if this is the default key policy.
	IsDefault *bool `json:"IsDefault,omitempty"`

	KeyPolicyType KeyPolicyV100KeyPolicyType `json:"KeyPolicyType,omitempty"`

	NVMeoF KeyPolicyV100NvMeoF `json:"NVMeoF,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertKeyPolicyV100KeyPolicyRequired checks if the required fields are not zero-ed
func AssertKeyPolicyV100KeyPolicyRequired(obj KeyPolicyV100KeyPolicy) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertKeyPolicyV100ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertKeyPolicyV100NvMeoFRequired(obj.NVMeoF); err != nil {
		return err
	}
	return nil
}

// AssertKeyPolicyV100KeyPolicyConstraints checks if the values respects the defined constraints
func AssertKeyPolicyV100KeyPolicyConstraints(obj KeyPolicyV100KeyPolicy) error {
	if err := AssertKeyPolicyV100ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertKeyPolicyV100NvMeoFConstraints(obj.NVMeoF); err != nil {
		return err
	}
	return nil
}