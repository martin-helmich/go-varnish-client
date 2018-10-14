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

	assert.Len(t, params, 10)
}