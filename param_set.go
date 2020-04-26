package varnishclient

import (
	"context"
	"fmt"
	"strconv"
)

// SetParameter sets a parameter to the specified value.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#param-set-param-value
func (c *Client) SetParameter(ctx context.Context, name, value string) error {
	args := []string{name, strconv.Quote(value)}
	resp, err := c.roundtrip.Execute(ctx, &Request{"param.set", args})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("could not set parameter '%s' (code %d): %s", name, resp.Code, string(resp.Body))
	}

	return nil
}
