package metrc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocationsById_Integration(t *testing.T) {
	resp, err := m.GetLocationsById(144720, &licenseNumber)
	assert.NoError(t, err)
	t.Log(resp)
}

func TestLocationsGetActive_Integration(t *testing.T) {
	activeLocations, err := m.GetLocationsActive(&licenseNumber)
	assert.NoError(t, err)

	// Check that we can find the desired ID.
	wantId := 144720
	foundId := false
	for _, loc := range activeLocations {
		if loc.Id == wantId {
			foundId = true
			break
		}
	}
	assert.True(t, foundId)
}

func TestLocationsGetTypes_Integration(t *testing.T) {
	locs, err := m.GetLocationsActive(&licenseNumber)
	assert.NoError(t, err)

	// Check that we can find the desired ID.
	wantId := 144720
	foundId := false
	for _, loc := range locs {
		if loc.Id == wantId {
			foundId = true
			break
		}
	}
	assert.True(t, foundId)
}

func TestLocationsCreateUpdateDelete_Integration(t *testing.T) {
	// Generate a random name for a new Location.
	rand.Seed(time.Now().Unix())
	name := fmt.Sprintf("%d", rand.Int())

	// Create a Location (with said Name).
	locs := []LocationPost{
		{
			Name:     name,
			TypeName: "Default Location Type",
		},
	}
	_, err := m.CreateLocations(locs, &licenseNumber)
	assert.NoError(t, err)

	// Get all active Locations, and then find the Id of the new Location.
	activeLocs, err := m.GetLocationsActive(&licenseNumber)
	assert.NoError(t, err)
	var locId int
	for _, al := range activeLocs {
		if al.Name == name {
			locId = al.Id
			break
		}
	}

	// Rename the Location using update.
	locs = []LocationPost{
		{
			Id:       locId,
			Name:     fmt.Sprintf("%d", rand.Int()),
			TypeName: "Default Location Type",
		},
	}
	_, err = m.UpdateLocations(locs, &licenseNumber)
	assert.NoError(t, err)

	// Delete the Location using the ID.
	_, err = m.DeleteLocationById(locId, &licenseNumber)
	assert.NoError(t, err)
}
