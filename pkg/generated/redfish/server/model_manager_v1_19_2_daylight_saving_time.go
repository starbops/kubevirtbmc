// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2024.3
 */

package server

import (
	"time"
)

// ManagerV1192DaylightSavingTime - The daylight saving time settings for a manager.
type ManagerV1192DaylightSavingTime struct {

	// The end date and time with UTC offset of daylight saving time.
	EndDateTime time.Time `json:"EndDateTime,omitempty"`

	// The daylight saving time offset in minutes.
	OffsetMinutes int64 `json:"OffsetMinutes,omitempty"`

	// The start date and time with UTC offset of daylight saving time.
	StartDateTime time.Time `json:"StartDateTime,omitempty"`

	// The time zone of the manager when daylight saving time is in effect.
	TimeZoneName string `json:"TimeZoneName,omitempty"`
}

// AssertManagerV1192DaylightSavingTimeRequired checks if the required fields are not zero-ed
func AssertManagerV1192DaylightSavingTimeRequired(obj ManagerV1192DaylightSavingTime) error {
	return nil
}

// AssertManagerV1192DaylightSavingTimeConstraints checks if the values respects the defined constraints
func AssertManagerV1192DaylightSavingTimeConstraints(obj ManagerV1192DaylightSavingTime) error {
	return nil
}