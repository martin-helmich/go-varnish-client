package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BanList", func() {
	It("should return a list of bans", func() {
		client := buildTestClient()

		bans, err := client.ListBans(ctx)

		Expect(err).ToNot(HaveOccurred())
		Expect(bans).To(HaveLen(1))
		Expect(bans[0].Time).ToNot(BeZero())
		Expect(bans[0].Objects).To(BeEquivalentTo(0))
		Expect(bans[0].Status).To(Equal(varnishclient.BanComplete))
	})
})
