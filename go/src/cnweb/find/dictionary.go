/* 
Functions to load Chinese-English dictionary from database
*/
package find

import (
	"context"
	"cnweb/applog"
	"database/sql"
	"encoding/csv"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"time"

	"cnweb/webconfig"
)

const DICT_FILE string = "data/words.txt"

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

func initEnglishQuery() error {
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
        applog.Error("find.initEnglishQuery() Error preparing fwstmt: ", err)
        return err
    }
    findEnglishStmt = fwstmt

    return nil
}

// Returns the word senses with English approximate or Pinyin exact match
func findWordsByEnglish(query string) ([]WordSense, error) {
	applog.Info("findWordsByEnglish, query = ", query)
	if findEnglishStmt == nil {
		applog.Error("findWordsByEnglish, findEnglishStmt == nil")
		// Re-initialize
		initFind()
		if simBM251Stmt == nil {
			applog.Error("findBodyBM25, still simBM251Stmt == nil")
		  return []WordSense{}, nil
		}
	}
	ctx := context.Background()
	likeEnglish := "%" + query + "%"
	results, err := findEnglishStmt.QueryContext(ctx, query, likeEnglish)
	if err != nil {
		applog.Error("findWordsByEnglish, Error for query: ", query, err)
		// Re-initialize the app
		initFind()
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
	start := time.Now()
	wdict := map[string]Word{}
	conString := webconfig.DBConfig()
	database, err := sql.Open("mysql", conString)
  if err != nil {
        applog.Error("find.load_dict connecting to database: ", err)
        return loadDictFile()
	}
	ctx := context.Background()
	stmt, err := database.PrepareContext(ctx, 
		"SELECT id, simplified, traditional, pinyin, english, notes, headword FROM words")
    if err != nil {
        applog.Error("find.load_dict Error preparing stmt: ", err)
        return loadDictFile()
    }
	results, err := stmt.QueryContext(ctx)
	if err != nil {
		applog.Error("find.load_dict, Error for query: ", err)
        return loadDictFile()
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
	applog.Info("LoadDict, loading time: ", time.Since(start))
	return wdict, nil
}

// Loads all words from a static file included in the Docker image
func loadDictFile() (map[string]Word, error) {
	applog.Info("loadDictFile, enter")
	start := time.Now()
	wdict := map[string]Word{}
	wsFilenames := []string{DICT_FILE}
	cnReaderHome := webconfig.GetCnReaderHome()
	for _, wsfilename := range wsFilenames {
		fName := cnReaderHome + "/" + wsfilename
		applog.Info("dictionary.loadDictFile: fName: ", fName)
		wsfile, err := os.Open(fName)
		if err != nil {
			applog.Error("dictionary.loadDictFile, error: ", err)
			return wdict, err
		}
		defer wsfile.Close()
		reader := csv.NewReader(wsfile)
		reader.FieldsPerRecord = -1
		reader.Comma = rune('\t')
		reader.Comment = '#'
		rawCSVdata, err := reader.ReadAll()
		if err != nil {
			applog.Error("Could not parse lexical units file", err)
			return wdict, err
		}
		for i, row := range rawCSVdata {
			id, err := strconv.ParseInt(row[0], 10, 0)
			if err != nil {
				applog.Error("Could not parse word id for word ", i, err)
				return wdict, err
			}
			simp := row[1]
			trad := row[2]
			pinyin := row[3]
			english := row[4]
			grammar := row[5]
			notes := row[14]
			if notes == "\\N" {
				notes = ""
			}
			hwId := 0
			if len(row) == 16 {
				hwIdInt, err := strconv.ParseInt(row[15], 10, 0)
				if err != nil {
					applog.Info("loadDictFile, id: %d, simp: %s, trad: %s, " + 
						"pinyin: %s, english: %s, grammar: %s\n",
						id, simp, trad, pinyin, english, grammar,)
					applog.Error("loadDictFile: Could not parse headword id for word ",
						id, err)
				}
				hwId = int(hwIdInt)
			} else {
				applog.Info("loadDictFile, No. cols: %d\n",len(row))
				applog.Info("loadDictFile, id: %d, simp: %s, trad: %s, pinyin: %s, " +
					"english: %s, grammar: %s\n",
					id, simp, trad, pinyin, english, grammar)
				applog.Error("loadDictFile wrong number of columns ", id, err)
			}
			ws := WordSense{}
			ws.Id = hwId
			ws.Simplified =simp
			ws.HeadwordId = hwId
			ws.Traditional = trad
			ws.Pinyin = pinyin
			ws.English = english
			ws.Notes = notes
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
			if trad != "\\N" {
				word1, ok1 := wdict[trad]
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
	}
	applog.Info("loadDictFile, loading time: ", time.Since(start))
	return wdict, nil
}
