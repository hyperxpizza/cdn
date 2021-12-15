package filebrowser

import (
	"fmt"
	"os"
	"sync"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/customErrors"
)

type FileBrowser struct {
	mutex    sync.RWMutex
	rootPath string
}

func NewFileBrowser(c config.Config) *FileBrowser {
	return &FileBrowser{sync.RWMutex{}, c.FileBrowser.Rootpath}
}

func (fb *FileBrowser) FindFile(name, bucket string) (*File, error) {
	fb.mutex.RLock()
	defer fb.mutex.RUnlock()

	var file File
	if !fb.CheckIfBucketExists(bucket) {
		return nil, customErrors.Wrap(customErrors.ErrBucketNotFound)
	}

	return &file, nil
}

func (fb *FileBrowser) DeleteFile(name, bucket string) error {
	fb.mutex.Lock()
	defer fb.mutex.Unlock()

	if !fb.CheckIfBucketExists(bucket) {
		return customErrors.Wrap(customErrors.ErrBucketNotFound)
	}

	return nil
}

func (fb *FileBrowser) SaveFile(data []byte, name, bucket string) error {
	fb.mutex.Lock()
	defer fb.mutex.Unlock()

	if !fb.CheckIfBucketExists(bucket) {
		err := fb.CreateBucket(bucket)
		if err != nil {
			return err
		}
	}

	//check if file exists

	name = name + ".gz"
	fullPath := fmt.Sprintf("%s/%s/%s", fb.rootPath, bucket, name)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (fb *FileBrowser) GetFile(name, bucket string) (*os.File, error) {
	var file os.File

	fb.mutex.RLock()
	defer fb.mutex.RUnlock()

	if !fb.CheckIfBucketExists(bucket) {
		return nil, customErrors.Wrap(customErrors.ErrBucketNotFound)
	}

	return &file, nil
}

func (fb *FileBrowser) CreateBucket(name string) error {
	if fb.CheckIfBucketExists(name) {
		return customErrors.Wrap(customErrors.ErrBucketAlreadyExists)
	}

	return nil
}

func (fb *FileBrowser) CheckIfFileExists(name, bucket string) bool {

}

func (fb *FileBrowser) CheckIfBucketExists(name string) bool {
	path := fmt.Sprintf("%s/%s", fb.rootPath, name)
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
