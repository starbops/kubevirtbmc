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

type LogEntryV1170SensorType string

// List of LogEntryV1170SensorType
const (
	LOGENTRYV1170SENSORTYPE_PLATFORM_SECURITY_VIOLATION_ATTEMPT LogEntryV1170SensorType = "Platform Security Violation Attempt"
	LOGENTRYV1170SENSORTYPE_TEMPERATURE                         LogEntryV1170SensorType = "Temperature"
	LOGENTRYV1170SENSORTYPE_VOLTAGE                             LogEntryV1170SensorType = "Voltage"
	LOGENTRYV1170SENSORTYPE_CURRENT                             LogEntryV1170SensorType = "Current"
	LOGENTRYV1170SENSORTYPE_FAN                                 LogEntryV1170SensorType = "Fan"
	LOGENTRYV1170SENSORTYPE_PHYSICAL_CHASSIS_SECURITY           LogEntryV1170SensorType = "Physical Chassis Security"
	LOGENTRYV1170SENSORTYPE_PROCESSOR                           LogEntryV1170SensorType = "Processor"
	LOGENTRYV1170SENSORTYPE_POWER_SUPPLY___CONVERTER            LogEntryV1170SensorType = "Power Supply / Converter"
	LOGENTRYV1170SENSORTYPE_POWER_UNIT                          LogEntryV1170SensorType = "PowerUnit"
	LOGENTRYV1170SENSORTYPE_COOLING_DEVICE                      LogEntryV1170SensorType = "CoolingDevice"
	LOGENTRYV1170SENSORTYPE_OTHER_UNITS_BASED_SENSOR            LogEntryV1170SensorType = "Other Units-based Sensor"
	LOGENTRYV1170SENSORTYPE_MEMORY                              LogEntryV1170SensorType = "Memory"
	LOGENTRYV1170SENSORTYPE_DRIVE_SLOT_BAY                      LogEntryV1170SensorType = "Drive Slot/Bay"
	LOGENTRYV1170SENSORTYPE_POST_MEMORY_RESIZE                  LogEntryV1170SensorType = "POST Memory Resize"
	LOGENTRYV1170SENSORTYPE_SYSTEM_FIRMWARE_PROGRESS            LogEntryV1170SensorType = "System Firmware Progress"
	LOGENTRYV1170SENSORTYPE_EVENT_LOGGING_DISABLED              LogEntryV1170SensorType = "Event Logging Disabled"
	LOGENTRYV1170SENSORTYPE_SYSTEM_EVENT                        LogEntryV1170SensorType = "System Event"
	LOGENTRYV1170SENSORTYPE_CRITICAL_INTERRUPT                  LogEntryV1170SensorType = "Critical Interrupt"
	LOGENTRYV1170SENSORTYPE_BUTTON_SWITCH                       LogEntryV1170SensorType = "Button/Switch"
	LOGENTRYV1170SENSORTYPE_MODULE_BOARD                        LogEntryV1170SensorType = "Module/Board"
	LOGENTRYV1170SENSORTYPE_MICROCONTROLLER_COPROCESSOR         LogEntryV1170SensorType = "Microcontroller/Coprocessor"
	LOGENTRYV1170SENSORTYPE_ADD_IN_CARD                         LogEntryV1170SensorType = "Add-in Card"
	LOGENTRYV1170SENSORTYPE_CHASSIS                             LogEntryV1170SensorType = "Chassis"
	LOGENTRYV1170SENSORTYPE_CHIP_SET                            LogEntryV1170SensorType = "ChipSet"
	LOGENTRYV1170SENSORTYPE_OTHER_FRU                           LogEntryV1170SensorType = "Other FRU"
	LOGENTRYV1170SENSORTYPE_CABLE_INTERCONNECT                  LogEntryV1170SensorType = "Cable/Interconnect"
	LOGENTRYV1170SENSORTYPE_TERMINATOR                          LogEntryV1170SensorType = "Terminator"
	LOGENTRYV1170SENSORTYPE_SYSTEM_BOOT_RESTART                 LogEntryV1170SensorType = "SystemBoot/Restart"
	LOGENTRYV1170SENSORTYPE_BOOT_ERROR                          LogEntryV1170SensorType = "Boot Error"
	LOGENTRYV1170SENSORTYPE_BASE_OS_BOOT_INSTALLATION_STATUS    LogEntryV1170SensorType = "BaseOSBoot/InstallationStatus"
	LOGENTRYV1170SENSORTYPE_OS_STOP_SHUTDOWN                    LogEntryV1170SensorType = "OS Stop/Shutdown"
	LOGENTRYV1170SENSORTYPE_SLOT_CONNECTOR                      LogEntryV1170SensorType = "Slot/Connector"
	LOGENTRYV1170SENSORTYPE_SYSTEM_ACPI_POWER_STATE             LogEntryV1170SensorType = "System ACPI PowerState"
	LOGENTRYV1170SENSORTYPE_WATCHDOG                            LogEntryV1170SensorType = "Watchdog"
	LOGENTRYV1170SENSORTYPE_PLATFORM_ALERT                      LogEntryV1170SensorType = "Platform Alert"
	LOGENTRYV1170SENSORTYPE_ENTITY_PRESENCE                     LogEntryV1170SensorType = "Entity Presence"
	LOGENTRYV1170SENSORTYPE_MONITOR_ASIC_IC                     LogEntryV1170SensorType = "Monitor ASIC/IC"
	LOGENTRYV1170SENSORTYPE_LAN                                 LogEntryV1170SensorType = "LAN"
	LOGENTRYV1170SENSORTYPE_MANAGEMENT_SUBSYSTEM_HEALTH         LogEntryV1170SensorType = "Management Subsystem Health"
	LOGENTRYV1170SENSORTYPE_BATTERY                             LogEntryV1170SensorType = "Battery"
	LOGENTRYV1170SENSORTYPE_SESSION_AUDIT                       LogEntryV1170SensorType = "Session Audit"
	LOGENTRYV1170SENSORTYPE_VERSION_CHANGE                      LogEntryV1170SensorType = "Version Change"
	LOGENTRYV1170SENSORTYPE_FRU_STATE                           LogEntryV1170SensorType = "FRUState"
	LOGENTRYV1170SENSORTYPE_OEM                                 LogEntryV1170SensorType = "OEM"
)

