/*
Library for bigram type
*/
package ngram

import (
	"fmt"
	"github.com/alexamies/cnreader/dictionary"
	"log"
)

var NULL_BIGRAM_PTR *Bigram

// A struct to hold an instance of a Bigram
// Since they could be either simplified or traditional, index by the headword
//ids. Also, include an example of the bigram so that usage context can be
// investigated
type Bigram struct {
	HeadwordDef1 *dictionary.HeadwordDef  // First headword
	HeadwordDef2 *dictionary.HeadwordDef  // Second headword
	Example, ExFile, ExDocTitle, ExColTitle *string
}

// Constructor for a Bigram struct
func NewBigram(hw1, hw2 dictionary.HeadwordDef,
		example, exFile, exDocTitle, exColTitle string) *Bigram {
	if hw1.WordSenses == nil || hw2.WordSenses == nil {
		msg := "bigram.NewBigram: nil reference"
		if hw1.Simplified != nil {
			msg += fmt.Sprintf(", Simplified1 = %s", *hw1.Simplified)
		}
		if hw2.Simplified != nil {
			msg += fmt.Sprintf(", Simplified2 = %s", *hw2.Simplified)
		}
		if hw1.Traditional != nil {
			msg += fmt.Sprintf(", Traditional1 = %s", *hw1.Traditional)
		}
		if hw2.Traditional != nil {
			msg += fmt.Sprintf(", Traditional2 = %s", *hw2.Traditional)
		}
		msg += fmt.Sprintf("in %s, %s", exFile, exColTitle)
		log.Printf(msg)
	}
	return &Bigram{
		HeadwordDef1: &hw1,
		HeadwordDef2: &hw2,
		Example: &example,
		ExFile: &exFile,
		ExDocTitle: &exDocTitle,
		ExColTitle: &exColTitle,
	}
}

func NullBigram() *Bigram {
	if NULL_BIGRAM_PTR == nil {
		s1 := ""
		s2 := ""
		hw1 := dictionary.HeadwordDef{
			Id: 0,
			Simplified: &s1, 
			Traditional: &s2,
			Pinyin: []string{},
			WordSenses: &[]dictionary.WordSenseEntry{},
		}
		NULL_BIGRAM_PTR = NewBigram(hw1, hw1, "", "", "", "")
	}
	return NULL_BIGRAM_PTR
}

// For comparison of bigrams
func bigramKey(id1, id2 int) string {
	return fmt.Sprintf("%.7d %.7d", id1, id2)
}

// Bigrams that contain function words should be excluded
func (bigram *Bigram) ContainsFunctionWord() bool {
	if bigram.HeadwordDef1.WordSenses == nil || bigram.HeadwordDef2.WordSenses == nil {
		msg := "bigram.ContainsFunctionWord: nil reference"
		if bigram.HeadwordDef1.Simplified != nil {
			msg += fmt.Sprintf(", Simplified1 = %s",  
				*bigram.HeadwordDef1.Simplified)
		}
		if bigram.HeadwordDef2.Simplified != nil {
			msg += fmt.Sprintf(", Simplified2 = %s",  
				*bigram.HeadwordDef1.Simplified)
		}
		log.Printf(msg)
		return false
	}
	ws1 := (*bigram.HeadwordDef1.WordSenses)[0]
	ws2 := (*bigram.HeadwordDef2.WordSenses)[0]
	return ws1.IsFunctionWord() || ws2.IsFunctionWord()
}

// The simplified text of the bigram
func (bigram *Bigram) Simplified() string {
	if bigram.HeadwordDef1.Simplified == nil || bigram.HeadwordDef2.Simplified == nil {
		msg := "bigram.Simplified nil value"
		log.Printf(msg)
		return msg
	}
	return fmt.Sprintf("%s%s", *bigram.HeadwordDef1.Simplified,
		*bigram.HeadwordDef2.Simplified)
}

// Override string method for comparison
func (bigram *Bigram) String() string {
	return bigramKey(bigram.HeadwordDef1.Id, bigram.HeadwordDef2.Id)
}

// The traditional text of the bigram
func (bigram *Bigram) Traditional() string {
	if bigram.HeadwordDef1.Traditional == nil || bigram.HeadwordDef2.Traditional == nil {
		msg := "bigram.Traditional(): nil reference"
		log.Printf(msg)
		return msg
	}
	t1 := *bigram.HeadwordDef1.Traditional
	if t1 == "\\N" {
		t1 = *bigram.HeadwordDef1.Simplified
	}
	t2 := *bigram.HeadwordDef2.Traditional
	if t2 == "\\N" {
		t2 = *bigram.HeadwordDef2.Simplified
	}
	return fmt.Sprintf("%s%s", t1, t2)
}
