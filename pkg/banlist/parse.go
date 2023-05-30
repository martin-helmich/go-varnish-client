package banlist

import (
	"strconv"
	"strings"
	"time"
)

func Parse(list string) (BanList, error) {
	lines := strings.Split(list, "\n")
	response := make(BanList, 0, len(lines))
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
