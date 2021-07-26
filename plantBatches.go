package metrc

import (
	"encoding/json"
	"fmt"
)

// PlantBatch represents a Plant Batch in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.get_plantbatches_v1_{id}.GET
type PlantBatch struct {
	Id                   *int    `json:"Id"`
	Name                 string  `json:"Name"`
	Type                 string  `json:"Type"`
	LocationId           *int    `json:"LocationId"`
	LocationName         *string `json:"LocationName"`
	LocationTypeName     *string `json:"LocationTypeName"`
	StrainId             *int    `json:"StrainId"`
	StrainName           *string `json:"StrainName"`
	PatientLicenseNumber *int    `json:"PatientLicenseNumber"`
	UntrackedCount       int     `json:"UntrackedCount"`
	TrackedCount         int     `json:"TrackedCount"`
	PackagedCount        int     `json:"PackagedCount"`
	HarvestedCount       int     `json:"HarvestedCount"`
	DestroyedCount       int     `json:"DestroyedCount"`
	SourcePackageId      *int    `json:"SourcePackageId"`
	SourcePlantId        *int    `json:"SourcePlantId"`
	SourcePlantLabel     *string `json:"SourcePlantLabel"`
	PlantedDate          string  `json:"PlantedDate"`
	LastModified         string  `json:"LastModified"`
}

// GetPlantBatchesById gets the Plant Batch with the specified ID.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.get_plantbatches_v1_{id}.GET
func (m *Metrc) GetPlantBatchesById(id int, licenseNumber *string) (PlantBatch, error) {
	endpoint := fmt.Sprintf("plantbatches/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var pbr PlantBatch
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return pbr, fmt.Errorf("could not get plant batches by id response: %s", err)
	}

	err = json.Unmarshal(responseBody, &pbr)
	if err != nil {
		return pbr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pbr, nil
}

func (m *Metrc) getPlantBatchesByStatus(status string, licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PlantBatch, error) {
	endpoint := fmt.Sprintf("plantbatches/v1/%s", status)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("?lastModifiedStart=%s", *lastModifiedStart)
	}
	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("?lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var pbr []PlantBatch
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return pbr, fmt.Errorf("could not get plant batches by status response: %s", err)
	}

	err = json.Unmarshal(responseBody, &pbr)
	if err != nil {
		return pbr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pbr, nil
}

// GetPlantBatchesActive gets the PlantBatches that are active.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.get_plantbatches_v1_active.GET
func (m *Metrc) GetPlantBatchesActive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PlantBatch, error) {
	batches, err := m.getPlantBatchesByStatus("active", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []PlantBatch{}, fmt.Errorf("could not get active plant batches: %s", err)
	}

	return batches, nil
}

// GetPlantBatchesInactive gets the Plant Batches that are inactive.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.get_plantbatches_v1_inactive.GET
func (m *Metrc) GetPlantBatchesInactive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PlantBatch, error) {
	batches, err := m.getPlantBatchesByStatus("inactive", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []PlantBatch{}, fmt.Errorf("could not get inactive plant batches: %s", err)
	}

	return batches, nil
}

// GetPlantBatchesTypes retrieves the types of Plant Batches.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.get_plantbatches_v1_types.GET
func (m *Metrc) GetPlantBatchesTypes() ([]string, error) {
	endpoint := "plantbatches/v1/types"

	var types []string
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return types, fmt.Errorf("could not get plant batches by status response: %s", err)
	}

	err = json.Unmarshal(responseBody, &types)
	if err != nil {
		return types, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return types, nil
}

// PlantBatchPlanting represents a new Planting of a PlantBatch.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_createplantings.POST
type PlantBatchPlanting struct {
	Name                 string  `json:"Name"`
	Type                 string  `json:"Type"`
	Count                int     `json:"Count"`
	Strain               string  `json:"Strain"`
	Location             *string `json:"Location"`
	PatientLicenseNumber string  `json:"PatientLicenseNumber"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantBatchesCreatePlantings creates a planting of a new plant batch.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_createplantings.POST
func (m *Metrc) PostPlantBatchesCreatePlantings(batches []PlantBatchPlanting, licenseNumber string) ([]byte, error) {
	endpoint := "plantbatches/v1/createplantings"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(batches)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal batches: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post plant batches create plantings: %s", err)
	}

	return resp, nil
}

// PlantBatchPackage represents a package of a plant batch.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_createpackages.POST
type PlantBatchPackage struct {
	Id                   *int    `json:"Id"`
	PlantBatch           string  `json:"PlantBatch"`
	Count                int     `json:"Count"`
	Location             *string `json:"Location"`
	Item                 string  `json:"Item"`
	Tag                  string  `json:"Tag"`
	PatientLicenseNumber string  `json:"PatientLicenseNumber"`
	Note                 string  `json:"Note"`
	IsTradeSample        bool    `json:"IsTradeSample"`
	IsDonation           bool    `json:"IsDonation"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantBatchesCreatePackages creates a package of a Plant Batch in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_createpackages.POST
func (m *Metrc) PostPlantBatchesCreatePackages(packages []PlantBatchPackage, licenseNumber string, isFromMotherPlant *bool) ([]byte, error) {
	endpoint := "plantbatches/v1/createpackages"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if isFromMotherPlant != nil {
		endpoint += fmt.Sprintf("?isFromMotherPlant=%t", *isFromMotherPlant)
	}

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post packages: %s", err)
	}

	return resp, nil
}

