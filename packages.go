package metrc

import (
	"encoding/json"
	"fmt"
)

// PackageGet represents a Package in a GET request from Metrc.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_{id}
type PackageGet struct {
	Id                                  int      `json:"Id"`
	Label                               string   `json:"Label"`
	PackageType                         string   `json:"PackageType"`
	SourceHarvestNames                  *string  `json:"SourceHarvestNames"`
	LocationId                          *int     `json:"LocationId"`
	LocationName                        *string  `json:"LocationName"`
	LocationTypeName                    *string  `json:"LocationTypeName"`
	Quantity                            float64  `json:"Quantity"`
	UnitOfMeasureName                   string   `json:"UnitOfMeasureName"`
	UnitOfMeasureAbbreviation           string   `json:"UnitOfMeasureAbbreviation"`
	PatientLicenseNumber                *int     `json:"PatientLicenseNumber"`
	ProductId                           int      `json:"ProductId"`
	ProductName                         string   `json:"ProductName"`
	ProductCategoryName                 string   `json:"ProductCategoryName"`
	ItemFromFacilityLicenseNumber       *string  `json:"ItemFromFacilityLicenseNumber"`
	ItemFromFacilityName                *string  `json:"ItemFromFacilityName"`
	ItemStrainName                      *string  `json:"ItemStrainName"`
	ItemUnitCbdPercent                  *float64 `json:"ItemUnitCbdPercent"`
	ItemUnitCbdContent                  *float64 `json:"ItemUnitCbdContent"`
	ItemUnitCbdContentUnitOfMeasureName *string  `json:"ItemUnitCbdContentUnitOfMeasureName"`
	ItemUnitThcPercent                  *float64 `json:"ItemUnitThcPercent"`
	ItemUnitThcContent                  *float64 `json:"ItemUnitThcContent"`
	ItemUnitThcContentUnitOfMeasureName *string  `json:"ItemUnitThcContentUnitOfMeasureName"`
	ItemUnitVolume                      *float64 `json:"ItemUnitVolume"`
	ItemUnitVolumeUnitOfMeasureName     *string  `json:"ItemUnitVolumeUnitOfMeasureName"`
	ItemUnitWeight                      *float64 `json:"ItemUnitWeight"`
	ItemUnitWeightUnitOfMeasureName     *string  `json:"ItemUnitWeightUnitOfMeasureName"`
	ItemServingSize                     *float64 `json:"ItemServingSize"`
	ItemSupplyDurationDays              *int     `json:"ItemSupplyDurationDays"`
	ItemUnitQuantity                    *float64 `json:"ItemUnitQuantity"`
	ItemUnitQuantityUnitOfMeasureName   *string  `json:"ItemUnitQuantityUnitOfMeasureName"`
	Note                                *string  `json:"Note"`
	PackagedDate                        string   `json:"PackagedDate"`
	InitialLabTestingState              string   `json:"InitialLabTestingState"`
	LabTestingState                     string   `json:"LabTestingState"`
	LabTestingStateDate                 string   `json:"LabTestingStateDate"`
	IsProductionBatch                   bool     `json:"IsProductionBatch"`
	ProductionBatchNumber               *int     `json:"ProductionBatchNumber"`
	IsTradeSample                       bool     `json:"IsTradeSample"`
	IsDonation                          bool     `json:"IsDonation"`
	IsDonationPersistent                bool     `json:"IsDonationPersistent"`
	SourcePackageIsDonation             bool     `json:"SourcePackageIsDonation"`
	IsTestingSample                     bool     `json:"IsTestingSample"`
	IsProcessValidationTestingSample    bool     `json:"IsProcessValidationTestingSample"`
	ProductRequiresRemediation          bool     `json:"ProductRequiresRemediation"`
	ContainsRemediatedProduct           bool     `json:"ContainsRemediatedProduct"`
	RemediationDate                     *string  `json:"RemediationDate"`
	ReceivedDateTime                    *string  `json:"ReceivedDateTime"`
	ReceivedFromManifestNumber          *string  `json:"ReceivedFromManifestNumber"`
	ReceivedFromFacilityLicenseNumber   *string  `json:"ReceivedFromFacilityLicenseNumber"`
	ReceivedFromFacilityName            *string  `json:"ReceivedFromFacilityName"`
	IsOnHold                            bool     `json:"IsOnHold"`
	ArchivedDate                        *string  `json:"ArchivedDate"`
	FinishedDate                        *string  `json:"FinishedDate"`
	LastModified                        string   `json:"LastModified"`
	Item                                ItemGet  `json:"Item"`
}

// PackageAdjustReasons represents a reason to adjust a package.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_adjust_reasons
type PackageAdjustReasons struct {
	Name         string `json:"Name"`
	RequiresNote bool   `json:"RequiresNote"`
}

