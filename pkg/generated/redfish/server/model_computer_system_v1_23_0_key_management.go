// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// ComputerSystemV1230KeyManagement - The key management settings of a computer system.
type ComputerSystemV1230KeyManagement struct {
	KMIPCertificates OdataV4IdRef `json:"KMIPCertificates,omitempty"`

	// The KMIP servers to which this computer system is subscribed.
	KMIPServers []ComputerSystemV1230KeyManagementKmipServersInner `json:"KMIPServers,omitempty"`
}

// AssertComputerSystemV1230KeyManagementRequired checks if the required fields are not zero-ed
func AssertComputerSystemV1230KeyManagementRequired(obj ComputerSystemV1230KeyManagement) error {
	if err := AssertOdataV4IdRefRequired(obj.KMIPCertificates); err != nil {
		return err
	}
	for _, el := range obj.KMIPServers {
		if err := AssertComputerSystemV1230KeyManagementKmipServersInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertComputerSystemV1230KeyManagementConstraints checks if the values respects the defined constraints
func AssertComputerSystemV1230KeyManagementConstraints(obj ComputerSystemV1230KeyManagement) error {
	if err := AssertOdataV4IdRefConstraints(obj.KMIPCertificates); err != nil {
		return err
	}
	for _, el := range obj.KMIPServers {
		if err := AssertComputerSystemV1230KeyManagementKmipServersInnerConstraints(el); err != nil {
			return err
		}
	}
	return nil
}