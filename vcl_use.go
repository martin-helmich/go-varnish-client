package varnishclient

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (c *Client) UseVCL(ctx context.Context, configname string) error {
	resp, err := c.roundtrip.Execute(ctx, &Request{"vcl.use", []string{strconv.Quote(configname)}})
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("error while activating VCL (code %d): %s", resp.Code, strings.TrimSpace(string(resp.Body)))
	}

	return nil
}
