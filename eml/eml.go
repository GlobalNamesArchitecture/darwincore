// EmlXML package reads and creates eml.xml files that contain provenance
// metadata from DarwinCore Archive files. EML (Ecological Metadata Language)
// is a metadata standard originally developed for the ecoloty discipline.
// GBIF uses it to keep metadata about DarwinCore Archive packages.
package emlxml

import (
	"encoding/xml"

	"github.com/dimus/darwincore/util"
)

type EML struct {
	XMLName  xml.Name `xml:"eml://ecoinformatics.org/eml-2.1.0 eml"`
	Dataset  Dataset  `xml:"dataset"`
	Citation string   `xml:"additionalMetadata>metadata>citation,omitempty"`
}

type Dataset struct {
	ID                string             `xml:"id,attr"`
	Title             string             `xml:"title"`
	Creators          []Creator          `xml:"creator,omitempty"`
	MetadataProviders []MetadataProvider `xml:"metadataProvider,omitempty"`
	Distributions     []Distribution     `xml:"distribution,omitempty"`
	Publisher         string             `xml:"publisher>organizationName,omitempty"`
}

type Creator struct {
	ID        string `xml:"id,attr"`
	Scope     string `xml:"scope,attr"`
	GivenName string `xml:"individualName>givenName"`
	SurName   string `xml:"individualName>surName"`
	Email     string `xml:"electronicMailAddress"`
}

type MetadataProvider struct {
	Organization string `xml:"organizationName"`
}

type Distribution struct {
	URL string `xml:"online>url"`
}

func ReadEML(e []byte) EML {
	var eml EML
	err := xml.Unmarshal(e, &eml)
	util.Check(err)
	return eml
}
