package disclosure

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb/v2"
)

const DefaultDatabasePath = "il-campaign-disclosures.db"

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
