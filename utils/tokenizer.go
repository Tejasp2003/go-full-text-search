package utils

import (
	"strings"
	"unicode"
	

)

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


