package metrc

import (
	"encoding/json"
	"fmt"
)

// Harvest is the core data structure for a harvest on Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_{id}.GET
type Harvest struct {
	Id                     *int     `json:"Id"`
	Name                   string   `json:"Name"`
	HarvestType            string   `json:"HarvestType"`
	SourceStrainCount      int      `json:"SourceStrainCount"`
	SourceStrainNames      *string  `json:"SourceStrainNames"` // TODO: Check data type.
	Strains                []string `json:"Strains"`
	DryingLocationId       int      `json:"DryingLocationId"`
	DryingLocationName     string   `json:"DryingLocationName"`
	DryingLocationTypeName *string  `json:"DryingLocationTypeName"`
	PatientLicenseNumber   *string  `json:"PatientLicenseNumber"`
	CurrentWeight          float64  `json:"CurrentWeight"`
	TotalWasteWeight       float64  `json:"TotalWasteWeight"`
	PlantCount             int      `json:"PlantCount"`
	TotalWetWeight         float64  `json:"TotalWetWeight"`
	TotalRestoredWeight    float64  `json:"TotalRestoredWeight"`
	PackageCount           int      `json:"PackageCount"`
	TotalPackagedWeight    float64  `json:"TotalPackagedWeight"`
	UnitOfWeightName       string   `json:"UnitOfWeightName"`
	LabTestingState        *string  `json:"LabTestingState"`
	LabTestingStateDate    *string  `json:"LabTestingStateDate"`
	IsOnHold               bool     `json:"IsOnHold"`
	HarvestStartDate       string   `json:"HarvestStateDate"`
	FinishedDate           *string  `json:"FinishedDate"`
	ArchivedDate           *string  `json:"ArchivedDate"`
	LastModified           string   `json:"LastModified"`
}

// GetHarvestsById gets the Harvest with the specified ID.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_{id}.GET
func (m *Metrc) GetHarvestsById(id int, licenseNumber *string) (Harvest, error) {
	endpoint := fmt.Sprintf("harvests/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var hr Harvest
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return hr, fmt.Errorf("could not get harvests by id response: %s", err)
	}

	err = json.Unmarshal(responseBody, &hr)
	if err != nil {
		return hr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return hr, nil
}

// helper method to get harvests with the "status" endpoint
func (m *Metrc) getHarvestsByStatus(status string, licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Harvest, error) {
	endpoint := fmt.Sprintf("harvests/v1/%s", status)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("?lastModifiedStart=%s", *lastModifiedStart)
	}
	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("?lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var hr []Harvest
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return hr, fmt.Errorf("could not get harvests by status response: %s", err)
	}

	err = json.Unmarshal(responseBody, &hr)
	if err != nil {
		return hr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return hr, nil
}

// GetHarvestsActive gets all active harvests from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_active.GET
func (m *Metrc) GetHarvestsActive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Harvest, error) {
	harvests, err := m.getHarvestsByStatus("active", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []Harvest{}, fmt.Errorf("could not get active harvests: %s", err)
	}

	return harvests, nil
}

// GetHarvestsOnHold gets all on hold harvests from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_onhold.GET
func (m *Metrc) GetHarvestsOnHold(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Harvest, error) {
	harvests, err := m.getHarvestsByStatus("onhold", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []Harvest{}, fmt.Errorf("could not get on hold harvests: %s", err)
	}

	return harvests, nil
}

// GetHarvestsInactive gets all inactive harvests from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_inactive.GET
func (m *Metrc) GetHarvestsInactive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Harvest, error) {
	harvests, err := m.getHarvestsByStatus("inactive", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []Harvest{}, fmt.Errorf("could not get inactive harvests: %s", err)
	}

	return harvests, nil
}

// HarvestWasteType represents a type of waste for a harvest.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_waste_types.GET
type HarvestWasteType struct {
	Name string `json:"Name"`
}

// GetHarvestWasteTypes gets the type of waste for a harvest
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.get_harvests_v1_waste_types.GET
func (m *Metrc) GetHarvestsWasteTypes() ([]HarvestWasteType, error) {
	endpoint := "harvests/v1/waste/types"

	var hwt []HarvestWasteType
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return hwt, fmt.Errorf("could not get harvest waste types response: %s", err)
	}

	err = json.Unmarshal(responseBody, &hwt)
	if err != nil {
		return hwt, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return hwt, nil
}

// HarvestIngredient represents an ingredient in a harvest.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_create_packages_testing.POST
type HarvestIngredient struct {
	HarvestId    *int    `json:"HarvestId"`
	HarvestName  *string `json:"HarvestName"`
	Weight       float64 `json:"Weight"`
	UnitOfWeight string  `json:"UnitOfWeight"`
}

