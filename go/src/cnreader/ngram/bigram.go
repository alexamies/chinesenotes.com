/*
Library for bigram type
*/
package ngram

import (
	"cnreader/dictionary"
	"fmt"
)

// A struct to hold an instance of a Bigram
// Since they could be either simplified or traditional, index by the headword
//ids. Also, include an example of the bigram so that usage context can be
// investigated
type Bigram struct {
	HeadwordDef1 *dictionary.HeadwordDef  // First headword
	HeadwordDef2 *dictionary.HeadwordDef  // Second headword
	Example, ExFile, ExDocTitle, ExColTitle string
}

// Bigrams that contain function words should be excluded
func (bigram *Bigram) ContainsFunctionWord() bool {
	ws1 := (*bigram.HeadwordDef1.WordSenses)[0]
	ws2 := (*bigram.HeadwordDef2.WordSenses)[0]
	return ws1.IsFunctionWord() || ws2.IsFunctionWord()
}

// The simplified text of the bigram
func (bigram *Bigram) Simplified() string {
	return fmt.Sprintf("%s%s", bigram.HeadwordDef1.Simplified,
		bigram.HeadwordDef2.Simplified)
}

// Override string method for comparison
func (bigram *Bigram) String() string {
	return fmt.Sprintf("%.7d %.7d", bigram.HeadwordDef1.Id,
		bigram.HeadwordDef2.Id)
}

// The traditional text of the bigram
func (bigram *Bigram) Traditional() string {
	t1 := bigram.HeadwordDef1.Traditional
	if t1 == "\\N" {
		t1 = bigram.HeadwordDef1.Simplified
	}
	t2 := bigram.HeadwordDef2.Traditional
	if t2 == "\\N" {
		t2 = bigram.HeadwordDef2.Simplified
	}
	return fmt.Sprintf("%s%s", t1, t2)
}
