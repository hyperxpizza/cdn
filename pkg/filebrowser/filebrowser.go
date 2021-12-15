package filebrowser

import (
	"sync"
)

type FileBrowser struct {
	mutex sync.Mutex
}

func (fb *FileBrowser) FindFile(name string) (*File, error) {
	fb.mutex.Lock()
	defer fb.mutex.Unlock()

	var file File

	return &file, nil
}

func (fb *FileBrowser) DeleteFile(id int) error {
	fb.mutex.Lock()
	defer fb.mutex.Unlock()

	return nil
}

func (fb *FileBrowser) SaveFile(data []byte, name, bucket string) error {
	return nil
}

func (fb *FileBrowser) GetFile(id int) {

}
