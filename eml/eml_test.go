package emlxml

import (
	"encoding/xml"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/dimus/darwincore/util"
)

func testData() []byte {
	eml, err := ioutil.ReadFile("../testdata/eml.xml")
	util.Check(err)
	return eml
}

func TestRead(t *testing.T) {
	eml := testData()
	res := ReadEML(eml)
	expected := EML{
		XMLName: xml.Name{Space: "eml://ecoinformatics.org/eml-2.1.0",
			Local: "eml"},
		Dataset: Dataset{ID: "leptogastrinae:version:2.5",
			Title: "Leptogastrinae (Diptera: Asilidae) Classification",
			Creators: []Creator{
				{"10", "document", "Keith", "Bayless", "keith.bayless@gmail.com"},
				{"5", "document", "Torsten", "Dikow", "dshorthouse@eol.org"},
			},
			MetadataProviders: []MetadataProvider{
				{Organization: "Encyclopedia of Life: LifeDesks (http://www.lifedesks.org)"},
			},
			Distributions: []Distribution{
				{URL: "http://leptogastrinae.lifedesks.org/files/leptogastrinae/classification_export/shared/leptogastrinae.tar.gz"},
			},
			Publisher: "The Marine Biological Laboratory",
		},
		Citation: "Dikow, Torsten. 2010. The Leptogastrinae classification.",
	}
	if !reflect.DeepEqual(res, expected) {
		t.Error("Result:", res)
		t.Error("Expect:", expected)
	}
}
