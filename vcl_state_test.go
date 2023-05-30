package varnishclient_test

import (
	"fmt"
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
)

func ExampleClient_SetVCLState() {
	if err := client.SetVCLState(ctx, "boot", varnishclient.VCLStateCold); err != nil {
		// handle error
	}
}

var _ = Describe("SetVCLState", func() {
	When("setting the VCL state", Ordered, func() {
		var client *varnishclient.Client

		name := fmt.Sprintf("state-%04d", rand.Intn(9999))

		BeforeAll(func() {
			client = buildTestClient()
			Expect(client.DefineInlineVCL(ctx, name, []byte("vcl 4.0; backend default { .host = \"127.0.0.1\"; }"), varnishclient.VCLStateAuto)).To(Succeed())
		})

		It("should succeed", func() {
			Expect(client.SetVCLState(ctx, name, varnishclient.VCLStateAuto)).To(Succeed())
		})
	})
})
