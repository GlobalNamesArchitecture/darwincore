// Package archive is responsible for packing and unpacking data kept in a
// DarwinCore archive. It allows tar and zip files.
package archive

import (
	"github.com/dimus/gnidump/util"
	filetype "gopkg.in/h2non/filetype.v1"
	types "gopkg.in/h2non/filetype.v1/types"
)

type FilePath string

// Extracts files from the archive.
func Extract(p FilePath) {
}

func FileType(p FilePath) types.Type {
	t, err := filetype.MatchFile(string(p))
	util.Check(err)
	return t
}
