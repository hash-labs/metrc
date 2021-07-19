package metrc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStrainsGetByID_Integration(t *testing.T) {
	_, err := m.GetStrainsById(204001, &licenseNumber)
	assert.NoError(t, err)
}

func TestStrainsGetActive_Integration(t *testing.T) {
	activeStrains, err := m.GetStrainsActive(&licenseNumber)
	assert.NoError(t, err)

	wantName := "1465247066340200968"
	foundStrain := false
	for _, strain := range activeStrains {
		if strain.Name == wantName {
			foundStrain = true
		}
	}
	assert.True(t, foundStrain)
}

func TestStrainsCreateUpdateDelete_Integration(t *testing.T) {
	// Generate a random name for a new Strain.
	rand.Seed(time.Now().Unix())
	name := fmt.Sprintf("%d", rand.Int())

	// Create a Strain (with said Name).
	strains := []Strain{
		{
			Name:             name,
			TestingStatus:    "None",
			ThcLevel:         0.2,
			CbdLevel:         0.2,
			IndicaPercentage: 25.0,
			SativaPercentage: 75.0,
		},
	}
	_, err := m.PostStrainsCreate(strains, &licenseNumber)
	assert.NoError(t, err)

	// Get all active Strains, and then find the Id of the new Strain.
	activeStrains, err := m.GetStrainsActive(&licenseNumber)
	assert.NoError(t, err)
	var strainId int
	for _, as := range activeStrains {
		if as.Name == name {
			strainId = as.Id
			break
		}
	}

	// Rename the Strain using update.
	strains = []Strain{
		{
			Id:               strainId,
			Name:             fmt.Sprintf("%d", rand.Int()),
			TestingStatus:    "None",
			ThcLevel:         0.2,
			CbdLevel:         0.2,
			IndicaPercentage: 25.0,
			SativaPercentage: 75.0,
		},
	}
	_, err = m.UpdateStrains(strains, &licenseNumber)
	assert.NoError(t, err)
}
