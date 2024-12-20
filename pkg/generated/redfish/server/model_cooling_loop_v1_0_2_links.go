// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CoolingLoopV102Links - The links to other resources that are related to this resource.
type CoolingLoopV102Links struct {
	Chassis OdataV4IdRef `json:"Chassis,omitempty"`

	Facility OdataV4IdRef `json:"Facility,omitempty"`

	// An array of links to the managers responsible for managing this equipment.
	ManagedBy []OdataV4IdRef `json:"ManagedBy,omitempty"`

	// The number of items in a collection.
	ManagedByodataCount int64 `json:"ManagedBy@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCoolingLoopV102LinksRequired checks if the required fields are not zero-ed
func AssertCoolingLoopV102LinksRequired(obj CoolingLoopV102Links) error {
	if err := AssertOdataV4IdRefRequired(obj.Chassis); err != nil {
		return err
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

// AssertCoolingLoopV102LinksConstraints checks if the values respects the defined constraints
func AssertCoolingLoopV102LinksConstraints(obj CoolingLoopV102Links) error {
	if err := AssertOdataV4IdRefConstraints(obj.Chassis); err != nil {
		return err
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
