/* 
Functions for parsing a search query
*/
package find

import (
	"github.com/alexamies/cnweb/dictionary"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Parses input queries into a slice of text segments
type QueryParser interface {
	ParseQuery(query string) []TextSegment
}

type DictQueryParser struct{WDict map[string]dictionary.Word}

// A text segment contains the QueryText searched for and possibly a matching
// dictionary entry. There will only be matching dictionary entries for 
// Chinese words in the dictionary. Non-Chinese text, punctuation, and unknown
// Chinese words will have nil DictEntry values and matching values will be
// included in the Senses field.
type TextSegment struct{
	QueryText string
	DictEntry dictionary.Word
	Senses []dictionary.WordSense
}

// The method for parsing the query text in this function is based on dictionary
// lookups
func (parser DictQueryParser) ParseQuery(query string) []TextSegment {
	terms := []TextSegment{}
	if query != "" {
		seg := TextSegment{}
		seg.QueryText = query
		terms = append(terms, seg)
	}
	return parser.get_chunks(query)
}

// Segments the text string into chunks that are CJK and non-CJK or puncuation
func (parser DictQueryParser) get_chunks(text string) []TextSegment {
	chunks := []TextSegment{}
	cjk := ""
	noncjk := ""
	for _, character := range text {
		if is_cjk(character) {
			if noncjk != "" {
				seg := TextSegment{}
				seg.QueryText = noncjk
				chunks = append(chunks, seg)
				noncjk = ""
			}
			cjk += string(character)
		} else if cjk != "" {
			segments := parser.parse_chinese(cjk)
			for _, s := range segments {
				chunks = append(chunks, s)
			}
			cjk = ""
			noncjk += string(character)
		} else {
			noncjk += string(character)
		}
	}
	if cjk != "" {
		segments := parser.parse_chinese(cjk)
		for _, s := range segments {
			chunks = append(chunks, s)
		}
	}
	if noncjk != "" {
		seg := TextSegment{}
		seg.QueryText = noncjk
		chunks = append(chunks, seg)
	}
	return chunks
}

// Tests whether the symbol is a CJK character, excluding punctuation
// Only looks at the first charater in the string
func is_cjk(r rune) bool {
	return unicode.Is(unicode.Han, r) && !unicode.IsPunct(r)
}

// Segments Chinese text based on dictionary entries
func (parser DictQueryParser) parse_chinese(text string) []TextSegment {
	terms := []TextSegment{}
	characters := strings.Split(text, "")
	for i := 0; i < len(characters); i++ {
		for j := len(characters); j > 0; j-- {
			w := strings.Join(characters[i:j], "")
			if entry, ok := parser.WDict[w]; ok {
				seg := TextSegment{w, entry, []dictionary.WordSense{}}
				terms = append(terms, seg)
				i = j - 1
				j = 0
			} else if utf8.RuneCountInString(w) == 1 {
				log.Printf("parse_chinese: found unknown character %s\n", w)
				seg := TextSegment{}
				seg.QueryText = w
				terms = append(terms, seg)
				break
			}
		}
	}
	return terms
}

func toQueryTerms(terms []TextSegment) []string {
	queryTerms := []string{}
	for _, term := range terms {
		queryTerms = append(queryTerms, term.QueryText)
	}
	return queryTerms
}
