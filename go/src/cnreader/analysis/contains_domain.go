/*
Library for Chinese vocabulary analysis
*/
package analysis

import (
	"github.com/alexamies/cnreader/config"
	"github.com/alexamies/cnreader/dictionary"
	"strings"
)

// Max number of words to display for contained in
const MAX_CONTAINED_BY = 50

// Filters the list of headwords to only those in the configured domain
func ContainsByDomain(contains []dictionary.HeadwordDef) []dictionary.HeadwordDef {
	domains := config.GetVar("ContainsByDomain")
	containsBy := []dictionary.HeadwordDef{}
	containsSet := make(map[int]bool)
	count := 0
	for _, hw := range contains {
		for _, ws := range *hw.WordSenses {
		  _, ok := containsSet[hw.Id]
			if !ok && count < MAX_CONTAINED_BY && strings.Contains(domains, ws.Topic_en) {
				containsBy = append(containsBy, hw)
				containsSet[hw.Id] = true  // Do not add it twice
				count++ // don't go over max number
			}
		}
	}
	return containsBy
}

// Subtract the items in the second list from the first
func Subtract(headwords, subtract []dictionary.HeadwordDef) []dictionary.HeadwordDef {
	subtracted := []dictionary.HeadwordDef{}
	subtractSet := make(map[int]bool)
	for _, hw := range subtract {
    subtractSet[hw.Id] = true
	}
	for _, hw := range headwords {
		if _, ok := subtractSet[hw.Id]; !ok {
			subtracted = append(subtracted, hw)
		}
	}
	return subtracted
}
