// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VLanNetworkInterfaceV130Actions - The available actions for this resource.
type VLanNetworkInterfaceV130Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertVLanNetworkInterfaceV130ActionsRequired checks if the required fields are not zero-ed
func AssertVLanNetworkInterfaceV130ActionsRequired(obj VLanNetworkInterfaceV130Actions) error {
	return nil
}

// AssertVLanNetworkInterfaceV130ActionsConstraints checks if the values respects the defined constraints
func AssertVLanNetworkInterfaceV130ActionsConstraints(obj VLanNetworkInterfaceV130Actions) error {
	return nil
}