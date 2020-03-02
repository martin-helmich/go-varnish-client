package varnishclient

import (
	"fmt"
	"strconv"
	"strings"
)

type VCLState string

const (
	VCLStateAuto VCLState = "auto"
	VCLStateWarm          = "warm"
	VCLStateCold          = "cold"
)

// SetVCLState can be used to force a loaded VCL file to a specific state.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-state-configname-auto-cold-warm
func (c *client) SetVCLState(configname string, state VCLState) error {
	resp, err := c.sendRequest("vcl.state", strconv.Quote(configname), strconv.Quote(string(state)))
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while setting VCL state (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
