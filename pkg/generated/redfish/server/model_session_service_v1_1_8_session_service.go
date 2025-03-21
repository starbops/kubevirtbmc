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
	"errors"
)

// SessionServiceV118SessionService - The SessionService schema describes the session service and its properties, with links to the actual list of sessions.
type SessionServiceV118SessionService struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions SessionServiceV118Actions `json:"Actions,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An indication of whether this service is enabled.  If `true`, this service is enabled.  If `false`, it is disabled, and new sessions cannot be created, old sessions cannot be deleted, and established sessions can continue operating.
	ServiceEnabled *bool `json:"ServiceEnabled,omitempty"`

	// The number of seconds of inactivity that a session can have before the session service closes the session due to inactivity.
	SessionTimeout int64 `json:"SessionTimeout,omitempty"`

	Sessions OdataV4IdRef `json:"Sessions,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`
}

// AssertSessionServiceV118SessionServiceRequired checks if the required fields are not zero-ed
func AssertSessionServiceV118SessionServiceRequired(obj SessionServiceV118SessionService) error {
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

	if err := AssertSessionServiceV118ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Sessions); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	return nil
}

// AssertSessionServiceV118SessionServiceConstraints checks if the values respects the defined constraints
func AssertSessionServiceV118SessionServiceConstraints(obj SessionServiceV118SessionService) error {
	if err := AssertSessionServiceV118ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if obj.SessionTimeout < 30 {
		return &ParsingError{Param: "SessionTimeout", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SessionTimeout > 86400 {
		return &ParsingError{Param: "SessionTimeout", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertOdataV4IdRefConstraints(obj.Sessions); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	return nil
}
