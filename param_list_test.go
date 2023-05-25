package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListParameters", func() {
	When("listing parameters", Ordered, func() {
		var client *varnishclient.Client
		var resp varnishclient.ParametersResponse
		var err error

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			resp, err = client.ListParameters(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should contain a plausible number of parameters", func() {
			// Don't assert on the exact number of parameters, as this may change
			// between varnish versions
			Expect(len(resp)).To(BeNumerically(">=", 100))
		})

		It("should contain the acceptor_sleep_incr parameter", func() {
			Expect(resp).To(ContainElement(And(
				HaveField("Name", Equal("acceptor_sleep_incr")),
				HaveField("Value", Equal("0.000")),
				HaveField("Unit", Equal("seconds")),
				HaveField("IsDefault", BeTrue()),
			)))
		})
	})
})
