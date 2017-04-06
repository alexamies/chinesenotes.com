/*
Functions for finding collections by partial match on collection title
*/
package find

import (
	"cnreader/config"
	"cnweb/applog"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

type Word struct {
	Simplified, Traditional, Pinyin, English string
	HeadwordId int
}

type QueryResults struct {
	NumCollections, NumDocuments int
	Collections []Collection
	Documents []Document
	Words []Word
}

// Open database connection and prepare statements
func init() {
	dbhost := config.GetVar("DBHost")
	dbport := config.GetVar("DBPort")
	dbuser := config.GetVar("DBUser")
	dbpass := config.GetVar("DBPass")
	dbname := config.GetVar("DBName")
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost,
		dbport, dbname)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Fatal("FATAL: could not connect to the database, ",
			err)
		panic(err.Error())
	}
	database = db

	stmt, err := database.Prepare("SELECT title, gloss_file FROM collection WHERE title LIKE ? LIMIT 50")
    if err != nil {
        log.Fatal("find.init() Error preparing stmt: ", err)
    }
    findColStmt = stmt

	cstmt, err := database.Prepare("SELECT count(title) FROM collection WHERE title LIKE ?")
    if err != nil {
        log.Fatal("find.init() Error preparing cstmt: ", err)
    }
    countColStmt = cstmt

	dstmt, err := database.Prepare("SELECT title, gloss_file FROM document WHERE title LIKE ? LIMIT 50")
    if err != nil {
        log.Fatal("find.init() Error preparing dstmt: ", err)
    }
    findDocStmt = dstmt

	cdstmt, err := database.Prepare("SELECT count(title) FROM document WHERE title LIKE ?")
    if err != nil {
        log.Fatal("find.init() Error preparing cDocStmt: ", err)
    }
    countDocStmt = cdstmt    

	fwstmt, err := database.Prepare("SELECT simplified, traditional, pinyin, english, headword FROM words WHERE simplified = ? OR traditional = ? LIMIT 1")
    if err != nil {
        log.Fatal("find.init() Error preparing fwstmt: ", err)
    }
    findWordStmt = fwstmt
}

func countCollections(query string) int {
	var count int
	results, err := countColStmt.Query("%" + query + "%")
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
	results, err := countDocStmt.Query("%" + query + "%")
	results.Next()
	results.Scan(&count)
	if err != nil {
		applog.Error("countDocuments: Error for query: ", query, err)
	}
	results.Close()
	return count
}

func findCollections(query string) []Collection {
	results, err := findColStmt.Query("%" + query + "%")
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

func findDocuments(query string) []Document {
	results, err := findDocStmt.Query("%" + query + "%")
	if err != nil {
		applog.Error("findDocuments, Error for query: ", query, err)
	}
	defer results.Close()

	documents := []Document{}
	for results.Next() {
		doc := Document{}
		results.Scan(&doc.Title, &doc.GlossFile)
		documents = append(documents, doc)
	}
	return documents
}

func FindDocuments(query string) QueryResults {
	applog.Info("FindDocuments, ", query)
	words := findWords(query)
	nCol := countCollections(query)
	nDoc := countDocuments(query)
	collections := findCollections(query)
	documents := findDocuments(query)
	applog.Info("FindDocuments, collection, doc count: ", nCol, nDoc)
	return QueryResults{nCol, nDoc, collections, documents, words}
}

// Returns the headword words in the query (only a single word at the moment)
func findWords(query string) []Word {
	results, err := findWordStmt.Query(query, query)
	if err != nil {
		applog.Error("findWords, Error for query: ", query, err)
	}
	words := []Word{}
	for results.Next() {
		word := Word{}
		var hw sql.NullInt64
		var trad sql.NullString
		results.Scan(&word.Simplified, &trad, &word.Pinyin, &word.English, &hw)
		applog.Error("findWords, simplified, headword = ", word.Simplified, hw)
		if trad.Valid {
			word.Traditional = trad.String
		}
		if hw.Valid {
			word.HeadwordId = int(hw.Int64)
		}
		words = append(words, word)
	}
	return words
}
