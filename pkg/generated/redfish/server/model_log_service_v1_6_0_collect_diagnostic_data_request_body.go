// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// LogServiceV160CollectDiagnosticDataRequestBody - The action to collect the diagnostic data for the given type.  When the diagnostic data is collected, a new log entry will be created and the additional data referenced by the new log entry will contain the diagnostic data.
type LogServiceV160CollectDiagnosticDataRequestBody struct {
	DiagnosticDataType LogServiceV160LogDiagnosticDataTypes `json:"DiagnosticDataType"`

	// The OEM-defined type of diagnostic data to collect.
	OEMDiagnosticDataType string `json:"OEMDiagnosticDataType,omitempty"`

	// The password to access the URI specified by the TargetURI parameter.
	Password string `json:"Password,omitempty"`

	// The URI to access when sending the diagnostic data.
	TargetURI string `json:"TargetURI,omitempty"`

	TransferProtocol LogServiceV160TransferProtocolType `json:"TransferProtocol,omitempty"`

	// The user name to access the URI specified by the TargetURI parameter.
	UserName string `json:"UserName,omitempty"`
}

// AssertLogServiceV160CollectDiagnosticDataRequestBodyRequired checks if the required fields are not zero-ed
func AssertLogServiceV160CollectDiagnosticDataRequestBodyRequired(obj LogServiceV160CollectDiagnosticDataRequestBody) error {
	elements := map[string]interface{}{
		"DiagnosticDataType": obj.DiagnosticDataType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertLogServiceV160CollectDiagnosticDataRequestBodyConstraints checks if the values respects the defined constraints
func AssertLogServiceV160CollectDiagnosticDataRequestBodyConstraints(obj LogServiceV160CollectDiagnosticDataRequestBody) error {
	return nil
}
