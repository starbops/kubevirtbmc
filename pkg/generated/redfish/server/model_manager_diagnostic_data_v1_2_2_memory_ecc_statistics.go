// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ManagerDiagnosticDataV122MemoryEccStatistics - The memory ECC statistics of a manager.
type ManagerDiagnosticDataV122MemoryEccStatistics struct {

	// The number of correctable errors since reset.
	CorrectableECCErrorCount *int64 `json:"CorrectableECCErrorCount,omitempty"`

	// The number of uncorrectable errors since reset.
	UncorrectableECCErrorCount *int64 `json:"UncorrectableECCErrorCount,omitempty"`
}

// AssertManagerDiagnosticDataV122MemoryEccStatisticsRequired checks if the required fields are not zero-ed
func AssertManagerDiagnosticDataV122MemoryEccStatisticsRequired(obj ManagerDiagnosticDataV122MemoryEccStatistics) error {
	return nil
}

// AssertManagerDiagnosticDataV122MemoryEccStatisticsConstraints checks if the values respects the defined constraints
func AssertManagerDiagnosticDataV122MemoryEccStatisticsConstraints(obj ManagerDiagnosticDataV122MemoryEccStatistics) error {
	return nil
}