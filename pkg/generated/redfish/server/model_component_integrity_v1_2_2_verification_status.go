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

type ComponentIntegrityV122VerificationStatus string

// List of ComponentIntegrityV122VerificationStatus
const (
	COMPONENTINTEGRITYV122VERIFICATIONSTATUS_SUCCESS ComponentIntegrityV122VerificationStatus = "Success"
	COMPONENTINTEGRITYV122VERIFICATIONSTATUS_FAILED  ComponentIntegrityV122VerificationStatus = "Failed"
)

// AllowedComponentIntegrityV122VerificationStatusEnumValues is all the allowed values of ComponentIntegrityV122VerificationStatus enum
var AllowedComponentIntegrityV122VerificationStatusEnumValues = []ComponentIntegrityV122VerificationStatus{
	"Success",
	"Failed",
}

// validComponentIntegrityV122VerificationStatusEnumValue provides a map of ComponentIntegrityV122VerificationStatuss for fast verification of use input
var validComponentIntegrityV122VerificationStatusEnumValues = map[ComponentIntegrityV122VerificationStatus]struct{}{
	"Success": {},
	"Failed":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComponentIntegrityV122VerificationStatus) IsValid() bool {
	_, ok := validComponentIntegrityV122VerificationStatusEnumValues[v]
	return ok
}

// NewComponentIntegrityV122VerificationStatusFromValue returns a pointer to a valid ComponentIntegrityV122VerificationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComponentIntegrityV122VerificationStatusFromValue(v string) (ComponentIntegrityV122VerificationStatus, error) {
	ev := ComponentIntegrityV122VerificationStatus(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ComponentIntegrityV122VerificationStatus: valid values are %v", v, AllowedComponentIntegrityV122VerificationStatusEnumValues)
}

// AssertComponentIntegrityV122VerificationStatusRequired checks if the required fields are not zero-ed
func AssertComponentIntegrityV122VerificationStatusRequired(obj ComponentIntegrityV122VerificationStatus) error {
	return nil
}

// AssertComponentIntegrityV122VerificationStatusConstraints checks if the values respects the defined constraints
func AssertComponentIntegrityV122VerificationStatusConstraints(obj ComponentIntegrityV122VerificationStatus) error {
	return nil
}
