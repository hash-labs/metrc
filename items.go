package metrc

import (
	"encoding/json"
	"fmt"
)

// TODO: Check null fields in documentation across all structs.

// ItemGet represents an Item in Metrc in the GET requests. This corresponds to Products/SKUs in our codebase.
// See: https://api-ca.metrc.com/Documentation/#Items.get_items_v1_{id}
type ItemGet struct {
	Id                                  int     `json:"Id"`
	Name                                string  `json:"Name"`
	ProductCategoryName                 string  `json:"ProductCategoryName"`
	ProductCategoryType                 string  `json:"ProductCategoryType"`
	QuantityType                        string  `json:"QuantityType"`
	DefaultLabTestingState              string  `json:"DefaultLabTestingState"`
	UnitOfMeasureName                   string  `json:"UnitOfMeasureName"`
	ApprovalStatus                      string  `json:"ApprovalStatus"`
	ApprovalStatusDateTime              string  `json:"ApprovalStatusDateTime"`
	StrainId                            int     `json:"StrainId"`
	StrainName                          string  `json:"StrainName"`
	AdministrationMethod                string  `json:"AdministrationMethod"`
	UnitCbdPercent                      float64 `json:"UnitCbdPercent"`
	UnitCbdContent                      float64 `json:"UnitCbdContent"`
	UnitCbdContentUnitOfMeasureName     float64 `json:"UnitCbdContentUnitOfMeasureName"`
	UnitCbdContentDose                  float64 `json:"UnitCbdContentDose"`
	UnitCbdContentDoseUnitOfMeasureName string  `json:"UnitCbdContentDoseUnitOfMeasureName"`
	UnitThcPercent                      float64 `json:"UnitThcPercent"`
	UnitThcContent                      float64 `json:"UnitThcContent"`
	UnitThcContentUnitOfMeasureName     string  `json:"UnitThcContentUnitOfMeasureName"`
	UnitThcContentDose                  float64 `json:"UnitThcContentDose"`
	UnitThcContentDoseUnitOfMeasureName string  `json:"UnitThcContentDoseUnitOfMeasureName"`
	UnitVolume                          float64 `json:"UnitVolume"`
	UnitVolumeUnitOfMeasureName         string  `json:"UnitVolumeUnitOfMeasureName"`
	UnitWeight                          float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasureName         string  `json:"UnitWeightUnitOfMeasureName"`
	ServingSize                         string  `json:"ServingSize"`
	SupplyDurationDays                  int     `json:"SupplyDurationDays"`
	NumberOfDoses                       int     `json:"NumberOfDoses"`
	UnitQuantity                        int     `json:"UnitQuantity"`
	UnitQuantityUnitOfMeasureName       string  `json:"UnitQuantityUnitOfMeasureName"`
	Ingredients                         string  `json:"Ingredients"`
	Description                         string  `json:"Description"`
	IsUsed                              bool    `json:"IsUsed"`
}

// ItemPost represents an Item in Metrc in the POST requests. This corresponds to Products/SKUs in our codebase.
// See: https://api-ca.metrc.com/Documentation/#Items.post_items_v1_create
type ItemPost struct {
	Id                              int     `json:"Id,omitempty"`
	ItemCategory                    string  `json:"ItemCategory"`
	Name                            string  `json:"Name"`
	UnitOfMeasure                   string  `json:"UnitOfMeasure"`
	Strain                          string  `json:"Strain"`
	ItemBrand                       string  `json:"ItemBrand"`
	AdministrationMethod            string  `json:"AdministrationMethod"`
	UnitCbdPercent                  float64 `json:"UnitCbdPercent"`
	UnitCbdContent                  float64 `json:"UnitCbdContent"`
	UnitCbdContentUnitOfMeasure     float64 `json:"UnitCbdContentUnitOfMeasure"`
	UnitCbdContentDose              float64 `json:"UnitCbdContentDose"`
	UnitCbdContentDoseUnitOfMeasure float64 `json:"UnitCbdContentDoseUnitOfMeasure"`
	UnitThcPercent                  float64 `json:"UnitThcPercent"`
	UnitThcContent                  float64 `json:"UnitThcContent"`
	UnitThcContentUnitOfMeasure     string  `json:"UnitThcContentUnitOfMeasure"`
	UnitThcContentDose              float64 `json:"UnitThcContentDose"`
	UnitThcContentDoseUnitOfMeasure string  `json:"UnitThcContentDoseUnitOfMeasure"`
	UnitVolume                      float64 `json:"UnitVolume"`
	UnitVolumeUnitOfMeasure         string  `json:"UnitVolumeUnitOfMeasure"`
	UnitWeight                      float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasure         string  `json:"UnitWeightUnitOfMeasure"`
	ServingSize                     float64 `json:"ServingSize"`
	SupplyDurationDays              int     `json:"SupplyDurationDays"`
	NumberOfDoses                   int     `json:"NumberOfDoses"`
	Ingredients                     string  `json:"Ingredients"`
	Description                     string  `json:"Description"`
}

