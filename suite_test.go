package varnishclient_test

import (
	"context"
	varnishclient "github.com/martin-helmich/go-varnish-client"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func teardown() error {
	down := exec.Command("docker-compose", "down", "-v")
	down.Stdout = GinkgoWriter
	down.Stderr = GinkgoWriter
	return down.Run()
}

func setup() error {
	up := exec.Command("docker-compose", "up", "-d")
	up.Stdout = GinkgoWriter
	up.Stderr = GinkgoWriter
	return up.Run()
}

var _ = BeforeSuite(func() {
	teardown()
	Expect(setup()).To(Succeed())

	for i := 0; i < 50; i++ {
		_, err := varnishclient.DialTCP(ctx, "0.0.0.0:6082")
		if err == nil {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}
})

var _ = AfterSuite(func() {
	Expect(teardown()).To(Succeed())
})

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
