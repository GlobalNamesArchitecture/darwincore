package darwincore_test

import (
	"encoding/xml"

	"github.com/GlobalNamesArchitecture/darwincore/eml"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Eml", func() {
	Describe("ReadEML", func() {
		expectedEML := eml.EML{
			XMLName: xml.Name{Space: "eml://ecoinformatics.org/eml-2.1.0",
				Local: "eml"},
			Dataset: eml.Dataset{ID: "leptogastrinae:version:2.5",
				Title: "Leptogastrinae (Diptera: Asilidae) Classification",
				Creators: []eml.Creator{
					{"10", "document", "Keith", "Bayless", "keith.bayless@gmail.com"},
					{"5", "document", "Torsten", "Dikow", "dshorthouse@eol.org"},
				},
				MetadataProviders: []eml.MetadataProvider{
					{Organization: "Encyclopedia of Life: LifeDesks (http://www.lifedesks.org)"},
				},
				Distributions: []eml.Distribution{
					{URL: "http://leptogastrinae.lifedesks.org/files/leptogastrinae/classification_export/shared/leptogastrinae.tar.gz"},
				},
				Publisher: "The Marine Biological Laboratory",
			},
			Citation: "Dikow, Torsten. 2010. The Leptogastrinae classification.",
		}

		It("injests eml data", func() {
			emlBody := testData("./testdata/eml.xml")
			Expect(eml.Read(emlBody)).To(Equal(expectedEML))
		})
	})
})
