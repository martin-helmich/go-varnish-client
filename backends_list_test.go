package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListBackends", func() {
	When("listing all backends", func() {
		It("should return a list of backends for the active vcl", func() {
			c := buildTestClient()

			vcls, err := c.ListVCL(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(vcls).NotTo(BeNil())

			var activeVCL string
			for _, vcl := range vcls {
				if vcl.Status == varnishclient.VCLActive {
					activeVCL = vcl.Name
				}
			}

			backends, err := c.ListBackends(ctx, "")

			Expect(err).NotTo(HaveOccurred())
			Expect(backends).NotTo(BeNil())
			Expect(backends).NotTo(HaveLen(0))
			Expect(backends).To(ContainElement(HaveField("Name", Equal(activeVCL+".default"))))
		})
	})

	When("listing backends with a pattern", func() {
		It("should return an empty list when the pattern does not match", func() {
			c := buildTestClient()

			backends, err := c.ListBackends(ctx, "nonexistent.*")

			Expect(err).NotTo(HaveOccurred())
			Expect(backends).NotTo(BeNil())
			Expect(backends).To(HaveLen(0))
		})

		It("should return matching backends when the pattern matches", func() {
			c := buildTestClient()

			backends, err := c.ListBackends(ctx, "boot.*")

			Expect(err).NotTo(HaveOccurred())
			Expect(backends).NotTo(BeNil())
			Expect(backends).To(HaveLen(1))
			Expect(backends[0].Name).To(Equal("boot.default"))
		})
	})
})
