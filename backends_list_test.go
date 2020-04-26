package varnishclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListBackends(t *testing.T) {
	c := buildTestClient(t)

	backends, err := c.ListBackends("")

	require.NoError(t, err)
	require.Len(t, backends, 1)
	require.Equal(t, "boot.default", backends[0].Name)
}

func TestListBackendsWithPattern(t *testing.T) {
	c := buildTestClient(t)

	backends, err := c.ListBackends("nonexistent.*")

	require.NoError(t, err)
	require.Len(t, backends, 0)
}
