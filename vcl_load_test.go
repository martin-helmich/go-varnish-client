package varnishclient_test

import (
	"fmt"
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
)

var _ = Describe("LoadVCL", Ordered, func() {
	When("loading a VCL from file system", func() {
		var client *varnishclient.Client
		var name = fmt.Sprintf("load-%04d", rand.Intn(10000))

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			Expect(client.LoadVCL(ctx, name, "/etc/varnish/default.vcl", varnishclient.VCLStateAuto)).To(Succeed())
		})

		It("should be retrievable", func() {
			vcls, err := client.ListVCL(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(vcls).To(ContainElement(HaveField("Name", Equal(name))))
		})
	})
})
