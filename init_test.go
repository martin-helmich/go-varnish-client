package varnishclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func buildTestClient(t *testing.T) *Client {
	client, err := DialTCP("0.0.0.0:6082")

	require.NoError(t, err)
	require.True(t, client.AuthenticationRequired())

	err = client.Authenticate([]byte("72be6aba-00c4-4908-a99f-0e4eb7cc86ca\n"))
	require.NoError(t, err)

	return client
}
