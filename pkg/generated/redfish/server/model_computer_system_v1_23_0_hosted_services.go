// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// ComputerSystemV1230HostedServices - The services that might be running or installed on the system.
type ComputerSystemV1230HostedServices struct {

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	StorageServices OdataV4IdRef `json:"StorageServices,omitempty"`
}

// AssertComputerSystemV1230HostedServicesRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1230HostedServicesRequired(obj ComputerSystemV1230HostedServices) error {
	if err := AssertOdataV4IdRefRequired(obj.StorageServices); err != nil {
		return err
	}
	return nil
}

// AssertComputerSystemV1230HostedServicesConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1230HostedServicesConstraints(obj ComputerSystemV1230HostedServices) error {
	if err := AssertOdataV4IdRefConstraints(obj.StorageServices); err != nil {
		return err
	}
	return nil
}