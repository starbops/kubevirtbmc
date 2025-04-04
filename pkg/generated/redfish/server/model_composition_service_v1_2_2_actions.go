// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CompositionServiceV122Actions - The available actions for this resource.
type CompositionServiceV122Actions struct {
	CompositionServiceCompose CompositionServiceV122Compose `json:"#CompositionService.Compose,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertCompositionServiceV122ActionsRequired checks if the required fields are not zero-ed
func AssertCompositionServiceV122ActionsRequired(obj CompositionServiceV122Actions) error {
	if err := AssertCompositionServiceV122ComposeRequired(obj.CompositionServiceCompose); err != nil {
		return err
	}
	return nil
}

// AssertCompositionServiceV122ActionsConstraints checks if the values respects the defined constraints
func AssertCompositionServiceV122ActionsConstraints(obj CompositionServiceV122Actions) error {
	if err := AssertCompositionServiceV122ComposeConstraints(obj.CompositionServiceCompose); err != nil {
		return err
	}
	return nil
}
