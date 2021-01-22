package metrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var m *Metrc = MakeIntegrationMetrc()

func TestFacilities_Integration(t *testing.T) {
	resp, err := m.Facilities()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(resp), 1)
}

// TODO: Add golden file and unit test for parseFacilitiesResponse.
// Note that this isn't top priority, since it's mostly checking correctness of our data structures and the JSON package.
// We do need to figure out a good strategy for this over time...
