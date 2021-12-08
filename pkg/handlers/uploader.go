package handlers

import "github.com/hyperxpizza/cdn/pkg/filebrowser"

type Uploader struct {
	fb *filebrowser.FileBrowser
}

func (u *Uploader) Upload(file *filebrowser.File) error {
	return nil
}
