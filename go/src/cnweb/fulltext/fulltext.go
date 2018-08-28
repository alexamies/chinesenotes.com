/*
Package for working with the plain, full text of corpus documents
*/
package fulltext

import (
	"cnweb/applog"
	"database/sql"
	"context"
	"io/ioutil"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"cloud.google.com/go/storage"

	"cnweb/webconfig"
)
const (
	SNIPPET_LEN = 40
)

var (
	database *sql.DB
	docListStmt *sql.Stmt
	docFileMap map[string]string
)

// Details of best matching text for the query terms
type MatchingText struct{
	Snippet, LongestMatch string
	ExactMatch bool
}

// Interface for plain text retrieval
type TextLoader interface {

	// Get the document text
	// param:
	//   plainTextFile - file containing plain text of the document
	//   , queryTerms - an array of query terms
	GetMatching(plainTextFile string,
		queryTerms []string) (MatchingText, error)
}

// Cache the plain text file names
func cacheDocFileMap() map[string]string {
	docFileMap := map[string]string{}
	ctx := context.Background()
	results, err := docListStmt.QueryContext(ctx)
	if err != nil {
		applog.Error("cacheDocFileMap, Error for query: ", err)
		return docFileMap
	}
	defer results.Close()

	for results.Next() {
		plainTextFile := ""
		glossFile := ""
		results.Scan(&plainTextFile, &glossFile)
		docFileMap[glossFile] = plainTextFile
	}
	return docFileMap
}

// Implements the TextLoader interface, loads the text from a Google Cloud
// Storage.
// Params:
//   Bucket - The base URL for the location of the plain text files
type GCSLoader struct{Bucket string}

// Gets the matching text from a local file and find the best match
func (loader GCSLoader) GetMatching(plainTextFile string,
		queryTerms []string) (MatchingText, error) {
	applog.Info("GCSLoader.GetMatching ", plainTextFile)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
    	applog.Info("GCSLoader.GetMatching error getting client ", err)
    	return MatchingText{}, err 
	}
	r, err := client.Bucket(loader.Bucket).Object(plainTextFile).NewReader(ctx)
	if err != nil {
        return MatchingText{}, err
	}
	defer r.Close()

	bs, err := ioutil.ReadAll(r)
	if err != nil {
        return MatchingText{}, err
	}
	txt := string(bs)
	applog.Info("GCSLoader.GetMatching len(txt) ", len(txt))
	return getMatch(txt, queryTerms), nil
}

// Uses the environment variableS GOOGLE_APPLICATION_CREDENTIALS and TEXT_BUCKET
// to determine whether to load the files from the local file system or GCS.
func getLoader() TextLoader {
	if _, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS"); ok {
		if bucket, ok := os.LookupEnv("TEXT_BUCKET"); ok {
			return GCSLoader{bucket}
		}
	}
	if corpusDir, ok := os.LookupEnv("CORPUS_DIR"); ok {
		return LocalTextLoader{corpusDir}
	}
	return LocalTextLoader{"../../../corpus"}
}

// Given the already retrieved text body, find the best match
// Log errors and continue because it is not essential to user experience
func GetMatch(glossFile string, queryTerms []string) MatchingText {
	applog.Info("fulltext.GCSLoader.GetMatch enter ", glossFile)
	loader := getLoader()
	plainTextFile, ok := docFileMap[glossFile]
	if !ok {
		applog.Info("fulltext.GCSLoader.GetMatch could not find ", glossFile)
		return MatchingText{}
	}
	mt, err := loader.GetMatching(plainTextFile, queryTerms)
	if err != nil {
		applog.Info("fulltext.GCSLoader.GetMatch error getting match ", err)
	}
	applog.Info("fulltext.GCSLoader.GetMatch snippet ", mt.Snippet)
	return mt
}

// Given the already retrieved text body, find the best match
func getMatch(txt string, queryTerms []string) MatchingText {
	query := strings.Join(queryTerms, "")
	match := false
	snippet := ""
	i := strings.Index(txt, query)
	if i > -1 {
		match = true
		start := i - SNIPPET_LEN / 2
		if start < 0 {
			start = 0
		}
		end := i + SNIPPET_LEN / 2
		if end > (len(txt) - 1) {
			end = len(txt) - 1
		}
		snippet = txt[start:end]
	}
	mt := MatchingText{
		Snippet: 		snippet,
		LongestMatch:	query,
		ExactMatch:		match,
	}
	return	mt
}

// Open database connection and prepare statement
func init() {
	err := initStatements()
	if err != nil {
		applog.Error("fulltext/init: error for database statements, retrying",
			err)
		time.Sleep(60000 * time.Millisecond)
		err = initStatements()
		conString := webconfig.DBConfig()
		applog.Error("fulltext/init: error for database statements, giving up",
			conString, err)
	}
	docFileMap = cacheDocFileMap()
	applog.Info("fulltext.init len(docFileMap) ", len(docFileMap))
}

func initStatements() error {
	conString := webconfig.DBConfig()
	db, err := sql.Open("mysql", conString)
	if err != nil {
		return err
	}
	database = db

	ctx := context.Background()
	docListStmt, err = database.PrepareContext(ctx,
		"SELECT plain_text_file, gloss_file " +
		"FROM document")
    if err != nil {
        applog.Error("fulltext.initStatements() Error for doc stmt: ", err)
        return err
    }
    return nil
}

// Implements the TextLoader interface, loads the text from a local file
// mounted on the application server
// Params:
//   corpusDir - The top level directory for the plain text files
type LocalTextLoader struct{corpusDir string}

// Gets the matching text from a local file and find the best match
func (loader LocalTextLoader) GetMatching(plainTextFile string,
		queryTerms []string) (MatchingText, error) {
	fullPath := loader.corpusDir + "/" + plainTextFile
	bs, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return MatchingText{}, err
	}
	return getMatch(string(bs), queryTerms), nil
}