// PlantBatchSplit represents a split plant batch.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_split.POST
type PlantBatchSplit struct {
	PlantBatch           string  `json:"PlantBatch"`
	GroupName            string  `json:"GroupName"`
	Count                int     `json:"Count"`
	Location             *string `json:"Location"`
	Strain               string  `json:"Strain"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantBatchesSplit posts split plant batches.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_split.POST
func (m *Metrc) PostPlantBatchesSplit(batches []PlantBatchSplit, licenseNumber string) ([]byte, error) {
	endpoint := "plantbatches/v1/split"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(batches)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PostPlantBatchesCreatePackagesFromMotherPlant creates packages of plant batches from a mother plant.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_create_packages_frommotherplant.POST
func (m *Metrc) PostPlantBatchesCreatePackagesFromMotherPlant(packages []PlantBatchPackage, licenseNumber string) ([]byte, error) {
	endpoint := "plantbatches/v1/create/packages/frommotherplant"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal json: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PlantBatchGrowthPhase creates
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_changegrowthphase.POST
type PlantBatchGrowthPhase struct {
	Name                 string `json:"Name"`
	Count                int    `json:"Count"`
	StartingTag          string `json:"StartingTag"`
	GrowthPhase          string `json:"GrowthPhase"`
	NewLocation          string `json:"NewLocation"`
	GrowthDate           string `json:"GrowthDate"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
}

// PostPlantBatchesChangeGrowthPhase changes the growth phase of batches.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_changegrowthphase.POST
func (m *Metrc) PostPlantBatchesChangeGrowthPhase(batches []PlantBatchGrowthPhase, licenseNumber string) ([]byte, error) {
	endpoint := "plantbatches/v1/changegrowthphase"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(batches)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PlantBatchMove represents data to move a plant batch.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.put_plantbatches_v1_moveplantbatches.PUT
type PlantBatchMove struct {
	Name     string `json:"Name"`
	Location string `json:"Location"`
	MoveDate string `json:"MoveDate"`
}

// PutPlantBatchesMove moves a Plant Batch in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.put_plantbatches_v1_moveplantbatches.PUT
func (m *Metrc) PutPlantBatchesMove(batches []PlantBatchMove, licenseNumber string) ([]byte, error) {
	endpoint := "/plantbatches/v1/moveplantbatches"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(batches)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.Client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PlantBatchAdditive represents an additive for the plant branch.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_additives.POST
type PlantBatchAdditive struct {
	AdditiveType             string                               `json:"AdditiveType"`
	ProductTradeName         string                               `json:"ProductTradeName"`
	EpaRegistrationNumber    *int                                 `json:"EpaRegistrationNumber"`
	ProductSupplier          string                               `json:"ProductSupplier"`
	ApplicationDevice        string                               `json:"ApplicationDevice"`
	TotalAmountApplied       float64                              `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string                               `json:"TotalAmountUnitOfMeasure"`
	ActiveIngredients        []PlantBatchAdditiveActiveIngredient `json:"ActiveIngredients"`
	PlantBatchName           string                               `json:"PlantBatchName"`
	ActualDate               string                               `json:"ActualDate"`
}

// PlantBatchAdditiveActiveIngredient represents active ingredient in the plant batch additive.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_additives.POST
type PlantBatchAdditiveActiveIngredient struct {
	Name       string  `json:"Name"`
	Percentage float64 `json:"Percentage"`
}

// PostPlantBatchesAdditives posts plant batch additives.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_additives.POST
func (m *Metrc) PostPlantBatchesAdditives(additives []PlantBatchAdditive, licenseNumber string) ([]byte, error) {
	endpoint := "plantbatches/v1/additives"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(additives)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PlantBatchDestroy represents a plant batch to destroy.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_destroy.POST
type PlantBatchDestroy struct {
	PlantBatch string `json:"PlantBatch"`
	Count      int    `json:"Count"`
	ReasonNote string `json:"ReasonNote"`
	ActualDate string `json:"ActualDate"`
}

// PostPlantBatchesDestroy destroys plant batches.
// See: https://testing-api-ca.metrc.com/Documentation#PlantBatches.post_plantbatches_v1_destroy.POST
func (m *Metrc) PostPlantBatchesDestroy(batches []PlantBatchDestroy, licenseNumber string) ([]byte, error) {
	endpoint := "plantbatches/v1/destroy"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(batches)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}
