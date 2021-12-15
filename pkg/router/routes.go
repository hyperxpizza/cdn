package router

import (
	"net/http"

	"github.com/hyperxpizza/cdn/pkg/compressor"
)

func (a *API) download(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (a *API) upload(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseMultipartForm(a.cfg.Uploader.MaxFileSize); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileData, handler, err := req.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer fileData.Close()

	var data []byte
	_, err = fileData.Read(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileName := handler.Filename
	mimeTypeHeader := handler.Header
	mimetype := mimeTypeHeader.Get("Content-Type")
	sizeBeforeCompression := handler.Size

	//file := filebrowser.NewFile(fileName, uint64(sizeBeforeCompression), mimetype)

	//compress the file
	compressedSize, compressedData, err := compressor.CompressFile(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//save into the bucket

	//insert record into the database

	w.WriteHeader(http.StatusOK)

}

func (a *API) search(w http.ResponseWriter, req *http.Request) {
	req.URL.Query()

}
