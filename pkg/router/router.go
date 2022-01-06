package router

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

	fb := filebrowser.NewFileBrowser(c)

	return &API{
		db:  db,
		fb:  fb,
		cfg: c,
	}, nil
}

func Run(c *config.Config) error {

	api, err := NewApi(c)
	if err != nil {
		log.Println(err)
		return err
	}

	router := gin.Default()

	router.GET("/serach", api.search)
	router.POST("/upload/:bucket", api.upload)
	router.GET("/download/:bucket/:name", api.download)
	router.DELETE("/delete/:bucket/:name", api.delete)

	router.Use(corsMiddleware())

	router.Run(fmt.Sprintf(":%d", c.Rest.Port))

	return nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
