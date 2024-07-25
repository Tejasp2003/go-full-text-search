package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

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
