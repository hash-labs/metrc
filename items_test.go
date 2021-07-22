package metrc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var licenseNumber string = "C12-1000006-LIC"

func TestItemsGetById_Integration(t *testing.T) {
	wantId := 266502
	gotItem, err := m.GetItemsById(wantId, &licenseNumber)
	assert.NoError(t, err)
	assert.Equal(t, wantId, gotItem.Id)
}

func TestItemsGetActive_Integration(t *testing.T) {
	activeItems, err := m.GetItemsActive(&licenseNumber)
	assert.NoError(t, err)

	// Check that an item with the desired ID exists.
	wantId := 266502
	foundItem := false
	for _, item := range activeItems {
		if item.Id == wantId {
			foundItem = true
		}
	}
	assert.True(t, foundItem)
}

func TestItemsGetCategories_Integration(t *testing.T) {
	categories, err := m.GetItemsCategories(&licenseNumber)
	assert.NoError(t, err)

	// Check that an item with the desired Name exists.
	wantName := "Capsule (weight)"
	foundName := false
	for _, cat := range categories {
		if cat.Name == wantName {
			foundName = true
		}
	}
	assert.True(t, foundName)
}

func TestItemsGetBrands_Integration(t *testing.T) {
	brands, err := m.GetItemsBrands(licenseNumber)
	assert.NoError(t, err)
	fmt.Println(brands)
}

// Tests Create, Update, and Delete.
// You have to GetActiveItems, find the created test item, and Update then Delete by ID.
func TestItemsCreateUpdateDelete_Integration(t *testing.T) {

	// Generate a random name for a new Item.
	rand.Seed(time.Now().Unix())
	name := fmt.Sprintf("%d", rand.Int())

	// Create an Item (with said Name).
	item := ItemPost{
		ItemCategory:                    "Capsule (weight)",
		Name:                            name,
		UnitOfMeasure:                   "Ounces",
		Strain:                          "Spring Hill Kush",
		UnitThcContent:                  10.0,
		UnitThcContentUnitOfMeasure:     "Milligrams",
		UnitThcContentDose:              5.0,
		UnitThcContentDoseUnitOfMeasure: "Milligrams",
		UnitWeight:                      100.0,
		UnitWeightUnitOfMeasure:         "Milligrams",
		NumberOfDoses:                   2,
	}
	items := []ItemPost{item}
	_, err := m.PostItemsCreate(items, &licenseNumber)
	assert.NoError(t, err)

	// Get all active Items, and then find the Id of the new Item.
	activeItems, err := m.GetItemsActive(&licenseNumber)
	assert.NoError(t, err)
	var itemId int
	for _, ai := range activeItems {
		if ai.Name == name {
			itemId = ai.Id
			break
		}
	}

	// Rename the Item using update.
	item.Id = itemId
	item.Name = fmt.Sprintf("%d", rand.Int())
	items = []ItemPost{item}
	_, err = m.PostItemsUpdate(items, &licenseNumber)
	assert.NoError(t, err)

	// Delete the Item using the ID.
	_, err = m.DeleteItemById(itemId, &licenseNumber)
	assert.NoError(t, err)
}
