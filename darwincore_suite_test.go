package darwincore_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDarwincore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Darwincore Suite")
}
