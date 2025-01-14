package main

import (
	"io"
	"net/http"
	"os"
)

type quoteReplacer struct {
	reader io.Reader
}

func (r *quoteReplacer) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if err != nil {
		return n, err
	}
	for i, b := range p {
		if b == '"' {
			p[i] = '\''
		}
	}
	return n, nil
}

func downloadFile(filepath string, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Go's CSV reader, when in TSV mode, can't handle double quotes inside
	// of fields. To get around this, we replace all double quotes with
	// single quotes when downloading the file.
	quoteReplacer := &quoteReplacer{reader: resp.Body}

	_, err = io.Copy(out, quoteReplacer)
	if err != nil {
		return err
	}
	return nil
}
