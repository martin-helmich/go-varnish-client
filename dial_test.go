package varnishclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDialTCP(t *testing.T) {
	client, err := DialTCP("0.0.0.0:6082")

	require.NoError(t, err)
	require.True(t, client.AuthenticationRequired())
}
