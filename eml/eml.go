// Package eml reads and creates eml.xml files that contain provenance
// metadata from DarwinCore Archive files. EML (Ecological Metadata Language)
// is a metadata standard originally developed for the ecoloty discipline.
// GBIF uses it to keep metadata about DarwinCore Archive packages.
package eml

import (
	"encoding/xml"

	"github.com/GlobalNamesArchitecture/darwincore/util"
)

// EML keeps metadata of the package
type EML struct {
	XMLName  xml.Name `xml:"eml://ecoinformatics.org/eml-2.1.0 eml"`
	Dataset  Dataset  `xml:"dataset"`
	Citation string   `xml:"additionalMetadata>metadata>citation,omitempty"`
}

// Dataset most provenance information
type Dataset struct {
	ID                string             `xml:"id,attr"`
	Title             string             `xml:"title"`
	Creators          []Creator          `xml:"creator,omitempty"`
	MetadataProviders []MetadataProvider `xml:"metadataProvider,omitempty"`
	Distributions     []Distribution     `xml:"distribution,omitempty"`
	Publisher         string             `xml:"publisher>organizationName,omitempty"`
}

// Creator of the data in the archive
type Creator struct {
	ID        string `xml:"id,attr"`
	Scope     string `xml:"scope,attr"`
	GivenName string `xml:"individualName>givenName"`
	SurName   string `xml:"individualName>surName"`
	Email     string `xml:"electronicMailAddress"`
}

// MetadataProvider provides organiztion that was responsible
// for creation of the data in the archive
type MetadataProvider struct {
	Organization string `xml:"organizationName"`
}

// Distribution provides URL to the archive
type Distribution struct {
	URL string `xml:"online>url"`
}

// Read converts xml data to EML structure
func Read(e []byte) EML {
	var emlxml EML
	err := xml.Unmarshal(e, &emlxml)
	util.Check(err)
	return emlxml
}
