package utils

// Index is an inverted index. It maps tokens to document IDs.
type Index map[string][]int

// Add adds documents to the Index.
// index Index is a map from tokens to document IDs. In functional programming, this is called an inverted index.
// documents []Document is a slice of documents to add to the index.

func (index Index) Add(documents []Document) {
	for _, document := range documents {
		for _, token := range analyze(document.Text) {

			docIDs := index[token]

			if docIDs != nil && docIDs[len(docIDs)-1] == document.ID {
				// Don't add the same ID twice.
				continue
			}
			index[token] = append(docIDs, document.ID)
		}
	}
}

// Intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func Intersection(a []int, b []int) []int {
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
		if docIDs, exists := index[token]; exists {
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
