package filebrowser

import (
	"database/sql"
	"sync"
)

type FileBrowser struct {
	db    *sql.DB
	mutex sync.Mutex
}

func (fb *FileBrowser) FindFile(name string) {

}
