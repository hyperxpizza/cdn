package database

import (
	"context"
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

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	var id int
	err = tx.QueryRowContext(ctx, `insert into files(id, name, bucket, size, sizeAfterCompression, extension, mimeType, created, updated) values(default, $1, $2, $3, $4, $5, $6, $7, $8) returning id`).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `update files set files_token=to_tsvector($1) where id =$2`, f.Name, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) DeleteFile(id int) error {
	stmt, err := db.Prepare(`delete from files where id = $1`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) SearchFile(name string) ([]*filebrowser.File, error) {
	var files []*filebrowser.File

	rows, err := d.Query(`select id, name, bucket, size, sizeAfterCompression, extension, mimeType, created, updated from files where files_token @@ to_tsquery($1)`, name)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file filebrowser.File
		err := rows.Scan(
			&file.ID,
			&file.Name,
			&file.Bucket,
			&file.Size,
			&file.SizeAfterCompression,
			&file.Extension,
			&file.MimeType,
			&file.Created,
			&file.Updated,
		)
		if err != nil {
			return nil, err
		}

		files = append(files, &file)
	}

	return files, nil
}
