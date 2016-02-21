/*
Library for vocabulary analysis of a collection of texts
*/
package analysis

// A struct to hold the analysis results for the collection
type CollectionAResults struct {
	Vocab map[string]int
	Usage map[string]string
	WC int
	UnknownChars map[string]int
}

func (results *CollectionAResults) AddResults(more CollectionAResults) {

	for k, v := range more.Vocab {
    	results.Vocab[k] += v
	}
	for k, v := range more.Usage {
    	results.Usage[k] = v
	}
	results.WC += more.WC
	for k, v := range more.UnknownChars {
    	results.UnknownChars[k] += v
	}
}
