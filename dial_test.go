package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DialTCP", Ordered, func() {
	When("dialling the correct address", func() {
		var client *varnishclient.Client
		var err error

		It("should succeed", func() {
			client, err = varnishclient.DialTCP(ctx, "0.0.0.0:6082")
			Expect(err).ToNot(HaveOccurred())
		})

		It("should correctly report authentication requirements", func() {
			Expect(client.AuthenticationRequired()).To(BeTrue())
		})
	})
})
