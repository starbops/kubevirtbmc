// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// EndpointV181PciId - A PCI ID.
type EndpointV181PciId struct {

	// The Class Code, Subclass, and Programming Interface code of this PCIe function.
	ClassCode *string `json:"ClassCode,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){3}$"`

	// The Device ID of this PCIe function.
	DeviceId *string `json:"DeviceId,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){2}$"`

	// The PCI ID of the connected entity.
	FunctionNumber *int64 `json:"FunctionNumber,omitempty"`

	// The Subsystem ID of this PCIe function.
	SubsystemId *string `json:"SubsystemId,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){2}$"`

	// The Subsystem Vendor ID of this PCIe function.
	SubsystemVendorId *string `json:"SubsystemVendorId,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){2}$"`

	// The Vendor ID of this PCIe function.
	VendorId *string `json:"VendorId,omitempty" validate:"regexp=^0[xX](([a-fA-F]|[0-9]){2}){2}$"`
}

// AssertEndpointV181PciIdRequired checks if the required fields are not zero-ed
func AssertEndpointV181PciIdRequired(obj EndpointV181PciId) error {
	return nil
}

// AssertEndpointV181PciIdConstraints checks if the values respects the defined constraints
func AssertEndpointV181PciIdConstraints(obj EndpointV181PciId) error {
	return nil
}
