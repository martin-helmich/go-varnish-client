package varnishclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParameterList(t *testing.T) {
	client := buildTestClient(t)

	params, err := client.ListParameters(ctx)
	if err != nil {
		t.Error(err)
	}

	// Don't assert on the exact number of parameters, as this may change
	// between varnish versions
	assert.GreaterOrEqual(t, len(params), 100)
	assert.Equal(t, "acceptor_sleep_incr", params[2].Name)
	assert.Equal(t, "0.000", params[2].Value)
	assert.Equal(t, "seconds", params[2].Unit)
	assert.True(t, params[2].IsDefault)
}
