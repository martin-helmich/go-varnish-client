package varnishclient_test

import (
	"fmt"
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
)

var _ = Describe("UseVCL", func() {
	When("using a VCL", Ordered, func() {
		var client *varnishclient.Client
		var name = fmt.Sprintf("use-%04d", rand.Intn(10000))

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			Expect(client.DefineInlineVCL(ctx, name, []byte("vcl 4.0; backend default { .host = \"127.0.0.1\"; }"), varnishclient.VCLStateAuto)).To(Succeed())
			Expect(client.UseVCL(ctx, name)).To(Succeed())
		})
	})
})
