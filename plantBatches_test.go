package metrc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlantBatchesById_Integration(t *testing.T) {
	id := 0 // TODO: Modify.
	batch, err := m.GetPlantBatchesById(id, &licenseNumber)
	assert.NoError(t, err)
	assert.Equal(t, id, batch.Id)
}

func TestPlantBatchesActive_Integration(t *testing.T) {
	_, err := m.GetPlantBatchesActive(licenseNumber, nil, nil)
	assert.NoError(t, err)
}

func TestPlantBatchesInactive_Integration(t *testing.T) {
	_, err := m.GetPlantBatchesInactive(licenseNumber, nil, nil)
	assert.NoError(t, err)
}

func TestPlantBatchesTypes_Integration(t *testing.T) {
	batches, err := m.GetPlantBatchesTypes()
	assert.NoError(t, err)
	fmt.Printf("%#v\n", batches)
}

func TestPlantBatchesCreatePlantings_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesCreatePackages_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesSplit_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesCreatePackagesFromMotherPlant_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesChangeGrowthPhase_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesMove_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesAdditives_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantBatchesDestroy_Integration(t *testing.T) {
	// TODO: Implement.
}
