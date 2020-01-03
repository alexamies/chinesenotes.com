/*
Library for Chinese vocabulary analysis
*/
package analysis

import (
	"github.com/alexamies/cnreader/config"
	"github.com/alexamies/cnreader/dictionary"
	"strings"
)

// Filters the list of headwords to only those in the configured domain
func ContainsByDomain(contains []dictionary.HeadwordDef) []dictionary.HeadwordDef {
	domains := config.GetVar("ContainsByDomain")
	containsBy := []dictionary.HeadwordDef{}
	for _, hw := range contains {
		for _, ws := range *hw.WordSenses {
			if strings.Contains(domains, ws.Topic_en) {
				containsBy = append(containsBy, hw)
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
		if _, ok := subtractSet[hw.Id]; ok {
			subtracted = append(subtracted, hw)
		}
	}
	return subtracted
}