// ItemCategory represents an Item Category in Metrc.
// See: https://api-ca.metrc.com/Documentation/#Items.get_items_v1_categories
type ItemCategory struct {
	Name                         string `json:"Name"`
	ProductCategoryType          string `json:"ProductCategoryType"`
	QuantityType                 string `json:"QuantityType"`
	RequiresStrain               bool   `json:"RequiresStrain"`
	RequiresItemBrand            bool   `json:"RequiresItemBrand"`
	RequiresAdministrationMethod bool   `json:"RequiresAdministrationMethod"`
	RequiresUnitCbdPercent       bool   `json:"RequiresUnitCbdPercent"`
	RequiresUnitCbdContent       bool   `json:"RequiresUnitCbdContent"`
	RequiresUnitCbdContentDose   bool   `json:"RequiresUnitCbdContentDose"`
	RequiresUnitThcPercent       bool   `json:"RequiresUnitThcPercent"`
	RequiresUnitThcContent       bool   `json:"RequiresUnitThcContent"`
	RequiresUnitThcContentDose   bool   `json:"RequiresUnitThcContentDose"`
	RequiresUnitVolume           bool   `json:"RequiresUnitVolume"`
	RequiresUnitWeight           bool   `json:"RequiresUnitWeight"`
	RequiresServingSize          bool   `json:"RequiresServingSize"`
	RequiresSupplyDurationDays   bool   `json:"RequiresSupplyDurationDays"`
	RequiresNumberOfDoses        bool   `json:"RequiresNumberOfDoses"`
	RequiresIngredients          bool   `json:"RequiresIngredients"`
	RequiresDescription          bool   `json:"RequiresDescription"`
	RequiresProductPhotos        int    `json:"RequiresProductPhotos"`
	RequiresLabelPhotos          int    `json:"RequiresLabelPhotos"`
	RequiresPackagingPhotos      int    `json:"RequiresPackagingPhotos"`
	CanContainSeeds              bool   `json:"CanContainSeeds"`
	CanBeRemediated              bool   `json:"CanBeRemediated"`
}

// GetItemsById gets items with an ID.
// See: https://api-ca.metrc.com/Documentation/#Items.get_items_v1_{id}
func (m *Metrc) GetItemsById(id int, licenseNumber *string) (ItemGet, error) {
	var endpoint string
	if licenseNumber != nil {
		endpoint = fmt.Sprintf("items/v1/%d?licenseNumber=%s", id, *licenseNumber)
	} else {
		endpoint = fmt.Sprintf("items/v1/%d", id)
	}

	var ir ItemGet
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return ir, fmt.Errorf("could not get items by id response: %s", err)
	}

	err = json.Unmarshal(responseBody, &ir)
	if err != nil {
		return ir, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ir, nil
}

// GetItemsActive gets all active items.
// See: https://api-ca.metrc.com/Documentation/#Items.get_items_v1_active
func (m *Metrc) GetItemsActive(licenseNumber *string) ([]ItemGet, error) {
	endpoint := "items/v1/active"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var ir []ItemGet
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return ir, fmt.Errorf("could not get active items from metrc: %s", err)
	}

	err = json.Unmarshal(responseBody, &ir)
	if err != nil {
		return ir, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ir, nil
}

// GetItemsCategories retrieves the categories of items, either all or for the provided license number.
// See: https://api-ca.metrc.com/Documentation/#Items.get_items_v1_categories
func (m *Metrc) GetItemsCategories(licenseNumber *string) ([]ItemCategory, error) {
	endpoint := "items/v1/categories"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var ic []ItemCategory
	responseBody, err := m.Client.Get(endpoint)
	if err != nil {
		return ic, fmt.Errorf("could not get item categories from metrc: %s", err)
	}

	err = json.Unmarshal(responseBody, &ic)
	if err != nil {
		return ic, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ic, nil
}

// CreateItems creates new Items.
// See: https://api-ca.metrc.com/Documentation/#Items.post_items_v1_create
func (m *Metrc) CreateItems(items []ItemPost, licenseNumber *string) ([]byte, error) {
	endpoint := "items/v1/create"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	body, err := json.Marshal(items)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal items: %s", err)
	}

	// TODO: Add more robust response handling.
	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting items: %s", err)
	}
	return resp, nil
}

// UpdateItems updates existing Items. Note that for this endpoint, `ItemPost.Id` is required in each `ItemPost` in the input slice.
// See: https://api-ca.metrc.com/Documentation/#Items.post_items_v1_update
func (m *Metrc) UpdateItems(items []ItemPost, licenseNumber *string) ([]byte, error) {
	endpoint := "items/v1/update"
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	body, err := json.Marshal(items)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal items: %s", err)
	}

	// TODO: Add more robust response handling.
	resp, err := m.Client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting items: %s", err)
	}
	return resp, nil
}

// DeleteItemById deletes an existing item by ID.
// See: https://api-ca.metrc.com/Documentation/#Items.delete_items_v1_{id}
func (m *Metrc) DeleteItemById(id int, licenseNumber *string) ([]byte, error) {
	endpoint := fmt.Sprintf("items/v1/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	resp, err := m.Client.Delete(endpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("failed deleting item %d: %s", id, err)
	}
	return resp, nil
}
