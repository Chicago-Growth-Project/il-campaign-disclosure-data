package main

import "testing"

func TestTable(t *testing.T) {
	const expected = "CREATE TABLE IF NOT EXISTS candidates (id INTEGER NOT NULL, first_name STRING, last_name STRING, address1 STRING, address2 STRING, city STRING, state STRING, zip STRING, office STRING, district_type STRING, district STRING, residence_county STRING, party_affiliation STRING, redaction_requested BOOLEAN)"
	actual := Candidates.createTableSQL()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
