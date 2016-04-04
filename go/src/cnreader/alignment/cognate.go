/*
Library for tracking occurence of cognates in a multipart text. The kinds of
cognates tracked are proper nouns and numeric expressions.
*/
package alignment

import (
	"cnreader/corpus"
	"cnreader/dictionary"
)

// A slice of first occurences of proper nouns in a single document that is part
// of a larger collection
type CorpEntryCognates struct {
	CorpEntry corpus.CorpusEntry
	properNouns []dictionary.WordSenseEntry
	pnMap map[string]bool
	numericExp []dictionary.WordSenseEntry
	numExpMap map[string]bool
}

// Constructor
func NewCorpEntryCognates(corpusEntry corpus.CorpusEntry) CorpEntryCognates {
	properNouns := []dictionary.WordSenseEntry{}
	pnMap := map[string]bool{}
	numericExp := []dictionary.WordSenseEntry{}
	numExpMap := map[string]bool{}
	return CorpEntryCognates{
		CorpEntry: corpusEntry, 
		properNouns: properNouns,
		pnMap: pnMap,
		numericExp: numericExp,
		numExpMap: numExpMap,
	}
}

func (corpEntryCogs *CorpEntryCognates) AddCognate(ws *dictionary.WordSenseEntry) {
	if ws.IsProperNoun() && !corpEntryCogs.pnMap[ws.Simplified] {
		corpEntryCogs.pnMap[ws.Simplified] = true
		corpEntryCogs.properNouns = append(corpEntryCogs.properNouns, *ws)
	} else if ws.IsNumericExpression() && !corpEntryCogs.numExpMap[ws.Simplified] {
		corpEntryCogs.numExpMap[ws.Simplified] = true
		corpEntryCogs.numericExp = append(corpEntryCogs.numericExp, *ws)
	}
}

// Returns slice of first occurences of numeric expressions in a single document.
func (corpEntryCogs *CorpEntryCognates) GetNumerExp() []dictionary.WordSenseEntry {
	return corpEntryCogs.numericExp
}

// Returns slice of first occurences of proper nouns in a single document.
func (corpEntryCogs *CorpEntryCognates) GetProperNouns() []dictionary.WordSenseEntry {
	return corpEntryCogs.properNouns
}
