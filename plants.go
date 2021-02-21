package metrc

import (
	"encoding/json"
	"fmt"
)

// Plant represents a Plant in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_{id}.GET
type Plant struct {
	Id                         int     `json:"Id,omitempty"`
	Label                      string  `json:"Label"`
	State                      string  `json:"State"`
	GrowthPhase                string  `json:"GrowthPhase"`
	PlantBatchId               int     `json:"PlantBatchId"`
	PlantBatchName             string  `json:"PlantBatchName"`
	PlantBatchTypeName         string  `json:"PlantBatchTypeName"`
	StrainId                   int     `json:"StrainId"`
	StrainName                 string  `json:"StrainName"`
	LocationId                 int     `json:"LocationId"`
	LocationName               string  `json:"LocationName"`
	PatientLicenseNumber       int     `json:"PatientLicenseNumber"`
	HarvestId                  *int    `json:"HarvestId"`
	HarvestedUnionOfWeightName *string `json:"HarvestedUnionOfWeightName"`
	HarvestCount               int     `json:"HarvestCount"`
	IsOnHold                   bool    `json:"IsOnHold"`
	PlantedDate                string  `json:"PlantedDate"`
	VegetativeDate             string  `json:"VegetativeDate"`
	FloweringDate              *string `json:"FloweringDate"`
	HarvestedDate              *string `json:"HarvestedDate"`
	DestroyedDate              *string `json:"DestroyedDate"`
	DestroyedNote              *string `json:"DestroyedNote"`
	DestroyedByUserName        *string `json:"DestroyedByUserName"`
	LastModified               string  `json:"LastModified"`
}

