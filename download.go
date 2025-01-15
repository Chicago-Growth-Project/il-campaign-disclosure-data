package main

import (
	"io"
	"net/http"
	"os"
)

func downloadFile(filepath string, url string) error {
	outputFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
