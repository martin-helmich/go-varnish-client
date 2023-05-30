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
	// NOTE
	// Output for varnish 6:
	//
	// varnish> vcl.list
	// 200
	// available   auto/warm          0 boot (1 label)
	// available   auto/warm          0 state-7752
	// active      auto/warm          0 use-0132
	// available   auto/warm          0 inline-7036
	// available  label/warm          0 label-5509 -> boot
	//
	// Output for varnish 7:
	//
	// varnish> vcl.list
	// 200
	// available   auto     warm         0    boot           <-    (1 label)
	// active      auto     warm         0    use-7500
	// available   label    warm         0    label-0737     ->    boot
	// available   auto     warm         0    state-9344
	// available   auto     warm         0    inline-4907
	//
	// TODO: This is a hacky way to support both varnish 6 and 7. Why don't we
	// just use the "-j" flag and parse the JSON output?

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

		var name, status, activeBackendStr string

		if _, err := strconv.Atoi(fields[2]); err == nil {
			name = fields[3]
			status = fields[0]
			activeBackendStr = fields[2]
		} else {
			name = fields[4]
			status = fields[0]
			activeBackendStr = fields[3]
		}

		vclConfig := VCLConfig{Name: name}

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
		if backends, err := strconv.Atoi(activeBackendStr); err == nil {
			vclConfig.ActiveBackends = backends
		}
		vclConfigs = append(vclConfigs, vclConfig)
	}
	return vclConfigs, nil
}
