package router

import (
	"errors"
	"net/http"

	"github.com/hyperxpizza/cdn/pkg/compressor"
	"github.com/hyperxpizza/cdn/pkg/customErrors"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
)

func (a *API) download(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//check if file exists in the database

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
	err = a.fb.SaveFile(compressedData, fileName, "test-bucket")
	if err != nil {
		if errors.Is(customErrors.Wrap(customErrors.ErrBucketAlreadyExists), err) {
			w.WriteHeader(http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bucket := "test-bucket"
	file := filebrowser.NewFile(fileName+".gz", bucket, uint64(sizeBeforeCompression), uint64(compressedSize), mimetype)

	//insert record into the database
	err = a.db.AddFile(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (a *API) search(w http.ResponseWriter, req *http.Request) {
	req.URL.Query()

}