// PackagePost represents a Package in a POST request to Metrc.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_create
type PackagePost struct {
	Tag                        string       `json:"Tag"`
	Location                   *string      `json:"Location"`
	Item                       string       `json:"Item"`
	Quantity                   float64      `json:"Quantity"`
	UnitOfMeasure              string       `json:"UnitOfMeasure"`
	PatientLicenseNumber       string       `json:"PatientLicenseNumber"`
	Note                       string       `json:"Note"`
	IsProductionBatch          bool         `json:"IsProductionBatch"`
	ProductionBatchNumber      *int         `json:"ProductionBatchNumber"`
	IsDonation                 bool         `json:"IsDonation"`
	ProductRequiresRemediation bool         `json:"ProductRequiresRemediation"`
	UseSameItem                bool         `json:"UseSameItem"`
	ActualDate                 string       `json:"ActualDate"`
	Ingredients                []Ingredient `json:"Ingredients"`
}

// Ingredient represents an Ingredient within a PackagePost.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_create
type Ingredient struct {
	Package       string  `json:"Package"`
	Quantity      float64 `json:"Quantity"`
	UnitOfMeasure string  `json:"UnitOfMeasure"`
}

// GetPackagesById gets a Package by ID.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_{id}
func (m *Metrc) GetPackagesById(id int, licenseNumber *string) (PackageGet, error) {
	endpoint := fmt.Sprintf("packages/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var pr PackageGet
	body, err := m.Client.Get(endpoint)
	if err != nil {
		return pr, fmt.Errorf("could not get packages by id response: %s", err)
	}

	err = json.Unmarshal(body, &pr)
	if err != nil {
		return pr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pr, nil
}

// GetPackagesByLabel gets a Package by its `Label`.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_{label}
func (m *Metrc) GetPackagesByLabel(label string, licenseNumber *string) (PackageGet, error) {
	endpoint := fmt.Sprintf("packages/v1/%s", label)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var pr PackageGet
	body, err := m.Client.Get(endpoint)
	if err != nil {
		return pr, fmt.Errorf("could not get packages by label response: %s", err)
	}

	err = json.Unmarshal(body, &pr)
	if err != nil {
		return pr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return pr, nil
}

// GetPackagesActive gets all active Packages for a license number. Optional timestamps can be passed for filtering.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_active
func (m *Metrc) GetPackagesActive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error) {
	return m.getPackages("active", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

// GetPackagesOnHold gets all on hold Packages for a license number. Optional timestamps can be passed for filtering.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_onhold
func (m *Metrc) GetPackagesOnHold(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error) {
	return m.getPackages("onhold", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

// GetPackagesInactive gets all inactive Packages for a license number. Optional timestamps can be passed for filtering.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_inactive
func (m *Metrc) GetPackagesInactive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error) {
	return m.getPackages("inactive", licenseNumber, lastModifiedStart, lastModifiedEnd)
}

func (m *Metrc) getPackages(endpointName string, licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error) {
	endpoint := fmt.Sprintf("packages/v1/%s", endpointName)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("&lastModifiedStart=%s", *lastModifiedStart)
	}

	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("&lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var pr []PackageGet
	body, err := m.Client.Get(endpoint)
	if err != nil {
		return pr, fmt.Errorf("could not get packages from metrc: %s", err)
	}

	err = json.Unmarshal(body, &pr)
	if err != nil {
		return pr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return pr, nil
}

// GetPackagesTypes gets all types for Packages.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_types
func (m *Metrc) GetPackagesTypes() ([]string, error) {
	endpoint := "packages/v1/types"

	var types []string
	resp, err := m.Client.Get(endpoint)
	if err != nil {
		return types, fmt.Errorf("failed to get types from metrc: %s", err)
	}

	err = json.Unmarshal(resp, &types)
	if err != nil {
		return types, fmt.Errorf("could not unmarshap response: %s", err)
	}

	return types, nil
}

// GetPackagesAdjustReasons gets the reasons to adjust Packages.
// See: https://api-ca.metrc.com/Documentation/#Packages.get_packages_v1_adjust_reasons
func (m *Metrc) GetPackagesAdjustReasons(licenseNumber string) ([]PackageAdjustReasons, error) {
	endpoint := fmt.Sprintf("packages/v1/adjust/reasons?licenseNumber=%s", licenseNumber)

	var par []PackageAdjustReasons
	resp, err := m.Client.Get(endpoint)
	if err != nil {
		return par, fmt.Errorf("failed to get packages adjust reasons: %s", err)
	}

	err = json.Unmarshal(resp, &par)
	if err != nil {
		return par, fmt.Errorf("could not unmarshap response: %s", err)
	}

	return par, nil
}

// PostPackagesCreate creates new Packages.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_create
func (m *Metrc) PostPackagesCreate(packages []PackagePost, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/create?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal packages: %s", err)
	}

	// TODO: Add more robust response handling for posts.
	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting packages: %s", err)
	}
	return resp, nil
}

// PostPackagesCreateTesting creates Packages for testing.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_create_testing
func (m *Metrc) PostPackagesCreateTesting(packages []PackagePost, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/create/testing?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal packages: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting packages: %s", err)
	}
	return resp, nil
}

// PostPackagesCreatePlanting creates Packages from Planting.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_create_plantings
func (m *Metrc) PostPackagesCreatePlantings(packages []PackagePost, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/create/plantings?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal packages: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting packages: %s", err)
	}
	return resp, nil
}

// PackageItem represents a Package (via its Label) and Item.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_change_item
type PackageItem struct {
	Label string `json:"Label"`
	Item  string `json:"Item"`
}

// PostPackagesChangeItem changes the Item on a Package.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_change_item
func (m *Metrc) PostPackagesChangeItem(packageItems []PackageItem, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/change/item?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageItems)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-items: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting package-items: %s", err)
	}

	return resp, nil
}

