package varnishclient_test

import (
	"fmt"
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"math/rand"
)

var _ = Describe("AddLabelToVCL", func() {

	When("adding a label to a VCL", Ordered, func() {
		var label = fmt.Sprintf("label-%04d", rand.Intn(9999))
		var client *varnishclient.Client

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			Expect(client.AddLabelToVCL(ctx, label, "boot")).To(Succeed())
		})

		It("should be able to retrieve the VCL by label", func() {
			vcls, err := client.ListVCL(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(vcls).To(ContainElement(HaveField("Name", Equal(label))))
		})
	})
})
