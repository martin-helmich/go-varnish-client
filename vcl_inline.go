package varnishclient

import (
	"context"
	"fmt"
	"strconv"
)

// DefineInlineVCL compiles and loads a new VCL file with the file contents
// specified by the "vcl" parameter.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#vcl-inline-configname-quoted-vclstring-auto-cold-warm
func (c *Client) DefineInlineVCL(ctx context.Context, configname string, vcl []byte, mode VCLState) error {
	args := []string{strconv.Quote(configname), strconv.Quote(string(vcl)), string(mode)}
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.inline", args})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while loading VCL (code %d): %s", resp.Code, string(resp.Body))
	}

	return nil
}
