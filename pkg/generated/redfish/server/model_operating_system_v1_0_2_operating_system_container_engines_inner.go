// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

type OperatingSystemV102OperatingSystemContainerEnginesInner struct {

	// The URIs to manage this container engine.
	ManagementURIs []*string `json:"ManagementURIs,omitempty"`

	// The supported image types for this container engine.
	SupportedImageTypes []OperatingSystemV102ContainerEngineSupportedImageTypesInner `json:"SupportedImageTypes,omitempty"`

	Type OperatingSystemV102ContainerEngineTypes `json:"Type,omitempty"`

	// The version of this container engine.
	Version *string `json:"Version,omitempty"`
}

// AssertOperatingSystemV102OperatingSystemContainerEnginesInnerRequired checks if the required fields are not zero-ed
func AssertOperatingSystemV102OperatingSystemContainerEnginesInnerRequired(obj OperatingSystemV102OperatingSystemContainerEnginesInner) error {
	for _, el := range obj.SupportedImageTypes {
		if err := AssertOperatingSystemV102ContainerEngineSupportedImageTypesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertOperatingSystemV102OperatingSystemContainerEnginesInnerConstraints checks if the values respects the defined constraints
func AssertOperatingSystemV102OperatingSystemContainerEnginesInnerConstraints(obj OperatingSystemV102OperatingSystemContainerEnginesInner) error {
	for _, el := range obj.SupportedImageTypes {
		if err := AssertOperatingSystemV102ContainerEngineSupportedImageTypesInnerConstraints(el); err != nil {
			return err
		}
	}
	return nil
}