// AllowedLogEntryV1170SensorTypeEnumValues is all the allowed values of LogEntryV1170SensorType enum
var AllowedLogEntryV1170SensorTypeEnumValues = []LogEntryV1170SensorType{
	"Platform Security Violation Attempt",
	"Temperature",
	"Voltage",
	"Current",
	"Fan",
	"Physical Chassis Security",
	"Processor",
	"Power Supply / Converter",
	"PowerUnit",
	"CoolingDevice",
	"Other Units-based Sensor",
	"Memory",
	"Drive Slot/Bay",
	"POST Memory Resize",
	"System Firmware Progress",
	"Event Logging Disabled",
	"System Event",
	"Critical Interrupt",
	"Button/Switch",
	"Module/Board",
	"Microcontroller/Coprocessor",
	"Add-in Card",
	"Chassis",
	"ChipSet",
	"Other FRU",
	"Cable/Interconnect",
	"Terminator",
	"SystemBoot/Restart",
	"Boot Error",
	"BaseOSBoot/InstallationStatus",
	"OS Stop/Shutdown",
	"Slot/Connector",
	"System ACPI PowerState",
	"Watchdog",
	"Platform Alert",
	"Entity Presence",
	"Monitor ASIC/IC",
	"LAN",
	"Management Subsystem Health",
	"Battery",
	"Session Audit",
	"Version Change",
	"FRUState",
	"OEM",
}

// validLogEntryV1170SensorTypeEnumValue provides a map of LogEntryV1170SensorTypes for fast verification of use input
var validLogEntryV1170SensorTypeEnumValues = map[LogEntryV1170SensorType]struct{}{
	"Platform Security Violation Attempt": {},
	"Temperature":                         {},
	"Voltage":                             {},
	"Current":                             {},
	"Fan":                                 {},
	"Physical Chassis Security":           {},
	"Processor":                           {},
	"Power Supply / Converter":            {},
	"PowerUnit":                           {},
	"CoolingDevice":                       {},
	"Other Units-based Sensor":            {},
	"Memory":                              {},
	"Drive Slot/Bay":                      {},
	"POST Memory Resize":                  {},
	"System Firmware Progress":            {},
	"Event Logging Disabled":              {},
	"System Event":                        {},
	"Critical Interrupt":                  {},
	"Button/Switch":                       {},
	"Module/Board":                        {},
	"Microcontroller/Coprocessor":         {},
	"Add-in Card":                         {},
	"Chassis":                             {},
	"ChipSet":                             {},
	"Other FRU":                           {},
	"Cable/Interconnect":                  {},
	"Terminator":                          {},
	"SystemBoot/Restart":                  {},
	"Boot Error":                          {},
	"BaseOSBoot/InstallationStatus":       {},
	"OS Stop/Shutdown":                    {},
	"Slot/Connector":                      {},
	"System ACPI PowerState":              {},
	"Watchdog":                            {},
	"Platform Alert":                      {},
	"Entity Presence":                     {},
	"Monitor ASIC/IC":                     {},
	"LAN":                                 {},
	"Management Subsystem Health":         {},
	"Battery":                             {},
	"Session Audit":                       {},
	"Version Change":                      {},
	"FRUState":                            {},
	"OEM":                                 {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogEntryV1170SensorType) IsValid() bool {
	_, ok := validLogEntryV1170SensorTypeEnumValues[v]
	return ok
}

// NewLogEntryV1170SensorTypeFromValue returns a pointer to a valid LogEntryV1170SensorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogEntryV1170SensorTypeFromValue(v string) (LogEntryV1170SensorType, error) {
	ev := LogEntryV1170SensorType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for LogEntryV1170SensorType: valid values are %v", v, AllowedLogEntryV1170SensorTypeEnumValues)
}

// AssertLogEntryV1170SensorTypeRequired checks if the required fields are not zero-ed
func AssertLogEntryV1170SensorTypeRequired(obj LogEntryV1170SensorType) error {
	return nil
}

// AssertLogEntryV1170SensorTypeConstraints checks if the values respects the defined constraints
func AssertLogEntryV1170SensorTypeConstraints(obj LogEntryV1170SensorType) error {
	return nil
}