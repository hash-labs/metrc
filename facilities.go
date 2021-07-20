package metrc

import (
	"encoding/json"
	"fmt"
)

// Facilities represents an element in the array returned by `GET /facilities/v1`.
// See: https://api-ca.metrc.com/Documentation/#Facilities.get_facilities_v1
type Facilities struct {
	HireDate              string                 `json:"HireDate"`
	IsOwner               bool                   `json:"IsOwner"`
	IsManager             bool                   `json:"IsManager"`
	Occupations           []string               `json:"Occupations"`
	Name                  string                 `json:"Name"`
	Alias                 string                 `json:"Alias"`
	DisplayName           string                 `json:"DisplayName"`
	CredentialedDate      string                 `json:"CredentialedDate"`
	SupportActivationDate string                 `json:"SupportActivationDate"`
	SupportExpirationDate string                 `json:"SupportExpirationDate"`
	SupportLastPaidDate   string                 `json:"SupportLastPaidDate"`
	FacilityType          FacilitiesFacilityType `json:"FacilityType"`
	License               FacilitiesLicense      `json:"License"`
}

// FacilitiesLicense represents the License in the Facilities struct.
// See: https://api-ca.metrc.com/Documentation/#Facilities.get_facilities_v1
type FacilitiesLicense struct {
	Number      string `json:"Number"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	LicenseType string `json:"LicenseType"`
}

// FacilitiesFacilityType represents the FacilityType in the Facilities struct.
// See: https://api-ca.metrc.com/Documentation/#Facilities.get_facilities_v1
type FacilitiesFacilityType struct {
	IsMedical                                       bool `json:"IsMedical"`
	IsRetail                                        bool `json:"IsRetail"`
	IsHemp                                          bool `json:"IsHemp"`
	RestrictHarvestPlantRestoreTimeHours            int  `json:"RestrictHarvestPlantRestoreTimeHours"`
	TotalMemberPatientsAllowed                      int  `json:"TotalMemberPatientsAllowed"`
	CanGrowPlants                                   bool `json:"CanGrowPlants"`
	CanCreateOpeningBalancePlantBatches             bool `json:"CanCreateOpeningBalancePlantBatches"`
	CanTagPlantBatches                              bool `json:"CanTagPlantBatches"`
	CanAssignLocationsToPlantBatches                bool `json:"CanAssignLocationsToPlantBatches"`
	PlantsRequirePatientAffiliation                 bool `json:"PlantsRequirePatientAffiliation"`
	PlantBatchesCanContainMotherPlants              bool `json:"PlantBatchesCanContainMotherPlants"`
	CanUpdatePlantStrains                           bool `json:"CanUpdatePlantStrains"`
	CanTrackVegetativePlants                        bool `json:"CanTrackVegetativePlants"`
	CanCreateImmaturePlantPackagesFromPlants        bool `json:"CanCreateImmaturePlantPackagesFromPlants"`
	CanPackageVegetativePlants                      bool `json:"CanPackageVegetativePlants"`
	CanReportHarvestSchedules                       bool `json:"CanReportHarvestSchedules"`
	CanSubmitHarvestsForTesting                     bool `json:"CanSubmitHarvestsForTesting"`
	CanRequireHarvestSampleLabTestBatches           bool `json:"CanRequireHarvestSampleLabTestBatches"`
	CanCreateOpeningBalancePackages                 bool `json:"CanCreateOpeningBalancePackages"`
	CanCreateDerivedPackages                        bool `json:"CanCreateDerivedPackages"`
	CanAssignLocationsToPackages                    bool `json:"CanAssignLocationsToPackages"`
	CanUpdateLocationsOnPackages                    bool `json:"CanUpdateLocationsOnPackages"`
	PackagesRequirePatientAffiliation               bool `json:"PackagesRequirePatientAffiliation"`
	CanCreateTradeSamplePackages                    bool `json:"CanCreateTradeSamplePackages"`
	CanDonatePackages                               bool `json:"CanDonatePackages"`
	CanSubmitPackagesForTesting                     bool `json:"CanSubmitPackagesForTesting"`
	CanCreateProcessValidationPackages              bool `json:"CanCreateProcessValidationPackages"`
	CanRequirePackageSampleLabTestBatches           bool `json:"CanRequirePackageSampleLabTestBatches"`
	CanRequestProductRemediation                    bool `json:"CanRequestProductRemediation"`
	CanRemediatePackagesWithFailedLabResults        bool `json:"CanRemediatePackagesWithFailedLabResults"`
	CanInfuseProducts                               bool `json:"CanInfuseProducts"`
	CanTestPackages                                 bool `json:"CanTestPackages"`
	CanTransferFromExternalFacilities               bool `json:"CanTransferFromExternalFacilities"`
	CanSellToConsumers                              bool `json:"CanSellToConsumers"`
	CanSellToPatients                               bool `json:"CanSellToPatients"`
	CanSellToExternalPatients                       bool `json:"CanSellToExternalPatients"`
	CanSellToCaregivers                             bool `json:"CanSellToCaregivers"`
	AdvancedSales                                   bool `json:"AdvancedSales"`
	SalesRequirePatientNumber                       bool `json:"SalesRequirePatientNumber"`
	SalesRequireExternalPatientNumber               bool `json:"SalesRequireExternalPatientNumber"`
	SalesRequireExternalPatientIdentificationMethod bool `json:"SalesRequireExternalPatientIdentificationMethod"`
	SalesRequireCaregiverNumber                     bool `json:"SalesRequireCaregiverNumber"`
	SalesRequireCaregiverPatientNumber              bool `json:"SalesRequireCaregiverPatientNumber"`
	CanDeliverSalesToConsumers                      bool `json:"CanDeliverSalesToConsumers"`
	SalesDeliveryRequireConsumerId                  bool `json:"SalesDeliveryRequireConsumerId"`
	CanDeliverSalesToPatients                       bool `json:"CanDeliverSalesToPatients"`
	SalesDeliveryRequirePatientNumber               bool `json:"SalesDeliveryRequirePatientNumber"`
	SalesDeliveryRequireRecipientName               bool `json:"SalesDeliveryRequireRecipientName"`
	CanHaveMemberPatients                           bool `json:"CanHaveMemberPatients"`
	CanReportPatientCheckIns                        bool `json:"CanReportPatientCheckIns"`
	CanSpecifyPatientSalesLimitExemption            bool `json:"CanSpecifyPatientSalesLimitExemption"`
	CanReportPatientsAdverseResponses               bool `json:"CanReportPatientsAdverseResponses"`
	CanReportStrainProperties                       bool `json:"CanReportStrainProperties"`
}

// Facilities gets all facilities from Metrc.
// See: https://api-ca.metrc.com/Documentation/#Facilities.get_facilities_v1
func (m *Metrc) Facilities() ([]Facilities, error) {
	endpoint := "facilities/v1"

	var fr []Facilities
	body, err := m.Client.Get(endpoint)
	if err != nil {
		return fr, fmt.Errorf("could not get facilities response: %s", err)
	}

	err = json.Unmarshal(body, &fr)
	if err != nil {
		return fr, fmt.Errorf("could not unmarshal response: %s", err)
	}
	return fr, nil
}
