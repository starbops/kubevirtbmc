// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CertificateServiceV104ReplaceCertificateRequestBody - This action replaces a certificate.
type CertificateServiceV104ReplaceCertificateRequestBody struct {

	// The string for the certificate.
	CertificateString string `json:"CertificateString"`

	CertificateType CertificateCertificateType `json:"CertificateType"`

	CertificateUri OdataV4IdRef `json:"CertificateUri"`
}

// AssertCertificateServiceV104ReplaceCertificateRequestBodyRequired checks if the required fields are not zero-ed
func AssertCertificateServiceV104ReplaceCertificateRequestBodyRequired(obj CertificateServiceV104ReplaceCertificateRequestBody) error {
	elements := map[string]interface{}{
		"CertificateString": obj.CertificateString,
		"CertificateType":   obj.CertificateType,
		"CertificateUri":    obj.CertificateUri,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertOdataV4IdRefRequired(obj.CertificateUri); err != nil {
		return err
	}
	return nil
}

// AssertCertificateServiceV104ReplaceCertificateRequestBodyConstraints checks if the values respects the defined constraints
func AssertCertificateServiceV104ReplaceCertificateRequestBodyConstraints(obj CertificateServiceV104ReplaceCertificateRequestBody) error {
	if err := AssertOdataV4IdRefConstraints(obj.CertificateUri); err != nil {
		return err
	}
	return nil
}
