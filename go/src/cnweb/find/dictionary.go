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
	"time"
)

var (
	findEnglishStmt *sql.Stmt
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

func initQueries() error {
	conString := webconfig.DBConfig()
	db, err := sql.Open("mysql", conString)
	if err != nil {
		return err
	}
	database = db

	ctx := context.Background()
	fwstmt, err := database.PrepareContext(ctx, 
`SELECT simplified, traditional, pinyin, english, notes, headword
FROM words
WHERE pinyin = ? OR english LIKE ?
LIMIT 20`)
    if err != nil {
        applog.Error("find.init() Error preparing fwstmt: ", err)
        return err
    }
    findEnglishStmt = fwstmt

    return nil
}

// Returns the word senses with English approximate or Pinyin exact match
func findWordsByEnglish(query string) ([]WordSense, error) {
	if findEnglishStmt == nil {
		initQueries()
	}
	ctx := context.Background()
	results, err := findEnglishStmt.QueryContext(ctx, query, query)
	if err != nil {
		applog.Error("findWordsByEnglish, Error for query: ", query, err)
		// Sleep for a while, reinitialize, and retry
		time.Sleep(2000 * time.Millisecond)
		initQueries()
		results, err = findEnglishStmt.QueryContext(ctx, query, query)
		if err != nil {
			applog.Error("findWordsByEnglish, Give up after retry: ", query, err)
			return []WordSense{}, err
		}
	}
	senses := []WordSense{}
	for results.Next() {
		ws := WordSense{}
		var hw sql.NullInt64
		var trad, pinyin, english, notes sql.NullString
		results.Scan(&ws.Simplified, &trad, &pinyin, &english, &notes, &hw)
		applog.Info("findWordsByEnglish, simplified, headword = ",
			ws.Simplified, hw)
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
		senses = append(senses, ws)
	}
	applog.Info("findWordsByEnglish, len(senses): ", len(senses))
	return senses, nil
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
				wdict[word1.Traditional] = word1
			} else {
				word1 = Word{}
				word1.Simplified = ws.Simplified
				word1.Traditional = ws.Traditional
				word1.Pinyin = ws.Pinyin
				word1.HeadwordId = ws.HeadwordId
				word1.Senses = []WordSense{ws}
				wdict[word1.Traditional] = word1
			}
		}
	}
	return wdict, nil
}
