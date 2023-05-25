package varnishclient

import (
	"context"
	"fmt"
	"github.com/martin-helmich/go-varnish-client/pkg/banlist"
)

// ListBans lists the Current bans.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#ban-field-operator-arg-field-oper-arg
func (c *Client) ListBans(ctx context.Context) (BanListResponse, error) {
	resp, err := c.roundtrip.Execute(ctx, &Request{"ban.list", nil})
	if err != nil {
		return nil, err
	}

	if resp.Code != ResponseOK {
		return nil, fmt.Errorf("could not list VCL configurations (code %d): %s", resp.Code, string(resp.Body))
	}

	return banlist.Parse(string(resp.Body))
}
