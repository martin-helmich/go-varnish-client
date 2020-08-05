package varnishclient

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// ListVCL lists the compiled VCLs.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-list-j
func (c *Client) ListVCL(ctx context.Context) (VCLConfigsResponse, error) {
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.list", nil})
	if err != nil {
		return nil, err
	}

	if resp.Code != ResponseOK {
		return nil, fmt.Errorf("could not list VCL configurations (code %d): %s", resp.Code, string(resp.Body))
	}
	lines := strings.Split(string(resp.Body), "\n")
	vclConfigs := make(VCLConfigsResponse, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		name := fields[len(fields)-1]
		vclConfig := VCLConfig{Name: name}
		status := fields[0]
		switch status {
		case "active":
			vclConfig.Status = VCLActive
		case "disabled":
			vclConfig.Status = VCLDiscarded
		case "available":
			vclConfig.Status = VCLAvailable
		default:
			vclConfig.Status = VCLUnknown
		}
		if backends, err := strconv.Atoi(fields[len(fields)-2]); err == nil {
			vclConfig.ActiveBackends = backends
		}
		vclConfigs = append(vclConfigs, vclConfig)
	}
	return vclConfigs, nil
}
