package varnishclient

import (
	"context"
	"fmt"
	"strconv"
)

// LoadVCL compiles and loads the VCL file from the given file.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-load-configname-filename-auto-cold-warm
func (c *Client) LoadVCL(ctx context.Context, configname, filename string, mode VCLState) error {
	args := []string{configname, strconv.Quote(filename), string(mode)}
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.load", args})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while loading VCL (code %d): %s", resp.Code, string(resp.Body))
	}

	return nil
}
