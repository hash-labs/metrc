package metrc

import (
	"encoding/json"
	"fmt"
)

// UnitsOfMeasure represents a UnitOfMeasure in Metrc in the GET request.
// See: https://api-ca.metrc.com/Documentation/#UnitsOfMeasure.get_unitsofmeasure_v1_active
type UnitsOfMeasure struct {
	QuantityType string `json:"QuantityType"`
	Name         string `json:"Name"`
	Abbreviation string `json:"Abbreviation"`
}

// GetUnitsOfMeasure will get the units of measure.
// See: https://api-ca.metrc.com/Documentation/#UnitsOfMeasure.get_unitsofmeasure_v1_active
func (m *Metrc) GetUnitsOfMeasure() ([]UnitsOfMeasure, error) {
	endpoint := "unitsofmeasure/v1/active"

	var ur []UnitsOfMeasure
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return ur, fmt.Errorf("could not get units of measure: %s", err)
	}

	err = json.Unmarshal(responseBody, &ur)
	if err != nil {
		return ur, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ur, nil
}
