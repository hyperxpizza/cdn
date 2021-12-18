package filebrowser

import "time"

type File struct {
	ID                   int
	Name                 string
	Bucket               string
	Size                 uint64
	SizeAfterCompression uint64
	Extension            string
	MimeType             string
	Created              time.Time
	Updated              time.Time
}

func NewFile(name string, bucket string, size uint64, sizeAfterCompression uint64, mimeType string) File {
	return File{
		Name:                 name,
		Bucket:               bucket,
		Size:                 size,
		SizeAfterCompression: sizeAfterCompression,
		MimeType:             mimeType,
		Created:              time.Now(),
		Updated:              time.Now(),
	}
}
