// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ResourceLocation - The location of a resource.
type ResourceLocation struct {

	// The altitude of the resource in meters.
	AltitudeMeters *float32 `json:"AltitudeMeters,omitempty"`

	// An array of contact information.
	Contacts []ResourceLocationContactsInner `json:"Contacts,omitempty"`

	// The location of the resource.
	// Deprecated
	Info *string `json:"Info,omitempty"`

	// The format of the `Info` property.
	// Deprecated
	InfoFormat *string `json:"InfoFormat,omitempty"`

	// The latitude of the resource.
	Latitude *float32 `json:"Latitude,omitempty"`

	// The longitude of the resource in degree units.
	Longitude *float32 `json:"Longitude,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PartLocation ResourceV1200PartLocation `json:"PartLocation,omitempty"`

	// Human-readable string to enable differentiation between `PartLocation` values for parts in the same enclosure, which might include hierarchical information of containing `PartLocation` values for the part.
	PartLocationContext *string `json:"PartLocationContext,omitempty"`

	PhysicalAddress ResourceV1200PhysicalAddress `json:"PhysicalAddress,omitempty"`

	Placement ResourceV1200Placement `json:"Placement,omitempty"`

	// Deprecated
	PostalAddress ResourceV1200PostalAddress `json:"PostalAddress,omitempty"`
}

// AssertResourceLocationRequired checks if the required fields are not zero-ed
func AssertResourceLocationRequired(obj ResourceLocation) error {
	for _, el := range obj.Contacts {
		if err := AssertResourceLocationContactsInnerRequired(el); err != nil {
			return err
		}
	}
	if err := AssertResourceV1200PartLocationRequired(obj.PartLocation); err != nil {
		return err
	}
	if err := AssertResourceV1200PhysicalAddressRequired(obj.PhysicalAddress); err != nil {
		return err
	}
	if err := AssertResourceV1200PlacementRequired(obj.Placement); err != nil {
		return err
	}
	if err := AssertResourceV1200PostalAddressRequired(obj.PostalAddress); err != nil {
		return err
	}
	return nil
}

// AssertResourceLocationConstraints checks if the values respects the defined constraints
func AssertResourceLocationConstraints(obj ResourceLocation) error {
	for _, el := range obj.Contacts {
		if err := AssertResourceLocationContactsInnerConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertResourceV1200PartLocationConstraints(obj.PartLocation); err != nil {
		return err
	}
	if err := AssertResourceV1200PhysicalAddressConstraints(obj.PhysicalAddress); err != nil {
		return err
	}
	if err := AssertResourceV1200PlacementConstraints(obj.Placement); err != nil {
		return err
	}
	if err := AssertResourceV1200PostalAddressConstraints(obj.PostalAddress); err != nil {
		return err
	}
	return nil
}
