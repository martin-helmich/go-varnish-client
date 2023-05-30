package varnishclient_test

import (
	"fmt"
	"math/rand"

	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiscardVCL", func() {
	When("discarding a VCL", Ordered, func() {
		var client *varnishclient.Client

		name := fmt.Sprintf("discard-%04d", rand.Intn(9999))

		BeforeAll(func() {
			client = buildTestClient()
			Expect(client.DefineInlineVCL(ctx, name, []byte("vcl 4.0; backend default { .host = \"127.0.0.1\"; }"), varnishclient.VCLStateAuto)).To(Succeed())
		})

		It("should succeed", func() {
			Expect(client.DiscardVCL(ctx, name)).To(Succeed())
		})

		It("should not be able to retrieve the same value", func() {
			vcls, err := client.ListVCL(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(vcls).NotTo(ContainElement(HaveField("Name", Equal(name))))
		})
	})
})
