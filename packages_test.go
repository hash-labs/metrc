package metrc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPackagesGetId_Integration(t *testing.T) {
	// TODO: Modify the id and change this test to work, once a package is added.

	// id := 144720
	// p, err := m.GetPackagesById(id, &licenseNumber)
	// assert.NoError(t, err)

	// assert.Equal(t, 144720, p.Id)
	// assert.Equal(t, "Default Location", p.LocationTypeName)
}

func TestPackagesGetActive_Integration(t *testing.T) {
	actives, err := m.GetPackagesActive(licenseNumber, nil, nil)
	assert.NoError(t, err)

	// TODO: May need to add some active packages for testing.
	assert.Equal(t, 0, len(actives))
}

func TestPackagesGetOnhold_Integration(t *testing.T) {
	onhold, err := m.GetPackagesOnHold(licenseNumber, nil, nil)
	assert.NoError(t, err)

	// TODO: May need to add some onhold packages for testing.
	assert.Equal(t, 0, len(onhold))
}

func TestPackagesGetInactive_Integration(t *testing.T) {
	inactives, err := m.GetPackagesInactive(licenseNumber, nil, nil)
	assert.NoError(t, err)

	// TODO: May need to add some inactive packages for testing.
	assert.Equal(t, 0, len(inactives))
}

func TestPackagesTypes_Integration(t *testing.T) {
	types, err := m.GetPackagesTypes()
	assert.NoError(t, err)
	assert.Equal(t, "Product", types[0])
}

func TestPackagesAdjustReasons_Integration(t *testing.T) {
	pars, err := m.GetPackagesAdjustReasons(licenseNumber)
	assert.NoError(t, err)

	par := pars[0]
	assert.Equal(t, "Damage (BCC)", par.Name)
	assert.Equal(t, true, par.RequiresNote)
}

func TestPackagesCreate_Integration(t *testing.T) {
	// Generate a random label for the new package.
	rand.Seed(time.Now().Unix())
	label := fmt.Sprintf("%d", rand.Int())

	// TODO: Create packages testing.
	wantPackage := PackagePost{
		Tag:                        label,
		Location:                   nil,
		Item:                       "[FLW-DSD-ROLL] Do Si Dos Preroll",
		Quantity:                   2.0,
		UnitOfMeasure:              "Grams",
		PatientLicenseNumber:       "X0001",
		Note:                       "",
		IsProductionBatch:          false,
		ProductionBatchNumber:      nil,
		IsDonation:                 false,
		ProductRequiresRemediation: false,
		UseSameItem:                false,
		ActualDate:                 "2015-12-15",
		Ingredients: []Ingredient{
			{
				Package:       "testIngredient",
				Quantity:      16.0,
				UnitOfMeasure: "Ounces",
			},
		},
	}
	packages := []PackagePost{wantPackage}
	_, err := m.PostPackagesCreateTesting(packages, licenseNumber)
	assert.NoError(t, err)

	// Get the package with the label.
	gotPackage, err := m.GetPackagesByLabel(label, &licenseNumber)
	assert.Equal(t, wantPackage, gotPackage)
	fmt.Printf("%+v\n", gotPackage)

	// TODO: Change the packages item.

	// TODO: Change the packages note.
}
