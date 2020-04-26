package varnishclient

import (
	"fmt"
	"strconv"
)

// LoadVCL compiles and loads the VCL file from the given file.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-load-configname-filename-auto-cold-warm
func (c *Client) LoadVCL(configname, filename string, mode VCLState) error {
	resp, err := c.sendRequest("vcl.load", strconv.Quote(configname), strconv.Quote(filename), string(mode))
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while loading VCL (code %d): %s", resp.Code, string(resp.Body))
	}

	return nil
}
