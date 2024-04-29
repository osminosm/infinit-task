package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ZipFileReader struct {
	data []byte
}

func (zipFileReader ZipFileReader) GetFileContentByExtention(extentions []string) (string, error) {

	// Open the zip archive
	zipReader, err := zip.NewReader(bytes.NewReader(zipFileReader.data), int64(len(zipFileReader.data)))
	if err != nil {
		return "", err
	}

	content := ""

	for _, file := range zipReader.File {
		if !hasSuffixes(file.Name, extentions...) {
			continue
		}
		// Open the file from the zip archive
		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file in zip:", err)
			continue
		}
		defer fileReader.Close()

		// Read the content of the file
		fileContent, err := io.ReadAll(fileReader)
		if err != nil {
			fmt.Println("Error reading file in zip:", err)
			continue
		}

		content += string(fileContent)

	}

	return content, nil
}

func NewZipFileReader(url string) (ZipFileReader, error) {

	zipFileReader := ZipFileReader{}
	// Download the zip file
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error downloading zip file:", err)
		return zipFileReader, err
	}
	defer resp.Body.Close()

	// Read the content of the zip file into memory
	zipData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading zip file:", err)
		return zipFileReader, err
	}

	zipFileReader.data = zipData

	return zipFileReader, err
}

func hasSuffixes(str string, suffixes ...string) bool {
	if len(suffixes) == 0 {
		return true
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}
