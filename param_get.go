package varnishclient

import (
	"context"
	"fmt"
	"strings"
)

// GetParameter returns a single parameter
func (c *Client) GetParameter(ctx context.Context, name string) (*Parameter, error) {
	params, err := c.ListParameters(ctx)
	if err != nil {
		return nil, err
	}

	for i := range params {
		if params[i].Name == name {
			return &params[i], nil
		}
	}

	return nil, fmt.Errorf("parameter not found: %s", name)
}

func paramFromLine(line string) Parameter {
	param := Parameter{}

	items := strings.SplitN(line, " ", 2)
	param.Name = items[0]

	if len(items) > 1 {
		val := strings.TrimSpace(items[1])

		if defaultRegex.MatchString(val) {
			param.IsDefault = true
			val = defaultRegex.ReplaceAllString(val, "")
		}

		uMatches := unitRegex.FindStringSubmatch(val)
		if len(uMatches) >= 1 {
			param.Unit = uMatches[1]
			val = unitRegex.ReplaceAllString(val, "")
		}

		param.Value = val
	}

	return param
}
