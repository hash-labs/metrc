package metrc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Think about a testing strategy for the client that doesn't require Metrc calls.
// At its core we want to confirm that we can make requests that have the basic auth header.
var hc *HttpClient = MakeHttpClient()

func TestHttpClientGet(t *testing.T) {
	endpoint := "facilities/v1"
	rb, err := hc.Get(endpoint)
	assert.NoError(t, err)
	assert.NotEmpty(t, rb)
}

// TODO: Add tests for other client functions.
