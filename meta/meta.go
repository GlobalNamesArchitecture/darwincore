// Package meta reads and writes fields data of DarwinCore Archive
// file.
package meta

import (
	"encoding/xml"

	"github.com/dimus/gnidump/util"
)

// Meta represent of data stored in meta file
type Meta struct {
	XMLName    xml.Name    `xml:"http://rs.tdwg.org/dwc/text/ archive"`
	Core       Core        `xml:"core"`
	Extensions []Extension `xml:"extension,omitempty"`
}

// Core represents metadata for the taxon file
type Core struct {
	ID     Field   `xml:"id"`
	Fields []Field `xml:"field"`
}

type Extension struct {
	CoreID Field   `xml:"coreid"`
	Fields []Field `xml:"field"`
}

// Field represents a field (or core id) of the archive
type Field struct {
	Index string `xml:"index,attr"`
	Term  string `xml:"term,attr,omitempty"`
}

// Read converts data stored in xml metadata into Meta data structure
func Read(m []byte) Meta {
	var metaData Meta
	err := xml.Unmarshal(m, &metaData)
	util.Check(err)
	return metaData
}
