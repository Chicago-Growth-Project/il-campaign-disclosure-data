package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type ColumnType string

const (
	ColumnTypeString  ColumnType = "STRING"
	ColumnTypeInt     ColumnType = "INTEGER"
	ColumnTypeDecimal ColumnType = "DECIMAL(12,2)"
	ColumnTypeBool    ColumnType = "BOOLEAN"
)

type Column struct {
	Name        string
	RawName     string
	Type        ColumnType
	NotNullable bool // Inverted so that the default value is permissive
}

type Table struct {
	Name           string
	IndexedColumns []string // TODO: Possibly handle unique indexes
	URL            string
	Columns        []Column
}

/*
Creates a new table through the following steps:
* Download the TSV from the URL. As part of this, we replace all double quotes with single quotes.
* Convert the TSV to CSV and rename the headers. We can handle bad rows at this point.
* Create the table and indexes.
* Load the CSV into the table.
* Remove the temporary files.
*/
func (t *Table) Create(db *sql.DB) error {
	err := downloadFile(t.tempFilename(), t.URL)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	err = t.convertFile(t.tempFilename())
	if err != nil {
		return fmt.Errorf("failed to convert file: %w", err)
	}

	err = t.createTable(db)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	err = t.loadFile(t.newFilename())
	if err != nil {
		return fmt.Errorf("failed to load file: %w", err)
	}

	count, err := t.countRows(db)
	if err != nil {
		return fmt.Errorf("failed to count rows: %w", err)
	}
	fmt.Printf("Loaded %d rows into %s\n", count, t.Name)

	for _, filename := range []string{t.tempFilename(), t.newFilename()} {
		err = os.Remove(filename)
		if err != nil {
			return fmt.Errorf("failed to remove file %s: %w", filename, err)
		}
	}

	return nil
}

func (t *Table) convertFile(filepath string) error {
	oldFile, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open old file: %w", err)
	}
	defer oldFile.Close()

	newFile, err := os.Create(t.newFilename())
	if err != nil {
		return fmt.Errorf("failed to create new file: %w", err)
	}
	defer newFile.Close()

	csvWriter := csv.NewWriter(newFile)
	csvWriter.Write(t.csvHeaders())

	// Go's CSV reader, when in TSV mode, can't handle double quotes inside
	// of fields. To get around this, we replace all double quotes with
	// single quotes when downloading the file.
	quoteReplacer := &quoteReplacer{oldFile}

	tsvReader := csv.NewReader(quoteReplacer)
	tsvReader.Comma = '\t'
	tsvReader.FieldsPerRecord = len(t.Columns)
	firstLine := true

	for {
		rec, err := tsvReader.Read()
		if err == io.EOF {
			break
		} else if errors.Unwrap(err) == csv.ErrFieldCount {
			fmt.Printf("Wrong number of fields on row detected. Usually this indicates an unescaped tab in the data. Skipping row:\n %v\n", rec)
		} else if err != nil {
			return fmt.Errorf("failed to read old file: %w", err)
		} else if firstLine {
			firstLine = false
		} else {
			csvWriter.Write(rec)
		}
	}

	csvWriter.Flush()
	return nil
}

func (t *Table) createTable(db *sql.DB) error {
	_, err := db.Exec(t.createTableSQL())
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", t.Name, err)
	}

	for _, indexSql := range t.indexColumnSQLs() {
		_, err = db.Exec(indexSql)
		if err != nil {
			return fmt.Errorf("failed to create index for %s: %w", t.Name, err)
		}
	}
	return nil
}

func (t *Table) loadFile(filepath string) error {
	output, err := exec.Command("sqlite3", DatabasePath, ".mode csv", fmt.Sprintf(".import %s %s", filepath, t.Name)).Output()
	if err != nil {
		fmt.Println(string(output))
		return fmt.Errorf("failed to load file %s into table %s: %w", filepath, t.Name, err)
	}
	return nil
}

func (t *Table) countRows(db *sql.DB) (int, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM ?", t.Name)
	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count rows: %w", err)
	}
	return count, err
}

func (t *Table) tempFilename() string {
	return t.Name + ".tsv"
}

func (t *Table) newFilename() string {
	return t.Name + ".csv"
}

func (t *Table) createTableSQL() string {
	sql := "CREATE TABLE IF NOT EXISTS " + t.Name + " ("
	for i, column := range t.Columns {
		sql += column.Name + " " + string(column.Type)
		if column.NotNullable {
			sql += " NOT NULL"
		}
		if i < len(t.Columns)-1 {
			sql += ", "
		}
	}
	sql += ")"
	return sql
}

func (t *Table) indexColumnSQLs() []string {
	result := []string{}
	for _, column := range t.IndexedColumns {
		result = append(result, "CREATE INDEX IF NOT EXISTS "+t.Name+"_"+column+" ON "+t.Name+" ("+column+")")
	}
	return result
}

func (t *Table) csvHeaders() []string {
	result := []string{}
	for _, column := range t.Columns {
		result = append(result, column.Name)
	}
	return result
}
