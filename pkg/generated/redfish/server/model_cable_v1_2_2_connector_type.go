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

type CableV122ConnectorType string

// List of CableV122ConnectorType
const (
	CABLEV122CONNECTORTYPE_AC_POWER     CableV122ConnectorType = "ACPower"
	CABLEV122CONNECTORTYPE_DB9          CableV122ConnectorType = "DB9"
	CABLEV122CONNECTORTYPE_DC_POWER     CableV122ConnectorType = "DCPower"
	CABLEV122CONNECTORTYPE_DISPLAY_PORT CableV122ConnectorType = "DisplayPort"
	CABLEV122CONNECTORTYPE_HDMI         CableV122ConnectorType = "HDMI"
	CABLEV122CONNECTORTYPE_ICI          CableV122ConnectorType = "ICI"
	CABLEV122CONNECTORTYPE_IPASS        CableV122ConnectorType = "IPASS"
	CABLEV122CONNECTORTYPE_PCIE         CableV122ConnectorType = "PCIe"
	CABLEV122CONNECTORTYPE_PROPRIETARY  CableV122ConnectorType = "Proprietary"
	CABLEV122CONNECTORTYPE_RJ45         CableV122ConnectorType = "RJ45"
	CABLEV122CONNECTORTYPE_SATA         CableV122ConnectorType = "SATA"
	CABLEV122CONNECTORTYPE_SCSI         CableV122ConnectorType = "SCSI"
	CABLEV122CONNECTORTYPE_SLIM_SAS     CableV122ConnectorType = "SlimSAS"
	CABLEV122CONNECTORTYPE_SFP          CableV122ConnectorType = "SFP"
	CABLEV122CONNECTORTYPE_SFP_PLUS     CableV122ConnectorType = "SFPPlus"
	CABLEV122CONNECTORTYPE_USBA         CableV122ConnectorType = "USBA"
	CABLEV122CONNECTORTYPE_USBC         CableV122ConnectorType = "USBC"
	CABLEV122CONNECTORTYPE_QSFP         CableV122ConnectorType = "QSFP"
	CABLEV122CONNECTORTYPE_CDFP         CableV122ConnectorType = "CDFP"
	CABLEV122CONNECTORTYPE_OSFP         CableV122ConnectorType = "OSFP"
)

// AllowedCableV122ConnectorTypeEnumValues is all the allowed values of CableV122ConnectorType enum
var AllowedCableV122ConnectorTypeEnumValues = []CableV122ConnectorType{
	"ACPower",
	"DB9",
	"DCPower",
	"DisplayPort",
	"HDMI",
	"ICI",
	"IPASS",
	"PCIe",
	"Proprietary",
	"RJ45",
	"SATA",
	"SCSI",
	"SlimSAS",
	"SFP",
	"SFPPlus",
	"USBA",
	"USBC",
	"QSFP",
	"CDFP",
	"OSFP",
}

// validCableV122ConnectorTypeEnumValue provides a map of CableV122ConnectorTypes for fast verification of use input
var validCableV122ConnectorTypeEnumValues = map[CableV122ConnectorType]struct{}{
	"ACPower":     {},
	"DB9":         {},
	"DCPower":     {},
	"DisplayPort": {},
	"HDMI":        {},
	"ICI":         {},
	"IPASS":       {},
	"PCIe":        {},
	"Proprietary": {},
	"RJ45":        {},
	"SATA":        {},
	"SCSI":        {},
	"SlimSAS":     {},
	"SFP":         {},
	"SFPPlus":     {},
	"USBA":        {},
	"USBC":        {},
	"QSFP":        {},
	"CDFP":        {},
	"OSFP":        {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableV122ConnectorType) IsValid() bool {
	_, ok := validCableV122ConnectorTypeEnumValues[v]
	return ok
}

// NewCableV122ConnectorTypeFromValue returns a pointer to a valid CableV122ConnectorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableV122ConnectorTypeFromValue(v string) (CableV122ConnectorType, error) {
	ev := CableV122ConnectorType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CableV122ConnectorType: valid values are %v", v, AllowedCableV122ConnectorTypeEnumValues)
}

// AssertCableV122ConnectorTypeRequired checks if the required fields are not zero-ed
func AssertCableV122ConnectorTypeRequired(obj CableV122ConnectorType) error {
	return nil
}

// AssertCableV122ConnectorTypeConstraints checks if the values respects the defined constraints
func AssertCableV122ConnectorTypeConstraints(obj CableV122ConnectorType) error {
	return nil
}