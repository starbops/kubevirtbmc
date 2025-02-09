// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

import (
	"time"
)

// SoftwareInventoryV1101SoftwareInventory - The SoftwareInventory schema contains an inventory of software components.  This can include software components such as BIOS, BMC firmware, firmware for other devices, system drivers, or provider software.
type SoftwareInventoryV1101SoftwareInventory struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions SoftwareInventoryV1101Actions `json:"Actions,omitempty"`

	AdditionalVersions SoftwareInventoryV1101AdditionalVersions `json:"AdditionalVersions,omitempty"`

	AssociatedPhysicalContext PhysicalContextPhysicalContext `json:"AssociatedPhysicalContext,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The lowest supported version of this software.
	LowestSupportedVersion *string `json:"LowestSupportedVersion,omitempty"`

	// The manufacturer or producer of this software.
	Manufacturer *string `json:"Manufacturer,omitempty"`

	Measurement SoftwareInventoryMeasurementBlock `json:"Measurement,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The IDs of the Resources associated with this software inventory item.
	RelatedItem []OdataV4IdRef `json:"RelatedItem,omitempty"`

	// The number of items in a collection.
	RelatedItemodataCount int64 `json:"RelatedItem@odata.count,omitempty"`

	// The release date of this software.
	ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`

	ReleaseType SoftwareInventoryV1101ReleaseType `json:"ReleaseType,omitempty"`

	// The implementation-specific label that identifies this software.
	SoftwareId string `json:"SoftwareId,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The list of UEFI device paths of the components associated with this software inventory item.
	UefiDevicePaths []*string `json:"UefiDevicePaths,omitempty"`

	// An indication of whether the Update Service can update this software.
	Updateable *bool `json:"Updateable,omitempty"`

	// The version of this software.
	Version *string `json:"Version,omitempty"`

	VersionScheme SoftwareInventoryV1101VersionScheme `json:"VersionScheme,omitempty"`

	// Indicates if the software is write-protected.
	WriteProtected *bool `json:"WriteProtected,omitempty"`
}

// AssertSoftwareInventoryV1101SoftwareInventoryRequired checks if the required fields are not zero-ed
func AssertSoftwareInventoryV1101SoftwareInventoryRequired(obj SoftwareInventoryV1101SoftwareInventory) error {
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

	if err := AssertSoftwareInventoryV1101ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertSoftwareInventoryV1101AdditionalVersionsRequired(obj.AdditionalVersions); err != nil {
		return err
	}
	if err := AssertSoftwareInventoryMeasurementBlockRequired(obj.Measurement); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertSoftwareInventoryV1101SoftwareInventoryConstraints checks if the values respects the defined constraints
func AssertSoftwareInventoryV1101SoftwareInventoryConstraints(obj SoftwareInventoryV1101SoftwareInventory) error {
	if err := AssertSoftwareInventoryV1101ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertSoftwareInventoryV1101AdditionalVersionsConstraints(obj.AdditionalVersions); err != nil {
		return err
	}
	if err := AssertSoftwareInventoryMeasurementBlockConstraints(obj.Measurement); err != nil {
		return err
	}
	for _, el := range obj.RelatedItem {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}
