package varnishclient_test

import (
	"fmt"
	"math/rand"

	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DefineInlineVCL", func() {
	When("defining a syntactically correct inline VCL", Ordered, func() {
		var client *varnishclient.Client
		var err error

		name := fmt.Sprintf("inline-%04d", rand.Intn(9999))

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should succeed", func() {
			err = client.DefineInlineVCL(ctx, name, []byte("vcl 4.0; backend default { .host = \"127.0.0.1\"; }"), varnishclient.VCLStateAuto)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to retrieve the same value", func() {
			vcls, err := client.ListVCL(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(vcls).To(ContainElement(HaveField("Name", Equal(name))))
		})
	})

	When("defining a bad VCL", Ordered, func() {
		var client *varnishclient.Client
		var err error

		name := fmt.Sprintf("inline-%04d", rand.Intn(9999))

		BeforeAll(func() {
			client = buildTestClient()
		})

		It("should not succeed", func() {
			err = client.DefineInlineVCL(ctx, name, []byte("This is not a valid VCL file"), varnishclient.VCLStateAuto)
			Expect(err).To(HaveOccurred())
		})
	})
})
