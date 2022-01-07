package filebrowser

import "time"

type File struct {
	ID                   int64     `json:"id"`
	Name                 string    `json:"name"`
	BucketID             int64     `json:"bucket_id"`
	Size                 uint64    `json:"size"`
	SizeAfterCompression uint64    `json:"size_after_compression"`
	Extension            string    `json:"extension"`
	MimeType             string    `json:"mime_type"`
	Created              time.Time `json:"created"`
	Updated              time.Time `json:"updated"`
}

func NewFile(name string, bucketID int64, size uint64, sizeAfterCompression uint64, mimeType string) File {
	return File{
		Name:                 name,
		BucketID:             bucketID,
		Size:                 size,
		SizeAfterCompression: sizeAfterCompression,
		MimeType:             mimeType,
		Created:              time.Now(),
		Updated:              time.Now(),
	}
}

type Bucket struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
