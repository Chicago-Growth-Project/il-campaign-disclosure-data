package main

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
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

	err := t.createTable(db)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	count, err := t.countRows(db)
	if err != nil {
		return fmt.Errorf("failed to count rows: %w", err)
	}

	if count > 0 {
		fmt.Printf("Table %s found to have %d rows. Skipping import", t.Name, count)
		return nil
	}

	err = downloadFile(t.tempFilename(), t.URL)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	err = t.convertToUTF8(t.tempFilename(), t.tempFilename())

	if t.FileType == TSV {
		err = t.convertFile(t.tempFilename())
		if err != nil {
			return fmt.Errorf("failed to convert file: %w", err)
		}
	} else if t.FileType == CSV {
		// Create a newfile for compatiblity for now
		data, err := os.ReadFile(t.tempFilename())

		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		err = os.WriteFile(t.newFilename(), data, 0644)

		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	} else {
		return fmt.Errorf("Failed to have a parser available: %w", err)
	}

	err = t.loadFile(t.newFilename(), db)
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
			var cleanedRec []string
			for _, field := range rec {
				cleanedRec = append(cleanedRec, strings.TrimSpace(field))
			}
			csvWriter.Write(cleanedRec)
		}
	}

	csvWriter.Flush()
	return nil
}

// convertToUTF8 attempts to convert the content of inputFile to UTF-8
// and writes it to outputFile.
// It tries decoding as UTF-8 first, then as Windows-1252 if UTF-8 fails.
func (t *Table) convertToUTF8(inputFile, outputFile string) error {
	// Read the raw bytes from the input file
	rawBytes, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file %s: %w", inputFile, err)
	}

	var content string

	// 1. Try to interpret as UTF-8 directly
	if utf8.Valid(rawBytes) {
		fmt.Printf("File %s is already valid UTF-8 or ASCII.\n", inputFile)
		content = string(rawBytes)
	} else {
		fmt.Printf("File %s is not valid UTF-8. Attempting to decode as Windows-1252...\n", inputFile)
		// 2. Try to decode as Windows-1252 (Cp1252)
		decoder := charmap.Windows1252.NewDecoder()
		utf8Bytes, _, err := transform.Bytes(decoder, rawBytes)
		if err == nil {
			content = string(utf8Bytes)
			fmt.Println("Successfully decoded as Windows-1252 and converted to UTF-8.")
		} else {
			// 3. If Windows-1252 fails, you could try other encodings here
			// or return an error indicating unknown encoding.
			// For this example, we'll assume it might be ISO-8859-1 if not Windows-1252
			// Note: Windows-1252 is a superset of ISO-8859-1 for many characters,
			// but they differ in the 0x80-0x9F range.
			fmt.Printf("Failed to decode as Windows-1252. Attempting ISO-8859-1 (Latin1)... Error: %v\n", err)
			isoDecoder := charmap.ISO8859_1.NewDecoder()
			utf8BytesISO, _, errISO := transform.Bytes(isoDecoder, rawBytes)
			if errISO == nil {
				content = string(utf8BytesISO)
				fmt.Println("Successfully decoded as ISO-8859-1 and converted to UTF-8.")
			} else {
				return fmt.Errorf("failed to decode as Windows-1252 and ISO-8859-1. Unknown encoding or corrupt data. Last error: %w", errISO)
			}
		}
	}

	// Optional: Remove BOM (Byte Order Mark) if present, though it's less common for UTF-8
	// unless it came from a Windows source.
	// UTF-8 BOM is EF BB BF
	if strings.HasPrefix(content, "\xef\xbb\xbf") {
		content = strings.TrimPrefix(content, "\xef\xbb\xbf")
		fmt.Println("Removed UTF-8 BOM.")
	}

	// Write the UTF-8 content to the output file
	err = os.WriteFile(outputFile, []byte(content), 0666)
	if err != nil {
		return fmt.Errorf("failed to write output file %s: %w", outputFile, err)
	}

	fmt.Printf("Successfully converted %s to UTF-8 and saved as %s\n", inputFile, outputFile)
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
	query := fmt.Sprintf("COPY %s FROM '%s' (AUTO_DETECT TRUE);", t.Name, filepath)
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
