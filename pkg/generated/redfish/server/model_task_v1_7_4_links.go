// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// TaskV174Links - The links to other resources that are related to this resource.
type TaskV174Links struct {

	// An array of URIs referencing the resources created as the result of the operation that produced this task.
	CreatedResources []OdataV4IdRef `json:"CreatedResources,omitempty"`

	// The number of items in a collection.
	CreatedResourcesodataCount int64 `json:"CreatedResources@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertTaskV174LinksRequired checks if the required fields are not zero-ed
func AssertTaskV174LinksRequired(obj TaskV174Links) error {
	for _, el := range obj.CreatedResources {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertTaskV174LinksConstraints checks if the values respects the defined constraints
func AssertTaskV174LinksConstraints(obj TaskV174Links) error {
	for _, el := range obj.CreatedResources {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}