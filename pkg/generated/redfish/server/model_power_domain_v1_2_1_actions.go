// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PowerDomainV121Actions - The available actions for this resource.
type PowerDomainV121Actions struct {

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertPowerDomainV121ActionsRequired checks if the required fields are not zero-ed
func AssertPowerDomainV121ActionsRequired(obj PowerDomainV121Actions) error {
	return nil
}

// AssertPowerDomainV121ActionsConstraints checks if the values respects the defined constraints
func AssertPowerDomainV121ActionsConstraints(obj PowerDomainV121Actions) error {
	return nil
}