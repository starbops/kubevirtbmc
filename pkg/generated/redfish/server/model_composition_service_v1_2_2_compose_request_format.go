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

type CompositionServiceV122ComposeRequestFormat string

// List of CompositionServiceV122ComposeRequestFormat
const (
	COMPOSITIONSERVICEV122COMPOSEREQUESTFORMAT_MANIFEST CompositionServiceV122ComposeRequestFormat = "Manifest"
)

// AllowedCompositionServiceV122ComposeRequestFormatEnumValues is all the allowed values of CompositionServiceV122ComposeRequestFormat enum
var AllowedCompositionServiceV122ComposeRequestFormatEnumValues = []CompositionServiceV122ComposeRequestFormat{
	"Manifest",
}

// validCompositionServiceV122ComposeRequestFormatEnumValue provides a map of CompositionServiceV122ComposeRequestFormats for fast verification of use input
var validCompositionServiceV122ComposeRequestFormatEnumValues = map[CompositionServiceV122ComposeRequestFormat]struct{}{
	"Manifest": {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompositionServiceV122ComposeRequestFormat) IsValid() bool {
	_, ok := validCompositionServiceV122ComposeRequestFormatEnumValues[v]
	return ok
}

// NewCompositionServiceV122ComposeRequestFormatFromValue returns a pointer to a valid CompositionServiceV122ComposeRequestFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompositionServiceV122ComposeRequestFormatFromValue(v string) (CompositionServiceV122ComposeRequestFormat, error) {
	ev := CompositionServiceV122ComposeRequestFormat(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for CompositionServiceV122ComposeRequestFormat: valid values are %v", v, AllowedCompositionServiceV122ComposeRequestFormatEnumValues)
}

// AssertCompositionServiceV122ComposeRequestFormatRequired checks if the required fields are not zero-ed
func AssertCompositionServiceV122ComposeRequestFormatRequired(obj CompositionServiceV122ComposeRequestFormat) error {
	return nil
}

// AssertCompositionServiceV122ComposeRequestFormatConstraints checks if the values respects the defined constraints
func AssertCompositionServiceV122ComposeRequestFormatConstraints(obj CompositionServiceV122ComposeRequestFormat) error {
	return nil
}