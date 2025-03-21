// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// VirtualMediaV163Actions - The available actions for this resource.
type VirtualMediaV163Actions struct {
	VirtualMediaEjectMedia VirtualMediaV163EjectMedia `json:"#VirtualMedia.EjectMedia,omitempty"`

	VirtualMediaInsertMedia VirtualMediaV163InsertMedia `json:"#VirtualMedia.InsertMedia,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertVirtualMediaV163ActionsRequired checks if the required fields are not zero-ed
func AssertVirtualMediaV163ActionsRequired(obj VirtualMediaV163Actions) error {
	if err := AssertVirtualMediaV163EjectMediaRequired(obj.VirtualMediaEjectMedia); err != nil {
		return err
	}
	if err := AssertVirtualMediaV163InsertMediaRequired(obj.VirtualMediaInsertMedia); err != nil {
		return err
	}
	return nil
}

// AssertVirtualMediaV163ActionsConstraints checks if the values respects the defined constraints
func AssertVirtualMediaV163ActionsConstraints(obj VirtualMediaV163Actions) error {
	if err := AssertVirtualMediaV163EjectMediaConstraints(obj.VirtualMediaEjectMedia); err != nil {
		return err
	}
	if err := AssertVirtualMediaV163InsertMediaConstraints(obj.VirtualMediaInsertMedia); err != nil {
		return err
	}
	return nil
}
