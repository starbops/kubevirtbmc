// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ProcessorV1190Actions - The available actions for this resource.
type ProcessorV1190Actions struct {
	ProcessorReset ProcessorV1190Reset `json:"#Processor.Reset,omitempty"`

	ProcessorResetToDefaults ProcessorV1190ResetToDefaults `json:"#Processor.ResetToDefaults,omitempty"`

	// The available OEM-specific actions for this resource.
	Oem map[string]interface{} `json:"Oem,omitempty"`
}

// AssertProcessorV1190ActionsRequired checks if the required fields are not zero-ed
func AssertProcessorV1190ActionsRequired(obj ProcessorV1190Actions) error {
	if err := AssertProcessorV1190ResetRequired(obj.ProcessorReset); err != nil {
		return err
	}
	if err := AssertProcessorV1190ResetToDefaultsRequired(obj.ProcessorResetToDefaults); err != nil {
		return err
	}
	return nil
}

// AssertProcessorV1190ActionsConstraints checks if the values respects the defined constraints
func AssertProcessorV1190ActionsConstraints(obj ProcessorV1190Actions) error {
	if err := AssertProcessorV1190ResetConstraints(obj.ProcessorReset); err != nil {
		return err
	}
	if err := AssertProcessorV1190ResetToDefaultsConstraints(obj.ProcessorResetToDefaults); err != nil {
		return err
	}
	return nil
}
