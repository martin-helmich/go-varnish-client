package varnishclient

import (
	"fmt"
	"regexp"
	"strings"
)

var defaultRegex = regexp.MustCompile(`\s+\(default\)$`)
var unitRegex = regexp.MustCompile(`\s+\[(.*)]$`)

// ListParameters lists all Varnish parameters and their values.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#param-show-l-j-param-changed
func (c *Client) ListParameters() (ParametersResponse, error) {
	resp, err := c.sendRequest("param.show")
	if err != nil {
		return nil, err
	}

	if resp.Code != ResponseOK {
		return nil, fmt.Errorf("could not list parameters (code %d): %s", resp.Code, string(resp.Body))
	}

	lines := strings.Split(string(resp.Body), "\n")
	params := make(ParametersResponse, 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		param := paramFromLine(line)
		params = append(params, param)
	}

	return params, nil
}
