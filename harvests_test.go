package metrc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHarvestsById_Integration(t *testing.T) {
	id := 0
	h, err := m.GetHarvestsById(id, &licenseNumber)
	assert.NoError(t, err)
	assert.Equal(t, id, h.Id)
}

func TestHarvestsActive_Integration(t *testing.T) {
	h, err := m.GetHarvestsActive(licenseNumber, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", h)
}

func TestHarvestsOnHold_Integration(t *testing.T) {
	h, err := m.GetHarvestsOnHold(licenseNumber, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", h)
}

func TestHarvestsInactive_Integration(t *testing.T) {
	h, err := m.GetHarvestsInactive(licenseNumber, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", h)
}

func TestHarvestsWasteTypes_Integration(t *testing.T) {
	types, err := m.GetHarvestsWasteTypes()
	assert.NoError(t, err)

	foundType := false
	wantName := "Plant Material"
	for _, typ := range types {
		if typ.Name == wantName {
			foundType = true
		}
	}
	assert.True(t, foundType)
}

func TestHarvestsCreatePackages_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestHarvestsCreatePackagesTesting_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestHarvestsRemoveWaste_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestHarvestsRename_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestHarvestsFinish_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestHarvestsUnfinish_Integration(t *testing.T) {
	// TODO: Implement.
}
