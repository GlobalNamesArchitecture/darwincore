package darwincore_test

import (
	"fmt"
	"io/ioutil"

	"github.com/GlobalNamesArchitecture/darwincore/archive"
	"github.com/GlobalNamesArchitecture/darwincore/util"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Archive", func() {
	Describe("FileType()", func() {
		dir := "./testdata"
		files, err := ioutil.ReadDir(dir)
		util.Check(err)
		for _, v := range files {
			path := fmt.Sprintf("%s/%s", dir, v.Name())
			fmt.Println(archive.FileType(FilePath(path)))
		}

	})
})
