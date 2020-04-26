package varnishclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetParam(t *testing.T) {
	client := buildTestClient(t)

	err := client.SetParameter(ctx, "backend_idle_timeout", "300")
	require.NoError(t, err)

	p, err := client.GetParameter(ctx, "backend_idle_timeout")
	require.NoError(t, err)
	require.Equal(t, "300.000", p.Value)
}
