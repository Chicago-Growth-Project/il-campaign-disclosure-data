package main

import (
	"testing"
)

func TestCreateTableSQL(t *testing.T) {
	const expected = "CREATE TABLE IF NOT EXISTS candidates (id INTEGER NOT NULL, last_name STRING, first_name STRING, address1 STRING, address2 STRING, city STRING, state STRING, zip STRING, office STRING, district_type STRING, district STRING, residence_county STRING, party_affiliation STRING, redaction_requested BOOLEAN)"
	actual := Candidates.createTableSQL()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestIndexColumnSQLs(t *testing.T) {
	expectations := []string{
		"CREATE INDEX IF NOT EXISTS candidates_id ON candidates (id)",
		"CREATE INDEX IF NOT EXISTS candidates_last_name ON candidates (last_name)",
		"CREATE INDEX IF NOT EXISTS candidates_first_name ON candidates (first_name)",
	}
	actual := Candidates.indexColumnSQLs()
	for i, expected := range expectations {
		if expected != actual[i] {
			t.Errorf("Expected %s, got %s", expected, actual[i])
		}
	}
}
