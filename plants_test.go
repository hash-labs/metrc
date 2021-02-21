package metrc

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var lastModifiedFmt string = "2006-01-02T15:04:05Z"

func TestPlantsGetId_Integration(t *testing.T) {
	id := 0 // TODO: Find valid id.
	p, err := m.GetPlantsById(id, &licenseNumber)
	assert.NoError(t, err)

	assert.Equal(t, p.Id, id)
}

func TestPlantsGetLabel_Integration(t *testing.T) {

}

func TestPlantsGetVegetative_Integration(t *testing.T) {
	end := time.Now().Add(time.Hour * -24 * 90)
	endStr := end.Format(lastModifiedFmt)
	fmt.Println(endStr)

	start := end.Add(time.Hour * -24)
	startStr := start.Format(lastModifiedFmt)
	fmt.Println(startStr)

	p, err := m.GetPlantsVegetative(licenseNumber, &startStr, &endStr)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", p)
}

func TestPlantsGrowthPhases_Integration(t *testing.T) {
	phases, err := m.GetPlantsGrowthPhases(licenseNumber)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", phases)
}

func TestPlantsAdditiveTypes_Integration(t *testing.T) {
	types, err := m.GetPlantsAdditivesTypes(licenseNumber)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"Fertilizer", "Pesticide", "Other"}, types)
}

func TestPlantsWasteMethods_Integration(t *testing.T) {
	_, err := m.GetPlantsWasteMethods(licenseNumber)
	assert.NoError(t, err)
}

func TestPlantsWasteReasons_Integration(t *testing.T) {
	reasons, err := m.GetPlantsWasteReasons(licenseNumber)
	assert.NoError(t, err)
	fmt.Printf("%#v\n", reasons)

	// Validate that a desired reason exists in the sandbox.
	foundReason := false
	wantReason := PlantWasteReason{Name: "Pesticides", RequiresNote: true}
	for _, gotReason := range reasons {
		if gotReason == wantReason {
			foundReason = true
		}
	}
	assert.True(t, foundReason)
}

func TestPlantsMovePlants_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantsChangeGrowthPhases_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantsPostAdditives_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantsCreatePlantings_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantsManicure_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestPlantsHarvest_Integration(t *testing.T) {
	// TODO: Implement.
}
