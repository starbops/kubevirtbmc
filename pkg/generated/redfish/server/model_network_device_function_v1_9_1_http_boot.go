// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// NetworkDeviceFunctionV191HttpBoot - The HTTP and HTTPS boot capabilities, status, and configuration for a network device function.
type NetworkDeviceFunctionV191HttpBoot struct {

	// The URI of the boot media loaded with this network device function.
	BootMediaURI *string `json:"BootMediaURI,omitempty"`
}

// AssertNetworkDeviceFunctionV191HttpBootRequired checks if the required fields are not zero-ed
func AssertNetworkDeviceFunctionV191HttpBootRequired(obj NetworkDeviceFunctionV191HttpBoot) error {
	return nil
}

// AssertNetworkDeviceFunctionV191HttpBootConstraints checks if the values respects the defined constraints
func AssertNetworkDeviceFunctionV191HttpBootConstraints(obj NetworkDeviceFunctionV191HttpBoot) error {
	return nil
}