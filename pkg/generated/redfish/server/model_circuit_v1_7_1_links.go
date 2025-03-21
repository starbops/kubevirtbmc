// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// CircuitV171Links - The links to other resources that are related to this resource.
type CircuitV171Links struct {
	BranchCircuit OdataV4IdRef `json:"BranchCircuit,omitempty"`

	// An array of links to the circuits powered by this circuit.
	DistributionCircuits []OdataV4IdRef `json:"DistributionCircuits,omitempty"`

	// The number of items in a collection.
	DistributionCircuitsodataCount int64 `json:"DistributionCircuits@odata.count,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	// An array of references to the outlets contained by this circuit.
	Outlets []OdataV4IdRef `json:"Outlets,omitempty"`

	// The number of items in a collection.
	OutletsodataCount int64 `json:"Outlets@odata.count,omitempty"`

	PowerOutlet OdataV4IdRef `json:"PowerOutlet,omitempty"`

	SourceCircuit OdataV4IdRef `json:"SourceCircuit,omitempty"`
}

// AssertCircuitV171LinksRequired checks if the required fields are not zero-ed
func AssertCircuitV171LinksRequired(obj CircuitV171Links) error {
	if err := AssertOdataV4IdRefRequired(obj.BranchCircuit); err != nil {
		return err
	}
	for _, el := range obj.DistributionCircuits {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Outlets {
		if err := AssertOdataV4IdRefRequired(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefRequired(obj.PowerOutlet); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.SourceCircuit); err != nil {
		return err
	}
	return nil
}

// AssertCircuitV171LinksConstraints checks if the values respects the defined constraints
func AssertCircuitV171LinksConstraints(obj CircuitV171Links) error {
	if err := AssertOdataV4IdRefConstraints(obj.BranchCircuit); err != nil {
		return err
	}
	for _, el := range obj.DistributionCircuits {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Outlets {
		if err := AssertOdataV4IdRefConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertOdataV4IdRefConstraints(obj.PowerOutlet); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.SourceCircuit); err != nil {
		return err
	}
	return nil
}
