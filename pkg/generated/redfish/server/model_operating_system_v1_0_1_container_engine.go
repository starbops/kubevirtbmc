// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// OperatingSystemV101ContainerEngine - A container engine running in an operating system.
type OperatingSystemV101ContainerEngine struct {

	// The URIs to manage this container engine.
	ManagementURIs []*string `json:"ManagementURIs,omitempty"`

	// The supported image types for this container engine.
	SupportedImageTypes []ContainerImageImageTypes `json:"SupportedImageTypes,omitempty"`

	Type OperatingSystemV101ContainerEngineTypes `json:"Type,omitempty"`

	// The version of this container engine.
	Version *string `json:"Version,omitempty"`
}

// AssertOperatingSystemV101ContainerEngineRequired checks if the required fields are not zero-ed
func AssertOperatingSystemV101ContainerEngineRequired(obj OperatingSystemV101ContainerEngine) error {
	return nil
}

// AssertOperatingSystemV101ContainerEngineConstraints checks if the values respects the defined constraints
func AssertOperatingSystemV101ContainerEngineConstraints(obj OperatingSystemV101ContainerEngine) error {
	return nil
}
