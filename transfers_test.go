package metrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfersIncoming_Integration(t *testing.T) {
	_, err := m.GetTransfersIncoming(licenseNumber, nil, nil)
	assert.NoError(t, err)
}

func TestTransfersOutgoing_Integration(t *testing.T) {
	_, err := m.GetTransfersOutgoing(licenseNumber, nil, nil)
	assert.NoError(t, err)
}

func TestTransfersRejected_Integration(t *testing.T) {
	_, err := m.GetTransfersRejected(licenseNumber, nil, nil)
	assert.NoError(t, err)
}

func TestTransfersDeliveryPackagesStates_Integration(t *testing.T) {
	states, err := m.GetTransfersDeliveryPackagesStates()
	assert.NoError(t, err)
	expected := []string{"Shipped", "Rejected", "Accepted", "Returned"}
	assert.EqualValues(t, expected, states)
}
