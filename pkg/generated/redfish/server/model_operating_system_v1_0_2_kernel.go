// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

// OperatingSystemV102Kernel - The kernel information for an operating system.
type OperatingSystemV102Kernel struct {

	// The machine hardware name of the kernel.
	Machine *string `json:"Machine,omitempty"`

	// The name of the kernel.
	Name *string `json:"Name,omitempty"`

	// The release of the kernel.
	Release *string `json:"Release,omitempty"`

	// The version of the kernel.
	Version *string `json:"Version,omitempty"`
}

// AssertOperatingSystemV102KernelRequired checks if the required fields are not zero-ed
func AssertOperatingSystemV102KernelRequired(obj OperatingSystemV102Kernel) error {
	return nil
}

// AssertOperatingSystemV102KernelConstraints checks if the values respects the defined constraints
func AssertOperatingSystemV102KernelConstraints(obj OperatingSystemV102Kernel) error {
	return nil
}