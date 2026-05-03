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

func TestCreateWWindows1252file(t *testing.T) {
	testDatabasePath := "../data/test_data.db"

	tServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv; charset=windows-1252")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ID, Year, Name\n"))
		w.Write([]byte("1,2023,John Doe\n"))
		w.Write([]byte("02,2023,Jane Smith\n"))
		w.Write([]byte("\"3\",\"2023\",Emily Davis\n"))
		w.Write([]byte("4,,Null Year\n"))
		w.Write([]byte("4,       ,Franco Reyes\n"))
		w.Write([]byte("5, 1985  ,Franco Reyes\n"))
		w.Write([]byte("9,\"2021\",Connor O'Neil\n"))
		w.Write([]byte(",\"2021\",THIS IS NOT A NAME\n"))
		w.Write([]byte("  ,\"2021\",NOT A NAME\n"))
	}))

	defer tServer.Close()

	db, err := ConnectDb(testDatabasePath)
	if err != nil {
		fmt.Println("Error connecting to test database: ", err)
		return
	}
	defer db.Close()

	testTable := Table{
		Name:     "test_table",
		URL:      tServer.URL,
		FileType: CSV,
		Columns: []Column{
			{Name: "id", RawName: "ID", Type: ColumnTypeInt, NotNullable: true},
			{Name: "year", RawName: "Year", Type: ColumnTypeInt},
			{Name: "name", RawName: "Name", Type: ColumnTypeString},
		},
	}

	yarp, _ := os.Executable()
	fmt.Println(yarp)
	if err := testTable.Create(db); err != nil {
		t.Errorf("Expected table creation without error. Received error: %s,", err)
	}

	// Additional tests go here

	os.Remove(testDatabasePath)
}
