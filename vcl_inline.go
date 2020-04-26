package varnishclient

import (
	"fmt"
	"strconv"
)

// DefineInlineVCL compiles and loads a new VCL file with the file contents
// specified by the "vcl" parameter.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-inline-configname-quoted-vclstring-auto-cold-warm
func (c *Client) DefineInlineVCL(configname string, vcl []byte, mode VCLState) error {
	resp, err := c.sendRequest("vcl.inline", strconv.Quote(configname), strconv.Quote(string(vcl)), string(mode))
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while loading VCL (code %d): %s", resp.Code, string(resp.Body))
	}

	return nil
}
