package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetParam", func() {
	When("setting a parameter", Ordered, func() {
		var client *varnishclient.Client

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			Expect(client.SetParameter(ctx, "backend_idle_timeout", "300")).To(Succeed())
		})

		It("should be able to retrieve the same value", func() {
			p, err := client.GetParameter(ctx, "backend_idle_timeout")
			Expect(err).NotTo(HaveOccurred())
			Expect(p.Value).To(Equal("300.000"))
		})
	})
})
