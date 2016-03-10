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
//ids
type Bigram struct {
	HeadwordDef1 dictionary.HeadwordDef  // First headword
	HeadwordDef2 dictionary.HeadwordDef  // Second headword
}

func (bigram *Bigram) Simplified() string {
	return fmt.Sprintf("%s%s", bigram.HeadwordDef1.Simplified,
		bigram.HeadwordDef2.Simplified)
}

// Override string method for comparison
func (bigram *Bigram) String() string {
	return fmt.Sprintf("%.7d %.7d", bigram.HeadwordDef1.Id,
		bigram.HeadwordDef2.Id)
}

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
