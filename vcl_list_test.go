package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListVCL", func() {
	When("listing VCLs", Ordered, func() {
		var client *varnishclient.Client
		var resp varnishclient.VCLConfigsResponse
		var err error

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			resp, err = client.ListVCL(ctx)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not be empty", func() {
			Expect(len(resp)).To(BeNumerically(">", 0))
		})

		It("should contain the boot VCL entry", func() {
			Expect(resp).To(ContainElement(And(
				HaveField("Name", Equal("boot")),
			)))
		})
	})
})
