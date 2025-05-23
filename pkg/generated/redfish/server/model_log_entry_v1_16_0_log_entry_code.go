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

type LogEntryV1160LogEntryCode string

// List of LogEntryV1160LogEntryCode
const (
	LOGENTRYV1160LOGENTRYCODE_ASSERT                                                        LogEntryV1160LogEntryCode = "Assert"
	LOGENTRYV1160LOGENTRYCODE_DEASSERT                                                      LogEntryV1160LogEntryCode = "Deassert"
	LOGENTRYV1160LOGENTRYCODE_LOWER_NON_CRITICAL___GOING_LOW                                LogEntryV1160LogEntryCode = "Lower Non-critical - going low"
	LOGENTRYV1160LOGENTRYCODE_LOWER_NON_CRITICAL___GOING_HIGH                               LogEntryV1160LogEntryCode = "Lower Non-critical - going high"
	LOGENTRYV1160LOGENTRYCODE_LOWER_CRITICAL___GOING_LOW                                    LogEntryV1160LogEntryCode = "Lower Critical - going low"
	LOGENTRYV1160LOGENTRYCODE_LOWER_CRITICAL___GOING_HIGH                                   LogEntryV1160LogEntryCode = "Lower Critical - going high"
	LOGENTRYV1160LOGENTRYCODE_LOWER_NON_RECOVERABLE___GOING_LOW                             LogEntryV1160LogEntryCode = "Lower Non-recoverable - going low"
	LOGENTRYV1160LOGENTRYCODE_LOWER_NON_RECOVERABLE___GOING_HIGH                            LogEntryV1160LogEntryCode = "Lower Non-recoverable - going high"
	LOGENTRYV1160LOGENTRYCODE_UPPER_NON_CRITICAL___GOING_LOW                                LogEntryV1160LogEntryCode = "Upper Non-critical - going low"
	LOGENTRYV1160LOGENTRYCODE_UPPER_NON_CRITICAL___GOING_HIGH                               LogEntryV1160LogEntryCode = "Upper Non-critical - going high"
	LOGENTRYV1160LOGENTRYCODE_UPPER_CRITICAL___GOING_LOW                                    LogEntryV1160LogEntryCode = "Upper Critical - going low"
	LOGENTRYV1160LOGENTRYCODE_UPPER_CRITICAL___GOING_HIGH                                   LogEntryV1160LogEntryCode = "Upper Critical - going high"
	LOGENTRYV1160LOGENTRYCODE_UPPER_NON_RECOVERABLE___GOING_LOW                             LogEntryV1160LogEntryCode = "Upper Non-recoverable - going low"
	LOGENTRYV1160LOGENTRYCODE_UPPER_NON_RECOVERABLE___GOING_HIGH                            LogEntryV1160LogEntryCode = "Upper Non-recoverable - going high"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_IDLE                                            LogEntryV1160LogEntryCode = "Transition to Idle"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_ACTIVE                                          LogEntryV1160LogEntryCode = "Transition to Active"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_BUSY                                            LogEntryV1160LogEntryCode = "Transition to Busy"
	LOGENTRYV1160LOGENTRYCODE_STATE_DEASSERTED                                              LogEntryV1160LogEntryCode = "State Deasserted"
	LOGENTRYV1160LOGENTRYCODE_STATE_ASSERTED                                                LogEntryV1160LogEntryCode = "State Asserted"
	LOGENTRYV1160LOGENTRYCODE_PREDICTIVE_FAILURE_DEASSERTED                                 LogEntryV1160LogEntryCode = "Predictive Failure deasserted"
	LOGENTRYV1160LOGENTRYCODE_PREDICTIVE_FAILURE_ASSERTED                                   LogEntryV1160LogEntryCode = "Predictive Failure asserted"
	LOGENTRYV1160LOGENTRYCODE_LIMIT_NOT_EXCEEDED                                            LogEntryV1160LogEntryCode = "Limit Not Exceeded"
	LOGENTRYV1160LOGENTRYCODE_LIMIT_EXCEEDED                                                LogEntryV1160LogEntryCode = "Limit Exceeded"
	LOGENTRYV1160LOGENTRYCODE_PERFORMANCE_MET                                               LogEntryV1160LogEntryCode = "Performance Met"
	LOGENTRYV1160LOGENTRYCODE_PERFORMANCE_LAGS                                              LogEntryV1160LogEntryCode = "Performance Lags"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_OK                                              LogEntryV1160LogEntryCode = "Transition to OK"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_NON_CRITICAL_FROM_OK                            LogEntryV1160LogEntryCode = "Transition to Non-Critical from OK"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_CRITICAL_FROM_LESS_SEVERE                       LogEntryV1160LogEntryCode = "Transition to Critical from less severe"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_NON_RECOVERABLE_FROM_LESS_SEVERE                LogEntryV1160LogEntryCode = "Transition to Non-recoverable from less severe"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_NON_CRITICAL_FROM_MORE_SEVERE                   LogEntryV1160LogEntryCode = "Transition to Non-Critical from more severe"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_CRITICAL_FROM_NON_RECOVERABLE                   LogEntryV1160LogEntryCode = "Transition to Critical from Non-recoverable"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_NON_RECOVERABLE                                 LogEntryV1160LogEntryCode = "Transition to Non-recoverable"
	LOGENTRYV1160LOGENTRYCODE_MONITOR                                                       LogEntryV1160LogEntryCode = "Monitor"
	LOGENTRYV1160LOGENTRYCODE_INFORMATIONAL                                                 LogEntryV1160LogEntryCode = "Informational"
	LOGENTRYV1160LOGENTRYCODE_DEVICE_REMOVED___DEVICE_ABSENT                                LogEntryV1160LogEntryCode = "Device Removed / Device Absent"
	LOGENTRYV1160LOGENTRYCODE_DEVICE_INSERTED___DEVICE_PRESENT                              LogEntryV1160LogEntryCode = "Device Inserted / Device Present"
	LOGENTRYV1160LOGENTRYCODE_DEVICE_DISABLED                                               LogEntryV1160LogEntryCode = "Device Disabled"
	LOGENTRYV1160LOGENTRYCODE_DEVICE_ENABLED                                                LogEntryV1160LogEntryCode = "Device Enabled"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_RUNNING                                         LogEntryV1160LogEntryCode = "Transition to Running"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_IN_TEST                                         LogEntryV1160LogEntryCode = "Transition to In Test"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_POWER_OFF                                       LogEntryV1160LogEntryCode = "Transition to Power Off"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_ON_LINE                                         LogEntryV1160LogEntryCode = "Transition to On Line"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_OFF_LINE                                        LogEntryV1160LogEntryCode = "Transition to Off Line"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_OFF_DUTY                                        LogEntryV1160LogEntryCode = "Transition to Off Duty"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_DEGRADED                                        LogEntryV1160LogEntryCode = "Transition to Degraded"
	LOGENTRYV1160LOGENTRYCODE_TRANSITION_TO_POWER_SAVE                                      LogEntryV1160LogEntryCode = "Transition to Power Save"
	LOGENTRYV1160LOGENTRYCODE_INSTALL_ERROR                                                 LogEntryV1160LogEntryCode = "Install Error"
	LOGENTRYV1160LOGENTRYCODE_FULLY_REDUNDANT                                               LogEntryV1160LogEntryCode = "Fully Redundant"
	LOGENTRYV1160LOGENTRYCODE_REDUNDANCY_LOST                                               LogEntryV1160LogEntryCode = "Redundancy Lost"
	LOGENTRYV1160LOGENTRYCODE_REDUNDANCY_DEGRADED                                           LogEntryV1160LogEntryCode = "Redundancy Degraded"
	LOGENTRYV1160LOGENTRYCODE_NON_REDUNDANTSUFFICIENT_RESOURCES_FROM_REDUNDANT              LogEntryV1160LogEntryCode = "Non-redundant:Sufficient Resources from Redundant"
	LOGENTRYV1160LOGENTRYCODE_NON_REDUNDANTSUFFICIENT_RESOURCES_FROM_INSUFFICIENT_RESOURCES LogEntryV1160LogEntryCode = "Non-redundant:Sufficient Resources from Insufficient Resources"
	LOGENTRYV1160LOGENTRYCODE_NON_REDUNDANTINSUFFICIENT_RESOURCES                           LogEntryV1160LogEntryCode = "Non-redundant:Insufficient Resources"
	LOGENTRYV1160LOGENTRYCODE_REDUNDANCY_DEGRADED_FROM_FULLY_REDUNDANT                      LogEntryV1160LogEntryCode = "Redundancy Degraded from Fully Redundant"
	LOGENTRYV1160LOGENTRYCODE_REDUNDANCY_DEGRADED_FROM_NON_REDUNDANT                        LogEntryV1160LogEntryCode = "Redundancy Degraded from Non-redundant"
	LOGENTRYV1160LOGENTRYCODE_D0_POWER_STATE                                                LogEntryV1160LogEntryCode = "D0 Power State"
	LOGENTRYV1160LOGENTRYCODE_D1_POWER_STATE                                                LogEntryV1160LogEntryCode = "D1 Power State"
	LOGENTRYV1160LOGENTRYCODE_D2_POWER_STATE                                                LogEntryV1160LogEntryCode = "D2 Power State"
	LOGENTRYV1160LOGENTRYCODE_D3_POWER_STATE                                                LogEntryV1160LogEntryCode = "D3 Power State"
	LOGENTRYV1160LOGENTRYCODE_OEM                                                           LogEntryV1160LogEntryCode = "OEM"
)

