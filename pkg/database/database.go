package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
)

type Database struct {
	*sql.DB
}

func NewDatabase(c *config.Config) (*Database, error) {
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	log.Println(psqlInfo)

	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{database}, nil
}

func (db *Database) AddFile(f filebrowser.File) error {
	stmt, err := db.Prepare(`insert into files(id, name, path, size, sizeAfterCompression, extension, mimeType, created, updated) values(default, $1, $2, $3, $4, $5, $6, $7, $8`)
	if err != nil {
		return err
	}

	err = s
	return nil
}

func (d *Database) DeleteFile() error {
	return nil
}

func (d *Database) SearchFile() {}

func (d *Database) GetFiles() {}
