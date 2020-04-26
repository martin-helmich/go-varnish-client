package varnishclient

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

var client *Client

func ExampleClient_Authenticate() {
	if client.AuthenticationRequired() {
		secret, err := ioutil.ReadFile("/etc/varnish/secret")
		if err != nil {
			panic(err)
		}

		if err := client.Authenticate(ctx, secret); err != nil {
			panic(err)
		}

		// You're authenticated. Yay!
	}
}

func TestAuthenticate(t *testing.T) {
	client, err := DialTCP(ctx, "0.0.0.0:6082")

	require.NoError(t, err)
	require.True(t, client.AuthenticationRequired())

	err = client.Authenticate(ctx, []byte("72be6aba-00c4-4908-a99f-0e4eb7cc86ca\n"))
	require.NoError(t, err)
}

func TestAuthenticateWrongSecret(t *testing.T) {
	client, err := DialTCP(ctx, "0.0.0.0:6082")

	require.NoError(t, err)
	require.True(t, client.AuthenticationRequired())

	err = client.Authenticate(ctx, []byte("72be6aba-00c4-4908-a99f-0e4eb7cc86cb\n"))
	require.Error(t, err)
}