// PackageNote repreqents a Package (via its Label) and Note.
// See: https://api-ca.metrc.com/Documentation/#Packages.put_packages_v1_change_note
type PackageNote struct {
	Label string `json:"PackageLabel"`
	Note  string `json:"Note"`
}

// PutPackagesChangeNote changes the Note on a Package.
// See: https://api-ca.metrc.com/Documentation/#Packages.put_packages_v1_change_note
func (m *Metrc) PutPackagesChangeNote(packageNotes []PackageNote, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/change/note?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageNotes)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-notes: %s", err)
	}

	resp, err := m.Client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed putting package-notes: %s", err)
	}

	return resp, nil
}

// PackageLocation represents a Package (via its Label), location, and move date.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_change_locations
type PackageLocation struct {
	Label    string `json:"Label"`
	Location string `json:"Location"`
	MoveDate string `json:"MoveDate"`
}

// PostPackagesChangeLocations changes the Location on a Package.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_change_locations
func (m *Metrc) PostPackagesChangeLocations(packageLocations []PackageLocation, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/change/locations?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageLocations)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-locations: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting package-locations: %s", err)
	}

	return resp, nil
}

// PackageAdjust represents a Package (via its Label) and other adjustments (quantity, unit, reason/date).
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_adjust
type PackageAdjust struct {
	Label            string  `json:"Label"`
	Quantity         float64 `json:"Quantity"`
	UnitOfMeasure    string  `json:"UnitOfMeasure"`
	AdjustmentReason string  `json:"AdjustmentReason"`
	AdjustmentDate   string  `json:"AdjustmentDate"`
	ReasonNote       *string `json:"ReasonNote"`
}

// PostPackagesAdjust changes adjustment metadata on a Package.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_adjust
func (m *Metrc) PostPackagesAdjust(packageAdjusts []PackageAdjust, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/adjust?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageAdjusts)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-adjusts: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting package-adjusts: %s", err)
	}

	return resp, nil
}

// PackageFinish represents when a Package was finished.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_finish
type PackageFinish struct {
	Label      string `json:"Label"`
	ActualDate string `json:"ActualDate"`
}

// PostPackagesFinish finishes a Package.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_finish
func (m *Metrc) PostPackagesFinish(packageFinishes []PackageFinish, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/finish?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageFinishes)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-finishes: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting package-finishes: %s", err)
	}

	return resp, nil
}

// PackageUnfinish represents when a Package was unfinished.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_unfinish
type PackageUnfinish struct {
	Label string `json:"Label"`
}

// PostPackagesUnfinish unfinishes Packages.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_unfinish
func (m *Metrc) PostPackagesUnfinish(packageUnfinishes []PackageUnfinish, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/unfinish?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageUnfinishes)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-unfinishes: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting package-unfinishes: %s", err)
	}

	return resp, nil
}

// PackageRemediate contains the information to remediate a package.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_remediate
type PackageRemediate struct {
	Label      string `json:"PackageLabel"`
	MethodName string `json:"RemediationmethodName"`
	Date       string `json:"RemediationDate"`
	Steps      string `json:"RemediationSteps"`
}

// PostPackagesRemediate remediates Packages.
// See: https://api-ca.metrc.com/Documentation/#Packages.post_packages_v1_remediate
func (m *Metrc) PostPackagesRemediate(packageRemediates []PackageRemediate, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("packages/v1/remediate?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packageRemediates)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal package-remediates: %s", err)
	}

	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting package-remediates: %s", err)
	}

	return resp, nil
}
