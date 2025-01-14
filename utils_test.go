package main

import (
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
