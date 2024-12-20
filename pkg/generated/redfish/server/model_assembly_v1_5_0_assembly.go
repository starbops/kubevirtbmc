// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AssemblyV150Assembly - The Assembly schema defines an assembly.  Assembly information contains details about a device, such as part number, serial number, manufacturer, and production date.  It also provides access to the original data for the assembly.
type AssemblyV150Assembly struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions AssemblyV150Actions `json:"Actions,omitempty"`

	// The assembly records.
	Assemblies []AssemblyV150AssemblyData `json:"Assemblies,omitempty"`

	// The number of items in a collection.
	AssembliesodataCount int64 `json:"Assemblies@odata.count,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertAssemblyV150AssemblyRequired checks if the required fields are not zero-ed
func AssertAssemblyV150AssemblyRequired(obj AssemblyV150Assembly) error {
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

	if err := AssertAssemblyV150ActionsRequired(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.Assemblies {
		if err := AssertAssemblyV150AssemblyDataRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAssemblyV150AssemblyConstraints checks if the values respects the defined constraints
func AssertAssemblyV150AssemblyConstraints(obj AssemblyV150Assembly) error {
	if err := AssertAssemblyV150ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.Assemblies {
		if err := AssertAssemblyV150AssemblyDataConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
