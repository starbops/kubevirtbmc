// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// PortMetricsV151Networking - The port metrics for network ports, including Ethernet, Fibre Channel, and InfiniBand, that are not specific to one of these protocols.
type PortMetricsV151Networking struct {

	// The total number of RDMA protection errors.
	RDMAProtectionErrors *int64 `json:"RDMAProtectionErrors,omitempty"`

	// The total number of RDMA protocol errors.
	RDMAProtocolErrors *int64 `json:"RDMAProtocolErrors,omitempty"`

	// The total number of RDMA bytes received on a port since reset.
	RDMARXBytes *int64 `json:"RDMARXBytes,omitempty"`

	// The total number of RDMA requests received on a port since reset.
	RDMARXRequests *int64 `json:"RDMARXRequests,omitempty"`

	// The total number of RDMA bytes transmitted on a port since reset.
	RDMATXBytes *int64 `json:"RDMATXBytes,omitempty"`

	// The total number of RDMA read requests transmitted on a port since reset.
	RDMATXReadRequests *int64 `json:"RDMATXReadRequests,omitempty"`

	// The total number of RDMA requests transmitted on a port since reset.
	RDMATXRequests *int64 `json:"RDMATXRequests,omitempty"`

	// The total number of RDMA send requests transmitted on a port since reset.
	RDMATXSendRequests *int64 `json:"RDMATXSendRequests,omitempty"`

	// The total number of RDMA write requests transmitted on a port since reset.
	RDMATXWriteRequests *int64 `json:"RDMATXWriteRequests,omitempty"`

	// The total number of valid broadcast frames received on a port since reset.
	RXBroadcastFrames *int64 `json:"RXBroadcastFrames,omitempty"`

	// The total number of frames discarded in a port's receive path since reset.
	RXDiscards *int64 `json:"RXDiscards,omitempty"`

	// The total number of frames received with frame check sequence (FCS) errors on a port since reset.
	RXFCSErrors *int64 `json:"RXFCSErrors,omitempty"`

	// The total number of false carrier errors received from phy on a port since reset.
	RXFalseCarrierErrors *int64 `json:"RXFalseCarrierErrors,omitempty"`

	// The total number of frames received with alignment errors on a port since reset.
	RXFrameAlignmentErrors *int64 `json:"RXFrameAlignmentErrors,omitempty"`

	// The total number of frames received on a port since reset.
	RXFrames *int64 `json:"RXFrames,omitempty"`

	// The total number of valid multicast frames received on a port since reset.
	RXMulticastFrames *int64 `json:"RXMulticastFrames,omitempty"`

	// The total number of frames that exceed the maximum frame size.
	RXOversizeFrames *int64 `json:"RXOversizeFrames,omitempty"`

	// The total number of priority flow control (PFC) frames received on a port since reset.
	RXPFCFrames *int64 `json:"RXPFCFrames,omitempty"`

	// The total number of flow control frames from the network to pause transmission.
	RXPauseXOFFFrames *int64 `json:"RXPauseXOFFFrames,omitempty"`

	// The total number of flow control frames from the network to resume transmission.
	RXPauseXONFrames *int64 `json:"RXPauseXONFrames,omitempty"`

	// The total number of frames that are smaller than the minimum frame size of 64 bytes.
	RXUndersizeFrames *int64 `json:"RXUndersizeFrames,omitempty"`

	// The total number of valid unicast frames received on a port since reset.
	RXUnicastFrames *int64 `json:"RXUnicastFrames,omitempty"`

	// The total number of good broadcast frames transmitted on a port since reset.
	TXBroadcastFrames *int64 `json:"TXBroadcastFrames,omitempty"`

	// The total number of frames discarded in a port's transmit path since reset.
	TXDiscards *int64 `json:"TXDiscards,omitempty"`

	// The number of times a single transmitted frame encountered more than 15 collisions.
	TXExcessiveCollisions *int64 `json:"TXExcessiveCollisions,omitempty"`

	// The total number of frames transmitted on a port since reset.
	TXFrames *int64 `json:"TXFrames,omitempty"`

	// The total number of collisions that occurred after one slot time as defined by IEEE 802.3.
	TXLateCollisions *int64 `json:"TXLateCollisions,omitempty"`

	// The total number of good multicast frames transmitted on a port since reset.
	TXMulticastFrames *int64 `json:"TXMulticastFrames,omitempty"`

	// The times that a transmitted frame encountered 2-15 collisions.
	TXMultipleCollisions *int64 `json:"TXMultipleCollisions,omitempty"`

	// The total number of priority flow control (PFC) frames sent on a port since reset.
	TXPFCFrames *int64 `json:"TXPFCFrames,omitempty"`

	// The total number of XOFF frames transmitted to the network.
	TXPauseXOFFFrames *int64 `json:"TXPauseXOFFFrames,omitempty"`

	// The total number of XON frames transmitted to the network.
	TXPauseXONFrames *int64 `json:"TXPauseXONFrames,omitempty"`

	// The times that a successfully transmitted frame encountered a single collision.
	TXSingleCollisions *int64 `json:"TXSingleCollisions,omitempty"`

	// The total number of good unicast frames transmitted on a port since reset.
	TXUnicastFrames *int64 `json:"TXUnicastFrames,omitempty"`
}

// AssertPortMetricsV151NetworkingRequired checks if the required fields are not zero-ed
func AssertPortMetricsV151NetworkingRequired(obj PortMetricsV151Networking) error {
	return nil
}

// AssertPortMetricsV151NetworkingConstraints checks if the values respects the defined constraints
func AssertPortMetricsV151NetworkingConstraints(obj PortMetricsV151Networking) error {
	return nil
}
