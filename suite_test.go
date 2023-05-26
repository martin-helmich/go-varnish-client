package varnishclient_test

import (
	"context"
	"testing"

	varnishclient "github.com/martin-helmich/go-varnish-client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func buildTestClient() *varnishclient.Client {
	client, err := varnishclient.DialTCP(ctx, "0.0.0.0:6082")

	Expect(err).NotTo(HaveOccurred())
	Expect(client).NotTo(BeNil())
	Expect(client.AuthenticationRequired()).To(BeTrue())

	Expect(client.Authenticate(ctx, []byte("72be6aba-00c4-4908-a99f-0e4eb7cc86ca\n"))).To(Succeed())

	return client
}

func TestGoVarnishClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoVarnishClient Suite")
}
