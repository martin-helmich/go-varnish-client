package varnishclient

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// ListBackends returns the list of available backends.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#backend-list-j-p-backend-pattern
func (c *Client) ListBackends(ctx context.Context, pattern string) (BackendsResponse, error) {
	args := []string{}

	if pattern != "" {
		args = append(args, "-p", strconv.Quote(pattern))
	}

	resp, err := c.roundtrip.Execute(ctx, &Request{"backend.list", args})
	if err != nil {
		return nil, err
	}

	if resp.Code != ResponseOK {
		return nil, fmt.Errorf("could not list backends (code %d): %s", resp.Code, string(resp.Body))
	}

	lines := strings.Split(string(resp.Body), "\n")[1:]
	backends := make(BackendsResponse, 0, len(lines))

	for i := range lines {
		name := strings.Split(lines[i], " ")[0]
		if name == "" {
			continue
		}

		backend := Backend{Name: name}
		backends = append(backends, backend)
	}

	return backends, nil
}
