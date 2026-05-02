package internal

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb/v2"
)

const DefaultDatabasePath = "data/il-campaign-disclosures.db"

func ConnectDb(path string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", path)
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec("INSTALL spatial"); err != nil {
		return nil, err
	}

	if _, err = db.Exec("LOAD spatial"); err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectDbReadOnly opens an existing database at path without loading extensions.
func ConnectDbReadOnly(path string) (*sql.DB, error) {
	return sql.Open("duckdb", path)
}
