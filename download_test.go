package main

import (
	"encoding/csv"
	"io"
	"strings"
	"testing"
)

func TestQuoteReplacer(t *testing.T) {
	input := "foo\"bar\"baz"
	expected := "foo'bar'baz"

	replacer := &quoteReplacer{reader: strings.NewReader(input)}
	bytes, err := io.ReadAll(replacer)
	actual := string(bytes)
	if err != nil {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestQuoteReplacerWithTSV(t *testing.T) {
	input := "11845		True	6635	False	0	Zwemke G  	Citizens For Gail D Zwemke	1351 Davey Dr			Batavia	IL	60510	F	1998-01-27 00:00:00	1996-01-06 00:00:00	197.84	False	False	True	False		S	 		To Support The Candidacy Of Gail D Zwemke For Office Of Circuit Court Judge In The 16th Judicial Circuit - Additional Judgeship \"a\"."
	expectedLastField := "To Support The Candidacy Of Gail D Zwemke For Office Of Circuit Court Judge In The 16th Judicial Circuit - Additional Judgeship 'a'."

	replacer := &quoteReplacer{reader: strings.NewReader(input)}
	tsvReader := csv.NewReader(replacer)
	tsvReader.Comma = '\t'

	rec, err := tsvReader.Read()
	if err != nil {
		t.Errorf("Error reading record: %v", err)
	}

	if len(rec) != 27 {
		t.Errorf("Expected 27 fields, got %d", len(rec))
	}

	if rec[26] != expectedLastField {
		t.Errorf("Expected last field to be %s, got %s", expectedLastField, rec[26])
	}
}