// GetPlantsById gets Plants with an Id.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_{id}.GET
func (m *Metrc) GetPlantsById(id int, licenseNumber *string) (Plant, error) {
	endpoint := fmt.Sprintf("plants/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var pr Plant
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return pr, fmt.Errorf("could not get plants by id response: %s", err)
	}

	err = json.Unmarshal(responseBody, &pr)
	if err != nil {
		return pr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pr, nil
}

// GetPlantsByLabel gets Strains with a label.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_{label}.GET
func (m *Metrc) GetPlantsByLabel(label string, licenseNumber *string) (Plant, error) {
	endpoint := fmt.Sprintf("plants/v1/%s", label)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var pr Plant
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return pr, fmt.Errorf("could not get plants by label response: %s", err)
	}

	err = json.Unmarshal(responseBody, &pr)
	if err != nil {
		return pr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pr, nil
}

// helper function to get Plants of various status.
func (m *Metrc) getPlantsByStatus(status string, licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Plant, error) {
	endpoint := fmt.Sprintf("plants/v1/%s", status)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("?lastModifiedStart=%s", *lastModifiedStart)
	}
	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("?lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var pr []Plant
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return pr, fmt.Errorf("could not get plants: %s", err)
	}

	err = json.Unmarshal(responseBody, &pr)
	if err != nil {
		return pr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pr, nil
}

// GetPlantsVegetative returns all vegetating plants for a facility.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_vegetative.GET
func (m *Metrc) GetPlantsVegetative(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Plant, error) {
	return m.getPlantsByStatus("vegetative", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

// GetPlantsFlowering returns all flowering plants for a facility.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_flowering.GET
func (m *Metrc) GetPlantsFlowering(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Plant, error) {
	return m.getPlantsByStatus("flowering", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

// GetPlantsOnHold returns all plants on hold for a facility.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_onhold.GET
func (m *Metrc) GetPlantsOnHold(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Plant, error) {
	return m.getPlantsByStatus("onhold", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

// GetPlantsInactive returns all inactive plants for a facility.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_inactive.GET
func (m *Metrc) GetPlantsInactive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Plant, error) {
	return m.getPlantsByStatus("inactive", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

// PlantAdditiveGet represents a Plant Additive in Metrc's GET request.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_additives.GET
type PlantAdditiveGet struct {
	AdditiveTypeName      *string `json:"AdditiveTypeName"`
	ProductTradeName      string  `json:"ProductTradeName"`
	EpaRegistrationNumber *int    `json:"EpaRegistrationNumber"`
	ProductSupplier       string  `json:"ProductSupplier"`
	ApplicationDevice     string  `json:"ApplicationDevice"`
	AmountUnitOfMeasure   string  `json:"AmountUnitOfMeasure"`
	TotalAmountApplied    float64 `json:"TotalAmountApplied"`
	PlantBatchId          *int    `json:"PlantBatchId"`
	PlantBatchName        *string `json:"PlantBatchName"`
	PlantCount            int     `json:"PlantCount"`
}

// GetPlantsAdditives gets Plant Additives.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_additives.GET
func (m *Metrc) GetPlantsAdditives(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PlantAdditiveGet, error) {
	endpoint := fmt.Sprintf("plants/v1/additives")
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("?lastModifiedStart=%s", *lastModifiedStart)
	}
	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("?lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var par []PlantAdditiveGet
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return par, fmt.Errorf("could not get plant additives: %s", err)
	}

	err = json.Unmarshal(responseBody, &par)
	if err != nil {
		return par, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return par, nil
}

// GetPlantsGrowthPhases returns the Plants' growth phases for a facility.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_growthphases.GET
func (m *Metrc) GetPlantsGrowthPhases(licenseNumber string) ([]string, error) {
	endpoint := fmt.Sprintf("plants/v1/additives")
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	var phases []string
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return phases, fmt.Errorf("could not get growth phases: %s", err)
	}

	err = json.Unmarshal(responseBody, &phases)
	if err != nil {
		return phases, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return phases, nil
}

// GetPlantsAdditivesTypes returns the types of Plant Additives.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_additives_types.GET
func (m *Metrc) GetPlantsAdditivesTypes(licenseNumber string) ([]string, error) {
	endpoint := fmt.Sprintf("plants/v1/additives/types")
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	var types []string
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return types, fmt.Errorf("could not get additives types: %s", err)
	}

	err = json.Unmarshal(responseBody, &types)
	if err != nil {
		return types, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return types, nil
}

// PlantWasteMethod represents a method to waste a Plant in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_waste_methods.GET
type PlantWasteMethod struct {
	Name string `json:"Name"`
}

// GetPlantsWasteMethods returns the methods to waste a Plant.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_waste_methods.GET
func (m *Metrc) GetPlantsWasteMethods(licenseNumber string) ([]PlantWasteMethod, error) {
	endpoint := fmt.Sprintf("plants/v1/waste/methods")
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	var methods []PlantWasteMethod
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return methods, fmt.Errorf("could not get waste methods: %s", err)
	}

	err = json.Unmarshal(responseBody, &methods)
	if err != nil {
		return methods, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return methods, nil
}

// PlantWasteReason represents a reason to waste a Plant in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_waste_reasons.GET
type PlantWasteReason struct {
	Name         string `json:"Name"`
	RequiresNote bool   `json:"RequiresNote"`
}

// GetPlantsWasteReasons gets the reasons to waste a Plant from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.get_plants_v1_waste_reasons.GET
func (m *Metrc) GetPlantsWasteReasons(licenseNumber string) ([]PlantWasteReason, error) {
	endpoint := fmt.Sprintf("plants/v1/waste/reasons")
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	var reasons []PlantWasteReason
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return reasons, fmt.Errorf("could not get waste reasons: %s", err)
	}

	err = json.Unmarshal(responseBody, &reasons)
	if err != nil {
		return reasons, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return reasons, nil
}

// PlantMovePost is used in the request body to Move a Plant.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_moveplants.POST
type PlantMovePost struct {
	Id         *int    `json:"Id"`
	Label      *string `json:"Label"`
	Location   string  `json:"Location"`
	ActualDate string  `json:"ActualDate"`
}

// PostPlantsMovePlants changes Plants locations within Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_moveplants.POST
func (m *Metrc) PostPlantsMovePlants(movePlants []PlantMovePost, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/moveplants"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(movePlants)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-locations: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting move-plants: %s", err)
	}

	return resp, nil
}

// PlantChangeGrowthPhase is used in the request body to change a Plant's growth phase.
type PlantChangeGrowthPhase struct {
	Id          *int    `json:"Id"`
	Label       *string `json:"Label"`
	NewTag      string  `json:"NewTag"`
	GrowthPhase string  `json:"GrowthPhase"`
	NewLocation string  `json:"NewLocation"`
	GrowthDate  string  `json:"GrowthDate"`
}

// PostPlantsChangeGrowthPhases changes Plants growth phases within Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_changegrowthphases.POST
func (m *Metrc) PostPlantsChangeGrowthPhases(changes []PlantChangeGrowthPhase, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/changegrowthphases"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(changes)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal growth phase changes: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting growth phase changes: %s", err)
	}

	return resp, nil
}

// PlantDestroy is used in the request body to destroy a Plant.
type PlantDestroy struct {
	Id                     *int    `json:"Id"`
	Label                  *string `json:"Label"`
	WasteMethodName        string  `json:"WasteMethodName"`
	WasteMaterialMixed     string  `json:"WasteMaterialMixed"`
	WasteWeight            float64 `json:"WasteWeight"`
	WasteUnitOfMeasureName string  `json:"WasteUnitOfMeasureName"`
	WasteReasonName        string  `json:"WasteReasonName"`
	ReasonNote             string  `json:"ReasonNote"`
	ActualDate             string  `json:"ActualDate"`
}

// PostPlantsDestroy is used in the request body to destroy a Plant.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_destroyplants.POST
func (m *Metrc) PostPlantsDestroy(plants []PlantDestroy, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/destroyplants"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(plants)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal plants to destroy: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting growth phase changes: %s", err)
	}

	return resp, nil
}

// PlantAdditivePost is used in the request body to post a Plant Additive.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_additives.POST
type PlantAdditivePost struct {
	AdditiveType          string                    `json:"AdditiveType"`
	ProductTradeName      string                    `json:"ProductTradeName"`
	EpaRegistrationNumber *int                      `json:"EpaRegistrationNumber"`
	ProductSupplier       string                    `json:"ProductSupplier"`
	ApplicationDevice     string                    `json:"ApplicationDevice"`
	TotalAmountApplied    string                    `json:"TotalAmountApplied"`
	ActiveIngredients     []PlantAdditiveIngredient `json:"ActiveIngredients"`
	PlantLabels           []string                  `json:"PlantLabels"`
	ActualDate            string                    `json:"ActualDate"`
}

// PlantAdditiveIngredient is used when posting a Plant Additive.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_additives.POST
type PlantAdditiveIngredient struct {
	Name       string  `json:"Name"`
	Percentage float64 `json:"Percentage"`
}

// PostPlantsAdditives is used in the request body to change the additives for the Plants.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_additives.POST
func (m *Metrc) PostPlantsAdditives(additives []PlantAdditivePost, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/additives"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(additives)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal additives: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting additives: %s", err)
	}

	return resp, nil
}

