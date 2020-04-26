package varnishclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleClient Client
var ctx context.Context

func init() {
	ctx = context.Background()
}

func buildTestClient(t *testing.T) *Client {
	client, err := DialTCP(ctx, "0.0.0.0:6082")

	require.NoError(t, err)
	require.True(t, client.AuthenticationRequired())

	err = client.Authenticate(ctx, []byte("72be6aba-00c4-4908-a99f-0e4eb7cc86ca\n"))
	require.NoError(t, err)

	return client
}
