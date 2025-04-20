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

	return db, nil
}