// HarvestPackagePost is used in the request body to create a new harvest package.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_create_packages.POST
type HarvestPackagePost struct {
	Tag                        string              `json:"Tag"`
	Location                   *string             `json:"Location"`
	Item                       string              `json:"Item"`
	UnitOfWeight               string              `json:"UnitOfWeight"`
	PatientLicenseNumber       string              `json:"PatientLicenseNumber"`
	Note                       string              `json:"Note"`
	IsProductionBatch          bool                `json:"IsProductionBatch"`
	ProductionBatchNumber      *int                `json:"ProductionBatchNumber"`
	IsTradeSample              bool                `json:"IsTradeSample"`
	IsDonation                 bool                `json:"IsDonation"`
	ProductRequiresRemediation bool                `json:"ProductRequiresRemediation"`
	RemediateProduct           bool                `json:"RemediateProduct"`
	RemediationMethodId        *int                `json:"RemediationMethodId"`
	RemediationDate            *string             `json:"RemediationDate"`
	RemediationSteps           *[]string           `json:"RemediationSteps"`
	ActualDate                 string              `json:"ActualDate"`
	Ingredients                []HarvestIngredient `json:"Ingredients"`
}

// PostHarvestsCreatePackages creates a new package in a harvest in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_create_packages.POST
func (m *Metrc) PostHarvestsCreatePackages(packages []HarvestPackagePost, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/create/packages"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal harvest packages: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting packages: %s", err)
	}

	return resp, nil
}

// PostHarvestsCreatePackagesTesting creates a new package for testing in a harvest in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_create_packages_testing.POST
func (m *Metrc) PostHarvestsCreatePackagesTesting(packages []HarvestPackagePost, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/create/packages/testing"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal harvest testing packages: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting testing packages: %s", err)
	}

	return resp, nil
}

// HarvestMove represents a harvest to move.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.put_harvests_v1_move.PUT
type HarvestMove struct {
	Id             *int    `json:"Id"`
	HarvestName    *string `json:"HarvestName"`
	DryingLocation string  `json:"DryingLocation"`
	ActualDate     string  `json:"ActualDate"`
}

// PutHarvestsMove is used to move harvests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.put_harvests_v1_move.PUT
func (m *Metrc) PutHarvestsMove(harvests []HarvestMove, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/move"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(harvests)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal move harvests: %s", err)
	}

	resp, err := m.client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed putting harvests: %s", err)
	}

	return resp, nil
}

// HarvestRemoveWaste is used to remove wasted harvests.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_removewaste.POST
type HarvestRemoveWaste struct {
	Id           int     `json:"Id"`
	WasteType    string  `json:"WasteType"`
	UnitOfWeight string  `json:"UnitOfWeight"`
	WasteWeight  float64 `json:"WasteWeight"`
	ActualDate   string  `json:"ActualDate"`
}

// PostHarvestsRemoveWaste removes wastes from harvests.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_removewaste.POST
func (m *Metrc) PostHarvestsRemoveWaste(wastes []HarvestRemoveWaste, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/removewaste"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(wastes)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal remove wastes: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting remove wastes: %s", err)
	}

	return resp, nil
}

// HarvestRename is used in the request body to rename a harvest.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.put_harvests_v1_rename.PUT
type HarvestRename struct {
	Id      *int    `json:"Id"`
	OldName *string `json:"OldName"`
	NewName string  `json:"NewName"`
}

// PutHarvestsRename renames a list of harvests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.put_harvests_v1_rename.PUT
func (m *Metrc) PutHarvestsRename(harvests []HarvestRename, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/rename"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(harvests)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal rename harvests: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting rename harvests: %s", err)
	}

	return resp, nil
}

// HarvestFinish represents a harvest when finishing.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_finish.POST
type HarvestFinish struct {
	Id         int    `json:"Id"`
	ActualDate string `json:"ActualDate"`
}

// PostHarvestsFinish finishes a list of harvests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_finish.POST
func (m *Metrc) PostHarvestsFinish(harvests []HarvestFinish, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/finish"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(harvests)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal finish harvests: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting finish harvests: %s", err)
	}

	return resp, nil
}

// HarvestUnfinish represents a harvest when unfinishing.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_unfinish.POST
type HarvestUnfinish struct {
	Id int `json:"Id"`
}

// PostHarvestsUnfinish unfinishes a list of harvests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Harvests.post_harvests_v1_unfinish.POST
func (m *Metrc) PostHarvestsUnfinish(harvests []HarvestUnfinish, licenseNumber string) ([]byte, error) {
	endpoint := "harvests/v1/unfinish"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(harvests)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal unfinish harvests: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting unfinish harvests: %s", err)
	}

	return resp, nil
}
