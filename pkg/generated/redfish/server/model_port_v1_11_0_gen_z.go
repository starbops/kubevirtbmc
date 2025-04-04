// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PortV1110GenZ - This type defines Gen-Z specific port properties.
type PortV1110GenZ struct {
	LPRT OdataV4IdRef `json:"LPRT,omitempty"`

	MPRT OdataV4IdRef `json:"MPRT,omitempty"`

	VCAT OdataV4IdRef `json:"VCAT,omitempty"`
}

// AssertPortV1110GenZRequired checks if the required fields are not zero-ed
func AssertPortV1110GenZRequired(obj PortV1110GenZ) error {
	if err := AssertOdataV4IdRefRequired(obj.LPRT); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.MPRT); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.VCAT); err != nil {
		return err
	}
	return nil
}

// AssertPortV1110GenZConstraints checks if the values respects the defined constraints
func AssertPortV1110GenZConstraints(obj PortV1110GenZ) error {
	if err := AssertOdataV4IdRefConstraints(obj.LPRT); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.MPRT); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.VCAT); err != nil {
		return err
	}
	return nil
}
