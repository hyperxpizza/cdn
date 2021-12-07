package database

import "database/sql"

type Database struct {
	*sql.DB
}

func (d *Database) AddFile() error {
	return nil
}

func (d *Database) DeleteFile() error {
	return nil
}

func (d *Database) SearchFile() {}

func (d *Database) GetFiles() {}
