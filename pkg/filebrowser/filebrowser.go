package filebrowser

import (
	"sync"

	"github.com/hyperxpizza/cdn/pkg/database"
)

type FileBrowser struct {
	db    database.Database
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
