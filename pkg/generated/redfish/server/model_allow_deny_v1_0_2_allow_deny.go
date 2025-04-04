// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AllowDenyV102AllowDeny - The AllowDeny schema represents a set of allow or deny configurations.
type AllowDenyV102AllowDeny struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions AllowDenyV102Actions `json:"Actions,omitempty"`

	AllowType AllowDenyV102AllowType `json:"AllowType,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The TCP, UDP, or other destination port to which this rule begins application, inclusive.
	DestinationPortLower *int64 `json:"DestinationPortLower,omitempty"`

	// The TCP, UDP, or other destination port to which this rule ends application, inclusive.
	DestinationPortUpper *int64 `json:"DestinationPortUpper,omitempty"`

	Direction AllowDenyV102DataDirection `json:"Direction,omitempty"`

	// The IANA protocol number to which this permission applies.  For TCP, this is `6`.  For UDP, this is `17`.
	IANAProtocolNumber *int64 `json:"IANAProtocolNumber,omitempty"`

	// The lower IP address to which this permission applies.
	IPAddressLower *string `json:"IPAddressLower,omitempty"`

	IPAddressType AllowDenyV102IpAddressType `json:"IPAddressType,omitempty"`

	// The upper IP address to which this permission applies.
	IPAddressUpper *string `json:"IPAddressUpper,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The TCP, UDP, or other source port to which this rule begins application, inclusive.
	SourcePortLower *int64 `json:"SourcePortLower,omitempty"`

	// The TCP, UDP or other source port to which this rule ends application, inclusive.
	SourcePortUpper *int64 `json:"SourcePortUpper,omitempty"`

	// Indicates if this is a permission that only applies to stateful connections.
	StatefulSession *bool `json:"StatefulSession,omitempty"`
}

// AssertAllowDenyV102AllowDenyRequired checks if the required fields are not zero-ed
func AssertAllowDenyV102AllowDenyRequired(obj AllowDenyV102AllowDeny) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAllowDenyV102ActionsRequired(obj.Actions); err != nil {
		return err
	}
	return nil
}

// AssertAllowDenyV102AllowDenyConstraints checks if the values respects the defined constraints
func AssertAllowDenyV102AllowDenyConstraints(obj AllowDenyV102AllowDeny) error {
	if err := AssertAllowDenyV102ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	return nil
}
