// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AccelerationFunctionV104AccelerationFunction - The AccelerationFunction schema describes an acceleration function that a processor implements.  This can include functions such as audio processing, compression, encryption, packet inspection, packet switching, scheduling, or video processing.
type AccelerationFunctionV104AccelerationFunction struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	AccelerationFunctionType AccelerationFunctionV104AccelerationFunctionType `json:"AccelerationFunctionType,omitempty"`

	Actions AccelerationFunctionV104Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// An array of the reconfiguration slot identifiers of the FPGA that this acceleration function occupies.
	FpgaReconfigurationSlots []string `json:"FpgaReconfigurationSlots,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Links AccelerationFunctionV104Links `json:"Links,omitempty"`

	// The acceleration function code manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The acceleration function power consumption, in watt units.
	PowerWatts int64 `json:"PowerWatts,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	UUID string `json:"UUID,omitempty" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`

	// The acceleration function version.
	Version string `json:"Version,omitempty"`
}

// AssertAccelerationFunctionV104AccelerationFunctionRequired checks if the required fields are not zero-ed
func AssertAccelerationFunctionV104AccelerationFunctionRequired(obj AccelerationFunctionV104AccelerationFunction) error {
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

	if err := AssertAccelerationFunctionV104ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertAccelerationFunctionV104LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertAccelerationFunctionV104AccelerationFunctionConstraints checks if the values respects the defined constraints
func AssertAccelerationFunctionV104AccelerationFunctionConstraints(obj AccelerationFunctionV104AccelerationFunction) error {
	if err := AssertAccelerationFunctionV104ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertAccelerationFunctionV104LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}