package varnishclient

import (
	"context"
	"fmt"
	"strings"
)

// VCLState describes one of the three possible VCL states
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-state-configname-auto-cold-warm
type VCLState string

const (
	// VCLStateAuto means that Varnish should automatically switch the VCL state
	// from "warm" to "cold" and back
	VCLStateAuto VCLState = "auto"

	// VCLStateWarm means that the VCL should always be "warm"
	VCLStateWarm VCLState = "warm"

	// VCLStateCold means that the VCL should always be "cold"
	VCLStateCold VCLState = "cold"
)

// SetVCLState can be used to force a loaded VCL file to a specific state.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-state-configname-auto-cold-warm
func (c *Client) SetVCLState(ctx context.Context, configname string, state VCLState) error {
	args := []string{configname, string(state)}
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.state", args})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while setting VCL state (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
