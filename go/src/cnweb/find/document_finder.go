/*
Functions for finding collections by partial match on collection title
*/
package find

import (
	"cnreader/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var database *sql.DB

type Collection struct {
	CollectionFile string
	GlossFile      string
	Title          string
	Description    string
	Intro_file     string
	CorpusName     string
}

func init() {
	log.Println("cnweb.find.init() enter")
	dbhost := config.GetVar("DBHost")
	dbport := config.GetVar("DBPort")
	dbuser := config.GetVar("DBUser")
	dbpass := config.GetVar("DBPass")
	dbname := config.GetVar("DBName")
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost,
		dbport, dbname)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Println("cnweb.find.init() could not connect to the database, ",
			err)
		panic(err.Error())
	}
	database = db
	log.Println("cnweb.find.init() exit")
}

func FindDocuments(query string) string {
	log.Println("cnweb.find.FindDocuments(): ", query)
	stmt, err := database.Prepare("SELECT title, gloss_file FROM collection WHERE title LIKE ?")
    if err != nil {
        log.Println("cnweb.find.FindDocuments() Error preparing query: ", query,
        	err)
    }
	results, err := stmt.Query("%" + query + "%")
	if err != nil {
		log.Println("cnweb.find.FindDocuments() Error for query: ", query, err)
	}
	defer results.Close()

	json := "{\"collections\": ["
	for results.Next() {
		col := Collection{}
		results.Scan(&col.Title, &col.GlossFile)
		json += fmt.Sprintf("{\"title\":\"%s\", \"gloss_file\":\"%s\"},",
			col.Title, col.GlossFile)
	}
	json = strings.TrimSuffix(json, ",") + "]}"
	return json
}
