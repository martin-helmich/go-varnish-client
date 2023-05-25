package banlist_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"

	"github.com/martin-helmich/go-varnish-client/pkg/banlist"
)

var _ = Describe("Parse", func() {
	tests := []struct {
		name     string
		response string
		expected banlist.BanList
	}{
		{
			name: "two",
			response: `Present bans:
1619165086.072521     2 -  obj.http.url ~ test
1619102459.239601     0 C
`,
			expected: banlist.BanList{
				banlist.Ban{
					Time:    time.Unix(1619165086, 72521000),
					Objects: 2,
					Status:  banlist.BanActive,
					Spec:    "obj.http.url ~ test",
				},
				banlist.Ban{
					Time:    time.Unix(1619102459, 239601000),
					Objects: 0,
					Status:  banlist.BanComplete,
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
			expected: banlist.BanList{
				banlist.Ban{
					Time:    time.Unix(1619166708, 197636000),
					Objects: 1,
					Status:  banlist.BanActive,
					Spec:    "obj.http.url ~ test",
				},
				banlist.Ban{
					Time:    time.Unix(1619166701, 830047000),
					Objects: 0,
					Status:  banlist.BanComplete,
					Spec:    "",
				},
				banlist.Ban{
					Time:    time.Unix(1617106224, 959675000),
					Objects: 45,
					Status:  banlist.BanComplete,
					Spec:    "",
				},
			},
		},
		{
			name: "one",
			response: `Present bans:
1619102459.239601     0 C
`,
			expected: banlist.BanList{banlist.Ban{
				Time:    time.Unix(1619102459, 239601000),
				Objects: 0,
				Status:  banlist.BanComplete,
				Spec:    "",
			}},
		},
	}

	for _, test := range tests {
		It("should parse ban list", func() {
			banList, err := banlist.Parse(test.response)
			Expect(err).ToNot(HaveOccurred())
			Expect(banList).To(Equal(test.expected))
		})
	}
})
