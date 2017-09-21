package darwincore_test

import (
	"io/ioutil"

	"github.com/GlobalNamesArchitecture/darwincore/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDarwincore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Darwincore Suite")
}

func testData(path string) []byte {
	emlData, err := ioutil.ReadFile(path)
	util.Check(err)
	return emlData
}
