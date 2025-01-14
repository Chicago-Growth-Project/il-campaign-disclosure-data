package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const DatabasePath = "new_elections.db"

type ColumnType string

const (
	ColumnTypeString ColumnType = "STRING"
	ColumnTypeInt    ColumnType = "INTEGER"
	ColumnTypeBool   ColumnType = "BOOLEAN"
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

func (t *Table) Create() error {
	err := downloadFile(t.tempFilename(), t.URL)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}

	err = t.convertFile(t.tempFilename())
	if err != nil {
		return fmt.Errorf("failed to convert file: %w", err)
	}

	err = t.createTable()
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	err = t.loadFile(t.newFilename())
	if err != nil {
		return fmt.Errorf("failed to load file: %w", err)
	}

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

	tsvReader := csv.NewReader(oldFile)
	tsvReader.Comma = '\t'
	tsvReader.FieldsPerRecord = len(t.Columns)
	firstLine := true

	for {
		rec, err := tsvReader.Read()
		if err == io.EOF {
			break
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

func (t *Table) createTable() error {
	cmd := exec.Command("sqlite3", DatabasePath, t.createTableSQL())
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", t.Name, err)
	}

	for _, indexSql := range t.indexColumnSQLs() {
		cmd = exec.Command("sqlite3", DatabasePath, indexSql)
		err = cmd.Run()
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
