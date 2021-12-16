package router

import (
	"log"
	"net/http"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/database"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
)

type API struct {
	fb  *filebrowser.FileBrowser
	cfg *config.Config
	db  *database.Database
}

func NewApi(c *config.Config) (*API, error) {
	db, err := database.NewDatabase(c)
	if err != nil {
		return nil, err
	}

	fb := filebrowser.NewFileBrowser(*c)

	return &API{
		db:  db,
		fb:  fb,
		cfg: c,
	}, nil
}

func Run(c *config.Config) {

	api, err := NewApi(c)
	if err != nil {
		log.Fatal(err)
	}

	server := http.NewServeMux()

	server.HandleFunc("/download", api.download)
	server.HandleFunc("/upload", api.upload)
	server.HandleFunc("/search", api.search)

	if err := http.ListenAndServe(":8888", server); err != nil {
		log.Fatal(err)
	}
}