// AllowedLogEntryV1160LogEntryCodeEnumValues is all the allowed values of LogEntryV1160LogEntryCode enum
var AllowedLogEntryV1160LogEntryCodeEnumValues = []LogEntryV1160LogEntryCode{
	"Assert",
	"Deassert",
	"Lower Non-critical - going low",
	"Lower Non-critical - going high",
	"Lower Critical - going low",
	"Lower Critical - going high",
	"Lower Non-recoverable - going low",
	"Lower Non-recoverable - going high",
	"Upper Non-critical - going low",
	"Upper Non-critical - going high",
	"Upper Critical - going low",
	"Upper Critical - going high",
	"Upper Non-recoverable - going low",
	"Upper Non-recoverable - going high",
	"Transition to Idle",
	"Transition to Active",
	"Transition to Busy",
	"State Deasserted",
	"State Asserted",
	"Predictive Failure deasserted",
	"Predictive Failure asserted",
	"Limit Not Exceeded",
	"Limit Exceeded",
	"Performance Met",
	"Performance Lags",
	"Transition to OK",
	"Transition to Non-Critical from OK",
	"Transition to Critical from less severe",
	"Transition to Non-recoverable from less severe",
	"Transition to Non-Critical from more severe",
	"Transition to Critical from Non-recoverable",
	"Transition to Non-recoverable",
	"Monitor",
	"Informational",
	"Device Removed / Device Absent",
	"Device Inserted / Device Present",
	"Device Disabled",
	"Device Enabled",
	"Transition to Running",
	"Transition to In Test",
	"Transition to Power Off",
	"Transition to On Line",
	"Transition to Off Line",
	"Transition to Off Duty",
	"Transition to Degraded",
	"Transition to Power Save",
	"Install Error",
	"Fully Redundant",
	"Redundancy Lost",
	"Redundancy Degraded",
	"Non-redundant:Sufficient Resources from Redundant",
	"Non-redundant:Sufficient Resources from Insufficient Resources",
	"Non-redundant:Insufficient Resources",
	"Redundancy Degraded from Fully Redundant",
	"Redundancy Degraded from Non-redundant",
	"D0 Power State",
	"D1 Power State",
	"D2 Power State",
	"D3 Power State",
	"OEM",
}

