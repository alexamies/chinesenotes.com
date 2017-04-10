/*
Definition and sorting of headwords in Pinyin alphabetical order
*/
package dictionary

import (
	"strings"
)

// Defines a headword with mapping between simplified and traditional
type HeadwordDef struct {
	Id int // key for the headword
	Simplified, Traditional *string
	Pinyin []string
	WordSenses *[]WordSenseEntry //keys for the lexical units
}

// Defines a single sense of a Chinese word
type WordSenseEntry struct {
	Id, HeadwordId int
	Simplified, Traditional, Pinyin, English, Grammar, Concept_cn,
		Concept_en, Topic_cn, Topic_en, Parent_cn, Parent_en, Image,
		Mp3, Notes string
}

// Sorted into descending order with most frequent bigram first
type Headwords []HeadwordDef


func (hwArr Headwords) Len() int {
	return len(hwArr)
}

func (hwArr Headwords) Swap(i, j int) {
	hwArr[i], hwArr[j] = hwArr[j], hwArr[i]
}

func (hwArr Headwords) Less(i, j int) bool {
	noTones1 := normalizePinyin(hwArr[i].Pinyin[0])
	noTones2 := normalizePinyin(hwArr[j].Pinyin[0])
	return noTones1 < noTones2
}

// Removes the tone diacritics from a Pinyin string
func normalizePinyin(pinyin string) string {
	runes := []rune{}
	for _, r := range pinyin {
		n, ok := NORMAL[r]
		if ok {
			runes = append(runes, n)
		} else {
			runes = append(runes, r)
		}
	}
	return strings.ToLower(string(runes))
}

var NORMAL = map[rune]rune{
	'ā': 'a',
	'á': 'a',
	'ǎ': 'a',
	'à': 'a',
	'ē': 'e',
	'é': 'e',
	'ě': 'e',
	'è': 'e',
	'ī': 'i',
	'í': 'i',
	'ǐ': 'i',
	'ì': 'i',
	'ō': 'o',
	'ó': 'o',
	'ǒ': 'o',
	'ò': 'o',
	'ū': 'u',
	'ú': 'u',
	'ǔ': 'u',
	'ù': 'u',
	'Ā': 'a',
	'Á': 'a',
	'Ǎ': 'a',
	'À': 'a',
	'Ē': 'e',
	'É': 'e',
	'Ě': 'e',
	'È': 'e',
	'Ī': 'i',
	'Í': 'i',
	'Ǐ': 'i',
	'Ì': 'i',
	'Ō': 'o',
	'Ó': 'o',
	'Ǒ': 'o',
	'Ò': 'o',
	'Ū': 'u',
	'Ú': 'u',
	'Ǔ': 'u',
	'Ù': 'u',
}
