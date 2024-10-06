package file

import (
	"io/fs"
	"time"
)

type File struct {
	Name			 string
	Size 			 int64
	Mode 			 fs.FileMode
	ModificationTime time.Time
	Content			 string
}
