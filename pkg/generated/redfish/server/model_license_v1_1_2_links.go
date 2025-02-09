// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// LicenseV112Links - The links to other resources that are related to this resource.
type LicenseV112Links struct {

	// An array of links to the devices authorized by the license.
	AuthorizedDevices []OdataV4IdRef `json:"AuthorizedDevices,omitempty"`

	// The number of items in a collection.
	AuthorizedDevicesodataCount int64 `json:"AuthorizedDevices@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the managers where the license is installed.
	TargetServices []OdataV4IdRef `json:"TargetServices,omitempty"`

	// The number of items in a collection.
	TargetServicesodataCount int64 `json:"TargetServices@odata.count,omitempty"`
}

// AssertLicenseV112LinksRequired checks if the required fields are not zero-ed
func AssertLicenseV112LinksRequired(obj LicenseV112Links) error {
	for _, el := range obj.AuthorizedDevices {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TargetServices {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertLicenseV112LinksConstraints checks if the values respects the defined constraints
func AssertLicenseV112LinksConstraints(obj LicenseV112Links) error {
	for _, el := range obj.AuthorizedDevices {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TargetServices {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
