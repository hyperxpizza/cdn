package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
	_ "github.com/lib/pq"
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

func (db *Database) InsertBucket(name string) (int, error) {
	var id int

	stmt, err := db.Prepare(`insert into buckets(id, name, created, updated) values(default, $1, $2, $3`)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRow(name, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *Database) GetBucketByName(name string) (*filebrowser.Bucket, error) {
	var bucket filebrowser.Bucket

	err := db.QueryRow(`select * from buckets where name=$1`).Scan(
		&bucket.ID,
		&bucket.Name,
		&bucket.Created,
		&bucket.Updated,
	)
	if err != nil {
		return nil, err
	}

	return &bucket, nil
}

func (db *Database) UpdateBucketTime(name string) error {
	stmt, err := db.Prepare(`update buckets set updated = $1 where name = $2`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now(), name)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) DeleteBucket(name string) error {
	stmt, err := db.Prepare(`delete from buckets where name = $1`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) CheckIfFileExists(name, bucket string) error {
	var id int
	err := db.QueryRow(`select id from files where name = $1 and bucket = $2`).Scan(&id)
	if err != nil {
		return err
	}

	return nil
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

	_, err = tx.ExecContext(ctx, `update files set files_token=to_tsvector($1) where id = $2`, f.Name, id)
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

func (db *Database) DeleteFile(name, bucket string) error {
	stmt, err := db.Prepare(`delete from files where name = $1 and bucket = $2`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, bucket)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) SearchFiles(query string) ([]*filebrowser.File, error) {
	var files []*filebrowser.File

	rows, err := d.Query(`select id, name, bucket, size, sizeAfterCompression, extension, mimeType, created, updated from files where files_token @@ to_tsquery($1)`, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var file filebrowser.File
		err := rows.Scan(
			&file.ID,
			&file.Name,
			&file.BucketID,
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
