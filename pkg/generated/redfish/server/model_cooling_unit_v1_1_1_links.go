// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CoolingUnitV111Links - The links to other resources that are related to this resource.
type CoolingUnitV111Links struct {

	// An array of links to the chassis that contain this equipment.
	Chassis []OdataV4IdRef `json:"Chassis,omitempty"`

	// The number of items in a collection.
	ChassisodataCount int64 `json:"Chassis@odata.count,omitempty"`

	Facility OdataV4IdRef `json:"Facility,omitempty"`

	// An array of links to the managers responsible for managing this equipment.
	ManagedBy []OdataV4IdRef `json:"ManagedBy,omitempty"`

	// The number of items in a collection.
	ManagedByodataCount int64 `json:"ManagedBy@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCoolingUnitV111LinksRequired checks if the required fields are not zero-ed
func AssertCoolingUnitV111LinksRequired(obj CoolingUnitV111Links) error {
	for _, el := range obj.Chassis {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.Facility); err != nil {
		return err
	}
	for _, el := range obj.ManagedBy {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCoolingUnitV111LinksConstraints checks if the values respects the defined constraints
func AssertCoolingUnitV111LinksConstraints(obj CoolingUnitV111Links) error {
	for _, el := range obj.Chassis {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefConstraints(obj.Facility); err != nil {
		return err
	}
	for _, el := range obj.ManagedBy {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}