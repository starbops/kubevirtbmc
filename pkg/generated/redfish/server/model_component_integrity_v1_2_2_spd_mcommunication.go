// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ComponentIntegrityV122SpdMcommunication - Information about communication between two components.
type ComponentIntegrityV122SpdMcommunication struct {

	// The active sessions or communication channels between two components.
	Sessions []ComponentIntegrityV122SingleSessionInfo `json:"Sessions,omitempty"`
}

// AssertComponentIntegrityV122SpdMcommunicationRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122SpdMcommunicationRequired(obj ComponentIntegrityV122SpdMcommunication) error {
	for _, el := range obj.Sessions {
		if err := AssertComponentIntegrityV122SingleSessionInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertComponentIntegrityV122SpdMcommunicationConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122SpdMcommunicationConstraints(obj ComponentIntegrityV122SpdMcommunication) error {
	for _, el := range obj.Sessions {
		if err := AssertComponentIntegrityV122SingleSessionInfoConstraints(el); err != nil {
			return err
		}
	}
	return nil
}