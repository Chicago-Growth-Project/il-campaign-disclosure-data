package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestCreateWithWindows1252fileCSV(t *testing.T) {
	testDatabasePath := "data/test-data.db"
	tServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv; charset=windows-1252")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ID, Year, Name\n"))
		w.Write([]byte("1,2023,John Doe\n"))
		w.Write([]byte("2,2023,Jane Smith\n"))
		w.Write([]byte("3,2023,Emily Davis\n"))
		w.Write([]byte("4,       ,Franco Reyes\n"))
		w.Write([]byte("05,       ,Michael Brown\n"))
		w.Write([]byte("06,       ,Chris Johnson\n"))
		w.Write([]byte("7,\"2021\",Connor O'Neil\n"))
	}))

	defer tServer.Close()

	testTable := Table{
		Name:     "test_table",
		URL:      "http://example.com/test_file.csv",
		FileType: CSV,
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "year", RawName: "Year", Type: ColumnTypeInt, NotNullable: true},
			{Name: "name", RawName: "Name", Type: ColumnTypeString},
		},
	}

	db, err := ConnectDb(testDatabasePath)
	if err != nil {
		fmt.Println("Error connecting to test database:", err)
		return
	}
	defer db.Close()

	if err := testTable.Create(db); err != nil {
		t.Errorf("Expected table creation without error. Received error %s", err)
	}

	// Additional Tests go here

	os.Remove(testDatabasePath)
}
