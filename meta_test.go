package darwincore_test

import (
	"encoding/xml"

	"github.com/GlobalNamesArchitecture/darwincore/meta"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Meta", func() {
	Describe("ReadMeta()", func() {
		expectedMeta := meta.Meta{
			XMLName: xml.Name{Space: "http://rs.tdwg.org/dwc/text/",
				Local: "archive"},
			Core: meta.Core{
				ID: meta.Field{
					Index: "0",
					Term:  "http://rs.tdwg.org/dwc/terms/TaxonID",
				},
				Fields: []meta.Field{
					{
						Index: "1",
						Term:  "http://purl.org/dc/terms/source",
					},
					{
						Index: "2",
						Term:  "http://rs.tdwg.org/dwc/terms/ScientificName",
					},
					{
						Index: "3",
						Term:  "http://rs.tdwg.org/dwc/terms/HigherTaxonID",
					},
					{
						Index: "4",
						Term:  "http://rs.tdwg.org/dwc/terms/TaxonRank",
					},
					{
						Index: "5",
						Term:  "http://rs.tdwg.org/dwc/terms/TaxonomicStatus",
					},
				},
			},
			Extensions: []meta.Extension{
				{
					CoreID: meta.Field{Index: "0", Term: ""},
					Fields: []meta.Field{
						{
							Index: "1",
							Term:  "http://rs.gbif.org/ecat/terms/vernacularName",
						},
						{
							Index: "2",
							Term:  "http://rs.gbif.org/thesaurus/languageCode",
						},
					},
				},
				{
					CoreID: meta.Field{Index: "0", Term: ""},
					Fields: []meta.Field{
						{
							Index: "1",
							Term:  "http://rs.tdwg.org/dwc/terms/locationID",
						},
						{
							Index: "2",
							Term:  "http://rs.tdwg.org/dwc/terms/locality",
						},
						{
							Index: "3",
							Term:  "http://rs.tdwg.org/dwc/terms/country",
						},
						{
							Index: "4",
							Term:  "http://rs.tdwg.org/dwc/terms/countryCode",
						},
					},
				},
			},
		}

		It("Reads metadta", func() {
			d := testData("./testdata/meta.xml")
			Expect(meta.Read(d)).To(Equal(expectedMeta))
		})
	})
})
