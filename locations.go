package metrc

import (
	"encoding/json"
	"fmt"
)

// LocationGet represents a Location in Metrc in the GET requests.
// See: https://api-ca.metrc.com/Documentation/#Locations.get_locations_v1_{id}
type LocationGet struct {
	Id              int    `json:"Id"`
	Name            string `json:"Name"`
	TypeID          int    `json:"LocationTypeId"`
	TypeName        string `json:"LocationTypeName"`
	ForPlantBatches bool   `json:"ForPlantBatches"`
	ForPlants       bool   `json:"ForPlants"`
	ForHarvests     bool   `json:"ForHarvests"`
	ForPackages     bool   `json:"ForPackages"`
}

// LocationPost represents a Location in Metrc in the POST requests.
// See: https://api-ca.metrc.com/Documentation/#Locations.post_locations_v1_update
type LocationPost struct {
	Id       int    `json:"Id,omitempty"`
	Name     string `json:"Name"`
	TypeName string `json:"LocationTypeName"`
}

// GetLocationsById gets locations with an ID.
// See: https://api-ca.metrc.com/Documentation/#Locations.get_locations_v1_{id}
func (m *Metrc) GetLocationsById(id int, licenseNumber *string) (LocationGet, error) {
	endpoint := fmt.Sprintf("locations/v1/%d", id)

	if licenseNumber != nil {
		endpoint = fmt.Sprintf("%s?licenseNumber=%s", endpoint, *licenseNumber)
	}

	var lr LocationGet
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return lr, fmt.Errorf("could not get locations by id response: %s", err)
	}

	err = json.Unmarshal(responseBody, &lr)
	if err != nil {
		return lr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return lr, nil
}

// GetLocationsActive gets all active locations.
// See: https://api-ca.metrc.com/Documentation/#Locations.get_locations_v1_active
func (m *Metrc) GetLocationsActive(licenseNumber *string) ([]LocationGet, error) {
	endpoint := fmt.Sprintf("locations/v1/active")

	if licenseNumber != nil {
		endpoint = fmt.Sprintf("%s?licenseNumber=%s", endpoint, *licenseNumber)
	}

	var lr []LocationGet
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return lr, fmt.Errorf("could not get active locations from metrc: %s", err)
	}

	err = json.Unmarshal(responseBody, &lr)
	if err != nil {
		return lr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return lr, nil
}

// GetLocationsTypes gets the list of active location types.
// See: https://api-ca.metrc.com/Documentation/#Locations.get_locations_v1_types
func (m *Metrc) GetLocationsTypes(licenseNumber *string) ([]LocationGet, error) {
	endpoint := fmt.Sprintf("locations/v1/types")

	if licenseNumber != nil {
		endpoint = fmt.Sprintf("%s?licenseNumber=%s", endpoint, *licenseNumber)
	}

	var lr []LocationGet
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return lr, fmt.Errorf("could not get locations from types: %s", err)
	}

	err = json.Unmarshal(responseBody, &lr)
	if err != nil {
		return lr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return lr, nil
}

// PostLocationsCreate creates new locations.
// See: https://api-ca.metrc.com/Documentation/#Locations.post_locations_v1_create
func (m *Metrc) PostLocationsCreate(locs []LocationPost, licenseNumber *string) ([]byte, error) {
	endpoint := fmt.Sprintf("locations/v1/create")

	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	body, err := json.Marshal(locs)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal locations: %s", err)
	}

	// TODO: Add more robust response handling.
	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting locations: %s", err)
	}
	return resp, nil
}

// PostLocationsUpdate updates existing Locations. Note that for this endpoint, `LocationPost.Id` is required in each `LocationPost` in the input slice.
// See: https://api-ca.metrc.com/Documentation/#Locations.post_locations_v1_update
func (m *Metrc) PostLocationsUpdate(locs []LocationPost, licenseNumber *string) ([]byte, error) {
	endpoint := "locations/v1/update"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	body, err := json.Marshal(locs)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal locations: %s", err)
	}

	// TODO: Add more robust response handling.
	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting locations: %s", err)
	}
	return resp, nil
}

// DeleteLocationById deletes an existing location by ID.
// See: https://api-ca.metrc.com/Documentation/#Locations.delete_locations_v1_{id}
func (m *Metrc) DeleteLocationById(id int, licenseNumber *string) ([]byte, error) {
	endpoint := fmt.Sprintf("locations/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	resp, err := m.client.Delete(endpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("failed deleting location %d: %s", id, err)
	}
	return resp, nil
}
