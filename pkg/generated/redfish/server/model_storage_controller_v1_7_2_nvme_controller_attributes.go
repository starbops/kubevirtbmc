// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// StorageControllerV172NvmeControllerAttributes - The NVMe controller attributes for a storage controller.
type StorageControllerV172NvmeControllerAttributes struct {

	// Indicates whether or not the controller supports reporting of Namespace Granularity.
	ReportsNamespaceGranularity *bool `json:"ReportsNamespaceGranularity,omitempty"`

	// Indicates whether or not the controller supports reporting of a UUID list.
	ReportsUUIDList *bool `json:"ReportsUUIDList,omitempty"`

	// Indicates whether or not the controller supports a 128-bit Host Identifier.
	Supports128BitHostId *bool `json:"Supports128BitHostId,omitempty"`

	// Indicates whether or not the controller supports Endurance Groups.
	SupportsEnduranceGroups *bool `json:"SupportsEnduranceGroups,omitempty"`

	// Indicates whether or not the controller supports exceeding Power of Non-Operational State in order to execute controller-initiated background operations in a non-operational power state.
	SupportsExceedingPowerOfNonOperationalState *bool `json:"SupportsExceedingPowerOfNonOperationalState,omitempty"`

	// Indicates whether or not the controller supports NVM Sets.
	SupportsNVMSets *bool `json:"SupportsNVMSets,omitempty"`

	// Indicates whether or not the controller supports Predictable Latency Mode.
	SupportsPredictableLatencyMode *bool `json:"SupportsPredictableLatencyMode,omitempty"`

	// Indicates whether or not the controller supports Read Recovery Levels.
	SupportsReadRecoveryLevels *bool `json:"SupportsReadRecoveryLevels,omitempty"`

	// Indicates if the controller supports reservations.
	SupportsReservations *bool `json:"SupportsReservations,omitempty"`

	// Indicates whether or not the controller supports SQ Associations.
	SupportsSQAssociations *bool `json:"SupportsSQAssociations,omitempty"`

	// Indicates whether or not the controller supports restarting the Keep Alive Timer if traffic is processed from an admin command or I/O during a Keep Alive Timeout interval.
	SupportsTrafficBasedKeepAlive *bool `json:"SupportsTrafficBasedKeepAlive,omitempty"`
}

// AssertStorageControllerV172NvmeControllerAttributesRequired checks if the required fields are not zero-ed
func AssertStorageControllerV172NvmeControllerAttributesRequired(obj StorageControllerV172NvmeControllerAttributes) error {
	return nil
}

// AssertStorageControllerV172NvmeControllerAttributesConstraints checks if the values respects the defined constraints
func AssertStorageControllerV172NvmeControllerAttributesConstraints(obj StorageControllerV172NvmeControllerAttributes) error {
	return nil
}