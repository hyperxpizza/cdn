package filebrowser

import "time"

type File struct {
	ID                   int
	Name                 string
	Bucket               string
	Size                 uint64
	SizeAfterCompression uint64
	MimeType             string
	Created              time.Time
	Updated              time.Time
}

func NewFile(name string, size uint64, mimeType string) *File {
	return &File{
		Name:     name,
		Size:     size,
		MimeType: mimeType,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
}
