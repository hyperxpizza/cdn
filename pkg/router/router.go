package router

import (
	"log"
	"net/http"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
)

type API struct {
	fb  *filebrowser.FileBrowser
	cfg *config.Config
}

func NewApi() *API {
	return &API{}
}

func Run(c *config.Config) {

	api := NewApi()

	server := http.NewServeMux()

	server.HandleFunc("/download", api.download)
	server.HandleFunc("/upload", api.upload)
	server.HandleFunc("/search", api.search)

	if err := http.ListenAndServe(":8888", server); err != nil {
		log.Fatal(err)
	}
}
