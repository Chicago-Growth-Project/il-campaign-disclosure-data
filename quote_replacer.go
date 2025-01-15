package main

import "io"

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
