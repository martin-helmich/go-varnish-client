package varnishclient

import (
	"fmt"
	"strconv"
	"strings"
)

// AddLabelToVCL adds a label to a configuration file.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-label-label-configname
func (c *Client) AddLabelToVCL(label string, configname string) error {
	resp, err := c.sendRequest("vcl.label", strconv.Quote(label), strconv.Quote(configname))
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while labelling VCL (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
