package handlers

import "github.com/hyperxpizza/cdn/pkg/filebrowser"

type Downloader struct {
	fb *filebrowser.FileBrowser
}

func (d *Downloader) DownloadFile(path string) ([]byte, error) {
	var file []byte
	return file, nil
}
