package varnishclient

import (
	"context"
	"fmt"
	"strings"
)

// DiscardVCL unloads the VCL file specified by the given "configname".
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-discard-configname-label
func (c *Client) DiscardVCL(ctx context.Context, configname string) error {
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.discard", []string{configname}})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while discarding VCL (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
