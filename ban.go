package varnishclient

import (
	"context"
	"fmt"
)

// Ban creates a ban.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#ban-field-operator-arg-field-oper-arg
func (c *Client) Ban(ctx context.Context, args... string) error {
	resp, err := c.roundtrip.Execute(ctx, &Request{"ban", args})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("could not list VCL configurations (code %d): %s", resp.Code, string(resp.Body))
	}
	return nil
}
