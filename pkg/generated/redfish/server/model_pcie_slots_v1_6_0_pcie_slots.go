// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PcieSlotsV160PcieSlots - The PCIeSlots schema describes PCIe slot properties.
type PcieSlotsV160PcieSlots struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions PcieSlotsV160Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of PCI Slot information.
	Slots []PcieSlotsV160PcieSlot `json:"Slots,omitempty"`
}

// AssertPcieSlotsV160PcieSlotsRequired checks if the required fields are not zero-ed
func AssertPcieSlotsV160PcieSlotsRequired(obj PcieSlotsV160PcieSlots) error {
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

	if err := AssertPcieSlotsV160ActionsRequired(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.Slots {
		if err := AssertPcieSlotsV160PcieSlotRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPcieSlotsV160PcieSlotsConstraints checks if the values respects the defined constraints
func AssertPcieSlotsV160PcieSlotsConstraints(obj PcieSlotsV160PcieSlots) error {
	if err := AssertPcieSlotsV160ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	for _, el := range obj.Slots {
		if err := AssertPcieSlotsV160PcieSlotConstraints(el); err != nil {
			return err
		}
	}
	return nil
}