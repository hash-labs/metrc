package metrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitsOfMeasure_Integration(t *testing.T) {
	ums, err := m.GetUnitsOfMeasure()
	assert.NoError(t, err)

	foundUnit := false
	for _, um := range ums {
		if um.QuantityType == "CountBased" && um.Name == "Each" && um.Abbreviation == "ea" {
			foundUnit = true
		}
	}
	assert.True(t, foundUnit)
}
