// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// LeakDetectorLeakDetectorArrayExcerpt - The `LeakDetector` schema describes a state-based or digital-value leak detector and its properties.
type LeakDetectorLeakDetectorArrayExcerpt struct {

	// The link to the resource that provides the data for this leak detector.
	DataSourceUri *string `json:"DataSourceUri,omitempty"`

	DetectorState *ResourceHealth `json:"DetectorState,omitempty"`

	// The name of the device.
	DeviceName *string `json:"DeviceName,omitempty"`

	PhysicalContext *PhysicalContextPhysicalContext `json:"PhysicalContext,omitempty"`

	PhysicalSubContext *PhysicalContextPhysicalSubContext `json:"PhysicalSubContext,omitempty"`
}

// AssertLeakDetectorLeakDetectorArrayExcerptRequired checks if the required fields are not zero-ed
func AssertLeakDetectorLeakDetectorArrayExcerptRequired(obj LeakDetectorLeakDetectorArrayExcerpt) error {
	return nil
}

// AssertLeakDetectorLeakDetectorArrayExcerptConstraints checks if the values respects the defined constraints
func AssertLeakDetectorLeakDetectorArrayExcerptConstraints(obj LeakDetectorLeakDetectorArrayExcerpt) error {
	return nil
}