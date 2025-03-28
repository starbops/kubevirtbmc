// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// StorageV1151Links - The links to other resources that are related to this resource.
type StorageV1151Links struct {

	// An array of links to the chassis to which this storage subsystem is attached.
	Enclosures []OdataV4IdRef `json:"Enclosures,omitempty"`

	// The number of items in a collection.
	EnclosuresodataCount int64 `json:"Enclosures@odata.count,omitempty"`

	// The storage systems that host this storage subsystem.
	HostingStorageSystems []OdataV4IdRef `json:"HostingStorageSystems,omitempty"`

	// The number of items in a collection.
	HostingStorageSystemsodataCount int64 `json:"HostingStorageSystems@odata.count,omitempty"`

	// An array of links to the discovery subsystems that discovered this subsystem in an NVMe-oF environment.
	NVMeoFDiscoverySubsystems []OdataV4IdRef `json:"NVMeoFDiscoverySubsystems,omitempty"`

	// The number of items in a collection.
	NVMeoFDiscoverySubsystemsodataCount int64 `json:"NVMeoFDiscoverySubsystems@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	SimpleStorage OdataV4IdRef `json:"SimpleStorage,omitempty"`

	// An array of links to the storage services that connect to this storage subsystem.
	StorageServices []OdataV4IdRef `json:"StorageServices,omitempty"`

	// The number of items in a collection.
	StorageServicesodataCount int64 `json:"StorageServices@odata.count,omitempty"`
}

// AssertStorageV1151LinksRequired checks if the required fields are not zero-ed
func AssertStorageV1151LinksRequired(obj StorageV1151Links) error {
	for _, el := range obj.Enclosures {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.HostingStorageSystems {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NVMeoFDiscoverySubsystems {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.SimpleStorage); err != nil {
		return err
	}
	for _, el := range obj.StorageServices {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertStorageV1151LinksConstraints checks if the values respects the defined constraints
func AssertStorageV1151LinksConstraints(obj StorageV1151Links) error {
	for _, el := range obj.Enclosures {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.HostingStorageSystems {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NVMeoFDiscoverySubsystems {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefConstraints(obj.SimpleStorage); err != nil {
		return err
	}
	for _, el := range obj.StorageServices {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
