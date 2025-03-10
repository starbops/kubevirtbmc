// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// ResourceBlockV142Links - The links to other resources that are related to this resource.
type ResourceBlockV142Links struct {

	// An array of links to the chassis in which this resource block is contained.
	Chassis []OdataV4IdRef `json:"Chassis,omitempty"`

	// The number of items in a collection.
	ChassisodataCount int64 `json:"Chassis@odata.count,omitempty"`

	// An array of links to the computer systems that are composed from this resource block.
	ComputerSystems []OdataV4IdRef `json:"ComputerSystems,omitempty"`

	// The number of items in a collection.
	ComputerSystemsodataCount int64 `json:"ComputerSystems@odata.count,omitempty"`

	// An array of links to resource blocks that depend on this resource block.
	ConsumingResourceBlocks []OdataV4IdRef `json:"ConsumingResourceBlocks,omitempty"`

	// The number of items in a collection.
	ConsumingResourceBlocksodataCount int64 `json:"ConsumingResourceBlocks@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of links to resource blocks that this resource block depends on.
	SupplyingResourceBlocks []OdataV4IdRef `json:"SupplyingResourceBlocks,omitempty"`

	// The number of items in a collection.
	SupplyingResourceBlocksodataCount int64 `json:"SupplyingResourceBlocks@odata.count,omitempty"`

	// An array of links to the zones in which this resource block is bound.
	Zones []OdataV4IdRef `json:"Zones,omitempty"`

	// The number of items in a collection.
	ZonesodataCount int64 `json:"Zones@odata.count,omitempty"`
}

// AssertResourceBlockV142LinksRequired checks if the required fields are not zero-ed
func AssertResourceBlockV142LinksRequired(obj ResourceBlockV142Links) error {
	for _, el := range obj.Chassis {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ComputerSystems {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConsumingResourceBlocks {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SupplyingResourceBlocks {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Zones {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertResourceBlockV142LinksConstraints checks if the values respects the defined constraints
func AssertResourceBlockV142LinksConstraints(obj ResourceBlockV142Links) error {
	for _, el := range obj.Chassis {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ComputerSystems {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ConsumingResourceBlocks {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SupplyingResourceBlocks {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Zones {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
