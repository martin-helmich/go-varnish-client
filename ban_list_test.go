package varnishclient

import (
	"testing"
	"time"
)

func TestBanList(t *testing.T) {
	sut := &Client{}
	tests := []struct {
		name     string
		response string
		expected BanListResponse
	}{
		{
			name: "two",
			response: `Present bans:
1619165086.072521     2 -  obj.http.url ~ test
1619102459.239601     0 C
`,
			expected: BanListResponse{
				Ban{
					Time:    time.Unix(1619165086, 72521000),
					Objects: 2,
					Status:  BanActive,
					Spec:    "obj.http.url ~ test",
				},
				Ban{
					Time:    time.Unix(1619102459, 239601000),
					Objects: 0,
					Status:  BanComplete,
					Spec:    "",
				},
			},
		},
		{
			name: "three",
			response: `Present bans:
1619166708.197636     1 -  obj.http.url ~ test
1619166701.830047     0 C
1617106224.959675    45 C
`,
			expected: BanListResponse{
				Ban{
					Time:    time.Unix(1619166708, 197636000),
					Objects: 1,
					Status:  BanActive,
					Spec:    "obj.http.url ~ test",
				},
				Ban{
					Time:    time.Unix(1619166701, 830047000),
					Objects: 0,
					Status:  BanComplete,
					Spec:    "",
				},
				Ban{
					Time:    time.Unix(1617106224, 959675000),
					Objects: 45,
					Status:  BanComplete,
					Spec:    "",
				},
			},
		},
		{
			name: "one",
			response: `Present bans:
1619102459.239601     0 C
`,
			expected: BanListResponse{Ban{
				Time:    time.Unix(1619102459, 239601000),
				Objects: 0,
				Status:  BanComplete,
				Spec:    "",
			}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			banList, err := sut.parseBanList(test.response)
			if err != nil {
				t.Error(err)
				t.Fail()
				return
			}
			if len(banList) != len(test.expected) {
				t.Fail()
			}
			for i, ban := range test.expected {
				if !ban.Equals(banList[i]) {
					t.Log("bans are not equal")
					t.Fail()
				}
			}
		})
	}
}
