package main

import (
	"flag"
	"log"
	"time"

	"github.com/tejasp2003/go-full-text-search/utils"
)


func main() {

	// fmt.Print(tokenize("Testing the tokenizer function"))

	// Define flags for the dump path and search query
	var dumpPath, searchQuery string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&searchQuery, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Running Full Text Search")

	startTime := time.Now()
	// Load documents from the specified dump path
	documents, err := utils.LoadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(documents), time.Since(startTime))

	startTime = time.Now()
	// Create an inverted index
	invertedIndex := make(utils.Index)
	invertedIndex.Add(documents)
	log.Printf("Indexed %d documents in %v", len(documents), time.Since(startTime))

	startTime = time.Now()
	// Search the index for the given query
	matchedDocIDs := invertedIndex.Search(searchQuery)
	log.Printf("Search found %d documents in %v", len(matchedDocIDs), time.Since(startTime))

	// Print the matched documents
	for _, docID := range matchedDocIDs {
		document := documents[docID]
		log.Printf("%d\t%s\n", docID, document.Text)
	}
}
