/* 
Functions to load Chinese-English dictionary from database
*/
package find

import (
	"context"
	"cnweb/applog"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"cnweb/webconfig"
)

// A top level word structure that may include multiple word senses
type Word struct {
	Simplified, Traditional, Pinyin string
	HeadwordId int
	Senses []WordSense
}

// Defines a single sense of a Chinese word
type WordSense struct {
	Id, HeadwordId int
	Simplified, Traditional, Pinyin, English, Notes string
}

// Loads all words from the database
func LoadDict() (map[string]Word, error) {
	wdict := map[string]Word{}
	conString := webconfig.DBConfig()
	database, err := sql.Open("mysql", conString)
    if err != nil {
        applog.Error("find.load_dict connecting to database: ", err)
        return wdict, err
	}
	ctx := context.Background()
	stmt, err := database.PrepareContext(ctx, 
		"SELECT id, simplified, traditional, pinyin, english, notes, headword FROM words")
    if err != nil {
        applog.Error("find.load_dict Error preparing stmt: ", err)
        return wdict, err
    }
	results, err := stmt.QueryContext(ctx)
	if err != nil {
		applog.Error("find.load_dict, Error for query: ", err)
        return wdict, err
	}
	for results.Next() {
		ws := WordSense{}
		var wsId, hw sql.NullInt64
		var trad, notes, pinyin, english sql.NullString
		results.Scan(&wsId, &ws.Simplified, &trad, &pinyin, &english, &notes,
			&hw)
		if wsId.Valid {
			ws.Id = int(wsId.Int64)
		}
		if hw.Valid {
			ws.HeadwordId = int(hw.Int64)
		}
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
		word, ok := wdict[ws.Simplified]
		if ok {
			word.Senses = append(word.Senses, ws)
			wdict[word.Simplified] = word
		} else {
			word = Word{}
			word.Simplified = ws.Simplified
			word.Traditional = ws.Traditional
			word.Pinyin = ws.Pinyin
			word.HeadwordId = ws.HeadwordId
			word.Senses = []WordSense{ws}
			wdict[word.Simplified] = word
		}
		if trad.Valid {
			word1, ok1 := wdict[trad.String]
			if ok1 {
				word1.Senses = append(word1.Senses, ws)
				wdict[word1.Simplified] = word1
			} else {
				word1 = Word{}
				word1.Simplified = ws.Simplified
				word1.Traditional = ws.Traditional
				word1.HeadwordId = ws.HeadwordId
				word1.Senses = []WordSense{ws}
				wdict[word.Traditional] = word1
			}
		}
	}
	return wdict, nil
}
