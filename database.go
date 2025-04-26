package main

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("duckdb", DatabasePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSTALL spatial")

	if err != nil {
		return nil, err
	}

	_, err = db.Exec("LOAD spatial")

	if err != nil {
		return nil, err
	}

	return db, nil
}
