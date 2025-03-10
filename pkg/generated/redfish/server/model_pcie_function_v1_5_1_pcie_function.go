// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PcieFunctionV151PcieFunction - The PCIeFunction schema describes the properties of a PCIe function that is attached to a system.
type PcieFunctionV151PcieFunction struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions PcieFunctionV151Actions `json:"Actions,omitempty"`

	// The Class Code of this PCIe function.
	ClassCode *string `json:"ClassCode,omitempty" validate:"regexp=^0[xX]([0-9A-Fa-f]{2}){3}$"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	DeviceClass PcieFunctionV151DeviceClass `json:"DeviceClass,omitempty"`

	// The Device ID of this PCIe function.
	DeviceId *string `json:"DeviceId,omitempty" validate:"regexp=^0[xX]([0-9A-Fa-f]{2}){2}$"`

	// An indication of whether this PCIe device function is enabled.
	Enabled bool `json:"Enabled,omitempty"`

	// The PCIe function number.
	FunctionId *int64 `json:"FunctionId,omitempty"`

	FunctionProtocol PcieFunctionV151FunctionProtocol `json:"FunctionProtocol,omitempty"`

	FunctionType PcieFunctionV151FunctionType `json:"FunctionType,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Links PcieFunctionV151Links `json:"Links,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The Revision ID of this PCIe function.
	RevisionId *string `json:"RevisionId,omitempty" validate:"regexp=^0[xX]([0-9A-Fa-f]{2}){1}$"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The Subsystem ID of this PCIe function.
	SubsystemId *string `json:"SubsystemId,omitempty" validate:"regexp=^0[xX]([0-9A-Fa-f]{2}){2}$"`

	// The Subsystem Vendor ID of this PCIe function.
	SubsystemVendorId *string `json:"SubsystemVendorId,omitempty" validate:"regexp=^0[xX]([0-9A-Fa-f]{2}){2}$"`

	// The Vendor ID of this PCIe function.
	VendorId *string `json:"VendorId,omitempty" validate:"regexp=^0[xX]([0-9A-Fa-f]{2}){2}$"`
}

// AssertPcieFunctionV151PcieFunctionRequired checks if the required fields are not zero-ed
func AssertPcieFunctionV151PcieFunctionRequired(obj PcieFunctionV151PcieFunction) error {
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

	if err := AssertPcieFunctionV151ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertPcieFunctionV151LinksRequired(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertPcieFunctionV151PcieFunctionConstraints checks if the values respects the defined constraints
func AssertPcieFunctionV151PcieFunctionConstraints(obj PcieFunctionV151PcieFunction) error {
	if err := AssertPcieFunctionV151ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertPcieFunctionV151LinksConstraints(obj.Links); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}
