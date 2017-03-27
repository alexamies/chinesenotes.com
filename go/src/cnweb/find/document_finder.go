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
	"strings"
)

var (
	database *sql.DB
	findColStmt *sql.Stmt
	countColStmt *sql.Stmt
)

type Collection struct {
	CollectionFile string
	GlossFile      string
	Title          string
	Description    string
	Intro_file     string
	CorpusName     string
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

	stmt, err := database.Prepare("SELECT title, gloss_file FROM collection WHERE title LIKE ?")
    if err != nil {
        log.Fatal("cnweb.find.FindDocuments() Error preparing query: ", err)
    }
    findColStmt = stmt

	cstmt, err := database.Prepare("SELECT count(title) FROM collection WHERE title LIKE ?")
    if err != nil {
        log.Fatal("cnweb.find.FindDocuments() Error preparing query: ", err)
    }
    countColStmt = cstmt
}

func countDocuments(query string) int {
	var count int
	results, err := countColStmt.Query("%" + query + "%")
	results.Next()
	results.Scan(&count)
	if err != nil {
		applog.Error("countDocuments: Error for query: ", query, err)
	}
	results.Close()
	return count
}

func FindDocuments(query string) string {
	applog.Info("FindDocuments, ", query)
	count := countDocuments(query)
	applog.Info("FindDocuments, expect count: ", count)
	results, err := findColStmt.Query("%" + query + "%")
	if err != nil {
		applog.Error("FindDocuments, Error for query: ", query, err)
	}
	defer results.Close()

	json := "{\"collections\": ["
	i := 0
	for results.Next() {
		i++
		col := Collection{}
		results.Scan(&col.Title, &col.GlossFile)
		json += fmt.Sprintf("{\"title\":\"%s\", \"gloss_file\":\"%s\"},",
			col.Title, col.GlossFile)
	}
	json = strings.TrimSuffix(json, ",") + "]}"
	applog.Info("FindDocuments, num results returned: ", i)
	applog.Info("FindDocuments, results returned: ", json)
	return json
}