package main

import (
	"fmt"
	"strings"
	"unicode"

	snowballeng "github.com/kljensen/snowball/english"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

type Index map[string][]int

// lowercaseFilter returns a slice of tokens normalized to lower case.
func lowercaseFilter(tokens []string) []string {
	lowercasedTokens := make([]string, len(tokens))
	for i, token := range tokens {
		lowercasedTokens[i] = strings.ToLower(token)
	}
	return lowercasedTokens
}

// stopwordFilter returns a slice of tokens with stop words removed.
func stopwordFilter(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	filteredTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, isStopword := stopwords[token]; !isStopword {
			filteredTokens = append(filteredTokens, token)
		}
	}
	return filteredTokens
}

// stemmerFilter returns a slice of stemmed tokens.
func stemmerFilter(tokens []string) []string {
	stemmedTokens := make([]string, len(tokens))
	for i, token := range tokens {
		stemmedTokens[i] = snowballeng.Stem(token, false) // false means don't lowercase because we've already done that
	}
	return stemmedTokens
}

// tokenize splits the text into tokens based on non-letter and non-number characters.
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// analyze processes the text through a series of filters: tokenizing, lowercasing, removing stop words, and stemming.
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

func (index Index) Add(documents []Document) {
	for _, document := range documents {
		for _, token := range analyze(document.Text) {
			docIDs := index[token]
			// fmt.Println("For token: ", token, "docIDs: ", docIDs)
			if docIDs != nil && docIDs[len(docIDs)-1] == document.ID {
				// Don't add the same ID twice.
				continue
			}
			index[token] = append(docIDs, document.ID)
			// fmt.Println("Index: ", index)
		}
	}
}

func Intersection(a []int, b []int) []int {

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	intersectionResult := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			intersectionResult = append(intersectionResult, a[i])
			i++
			j++
		}
	}
	return intersectionResult
}

// Search queries the Index for the given text.
func (index Index) Search(queryText string) []int {
	var result []int
	for _, token := range analyze(queryText) {
		docIDs, exists := index[token]

		if exists {
			if result == nil {
				result = docIDs
			} else {
				result = Intersection(result, docIDs)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return result
}

func main() {

	documents := []Document{
		{ID: 0, Text: "the quick brown fox jumps over the lazy dog"},
		{ID: 1, Text: "the man went to the store to get a dog"},
		{ID: 2, Text: "lion is the king of the jungle"},
		{ID: 1, Text: "dog in store"},
	}

	index := make(Index)
	index.Add(documents)
	// fmt.Println(index)

	matchedDocIDs := index.Search("the quick brown dog")
	fmt.Println(matchedDocIDs)

}
