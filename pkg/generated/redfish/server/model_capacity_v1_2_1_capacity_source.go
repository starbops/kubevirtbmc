// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CapacityV121CapacitySource - A description of the type and source of storage.
type CapacityV121CapacitySource struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions CapacityV121Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	ProvidedCapacity CapacityV121Capacity `json:"ProvidedCapacity,omitempty"`

	ProvidedClassOfService OdataV4IdRef `json:"ProvidedClassOfService,omitempty"`

	ProvidingDrives OdataV4IdRef `json:"ProvidingDrives,omitempty"`

	ProvidingMemory OdataV4IdRef `json:"ProvidingMemory,omitempty"`

	ProvidingMemoryChunks OdataV4IdRef `json:"ProvidingMemoryChunks,omitempty"`

	ProvidingPools OdataV4IdRef `json:"ProvidingPools,omitempty"`

	ProvidingVolumes OdataV4IdRef `json:"ProvidingVolumes,omitempty"`
}

// AssertCapacityV121CapacitySourceRequired checks if the required fields are not zero-ed
func AssertCapacityV121CapacitySourceRequired(obj CapacityV121CapacitySource) error {
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

	if err := AssertCapacityV121ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertCapacityV121CapacityRequired(obj.ProvidedCapacity); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ProvidedClassOfService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ProvidingDrives); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ProvidingMemory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ProvidingMemoryChunks); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ProvidingPools); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.ProvidingVolumes); err != nil {
		return err
	}
	return nil
}

// AssertCapacityV121CapacitySourceConstraints checks if the values respects the defined constraints
func AssertCapacityV121CapacitySourceConstraints(obj CapacityV121CapacitySource) error {
	if err := AssertCapacityV121ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertCapacityV121CapacityConstraints(obj.ProvidedCapacity); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ProvidedClassOfService); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ProvidingDrives); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ProvidingMemory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ProvidingMemoryChunks); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ProvidingPools); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.ProvidingVolumes); err != nil {
		return err
	}
	return nil
}