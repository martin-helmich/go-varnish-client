package varnishclient

import (
	"context"
	"fmt"
	"strings"
)

// AddLabelToVCL adds a label to a configuration file.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-label-label-configname
func (c *Client) AddLabelToVCL(ctx context.Context, label string, configname string) error {
	args := []string{label, configname}
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.label", args})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while labelling VCL (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
