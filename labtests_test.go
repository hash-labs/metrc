package metrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabTestsStates_Integration(t *testing.T) {
	states, err := m.GetLabTestsStates()
	assert.NoError(t, err)

	foundState := false
	wantState := "NotSubmitted"
	for _, gotState := range states {
		if gotState == wantState {
			foundState = true
		}
	}
	assert.True(t, foundState)
}

func TestLabTestsTypes_Integration(t *testing.T) {
	_, err := m.GetLabTestsTypes()
	assert.NoError(t, err)

	// TODO: Define found type.
}

func TestLabTestsResults_Integration(t *testing.T) {
	id := 1 // TODO: Define proper number.
	_, err := m.GetLabTestsResults(id, licenseNumber)
	assert.NoError(t, err)

	// TODO: Define found result.
}

func TestLabTestsRecord_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestLabTestsDocument_Integration(t *testing.T) {
	// TODO: Implement.
}

func TestLabTestsRelease_Integration(t *testing.T) {
	// TODO: Implement.
}
