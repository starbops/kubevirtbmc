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

// ScheduleV125DayOfWeek : Days of the week.
type ScheduleV125DayOfWeek string

// List of ScheduleV125DayOfWeek
const (
	SCHEDULEV125DAYOFWEEK_MONDAY    ScheduleV125DayOfWeek = "Monday"
	SCHEDULEV125DAYOFWEEK_TUESDAY   ScheduleV125DayOfWeek = "Tuesday"
	SCHEDULEV125DAYOFWEEK_WEDNESDAY ScheduleV125DayOfWeek = "Wednesday"
	SCHEDULEV125DAYOFWEEK_THURSDAY  ScheduleV125DayOfWeek = "Thursday"
	SCHEDULEV125DAYOFWEEK_FRIDAY    ScheduleV125DayOfWeek = "Friday"
	SCHEDULEV125DAYOFWEEK_SATURDAY  ScheduleV125DayOfWeek = "Saturday"
	SCHEDULEV125DAYOFWEEK_SUNDAY    ScheduleV125DayOfWeek = "Sunday"
	SCHEDULEV125DAYOFWEEK_EVERY     ScheduleV125DayOfWeek = "Every"
)

// AllowedScheduleV125DayOfWeekEnumValues is all the allowed values of ScheduleV125DayOfWeek enum
var AllowedScheduleV125DayOfWeekEnumValues = []ScheduleV125DayOfWeek{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
	"Every",
}

// validScheduleV125DayOfWeekEnumValue provides a map of ScheduleV125DayOfWeeks for fast verification of use input
var validScheduleV125DayOfWeekEnumValues = map[ScheduleV125DayOfWeek]struct{}{
	"Monday":    {},
	"Tuesday":   {},
	"Wednesday": {},
	"Thursday":  {},
	"Friday":    {},
	"Saturday":  {},
	"Sunday":    {},
	"Every":     {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScheduleV125DayOfWeek) IsValid() bool {
	_, ok := validScheduleV125DayOfWeekEnumValues[v]
	return ok
}

// NewScheduleV125DayOfWeekFromValue returns a pointer to a valid ScheduleV125DayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScheduleV125DayOfWeekFromValue(v string) (ScheduleV125DayOfWeek, error) {
	ev := ScheduleV125DayOfWeek(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for ScheduleV125DayOfWeek: valid values are %v", v, AllowedScheduleV125DayOfWeekEnumValues)
}

// AssertScheduleV125DayOfWeekRequired checks if the required fields are not zero-ed
func AssertScheduleV125DayOfWeekRequired(obj ScheduleV125DayOfWeek) error {
	return nil
}

// AssertScheduleV125DayOfWeekConstraints checks if the values respects the defined constraints
func AssertScheduleV125DayOfWeekConstraints(obj ScheduleV125DayOfWeek) error {
	return nil
}
