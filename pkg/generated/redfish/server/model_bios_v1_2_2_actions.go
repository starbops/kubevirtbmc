// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// BiosV122Actions - The available actions for this resource.
type BiosV122Actions struct {
	BiosChangePassword BiosV122ChangePassword `json:"#Bios.ChangePassword,omitempty"`

	BiosResetBios BiosV122ResetBios `json:"#Bios.ResetBios,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertBiosV122ActionsRequired checks if the required fields are not zero-ed
func AssertBiosV122ActionsRequired(obj BiosV122Actions) error {
	if err := AssertBiosV122ChangePasswordRequired(obj.BiosChangePassword); err != nil {
		return err
	}
	if err := AssertBiosV122ResetBiosRequired(obj.BiosResetBios); err != nil {
		return err
	}
	return nil
}

// AssertBiosV122ActionsConstraints checks if the values respects the defined constraints
func AssertBiosV122ActionsConstraints(obj BiosV122Actions) error {
	if err := AssertBiosV122ChangePasswordConstraints(obj.BiosChangePassword); err != nil {
		return err
	}
	if err := AssertBiosV122ResetBiosConstraints(obj.BiosResetBios); err != nil {
		return err
	}
	return nil
}