package varnishclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParameterList(t *testing.T) {
	client, err := DialTCP("127.0.0.1:6082")
	if err != nil {
		t.Error(err)
	}

	err = client.Authenticate([]byte("72be6aba-00c4-4908-a99f-0e4eb7cc86ca\n"))
	if err != nil {
		t.Error(err)
	}

	params, err := client.ListParameters()
	if err != nil {
		t.Error(err)
	}

	assert.Len(t, params, 104)
	assert.Equal(t, "acceptor_sleep_incr", params[2].Name)
	assert.Equal(t, "0.000", params[2].Value)
	assert.Equal(t, "seconds", params[2].Unit)
	assert.True(t, params[2].IsDefault)
}
