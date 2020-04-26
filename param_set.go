package varnishclient

import (
	"fmt"
	"strconv"
)

// SetParameter sets a parameter to the specified value.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#param-set-param-value
func (c *Client) SetParameter(name, value string) error {
	resp, err := c.sendRequest("param.set", name, strconv.Quote(value))
	if err != nil {
		return err
	}

	if resp.Code != ResponseOK {
		return fmt.Errorf("could not set parameter '%s' (code %d): %s", name, resp.Code, string(resp.Body))
	}

	return nil
}
