package filebrowser

import "time"

type File struct {
	ID      int
	Data    []byte
	Size    uint
	Created time.Time
	Updated time.Time
}
