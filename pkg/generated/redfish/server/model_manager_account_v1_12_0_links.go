// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ManagerAccountV1120Links - The links to other resources that are related to this resource.
type ManagerAccountV1120Links struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	Role OdataV4IdRef `json:"Role,omitempty"`
}

// AssertManagerAccountV1120LinksRequired checks if the required fields are not zero-ed
func AssertManagerAccountV1120LinksRequired(obj ManagerAccountV1120Links) error {
	if err := AssertOdataV4IdRefRequired(obj.Role); err != nil {
		return err
	}
	return nil
}

// AssertManagerAccountV1120LinksConstraints checks if the values respects the defined constraints
func AssertManagerAccountV1120LinksConstraints(obj ManagerAccountV1120Links) error {
	if err := AssertOdataV4IdRefConstraints(obj.Role); err != nil {
		return err
	}
	return nil
}
