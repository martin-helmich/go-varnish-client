package varnishclient

import (
	"fmt"
	"strconv"
	"strings"
)

// DiscardVCL unloads the VCL file specified by the given "configname".
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-discard-configname-label
func (c *Client) DiscardVCL(configname string) error {
	resp, err := c.sendRequest("vcl.discard", strconv.Quote(configname))
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while discarding VCL (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
