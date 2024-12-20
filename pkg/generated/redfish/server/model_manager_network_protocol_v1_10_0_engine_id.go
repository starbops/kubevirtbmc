// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ManagerNetworkProtocolV1100EngineId - The engine ID.
type ManagerNetworkProtocolV1100EngineId struct {

	// The architecture identifier.
	ArchitectureId *string `json:"ArchitectureId,omitempty" validate:"regexp=^([A-Fa-f0-9]{2} ){0,27}[A-Fa-f0-9]{2}$"`

	// The enterprise-specific method.
	EnterpriseSpecificMethod *string `json:"EnterpriseSpecificMethod,omitempty" validate:"regexp=^([A-Fa-f0-9]{2} ){7}[A-Fa-f0-9]{2}$"`

	// The private enterprise ID.
	PrivateEnterpriseId *string `json:"PrivateEnterpriseId,omitempty" validate:"regexp=^([A-Fa-f0-9]{2} ){3}[A-Fa-f0-9]{2}$"`
}

// AssertManagerNetworkProtocolV1100EngineIdRequired checks if the required fields are not zero-ed
func AssertManagerNetworkProtocolV1100EngineIdRequired(obj ManagerNetworkProtocolV1100EngineId) error {
	return nil
}

// AssertManagerNetworkProtocolV1100EngineIdConstraints checks if the values respects the defined constraints
func AssertManagerNetworkProtocolV1100EngineIdConstraints(obj ManagerNetworkProtocolV1100EngineId) error {
	return nil
}
