// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ResourceV1200PhysicalAddress - The physical address for a resource.
type ResourceV1200PhysicalAddress struct {

	// City, township, or shi (JP).
	City *string `json:"City,omitempty"`

	// The country.
	Country *string `json:"Country,omitempty"`

	// The ISO 3166-1 country code.
	ISOCountryCode *string `json:"ISOCountryCode,omitempty" validate:"regexp=^([A-Z]{2}|[A-Z]{3})$"`

	// ISO 3166-2 subdivision code.
	ISOSubdivisionCode *string `json:"ISOSubdivisionCode,omitempty" validate:"regexp=^[A-Z0-9]{1,3}$"`

	// The postal code.
	PostalCode *string `json:"PostalCode,omitempty"`

	// State or province.
	StateOrProvince *string `json:"StateOrProvince,omitempty"`

	// The street-level address, including building, room, or other identifiers.
	StreetAddress *string `json:"StreetAddress,omitempty"`
}

// AssertResourceV1200PhysicalAddressRequired checks if the required fields are not zero-ed
func AssertResourceV1200PhysicalAddressRequired(obj ResourceV1200PhysicalAddress) error {
	return nil
}

// AssertResourceV1200PhysicalAddressConstraints checks if the values respects the defined constraints
func AssertResourceV1200PhysicalAddressConstraints(obj ResourceV1200PhysicalAddress) error {
	return nil
}
