package metrc

import (
	"encoding/json"
	"fmt"
)

// Strain represents a Strain in Metrc.
// See: https://api-ca.metrc.com/Documentation/#Strains.get_strains_v1_{id}
type Strain struct {
	Id               int     `json:"Id,omitempty"`
	Name             string  `json:"Name"`
	TestingStatus    string  `json:"TestingStatus"`
	ThcLevel         float64 `json:"ThcLevel"`
	CbdLevel         float64 `json:"CbdLevel"`
	IndicaPercentage float64 `json:"IndicaPercentage"`
	SativaPercentage float64 `json:"SativaPercentage"`
}

// GetStrainsById gets Strains with an Id.
// See: https://api-ca.metrc.com/Documentation/#Strains.get_strains_v1_{id}
func (m *Metrc) GetStrainsById(id int, licenseNumber *string) (Strain, error) {
	endpoint := fmt.Sprintf("strains/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var sr Strain
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return sr, fmt.Errorf("could not get strains by id response: %s", err)
	}

	err = json.Unmarshal(responseBody, &sr)
	if err != nil {
		return sr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return sr, nil
}

// GetStrainsActive gets all active Strains.
// See: https://api-ca.metrc.com/Documentation/#Strains.get_strains_v1_active
func (m *Metrc) GetStrainsActive(licenseNumber *string) ([]Strain, error) {
	endpoint := "strains/v1/active"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var sr []Strain
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return sr, fmt.Errorf("could not get active strains from metrc: %s", err)
	}

	err = json.Unmarshal(responseBody, &sr)
	if err != nil {
		return sr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return sr, nil
}

// PostStrainsCreate creates new Strains.
// See: https://api-ca.metrc.com/Documentation/#Strains.post_strains_v1_create
func (m *Metrc) PostStrainsCreate(strains []Strain, licenseNumber *string) ([]byte, error) {
	endpoint := "strains/v1/create"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	body, err := json.Marshal(strains)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal strains: %s", err)
	}

	// TODO: Add more robust response handling for posts.
	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting strains: %s", err)
	}
	return resp, nil
}

// PostStrainsUpdate updates existing Strains.
// See: https://api-ca.metrc.com/Documentation/#Strains.post_strains_v1_update
func (m *Metrc) PostStrainsUpdate(strains []Strain, licenseNumber *string) ([]byte, error) {
	endpoint := "strains/v1/update"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	body, err := json.Marshal(strains)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal strains: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting strains: %s", err)
	}
	return resp, nil
}

// DeleteStrainById deletes an existing Strain by Id.
// See: https://api-ca.metrc.com/Documentation/#Strains.delete_strains_v1_{id}
func (m *Metrc) DeleteStrainById(id int, licenseNumber *string) ([]byte, error) {
	endpoint := fmt.Sprintf("strains/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	resp, err := m.Client.Delete(endpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("failed deleting strain %d: %s", id, err)
	}
	return resp, nil
}
