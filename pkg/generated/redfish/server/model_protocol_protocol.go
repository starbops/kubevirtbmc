// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

import (
	"fmt"
)

type ProtocolProtocol string

// List of ProtocolProtocol
const (
	PROTOCOLPROTOCOL_PCIE              ProtocolProtocol = "PCIe"
	PROTOCOLPROTOCOL_AHCI              ProtocolProtocol = "AHCI"
	PROTOCOLPROTOCOL_UHCI              ProtocolProtocol = "UHCI"
	PROTOCOLPROTOCOL_SAS               ProtocolProtocol = "SAS"
	PROTOCOLPROTOCOL_SATA              ProtocolProtocol = "SATA"
	PROTOCOLPROTOCOL_USB               ProtocolProtocol = "USB"
	PROTOCOLPROTOCOL_NVME              ProtocolProtocol = "NVMe"
	PROTOCOLPROTOCOL_FC                ProtocolProtocol = "FC"
	PROTOCOLPROTOCOL_I_SCSI            ProtocolProtocol = "iSCSI"
	PROTOCOLPROTOCOL_FCO_E             ProtocolProtocol = "FCoE"
	PROTOCOLPROTOCOL_FCP               ProtocolProtocol = "FCP"
	PROTOCOLPROTOCOL_FICON             ProtocolProtocol = "FICON"
	PROTOCOLPROTOCOL_NVME_OVER_FABRICS ProtocolProtocol = "NVMeOverFabrics"
	PROTOCOLPROTOCOL_SMB               ProtocolProtocol = "SMB"
	PROTOCOLPROTOCOL_NFSV3             ProtocolProtocol = "NFSv3"
	PROTOCOLPROTOCOL_NFSV4             ProtocolProtocol = "NFSv4"
	PROTOCOLPROTOCOL_HTTP              ProtocolProtocol = "HTTP"
	PROTOCOLPROTOCOL_HTTPS             ProtocolProtocol = "HTTPS"
	PROTOCOLPROTOCOL_FTP               ProtocolProtocol = "FTP"
	PROTOCOLPROTOCOL_SFTP              ProtocolProtocol = "SFTP"
	PROTOCOLPROTOCOL_I_WARP            ProtocolProtocol = "iWARP"
	PROTOCOLPROTOCOL_RO_CE             ProtocolProtocol = "RoCE"
	PROTOCOLPROTOCOL_RO_CEV2           ProtocolProtocol = "RoCEv2"
	PROTOCOLPROTOCOL_I2_C              ProtocolProtocol = "I2C"
	PROTOCOLPROTOCOL_TCP               ProtocolProtocol = "TCP"
	PROTOCOLPROTOCOL_UDP               ProtocolProtocol = "UDP"
	PROTOCOLPROTOCOL_TFTP              ProtocolProtocol = "TFTP"
	PROTOCOLPROTOCOL_GEN_Z             ProtocolProtocol = "GenZ"
	PROTOCOLPROTOCOL_MULTI_PROTOCOL    ProtocolProtocol = "MultiProtocol"
	PROTOCOLPROTOCOL_INFINI_BAND       ProtocolProtocol = "InfiniBand"
	PROTOCOLPROTOCOL_ETHERNET          ProtocolProtocol = "Ethernet"
	PROTOCOLPROTOCOL_NV_LINK           ProtocolProtocol = "NVLink"
	PROTOCOLPROTOCOL_OEM               ProtocolProtocol = "OEM"
	PROTOCOLPROTOCOL_DISPLAY_PORT      ProtocolProtocol = "DisplayPort"
	PROTOCOLPROTOCOL_HDMI              ProtocolProtocol = "HDMI"
	PROTOCOLPROTOCOL_VGA               ProtocolProtocol = "VGA"
	PROTOCOLPROTOCOL_DVI               ProtocolProtocol = "DVI"
	PROTOCOLPROTOCOL_CXL               ProtocolProtocol = "CXL"
	PROTOCOLPROTOCOL_UPI               ProtocolProtocol = "UPI"
	PROTOCOLPROTOCOL_QPI               ProtocolProtocol = "QPI"
	PROTOCOLPROTOCOL_E_MMC             ProtocolProtocol = "eMMC"
)

// AllowedProtocolProtocolEnumValues is all the allowed values of ProtocolProtocol enum
var AllowedProtocolProtocolEnumValues = []ProtocolProtocol{
	"PCIe",
	"AHCI",
	"UHCI",
	"SAS",
	"SATA",
	"USB",
	"NVMe",
	"FC",
	"iSCSI",
	"FCoE",
	"FCP",
	"FICON",
	"NVMeOverFabrics",
	"SMB",
	"NFSv3",
	"NFSv4",
	"HTTP",
	"HTTPS",
	"FTP",
	"SFTP",
	"iWARP",
	"RoCE",
	"RoCEv2",
	"I2C",
	"TCP",
	"UDP",
	"TFTP",
	"GenZ",
	"MultiProtocol",
	"InfiniBand",
	"Ethernet",
	"NVLink",
	"OEM",
	"DisplayPort",
	"HDMI",
	"VGA",
	"DVI",
	"CXL",
	"UPI",
	"QPI",
	"eMMC",
}

// validProtocolProtocolEnumValue provides a map of ProtocolProtocols for fast verification of use input
var validProtocolProtocolEnumValues = map[ProtocolProtocol]struct{}{
	"PCIe":            {},
	"AHCI":            {},
	"UHCI":            {},
	"SAS":             {},
	"SATA":            {},
	"USB":             {},
	"NVMe":            {},
	"FC":              {},
	"iSCSI":           {},
	"FCoE":            {},
	"FCP":             {},
	"FICON":           {},
	"NVMeOverFabrics": {},
	"SMB":             {},
	"NFSv3":           {},
	"NFSv4":           {},
	"HTTP":            {},
	"HTTPS":           {},
	"FTP":             {},
	"SFTP":            {},
	"iWARP":           {},
	"RoCE":            {},
	"RoCEv2":          {},
	"I2C":             {},
	"TCP":             {},
	"UDP":             {},
	"TFTP":            {},
	"GenZ":            {},
	"MultiProtocol":   {},
	"InfiniBand":      {},
	"Ethernet":        {},
	"NVLink":          {},
	"OEM":             {},
	"DisplayPort":     {},
	"HDMI":            {},
	"VGA":             {},
	"DVI":             {},
	"CXL":             {},
	"UPI":             {},
	"QPI":             {},
	"eMMC":            {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProtocolProtocol) IsValid() bool {
	_, ok := validProtocolProtocolEnumValues[v]
	return ok
}

// NewProtocolProtocolFromValue returns a pointer to a valid ProtocolProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProtocolProtocolFromValue(v string) (ProtocolProtocol, error) {
	ev := ProtocolProtocol(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ProtocolProtocol: valid values are %v", v, AllowedProtocolProtocolEnumValues)
}

// AssertProtocolProtocolRequired checks if the required fields are not zero-ed
func AssertProtocolProtocolRequired(obj ProtocolProtocol) error {
	return nil
}

// AssertProtocolProtocolConstraints checks if the values respects the defined constraints
func AssertProtocolProtocolConstraints(obj ProtocolProtocol) error {
	return nil
}
