// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ResourceV1200ContactInfo - Contact information for this resource.
type ResourceV1200ContactInfo struct {

	// Name of this contact.
	ContactName *string `json:"ContactName,omitempty"`

	// Email address for this contact.
	EmailAddress *string `json:"EmailAddress,omitempty"`

	// Phone number for this contact.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
}

// AssertResourceV1200ContactInfoRequired checks if the required fields are not zero-ed
func AssertResourceV1200ContactInfoRequired(obj ResourceV1200ContactInfo) error {
	return nil
}

// AssertResourceV1200ContactInfoConstraints checks if the values respects the defined constraints
func AssertResourceV1200ContactInfoConstraints(obj ResourceV1200ContactInfo) error {
	return nil
}
