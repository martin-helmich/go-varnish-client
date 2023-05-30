package banlist_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBanlist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Banlist Suite")
}
