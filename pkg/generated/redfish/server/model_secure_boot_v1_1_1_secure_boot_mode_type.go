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

type SecureBootV111SecureBootModeType string

// List of SecureBootV111SecureBootModeType
const (
	SECUREBOOTV111SECUREBOOTMODETYPE_SETUP_MODE    SecureBootV111SecureBootModeType = "SetupMode"
	SECUREBOOTV111SECUREBOOTMODETYPE_USER_MODE     SecureBootV111SecureBootModeType = "UserMode"
	SECUREBOOTV111SECUREBOOTMODETYPE_AUDIT_MODE    SecureBootV111SecureBootModeType = "AuditMode"
	SECUREBOOTV111SECUREBOOTMODETYPE_DEPLOYED_MODE SecureBootV111SecureBootModeType = "DeployedMode"
)

// AllowedSecureBootV111SecureBootModeTypeEnumValues is all the allowed values of SecureBootV111SecureBootModeType enum
var AllowedSecureBootV111SecureBootModeTypeEnumValues = []SecureBootV111SecureBootModeType{
	"SetupMode",
	"UserMode",
	"AuditMode",
	"DeployedMode",
}

// validSecureBootV111SecureBootModeTypeEnumValue provides a map of SecureBootV111SecureBootModeTypes for fast verification of use input
var validSecureBootV111SecureBootModeTypeEnumValues = map[SecureBootV111SecureBootModeType]struct{}{
	"SetupMode":    {},
	"UserMode":     {},
	"AuditMode":    {},
	"DeployedMode": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecureBootV111SecureBootModeType) IsValid() bool {
	_, ok := validSecureBootV111SecureBootModeTypeEnumValues[v]
	return ok
}

// NewSecureBootV111SecureBootModeTypeFromValue returns a pointer to a valid SecureBootV111SecureBootModeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecureBootV111SecureBootModeTypeFromValue(v string) (SecureBootV111SecureBootModeType, error) {
	ev := SecureBootV111SecureBootModeType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for SecureBootV111SecureBootModeType: valid values are %v", v, AllowedSecureBootV111SecureBootModeTypeEnumValues)
}

// AssertSecureBootV111SecureBootModeTypeRequired checks if the required fields are not zero-ed
func AssertSecureBootV111SecureBootModeTypeRequired(obj SecureBootV111SecureBootModeType) error {
	return nil
}

// AssertSecureBootV111SecureBootModeTypeConstraints checks if the values respects the defined constraints
func AssertSecureBootV111SecureBootModeTypeConstraints(obj SecureBootV111SecureBootModeType) error {
	return nil
}
