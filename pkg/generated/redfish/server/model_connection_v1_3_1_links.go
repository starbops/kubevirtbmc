// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ConnectionV131Links - The links to other resources that are related to this resource.
type ConnectionV131Links struct {

	// An array of links to the initiator endpoint groups that are associated with this connection.
	InitiatorEndpointGroups []OdataV4IdRef `json:"InitiatorEndpointGroups,omitempty"`

	// The number of items in a collection.
	InitiatorEndpointGroupsodataCount int64 `json:"InitiatorEndpointGroups@odata.count,omitempty"`

	// An array of links to the initiator endpoints that are associated with this connection.
	InitiatorEndpoints []OdataV4IdRef `json:"InitiatorEndpoints,omitempty"`

	// The number of items in a collection.
	InitiatorEndpointsodataCount int64 `json:"InitiatorEndpoints@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to the target endpoint groups that are associated with this connection.
	TargetEndpointGroups []OdataV4IdRef `json:"TargetEndpointGroups,omitempty"`

	// The number of items in a collection.
	TargetEndpointGroupsodataCount int64 `json:"TargetEndpointGroups@odata.count,omitempty"`

	// An array of links to the target endpoints that are associated with this connection.
	TargetEndpoints []OdataV4IdRef `json:"TargetEndpoints,omitempty"`

	// The number of items in a collection.
	TargetEndpointsodataCount int64 `json:"TargetEndpoints@odata.count,omitempty"`
}

// AssertConnectionV131LinksRequired checks if the required fields are not zero-ed
func AssertConnectionV131LinksRequired(obj ConnectionV131Links) error {
	for _, el := range obj.InitiatorEndpointGroups {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.InitiatorEndpoints {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TargetEndpointGroups {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TargetEndpoints {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertConnectionV131LinksConstraints checks if the values respects the defined constraints
func AssertConnectionV131LinksConstraints(obj ConnectionV131Links) error {
	for _, el := range obj.InitiatorEndpointGroups {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.InitiatorEndpoints {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TargetEndpointGroups {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TargetEndpoints {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}