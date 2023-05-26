package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func ExampleClient_SetVCLState() {
	if err := client.SetVCLState(ctx, "boot", varnishclient.VCLStateCold); err != nil {
		// handle error
	}
}

var _ = Describe("SetVCLState", func() {
	When("setting the VCL state", Ordered, func() {
		var client *varnishclient.Client

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			Expect(client.SetVCLState(ctx, "boot", varnishclient.VCLStateAuto)).To(Succeed())
		})
	})
})
