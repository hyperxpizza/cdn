package router

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/cdn/pkg/compressor"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
)

func (a *API) download(c *gin.Context) {
	bucket := c.Param("bucket")
	name := c.Param("name")

	if bucket == "" || name == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//check if file exists in the database
	err := a.db.CheckIfFileExists(name, bucket)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	/*
		//get file from the filebrowser
		file, err := a.fb.GetFile(name, bucket)
		if err != nil {
			if errors.Is(err, customErrors.Wrap(customErrors.ErrBucketNotFound)) || errors.Is(err, customErrors.Wrap(customErrors.ErrFileNotFound)) {
				c.Status(http.StatusNotFound)
				return
			}

			c.Status(http.StatusInternalServerError)
			return
		}


			data, err := ioutil.ReadAll(file)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				return
			}

			c.File()
	*/
}

func (a *API) upload(c *gin.Context) {
	fileData, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	bucket := c.Param("bucket")
	if bucket == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	data, err := fileData.Open()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	defer data.Close()

	//extenstion := filepath.Ext(fileData.Filename)

	byteData, err := ioutil.ReadAll(data)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	size := len(byteData)
	compressedSize, compressedData, err := compressor.CompressFile(byteData)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	err = a.fb.SaveFile(compressedData, fileData.Filename, bucket)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	mimeType := mimetype.Detect(byteData)

	file := filebrowser.NewFile(fileData.Filename, bucket, uint64(size), uint64(compressedSize), mimeType.String())

	err = a.db.AddFile(file)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (a *API) search(c *gin.Context) {
	searchQuery := c.Query("q")
	if searchQuery == "" {
		c.Status(http.StatusNoContent)
		return
	}

	files, err := a.db.SearchFiles(searchQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

func (a *API) delete(c *gin.Context) {

}
