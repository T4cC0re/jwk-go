package okp_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOkp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Okp Suite")
}
