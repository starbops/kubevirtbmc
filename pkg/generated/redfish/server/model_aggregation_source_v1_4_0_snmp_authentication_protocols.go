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

type AggregationSourceV140SnmpAuthenticationProtocols string

// List of AggregationSourceV140SnmpAuthenticationProtocols
const (
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_NONE             AggregationSourceV140SnmpAuthenticationProtocols = "None"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_COMMUNITY_STRING AggregationSourceV140SnmpAuthenticationProtocols = "CommunityString"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_HMAC_MD5         AggregationSourceV140SnmpAuthenticationProtocols = "HMAC_MD5"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_HMAC_SHA96       AggregationSourceV140SnmpAuthenticationProtocols = "HMAC_SHA96"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_HMAC128_SHA224   AggregationSourceV140SnmpAuthenticationProtocols = "HMAC128_SHA224"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_HMAC192_SHA256   AggregationSourceV140SnmpAuthenticationProtocols = "HMAC192_SHA256"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_HMAC256_SHA384   AggregationSourceV140SnmpAuthenticationProtocols = "HMAC256_SHA384"
	AGGREGATIONSOURCEV140SNMPAUTHENTICATIONPROTOCOLS_HMAC384_SHA512   AggregationSourceV140SnmpAuthenticationProtocols = "HMAC384_SHA512"
)

// AllowedAggregationSourceV140SnmpAuthenticationProtocolsEnumValues is all the allowed values of AggregationSourceV140SnmpAuthenticationProtocols enum
var AllowedAggregationSourceV140SnmpAuthenticationProtocolsEnumValues = []AggregationSourceV140SnmpAuthenticationProtocols{
	"None",
	"CommunityString",
	"HMAC_MD5",
	"HMAC_SHA96",
	"HMAC128_SHA224",
	"HMAC192_SHA256",
	"HMAC256_SHA384",
	"HMAC384_SHA512",
}

// validAggregationSourceV140SnmpAuthenticationProtocolsEnumValue provides a map of AggregationSourceV140SnmpAuthenticationProtocolss for fast verification of use input
var validAggregationSourceV140SnmpAuthenticationProtocolsEnumValues = map[AggregationSourceV140SnmpAuthenticationProtocols]struct{}{
	"None":            {},
	"CommunityString": {},
	"HMAC_MD5":        {},
	"HMAC_SHA96":      {},
	"HMAC128_SHA224":  {},
	"HMAC192_SHA256":  {},
	"HMAC256_SHA384":  {},
	"HMAC384_SHA512":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AggregationSourceV140SnmpAuthenticationProtocols) IsValid() bool {
	_, ok := validAggregationSourceV140SnmpAuthenticationProtocolsEnumValues[v]
	return ok
}

// NewAggregationSourceV140SnmpAuthenticationProtocolsFromValue returns a pointer to a valid AggregationSourceV140SnmpAuthenticationProtocols
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAggregationSourceV140SnmpAuthenticationProtocolsFromValue(v string) (AggregationSourceV140SnmpAuthenticationProtocols, error) {
	ev := AggregationSourceV140SnmpAuthenticationProtocols(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for AggregationSourceV140SnmpAuthenticationProtocols: valid values are %v", v, AllowedAggregationSourceV140SnmpAuthenticationProtocolsEnumValues)
}

// AssertAggregationSourceV140SnmpAuthenticationProtocolsRequired checks if the required fields are not zero-ed
func AssertAggregationSourceV140SnmpAuthenticationProtocolsRequired(obj AggregationSourceV140SnmpAuthenticationProtocols) error {
	return nil
}

// AssertAggregationSourceV140SnmpAuthenticationProtocolsConstraints checks if the values respects the defined constraints
func AssertAggregationSourceV140SnmpAuthenticationProtocolsConstraints(obj AggregationSourceV140SnmpAuthenticationProtocols) error {
	return nil
}
