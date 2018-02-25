/* 
Functions for finding collections by partial match on collection title
*/
package find

import (
	"cnweb/applog"
	"database/sql"
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"cnweb/webconfig"
)

var (
	countColStmt *sql.Stmt
	countDocStmt *sql.Stmt
	database *sql.DB
	findColStmt *sql.Stmt
	findDocStmt *sql.Stmt
	findWordStmt *sql.Stmt
)

type Collection struct {
	GlossFile, Title string
}

type Document struct {
	GlossFile, Title string
}

type QueryResults struct {
	NumCollections, NumDocuments int
	Collections []Collection
	Documents []Document
	Terms []TextSegment
}

// Open database connection and prepare statements
func init() {
	err := initStatements()
	if err != nil {
		applog.Error("find/init: error preparing database statements, retrying",
			err)
		time.Sleep(60000 * time.Millisecond)
		err = initStatements()
		conString := webconfig.DBConfig()
		applog.Fatal("find/init: error preparing database statements, giving up",
			conString, err)
	}
	result := hello() 
	if !result {
		conString := webconfig.DBConfig()
		applog.Fatal("find/init: got error with findWords ", conString, err)
	}
}

func hello() bool {
	words, err := findWords("你好")
	if err != nil {
		conString := webconfig.DBConfig()
		applog.Error("find/hello: got error with findWords ", conString, err)
		return false
	}
	if len(words) != 1 {
		applog.Error("find/hello: could not find my word ", len(words))
		return false
	} 
	applog.Info("find/hello: Ready to go")
	return true
}

func initStatements() error {
	conString := webconfig.DBConfig()
	db, err := sql.Open("mysql", conString)
	if err != nil {
		return err
	}
	database = db

	ctx := context.Background()
	stmt, err := database.PrepareContext(ctx,
		"SELECT title, gloss_file FROM collection WHERE title LIKE ? LIMIT 50")
    if err != nil {
        applog.Error("find.init() Error preparing collection stmt: ", err)
        return err
    }
    findColStmt = stmt

	cstmt, err := database.PrepareContext(ctx,
		"SELECT count(title) FROM collection WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.init() Error preparing cstmt: ", err)
        return err
    }
    countColStmt = cstmt

	dstmt, err := database.PrepareContext(ctx,
		"SELECT title, gloss_file FROM document WHERE title LIKE ? LIMIT 50")
    if err != nil {
        applog.Error("find.init() Error preparing dstmt: ", err)
        return err
    }
    findDocStmt = dstmt

	cdstmt, err := database.PrepareContext(ctx,
		"SELECT count(title) FROM document WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.init() Error preparing cDocStmt: ", err)
        return err
    }
    countDocStmt = cdstmt    

	fwstmt, err := database.PrepareContext(ctx, 
		"SELECT simplified, traditional, pinyin, headword FROM words WHERE simplified = ? OR traditional = ? LIMIT 1")
    if err != nil {
        applog.Error("find.init() Error preparing fwstmt: ", err)
        return err
    }
    findWordStmt = fwstmt
    return nil
}

func countCollections(query string) int {
	var count int
	ctx := context.Background()
	results, err := countColStmt.QueryContext(ctx, "%" + query + "%")
	results.Next()
	results.Scan(&count)
	if err != nil {
		applog.Error("countCollections: Error for query: ", query, err)
	}
	results.Close()
	return count
}

func countDocuments(query string) int {
	var count int
	ctx := context.Background()
	results, err := countDocStmt.QueryContext(ctx, "%" + query + "%")
	results.Next()
	results.Scan(&count)
	if err != nil {
		applog.Error("countDocuments: Error for query: ", query, err)
	}
	results.Close()
	return count
}

func findCollections(query string) []Collection {
	ctx := context.Background()
	results, err := findColStmt.QueryContext(ctx, "%" + query + "%")
	if err != nil {
		applog.Error("findCollections, Error for query: ", query, err)
	}
	defer results.Close()

	collections := []Collection{}
	for results.Next() {
		col := Collection{}
		results.Scan(&col.Title, &col.GlossFile)
		collections = append(collections, col)
	}
	return collections
}

func findDocuments(query string) ([]Document, error) {
	ctx := context.Background()
	results, err := findDocStmt.QueryContext(ctx, "%" + query + "%")
	if err != nil {
		applog.Error("findDocuments, Error for query: ", query, err)
		return nil, err
	}
	defer results.Close()

	documents := []Document{}
	for results.Next() {
		doc := Document{}
		results.Scan(&doc.Title, &doc.GlossFile)
		documents = append(documents, doc)
	}
	return documents, nil
}

// Returns a QueryResults object containing matching collections, documents,
// and dictionary words. For dictionary lookup, a text segment will
// contains the QueryText searched for and possibly a matching
// dictionary entry. There will only be matching dictionary entries for 
// Chinese words in the dictionary. If there are no Chinese words in the query
// then the Chinese word senses matching the English or Pinyin will be included
// in the TextSegment.Senses field.
func FindDocuments(parser QueryParser, query string) (QueryResults, error) {
	if query == "" {
		applog.Error("FindDocuments, Empty query string")
		return QueryResults{}, errors.New("Empty query string")
	}
	terms := parser.ParseQuery(query)
	if (len(terms) == 1) && (terms[0].DictEntry.HeadwordId == 0) {
	    applog.Info("FindDocuments,Query does not contain Chinese, look for " +
	    	"English and Pinyin matches: ", query)
		senses, err := findWordsByEnglish(terms[0].QueryText)
		if err != nil {
			return QueryResults{}, err
		} else {
			terms[0].Senses = senses
		}
	}
	nCol := countCollections(query)
	nDoc := countDocuments(query)
	collections := findCollections(query)
	documents, err := findDocuments(query)
	if err != nil {
		// Got an error, see if we can connect and try again
		if hello() {
			documents, err = findDocuments(query)
		} // else do not try again, giveup and return the error
	}
	applog.Info("FindDocuments, query, nTerms, collection, doc count: ", query,
		len(terms), nCol, nDoc)
	return QueryResults{nCol, nDoc, collections, documents, terms}, err
}

// Returns the headword words in the query (only a single word based on Chinese
// query)
func findWords(query string) ([]Word, error) {
	ctx := context.Background()
	results, err := findWordStmt.QueryContext(ctx, query, query)
	if err != nil {
		applog.Error("findWords, Error for query: ", query, err)
		// Sleep for a while, reinitialize, and retry
		time.Sleep(2000 * time.Millisecond)
		initStatements()
		results, err = findWordStmt.QueryContext(ctx, query, query)
		if err != nil {
			applog.Error("findWords, Give up after retry: ", query, err)
			return []Word{}, err
		}
	}
	words := []Word{}
	for results.Next() {
		word := Word{}
		var hw sql.NullInt64
		var trad sql.NullString
		results.Scan(&word.Simplified, &trad, &word.Pinyin, &hw)
		applog.Info("findWords, simplified, headword = ", word.Simplified, hw)
		if trad.Valid {
			word.Traditional = trad.String
		}
		if hw.Valid {
			word.HeadwordId = int(hw.Int64)
		}
		words = append(words, word)
	}
	return words, nil
}
