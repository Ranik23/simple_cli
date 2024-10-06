package file

import (
	"os"
	"time"
)

type File struct {
	Name 			 string
	Size 			 int64
	Mode 			 os.FileMode
	ModificationTime time.Time
	Content			 string
}