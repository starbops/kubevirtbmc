// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PortMetricsV151PortMetrics - The PortMetrics schema contains usage and health statistics for a switch device or component port summary.
type PortMetricsV151PortMetrics struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions PortMetricsV151Actions `json:"Actions,omitempty"`

	CXL PortMetricsV151Cxl `json:"CXL,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	FibreChannel PortMetricsV151FibreChannel `json:"FibreChannel,omitempty"`

	GenZ PortMetricsV151GenZ `json:"GenZ,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	Networking PortMetricsV151Networking `json:"Networking,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	PCIeErrors PcieDevicePcieErrors `json:"PCIeErrors,omitempty"`

	// The total number of bytes received on a port since reset.
	RXBytes *int64 `json:"RXBytes,omitempty"`

	// The total number of received errors on a port since reset.
	RXErrors *int64 `json:"RXErrors,omitempty"`

	// The physical (phy) metrics for Serial Attached SCSI (SAS).  Each member represents a single phy.
	SAS []PortMetricsV151Sas `json:"SAS,omitempty"`

	// The total number of bytes transmitted on a port since reset.
	TXBytes *int64 `json:"TXBytes,omitempty"`

	// The total number of transmission errors on a port since reset.
	TXErrors *int64 `json:"TXErrors,omitempty"`

	// The metrics for the transceivers in this port.  Each member represents a single transceiver.
	Transceivers []PortMetricsV151Transceiver `json:"Transceivers,omitempty"`
}

// AssertPortMetricsV151PortMetricsRequired checks if the required fields are not zero-ed
func AssertPortMetricsV151PortMetricsRequired(obj PortMetricsV151PortMetrics) error {
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

	if err := AssertPortMetricsV151ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertPortMetricsV151CxlRequired(obj.CXL); err != nil {
		return err
	}
	if err := AssertPortMetricsV151FibreChannelRequired(obj.FibreChannel); err != nil {
		return err
	}
	if err := AssertPortMetricsV151GenZRequired(obj.GenZ); err != nil {
		return err
	}
	if err := AssertPortMetricsV151NetworkingRequired(obj.Networking); err != nil {
		return err
	}
	if err := AssertPcieDevicePcieErrorsRequired(obj.PCIeErrors); err != nil {
		return err
	}
	for _, el := range obj.SAS {
		if err := AssertPortMetricsV151SasRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Transceivers {
		if err := AssertPortMetricsV151TransceiverRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPortMetricsV151PortMetricsConstraints checks if the values respects the defined constraints
func AssertPortMetricsV151PortMetricsConstraints(obj PortMetricsV151PortMetrics) error {
	if err := AssertPortMetricsV151ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertPortMetricsV151CxlConstraints(obj.CXL); err != nil {
		return err
	}
	if err := AssertPortMetricsV151FibreChannelConstraints(obj.FibreChannel); err != nil {
		return err
	}
	if err := AssertPortMetricsV151GenZConstraints(obj.GenZ); err != nil {
		return err
	}
	if err := AssertPortMetricsV151NetworkingConstraints(obj.Networking); err != nil {
		return err
	}
	if err := AssertPcieDevicePcieErrorsConstraints(obj.PCIeErrors); err != nil {
		return err
	}
	for _, el := range obj.SAS {
		if err := AssertPortMetricsV151SasConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Transceivers {
		if err := AssertPortMetricsV151TransceiverConstraints(el); err != nil {
			return err
		}
	}
	return nil
}