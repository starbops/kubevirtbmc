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
	"time"
)

// CompositionReservationV101CompositionReservation - The CompositionReservation schema contains reservation information related to the Compose action defined in the CompositionService resource when the RequestType parameter contains the value `PreviewReserve`.
type CompositionReservationV101CompositionReservation struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	Actions CompositionReservationV101Actions `json:"Actions,omitempty"`

	// The client that owns the reservation.
	Client string `json:"Client,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	Manifest ManifestManifest `json:"Manifest,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// The date and time the service created the reservation.
	ReservationTime time.Time `json:"ReservationTime,omitempty"`

	// The array of links to the reserved resource blocks.
	ReservedResourceBlocks []OdataV4IdRef `json:"ReservedResourceBlocks,omitempty"`

	// The number of items in a collection.
	ReservedResourceBlocksodataCount int64 `json:"ReservedResourceBlocks@odata.count,omitempty"`
}

// AssertCompositionReservationV101CompositionReservationRequired checks if the required fields are not zero-ed
func AssertCompositionReservationV101CompositionReservationRequired(obj CompositionReservationV101CompositionReservation) error {
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

	if err := AssertCompositionReservationV101ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertManifestManifestRequired(obj.Manifest); err != nil {
		return err
	}
	for _, el := range obj.ReservedResourceBlocks {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCompositionReservationV101CompositionReservationConstraints checks if the values respects the defined constraints
func AssertCompositionReservationV101CompositionReservationConstraints(obj CompositionReservationV101CompositionReservation) error {
	if err := AssertCompositionReservationV101ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertManifestManifestConstraints(obj.Manifest); err != nil {
		return err
	}
	for _, el := range obj.ReservedResourceBlocks {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}