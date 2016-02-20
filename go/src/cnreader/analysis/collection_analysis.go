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

func (results *CollectionAResults) AddResults(vocab1 map[string]int, 
		usage1 map[string]string, wc1 int, unknown1 map[string]int) {

	for k, v := range vocab1 {
    	results.Vocab[k] += v
	}
	for k, v := range usage1 {
    	results.Usage[k] = v
	}
	results.WC += wc1
	for k, v := range unknown1 {
    	results.UnknownChars[k] += v
	}
}
