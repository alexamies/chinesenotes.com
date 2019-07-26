/*
 * Package for looking up words and multiword expressions.
 */
package dictionary

import (
	"cnweb/applog"
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	findSubstrStmt *sql.Stmt
)


// Encapsulates term lookup recults
type Results struct {
	Words []Word
}

// Used for grouping word senses by similar headwords in result sets
func addWordSense2Map(wmap map[string]Word, ws WordSense) {
	//applog.Info("dictionary.addWordSense2Map() ", ws.Simplified, ws.Traditional)
	word, ok := wmap[ws.Simplified]
	if ok {
		word.Senses = append(word.Senses, ws)
		wmap[word.Simplified] = word
	} else {
		word = Word{}
		word.Simplified = ws.Simplified
		word.Traditional = ws.Traditional
		word.Pinyin = ws.Pinyin
		word.HeadwordId = ws.HeadwordId
		word.Senses = []WordSense{ws}
		wmap[word.Simplified] = word
	}
}

func initSubtrQuery() error {
	ctx := context.Background()
	stmt, err := database.PrepareContext(ctx, 
`SELECT simplified, traditional, pinyin, english, notes, headword 
FROM words 
WHERE
  (simplified LIKE ? OR traditional LIKE ?)
  AND 
  topic_en = ? 
LIMIT 20`)
    if err != nil {
        applog.Error("dictionary.initSubtrQuery() Error preparing fwstmt: ", err)
        return err
    }
    findSubstrStmt = stmt
    return nil
}

// Lookup a term based on a substring and a topic
func LookupSubstr(query, topic_en string) (*Results, error) {
	if query == "" {
		return nil, errors.New("Query string is empty")
	}
	applog.Info("LookupSubstr, query, topic = ", query, topic_en)
	if findSubstrStmt == nil {
		applog.Error("LookupSubstr, findSubstr == nil")
		// Re-initialize
		initDictionary()
		if findSubstrStmt == nil {
			applog.Error("LookupSubstr, still findSubstr == nil")
		  return &Results{[]Word{}}, errors.New("Unable to look up term")
		}
	}
	ctx := context.Background()
	likeTerm := "%" + query + "%"
	results, err := findSubstrStmt.QueryContext(ctx, likeTerm, likeTerm, topic_en)
	if err != nil {
		applog.Error("LookupSubstr, Error for query: ", query, err)
		// Re-initialize the app
		initDictionary()
		results, err = findSubstrStmt.QueryContext(ctx, likeTerm, likeTerm, topic_en)
		if err != nil {
			applog.Error("LookupSubstr, Give up after retry: ", query, err)
			return &Results{[]Word{}}, err
		}
	}
	wmap := map[string]Word{}
	for results.Next() {
		ws := WordSense{}
		var hw sql.NullInt64
		var trad, pinyin, english, notes sql.NullString
		results.Scan(&ws.Simplified, &trad, &pinyin, &english, &notes, &hw)
		//applog.Info("LookupSubstr, simplified, headword = ", ws.Simplified, hw)
		if trad.Valid {
			ws.Traditional = trad.String
		}
		if pinyin.Valid {
			ws.Pinyin = pinyin.String
		}
		if english.Valid {
			ws.English = english.String
		}
		if notes.Valid {
			ws.Notes = notes.String
		}
		if hw.Valid {
			ws.HeadwordId = int(hw.Int64)
		}
		addWordSense2Map(wmap, ws)
	}
	applog.Info("LookupSubstr, len(wmap): ", len(wmap))
	words := wordMap2Array(wmap)
	return &Results{words}, nil
}

func wordMap2Array(wmap map[string]Word) []Word {
	words := []Word{}
	for _, w := range wmap {
		words = append(words, w)
	}
	return words
}
