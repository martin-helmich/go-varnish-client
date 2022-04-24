package varnishclient

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ListBan lists the Current bans.
// See https://varnish-cache.org/docs/trunk/reference/varnish-cli.html#ban-field-operator-arg-field-oper-arg
func (c *Client) ListBan(ctx context.Context) (BanListResponse, error) {
	resp, err := c.roundtrip.Execute(ctx, &Request{"ban.list", nil})
	if err != nil {
		return nil, err
	}

	if resp.Code != ResponseOK {
		return nil, fmt.Errorf("could not list VCL configurations (code %d): %s", resp.Code, string(resp.Body))
	}
	return c.parseBanList(string(resp.Body))
}

func (c *Client) parseBanList(list string) (BanListResponse, error) {
	lines := strings.Split(list, "\n")
	response := make(BanListResponse, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		lenfields := len(fields)
		if lenfields < 3 {
			continue
		}
		ts := strings.Split(fields[0], ".")
		seconds, err := strconv.ParseInt(ts[0], 10, 64)
		if err != nil {
			return nil, err
		}
		nseconds, err := strconv.ParseInt(ts[1], 10, 64)
		if err != nil {
			return nil, err
		}
		timestamp := time.Unix(seconds, nseconds*1000)
		objects, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return nil, err
		}
		status := BanActive
		switch fields[2] {
		case "G":
			status = BanGone
			break
		case "C":
			status = BanComplete
			break
		}
		spec := ""
		if len(fields) > 3 {
			spec = strings.Join(fields[3:], " ")
		}
		ban := Ban{
			Time:    timestamp,
			Spec:    spec,
			Objects: objects,
			Status:  status,
		}
		response = append(response, ban)
	}
	return response, nil
}
