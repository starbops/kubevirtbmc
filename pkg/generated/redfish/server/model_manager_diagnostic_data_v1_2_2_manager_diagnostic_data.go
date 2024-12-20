// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ManagerDiagnosticDataV122ManagerDiagnosticData - The ManagerDiagnosticData schema defines internal diagnostic data for a manager.  It contains information that might be used by vendors to collect debug information about the manager.  Clients should not make decisions for raising alerts, creating service events, or other actions based on information in this resource.
type ManagerDiagnosticDataV122ManagerDiagnosticData struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions ManagerDiagnosticDataV122Actions `json:"Actions,omitempty"`

	BootTimeStatistics ManagerDiagnosticDataV122BootTimeStatistics `json:"BootTimeStatistics,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The available storage space on this manager in kibibytes (KiB).
	FreeStorageSpaceKiB *int64 `json:"FreeStorageSpaceKiB,omitempty"`

	// The statistics of the I2C buses.
	I2CBuses []ManagerDiagnosticDataV122I2CBusStatistics `json:"I2CBuses,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	MemoryECCStatistics ManagerDiagnosticDataV122MemoryEccStatistics `json:"MemoryECCStatistics,omitempty"`

	MemoryStatistics ManagerDiagnosticDataV122MemoryStatistics `json:"MemoryStatistics,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	ProcessorStatistics ManagerDiagnosticDataV122ProcessorStatistics `json:"ProcessorStatistics,omitempty"`

	// The wall-clock time the service root hosted by this manager has been running in seconds.
	ServiceRootUptimeSeconds *float32 `json:"ServiceRootUptimeSeconds,omitempty"`

	// The statistics of the top processes of this manager.
	TopProcesses []ManagerDiagnosticDataV122ProcessStatistics `json:"TopProcesses,omitempty"`
}

// AssertManagerDiagnosticDataV122ManagerDiagnosticDataRequired checks if the required fields are not zero-ed
func AssertManagerDiagnosticDataV122ManagerDiagnosticDataRequired(obj ManagerDiagnosticDataV122ManagerDiagnosticData) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertManagerDiagnosticDataV122ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertManagerDiagnosticDataV122BootTimeStatisticsRequired(obj.BootTimeStatistics); err != nil {
		return err
	}
	for _, el := range obj.I2CBuses {
		if err := AssertManagerDiagnosticDataV122I2CBusStatisticsRequired(el); err != nil {
			return err
		}
	}
	if err := AssertManagerDiagnosticDataV122MemoryEccStatisticsRequired(obj.MemoryECCStatistics); err != nil {
		return err
	}
	if err := AssertManagerDiagnosticDataV122MemoryStatisticsRequired(obj.MemoryStatistics); err != nil {
		return err
	}
	if err := AssertManagerDiagnosticDataV122ProcessorStatisticsRequired(obj.ProcessorStatistics); err != nil {
		return err
	}
	for _, el := range obj.TopProcesses {
		if err := AssertManagerDiagnosticDataV122ProcessStatisticsRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertManagerDiagnosticDataV122ManagerDiagnosticDataConstraints checks if the values respects the defined constraints
func AssertManagerDiagnosticDataV122ManagerDiagnosticDataConstraints(obj ManagerDiagnosticDataV122ManagerDiagnosticData) error {
	if err := AssertManagerDiagnosticDataV122ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertManagerDiagnosticDataV122BootTimeStatisticsConstraints(obj.BootTimeStatistics); err != nil {
		return err
	}
	for _, el := range obj.I2CBuses {
		if err := AssertManagerDiagnosticDataV122I2CBusStatisticsConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertManagerDiagnosticDataV122MemoryEccStatisticsConstraints(obj.MemoryECCStatistics); err != nil {
		return err
	}
	if err := AssertManagerDiagnosticDataV122MemoryStatisticsConstraints(obj.MemoryStatistics); err != nil {
		return err
	}
	if err := AssertManagerDiagnosticDataV122ProcessorStatisticsConstraints(obj.ProcessorStatistics); err != nil {
		return err
	}
	for _, el := range obj.TopProcesses {
		if err := AssertManagerDiagnosticDataV122ProcessStatisticsConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