// PlantCreatePlanting is used in the request body to create a planting.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_create_plantings.POST
type PlantCreatePlanting struct {
	PlantLabel           *string `json:"PlantLabel"`
	PlantBatchName       string  `json:"PlantBatchName"`
	PlantBatchType       string  `json:"PlantBatchType"`
	PlantCount           int     `json:"PlantCount"`
	LocationName         *string `json:"LocationName"`
	StrainName           string  `json:"StrainName"`
	PatientLicenseNumber string  `json:"PatientLicenseNumber"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantsCreatePlantings creates new plants.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_create_plantings.POST
func (m *Metrc) PostPlantsCreatePlantings(plants []PlantCreatePlanting, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/create/plantings"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(plants)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal create plantings: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting create plantings: %s", err)
	}

	return resp, nil
}

// PlantCreatePlantBatchPackage represents a plant batch package to create.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_create_plantbatch_packages.POST
type PlantCreatePlantBatchPackage struct {
	PlantLabel           string  `json:"PlantLabel"`
	PackageTag           string  `json:"PackageTag"`
	PlantBatchType       string  `json:"PlantBatchType"`
	Item                 string  `json:"Item"`
	Location             *string `json:"Location"`
	Note                 *string `json:"Note"`
	IsTradeSample        bool    `json:"IsTradeSample"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	IsDonation           bool    `json:"IsDonation"`
	Count                int     `json:"Count"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantsCreateBatchPackages creates Plant batch packages.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_create_plantbatch_packages.POST
func (m *Metrc) PostPlantsCreateBatchPackages(batches []PlantCreatePlantBatchPackage, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/create/plantbatch/packages"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(batches)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal create batch packages: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting create batch packages: %s", err)
	}

	return resp, nil
}

// PlantManicure represents a plant to manicure.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_manicureplants.POST
type PlantManicure struct {
	Plant                string  `json:"Plant"`
	Weight               float64 `json:"Weight"`
	UnitOfWeight         string  `json:"UnitOfWeight"`
	DryingLocation       string  `json:"DryingLocation"`
	HarvestName          *string `json:"HarvestName"`
	PatientLicenseNumber string  `json:"PatientLicenseNumber"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantsManicure posts the plant to manicure.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_manicureplants.POST
func (m *Metrc) PostPlantsManicure(plants []PlantManicure, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/manicureplants"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(plants)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal manicure plants: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting manicure plants: %s", err)
	}

	return resp, nil
}

// PlantHarvest represents the plants to harvest.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_harvestplants.POST
type PlantHarvest struct {
	Plant                string  `json:"Plant"`
	Weight               float64 `json:"Weight"`
	UnitOfWeight         string  `json:"UnitOfWeight"`
	DryingLocation       string  `json:"DryingLocation"`
	HarvestName          string  `json:"HarvestName"`
	PatientLicenseNumber string  `json:"PatientLicenseNumber"`
	ActualDate           string  `json:"ActualDate"`
}

// PostPlantsHarvest posts the plants to harvest.
// See: https://testing-api-ca.metrc.com/Documentation#Plants.post_plants_v1_harvestplants.POST
func (m *Metrc) PostPlantsHarvest(plants []PlantHarvest, licenseNumber string) ([]byte, error) {
	endpoint := "plants/v1/harvestplants"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(plants)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal harvest plants: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting harvest plants: %s", err)
	}

	return resp, nil
}
