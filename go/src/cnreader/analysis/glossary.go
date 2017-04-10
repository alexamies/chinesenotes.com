/*
Library for constructing a glossary for a text or collection of texts

A glossary consists of all the words that occur in the collection of texts and
that match the given domain.
*/
package analysis

import (
	"cnreader/dictionary"
	"sort"
)

// The content for a corpus entry
type Glossary struct {
	Domain string
	Words []dictionary.HeadwordDef
}

// Makes a glossary by filtering by the domain label and sorting by Chinese
// pinyin.
func MakeGlossary(domain_en string, headwords []dictionary.HeadwordDef) Glossary {
	hws := dictionary.Headwords{}
	if domain_en == "" {
		return Glossary{domain_en, hws}
	}
	for _, hw := range headwords {
		gwsArray := []dictionary.WordSenseEntry{}
		for _, ws := range *hw.WordSenses {
			if ws.Topic_en == domain_en && ws.Grammar != "proper noun" {
				gwsArray = append(gwsArray, ws)
			}
		}
		if len(gwsArray) > 0 {
			ghw := dictionary.CloneHeadword(hw)
			ghw.WordSenses = &gwsArray
			hws = append(hws, ghw)
		}
	}
	sort.Sort(hws)
	return Glossary{domain_en, hws}
}