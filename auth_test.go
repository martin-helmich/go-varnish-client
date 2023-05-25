package varnishclient_test

import (
	varnishclient "github.com/martin-helmich/go-varnish-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var client *varnishclient.Client

func ExampleClient_Authenticate() {
	if client.AuthenticationRequired() {
		secret, err := ioutil.ReadFile("/etc/varnish/secret")
		if err != nil {
			panic(err)
		}

		if err := client.Authenticate(ctx, secret); err != nil {
			panic(err)
		}

		// You're authenticated. Yay!
	}
}

var _ = Describe("Authentication", func() {
	When("providing the correct secret", func() {
		It("should connect successfully", func() {
			client, err := varnishclient.DialTCP(ctx, "0.0.0.0:6082")

			Expect(err).NotTo(HaveOccurred())
			Expect(client.AuthenticationRequired()).To(BeTrue())

			Expect(client.Authenticate(ctx, []byte("72be6aba-00c4-4908-a99f-0e4eb7cc86ca\n"))).To(Succeed())
		})
	})

	When("providing the wrong secret", func() {
		It("should return an error", func() {
			client, err := varnishclient.DialTCP(ctx, "0.0.0.0:6082")

			Expect(err).NotTo(HaveOccurred())
			Expect(client.AuthenticationRequired()).To(BeTrue())

			Expect(client.Authenticate(ctx, []byte("72be6aba-00c4-4908-a99f-0e4eb7cc86cb\n"))).NotTo(Succeed())
		})
	})
})
