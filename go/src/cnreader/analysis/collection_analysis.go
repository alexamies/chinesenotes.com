/*
CollectionAResults type for vocabulary analysis of a collection of texts
*/
package analysis

import (
	//"cnreader/alignment"
	"cnreader/dictionary"
	"cnreader/index"
	"cnreader/ngram"
)

// A struct to hold the analysis results for the collection
type CollectionAResults struct {
	Vocab				map[string]int
	Bigrams				map[string]int
	Usage				map[string]string
	BigramFrequencies	ngram.BigramFreqMap
	Collocations		ngram.CollocationMap
	WC, CCount			int
	UnknownChars		map[string]int
	WFDocMap			index.TermFreqDocMap
	BigramDocMap		index.TermFreqDocMap
	DocFreq				index.DocumentFrequency
	BigramDF			index.DocumentFrequency
	DocLengthArray		[]index.DocLength
}

// Add more results to this set of results
func (results *CollectionAResults) AddResults(more CollectionAResults) {

	for k, v := range more.Vocab {
		results.Vocab[k] += v
	}

	for k, v := range more.Bigrams {
		results.Bigrams[k] += v
	}

	for k, v := range more.Usage {
		results.Usage[k] = v
	}

	results.BigramFrequencies.Merge(more.BigramFrequencies)

	results.Collocations.MergeCollocationMap(more.Collocations)

	results.WC += more.WC
	results.CCount += more.CCount

	for k, v := range more.UnknownChars {
		results.UnknownChars[k] += v
	}
	for _, dl := range more.DocLengthArray {
		results.DocLengthArray = append(results.DocLengthArray, dl)
	}
}

// Returns the subset of words that are lexical (content) words
func (results *CollectionAResults) GetLexicalWordFreq(sortedWords []index.SortedWordItem) []WFResult {

	wfResults := make([]WFResult, 0)
	for _, value := range sortedWords {
		ws, _ := dictionary.GetWordSense(value.Word)
		if !ws.IsFunctionWord() {
			wfResults = append(wfResults, WFResult{
				Freq:       value.Freq,
				HeadwordId: ws.HeadwordId,
				Chinese:    value.Word,
				Pinyin:     ws.Pinyin,
				English:    ws.English,
				Usage:      results.Usage[value.Word]})
		}
	}
	return wfResults
}

// Returns the subset of words that are lexical (content) words
func (results *CollectionAResults) GetHeadwords() []dictionary.HeadwordDef {
	headwords := make([]dictionary.HeadwordDef, 0, len(results.Vocab))
	for k := range results.Vocab {
		hw, _ := dictionary.GetHeadword(k)
		headwords = append(headwords, hw)
	}
	return headwords
}

// Returns the subset of words that are lexical (content) words
func (results *CollectionAResults) GetWordFreq(sortedWords []index.SortedWordItem) []WFResult {

	wfResults := make([]WFResult, 0)
	maxWFOutput := len(sortedWords)
	if maxWFOutput > MAX_WF_OUTPUT {
		maxWFOutput = MAX_WF_OUTPUT
	}
	for _, value := range sortedWords[:maxWFOutput] {
		ws, _ := dictionary.GetWordSense(value.Word)
		wfResults = append(wfResults, WFResult{
			Freq:       value.Freq,
			HeadwordId: ws.HeadwordId,
			Chinese:    value.Word,
			Pinyin:     ws.Pinyin,
			English:    ws.English,
			Usage:      results.Usage[value.Word]})
	}
	return wfResults
}

// Constructor for empty CollectionAResults
func NewCollectionAResults() CollectionAResults {
	return CollectionAResults{
		Vocab:				map[string]int{},
		Bigrams:				map[string]int{},
		Usage:				map[string]string{},
		BigramFrequencies:	ngram.BigramFreqMap{},
		Collocations:		ngram.CollocationMap{},
		WC:					0,
		UnknownChars:		map[string]int{},
		WFDocMap:			index.TermFreqDocMap{},
		BigramDocMap:		index.TermFreqDocMap{},
		DocFreq:			index.NewDocumentFrequency(),
		BigramDF:			index.NewDocumentFrequency(),
		DocLengthArray:		[]index.DocLength{},
	}
}

func setDocLength(wfDocMap index.TermFreqDocMap, wc int) {

}
