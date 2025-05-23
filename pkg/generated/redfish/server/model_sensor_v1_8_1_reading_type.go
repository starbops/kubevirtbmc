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

type SensorV181ReadingType string

// List of SensorV181ReadingType
const (
	SENSORV181READINGTYPE_TEMPERATURE       SensorV181ReadingType = "Temperature"
	SENSORV181READINGTYPE_HUMIDITY          SensorV181ReadingType = "Humidity"
	SENSORV181READINGTYPE_POWER             SensorV181ReadingType = "Power"
	SENSORV181READINGTYPE_ENERGYK_WH        SensorV181ReadingType = "EnergykWh"
	SENSORV181READINGTYPE_ENERGY_JOULES     SensorV181ReadingType = "EnergyJoules"
	SENSORV181READINGTYPE_ENERGY_WH         SensorV181ReadingType = "EnergyWh"
	SENSORV181READINGTYPE_CHARGE_AH         SensorV181ReadingType = "ChargeAh"
	SENSORV181READINGTYPE_VOLTAGE           SensorV181ReadingType = "Voltage"
	SENSORV181READINGTYPE_CURRENT           SensorV181ReadingType = "Current"
	SENSORV181READINGTYPE_FREQUENCY         SensorV181ReadingType = "Frequency"
	SENSORV181READINGTYPE_PRESSURE          SensorV181ReadingType = "Pressure"
	SENSORV181READINGTYPE_PRESSUREK_PA      SensorV181ReadingType = "PressurekPa"
	SENSORV181READINGTYPE_PRESSURE_PA       SensorV181ReadingType = "PressurePa"
	SENSORV181READINGTYPE_LIQUID_LEVEL      SensorV181ReadingType = "LiquidLevel"
	SENSORV181READINGTYPE_ROTATIONAL        SensorV181ReadingType = "Rotational"
	SENSORV181READINGTYPE_AIR_FLOW          SensorV181ReadingType = "AirFlow"
	SENSORV181READINGTYPE_AIR_FLOW_CMM      SensorV181ReadingType = "AirFlowCMM"
	SENSORV181READINGTYPE_LIQUID_FLOW       SensorV181ReadingType = "LiquidFlow"
	SENSORV181READINGTYPE_LIQUID_FLOW_LPM   SensorV181ReadingType = "LiquidFlowLPM"
	SENSORV181READINGTYPE_BAROMETRIC        SensorV181ReadingType = "Barometric"
	SENSORV181READINGTYPE_ALTITUDE          SensorV181ReadingType = "Altitude"
	SENSORV181READINGTYPE_PERCENT           SensorV181ReadingType = "Percent"
	SENSORV181READINGTYPE_ABSOLUTE_HUMIDITY SensorV181ReadingType = "AbsoluteHumidity"
	SENSORV181READINGTYPE_HEAT              SensorV181ReadingType = "Heat"
)

// AllowedSensorV181ReadingTypeEnumValues is all the allowed values of SensorV181ReadingType enum
var AllowedSensorV181ReadingTypeEnumValues = []SensorV181ReadingType{
	"Temperature",
	"Humidity",
	"Power",
	"EnergykWh",
	"EnergyJoules",
	"EnergyWh",
	"ChargeAh",
	"Voltage",
	"Current",
	"Frequency",
	"Pressure",
	"PressurekPa",
	"PressurePa",
	"LiquidLevel",
	"Rotational",
	"AirFlow",
	"AirFlowCMM",
	"LiquidFlow",
	"LiquidFlowLPM",
	"Barometric",
	"Altitude",
	"Percent",
	"AbsoluteHumidity",
	"Heat",
}

// validSensorV181ReadingTypeEnumValue provides a map of SensorV181ReadingTypes for fast verification of use input
var validSensorV181ReadingTypeEnumValues = map[SensorV181ReadingType]struct{}{
	"Temperature":      {},
	"Humidity":         {},
	"Power":            {},
	"EnergykWh":        {},
	"EnergyJoules":     {},
	"EnergyWh":         {},
	"ChargeAh":         {},
	"Voltage":          {},
	"Current":          {},
	"Frequency":        {},
	"Pressure":         {},
	"PressurekPa":      {},
	"PressurePa":       {},
	"LiquidLevel":      {},
	"Rotational":       {},
	"AirFlow":          {},
	"AirFlowCMM":       {},
	"LiquidFlow":       {},
	"LiquidFlowLPM":    {},
	"Barometric":       {},
	"Altitude":         {},
	"Percent":          {},
	"AbsoluteHumidity": {},
	"Heat":             {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SensorV181ReadingType) IsValid() bool {
	_, ok := validSensorV181ReadingTypeEnumValues[v]
	return ok
}

// NewSensorV181ReadingTypeFromValue returns a pointer to a valid SensorV181ReadingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSensorV181ReadingTypeFromValue(v string) (SensorV181ReadingType, error) {
	ev := SensorV181ReadingType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for SensorV181ReadingType: valid values are %v", v, AllowedSensorV181ReadingTypeEnumValues)
}

// AssertSensorV181ReadingTypeRequired checks if the required fields are not zero-ed
func AssertSensorV181ReadingTypeRequired(obj SensorV181ReadingType) error {
	return nil
}

// AssertSensorV181ReadingTypeConstraints checks if the values respects the defined constraints
func AssertSensorV181ReadingTypeConstraints(obj SensorV181ReadingType) error {
	return nil
}