// validLogEntryV1160LogEntryCodeEnumValue provides a map of LogEntryV1160LogEntryCodes for fast verification of use input
var validLogEntryV1160LogEntryCodeEnumValues = map[LogEntryV1160LogEntryCode]struct{}{
	"Assert":                                            {},
	"Deassert":                                          {},
	"Lower Non-critical - going low":                    {},
	"Lower Non-critical - going high":                   {},
	"Lower Critical - going low":                        {},
	"Lower Critical - going high":                       {},
	"Lower Non-recoverable - going low":                 {},
	"Lower Non-recoverable - going high":                {},
	"Upper Non-critical - going low":                    {},
	"Upper Non-critical - going high":                   {},
	"Upper Critical - going low":                        {},
	"Upper Critical - going high":                       {},
	"Upper Non-recoverable - going low":                 {},
	"Upper Non-recoverable - going high":                {},
	"Transition to Idle":                                {},
	"Transition to Active":                              {},
	"Transition to Busy":                                {},
	"State Deasserted":                                  {},
	"State Asserted":                                    {},
	"Predictive Failure deasserted":                     {},
	"Predictive Failure asserted":                       {},
	"Limit Not Exceeded":                                {},
	"Limit Exceeded":                                    {},
	"Performance Met":                                   {},
	"Performance Lags":                                  {},
	"Transition to OK":                                  {},
	"Transition to Non-Critical from OK":                {},
	"Transition to Critical from less severe":           {},
	"Transition to Non-recoverable from less severe":    {},
	"Transition to Non-Critical from more severe":       {},
	"Transition to Critical from Non-recoverable":       {},
	"Transition to Non-recoverable":                     {},
	"Monitor":                                           {},
	"Informational":                                     {},
	"Device Removed / Device Absent":                    {},
	"Device Inserted / Device Present":                  {},
	"Device Disabled":                                   {},
	"Device Enabled":                                    {},
	"Transition to Running":                             {},
	"Transition to In Test":                             {},
	"Transition to Power Off":                           {},
	"Transition to On Line":                             {},
	"Transition to Off Line":                            {},
	"Transition to Off Duty":                            {},
	"Transition to Degraded":                            {},
	"Transition to Power Save":                          {},
	"Install Error":                                     {},
	"Fully Redundant":                                   {},
	"Redundancy Lost":                                   {},
	"Redundancy Degraded":                               {},
	"Non-redundant:Sufficient Resources from Redundant": {},
	"Non-redundant:Sufficient Resources from Insufficient Resources": {},
	"Non-redundant:Insufficient Resources":                           {},
	"Redundancy Degraded from Fully Redundant":                       {},
	"Redundancy Degraded from Non-redundant":                         {},
	"D0 Power State":                                                 {},
	"D1 Power State":                                                 {},
	"D2 Power State":                                                 {},
	"D3 Power State":                                                 {},
	"OEM":                                                            {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogEntryV1160LogEntryCode) IsValid() bool {
	_, ok := validLogEntryV1160LogEntryCodeEnumValues[v]
	return ok
}

// NewLogEntryV1160LogEntryCodeFromValue returns a pointer to a valid LogEntryV1160LogEntryCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogEntryV1160LogEntryCodeFromValue(v string) (LogEntryV1160LogEntryCode, error) {
	ev := LogEntryV1160LogEntryCode(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for LogEntryV1160LogEntryCode: valid values are %v", v, AllowedLogEntryV1160LogEntryCodeEnumValues)
}

// AssertLogEntryV1160LogEntryCodeRequired checks if the required fields are not zero-ed
func AssertLogEntryV1160LogEntryCodeRequired(obj LogEntryV1160LogEntryCode) error {
	return nil
}

// AssertLogEntryV1160LogEntryCodeConstraints checks if the values respects the defined constraints
func AssertLogEntryV1160LogEntryCodeConstraints(obj LogEntryV1160LogEntryCode) error {
	return nil
}
