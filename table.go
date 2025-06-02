package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type ColumnType string
type FileType string

const (
	ColumnTypeString   ColumnType = "STRING"
	ColumnTypeInt      ColumnType = "INTEGER"
	ColumnTypeBigInt   ColumnType = "BIGINT"
	ColumnTypeDecimal  ColumnType = "DECIMAL(12,2)"
	ColumnTypeDouble   ColumnType = "DOUBLE"
	ColumnTypeBool     ColumnType = "BOOLEAN"
	ColumnTypeGeometry ColumnType = "GEOMETRY"
	CSV                FileType   = "CSV"
	TSV                FileType   = "TSV"
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
	FileType       FileType
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

	downloadOnly := os.Getenv("DOWNLOAD_ONLY") == "1"

	err := t.createTable(db)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	count, err := t.countRows(db)
	if err != nil {
		return fmt.Errorf("failed to count rows: %w", err)
	}

	if count > 0 && !downloadOnly {
		fmt.Printf("Table %s found to have %d rows. Skipping import", t.Name, count)
		return nil
	}

	/*
		fmt.Println("Downloading file from URL:", t.URL)
		err = downloadFile(t.tempFilename(), t.URL)
		if err != nil {
			return fmt.Errorf("failed to download file: %w", err)
		}
		fmt.Println("Done Downloading file")
	*/
	if downloadOnly {
		return nil
	}

	fmt.Println("Convert File to UTF-8")

	err = t.convertFileToUTF8(t.tempFilename())
	if err != nil {
		return fmt.Errorf("failed to convert file to UTF-8: %w", err)
	}

	fmt.Println("Clean File (trim)", t.tempFilename())
	err = t.cleanFile(t.tempFilename())
	if err != nil {
		return fmt.Errorf("failed to clean file: %w", err)
	}

	err = t.loadFile(t.tempFilename(), db)
	if err != nil {
		return fmt.Errorf("failed to load file: %w", err)
	}

	count, err = t.countRows(db)
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

	fmt.Println("Removed temporary files:", t.tempFilename(), t.newFilename())
	fmt.Println("Done with table:", t.Name)

	return nil
}

func (t *Table) convertFileToUTF8(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file for conversion: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file for conversion: %w", err)
	}

	// Convert to UTF-8
	utf8Data := []byte(strings.ReplaceAll(string(data), "\"", "'"))

	err = os.WriteFile(filepath, utf8Data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write converted file: %w", err)
	}

	return nil
}

func (t *Table) cleanFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file for cleaning: %w", err)
	}
	defer file.Close()

	var cleanedRecords [][]string

	// Go's CSV reader, when in TSV mode, can't handle double quotes inside
	// of fields. To get around this, we replace all double quotes with
	// single quotes when downloading the file.
	quoteReplacer := &quoteReplacer{file}
	csvReader := csv.NewReader(quoteReplacer)

	if t.FileType == TSV {
		csvReader.Comma = '\t'
	} else if t.FileType == CSV {
		csvReader.Comma = ','
	}

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("failed to read file for cleaning: %w", err)
		}

		var cleanedRec []string
		for _, field := range rec {
			cleanedRec = append(cleanedRec, strings.TrimSpace(field))
		}

		cleanedRecords = append(cleanedRecords, cleanedRec)
	}

	rewriteFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create cleaned file: %w", err)
	}
	defer rewriteFile.Close()

	csvWriter := csv.NewWriter(rewriteFile)
	defer csvWriter.Flush()

	if err = csvWriter.WriteAll(cleanedRecords); err != nil {
		return fmt.Errorf("failed to write cleaned data: %w", err)
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

func (t *Table) loadFile(filepath string, db *sql.DB) error {
	delim := ","
	if t.FileType == TSV {
		delim = "\\t"
	}
	query := fmt.Sprintf("COPY %s FROM '%s' (AUTO_DETECT TRUE, STORE_REJECTS TRUE, DELIM '%s');", t.Name, filepath, delim)
	//query := "COPY filed_docs FROM 'filed_docs.csv' (AUTO_DETECT TRUE, STORE_REJECTS TRUE, DELIM '\\t');"

	fmt.Println("QUERY:", query)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to load file %s into table %s: %w", filepath, t.Name, err)
	}
	return nil
}

func (t *Table) countRows(db *sql.DB) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", t.Name)
	row := db.QueryRow(query)
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
