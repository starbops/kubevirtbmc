// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// BatteryV122Links - The links to other resources that are related to this resource.
type BatteryV122Links struct {

	// An array of links to the memory devices to which this battery provides power during a power-loss event.
	Memory []OdataV4IdRef `json:"Memory,omitempty"`

	// The number of items in a collection.
	MemoryodataCount int64 `json:"Memory@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the storage controllers to which this battery provides power during a power-loss event.
	StorageControllers []OdataV4IdRef `json:"StorageControllers,omitempty"`

	// The number of items in a collection.
	StorageControllersodataCount int64 `json:"StorageControllers@odata.count,omitempty"`
}

// AssertBatteryV122LinksRequired checks if the required fields are not zero-ed
func AssertBatteryV122LinksRequired(obj BatteryV122Links) error {
	for _, el := range obj.Memory {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.StorageControllers {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertBatteryV122LinksConstraints checks if the values respects the defined constraints
func AssertBatteryV122LinksConstraints(obj BatteryV122Links) error {
	for _, el := range obj.Memory {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.StorageControllers {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}