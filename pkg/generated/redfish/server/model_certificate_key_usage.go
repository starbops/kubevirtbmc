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

// CertificateKeyUsage : The usages of a key contained within a certificate.
type CertificateKeyUsage string

// List of CertificateKeyUsage
const (
	CERTIFICATEKEYUSAGE_DIGITAL_SIGNATURE     CertificateKeyUsage = "DigitalSignature"
	CERTIFICATEKEYUSAGE_NON_REPUDIATION       CertificateKeyUsage = "NonRepudiation"
	CERTIFICATEKEYUSAGE_KEY_ENCIPHERMENT      CertificateKeyUsage = "KeyEncipherment"
	CERTIFICATEKEYUSAGE_DATA_ENCIPHERMENT     CertificateKeyUsage = "DataEncipherment"
	CERTIFICATEKEYUSAGE_KEY_AGREEMENT         CertificateKeyUsage = "KeyAgreement"
	CERTIFICATEKEYUSAGE_KEY_CERT_SIGN         CertificateKeyUsage = "KeyCertSign"
	CERTIFICATEKEYUSAGE_CRL_SIGNING           CertificateKeyUsage = "CRLSigning"
	CERTIFICATEKEYUSAGE_ENCIPHER_ONLY         CertificateKeyUsage = "EncipherOnly"
	CERTIFICATEKEYUSAGE_DECIPHER_ONLY         CertificateKeyUsage = "DecipherOnly"
	CERTIFICATEKEYUSAGE_SERVER_AUTHENTICATION CertificateKeyUsage = "ServerAuthentication"
	CERTIFICATEKEYUSAGE_CLIENT_AUTHENTICATION CertificateKeyUsage = "ClientAuthentication"
	CERTIFICATEKEYUSAGE_CODE_SIGNING          CertificateKeyUsage = "CodeSigning"
	CERTIFICATEKEYUSAGE_EMAIL_PROTECTION      CertificateKeyUsage = "EmailProtection"
	CERTIFICATEKEYUSAGE_TIMESTAMPING          CertificateKeyUsage = "Timestamping"
	CERTIFICATEKEYUSAGE_OCSP_SIGNING          CertificateKeyUsage = "OCSPSigning"
)

// AllowedCertificateKeyUsageEnumValues is all the allowed values of CertificateKeyUsage enum
var AllowedCertificateKeyUsageEnumValues = []CertificateKeyUsage{
	"DigitalSignature",
	"NonRepudiation",
	"KeyEncipherment",
	"DataEncipherment",
	"KeyAgreement",
	"KeyCertSign",
	"CRLSigning",
	"EncipherOnly",
	"DecipherOnly",
	"ServerAuthentication",
	"ClientAuthentication",
	"CodeSigning",
	"EmailProtection",
	"Timestamping",
	"OCSPSigning",
}

// validCertificateKeyUsageEnumValue provides a map of CertificateKeyUsages for fast verification of use input
var validCertificateKeyUsageEnumValues = map[CertificateKeyUsage]struct{}{
	"DigitalSignature":     {},
	"NonRepudiation":       {},
	"KeyEncipherment":      {},
	"DataEncipherment":     {},
	"KeyAgreement":         {},
	"KeyCertSign":          {},
	"CRLSigning":           {},
	"EncipherOnly":         {},
	"DecipherOnly":         {},
	"ServerAuthentication": {},
	"ClientAuthentication": {},
	"CodeSigning":          {},
	"EmailProtection":      {},
	"Timestamping":         {},
	"OCSPSigning":          {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CertificateKeyUsage) IsValid() bool {
	_, ok := validCertificateKeyUsageEnumValues[v]
	return ok
}

// NewCertificateKeyUsageFromValue returns a pointer to a valid CertificateKeyUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCertificateKeyUsageFromValue(v string) (CertificateKeyUsage, error) {
	ev := CertificateKeyUsage(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CertificateKeyUsage: valid values are %v", v, AllowedCertificateKeyUsageEnumValues)
}

// AssertCertificateKeyUsageRequired checks if the required fields are not zero-ed
func AssertCertificateKeyUsageRequired(obj CertificateKeyUsage) error {
	return nil
}

// AssertCertificateKeyUsageConstraints checks if the values respects the defined constraints
func AssertCertificateKeyUsageConstraints(obj CertificateKeyUsage) error {
	return nil
}