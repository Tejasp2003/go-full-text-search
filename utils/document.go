package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// Document represents a document with title, URL, abstract text, and ID.
type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// LoadDocuments loads documents from a gzipped XML file at the given path.
func LoadDocuments(filePath string) ([]Document, error) {
	// Open the gzipped file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Create a new XML decoder
	decoder := xml.NewDecoder(gzipReader)

	// Struct to hold the parsed documents
	var documentDump struct {
		Documents []Document `xml:"doc"`
	}

	// Decode the XML data
	if err := decoder.Decode(&documentDump); err != nil {
		return nil, err
	}

	// Assign IDs to each document
	for i := range documentDump.Documents {
		documentDump.Documents[i].ID = i
	}
	return documentDump.Documents, nil
}
