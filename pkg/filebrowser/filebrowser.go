package filebrowser

import (
	"sync"

	"github.com/hyperxpizza/cdn/pkg/database"
)

type FileBrowser struct {
	db    database.Database
	mutex sync.Mutex
}

func (fb *FileBrowser) FindFile(name string) {
	fb.mutex.Lock()
	defer fb.mutex.Unlock()
}
