/*
Package for working with the plain, full text of corpus documents
*/
package fulltext

import (
	"cnweb/applog"
	"context"
	"io/ioutil"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

const (
	SNIPPET_LEN = 200
)

// Details of best matching text for the query terms
type DocMatch struct{
	PlainTextFile string
	MT MatchingText
}

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

// Implements the TextLoader interface, loads the text from a Google Cloud
// Storage.
// Params:
//   Bucket - The base URL for the location of the plain text files
type GCSLoader struct{
	bucket string
	client *storage.Client
}

// Gets the matching text from a local file and find the best match
func (loader GCSLoader) GetMatching(plainTextFile string,
		queryTerms []string) (MatchingText, error) {
	applog.Info("GCSLoader.GetMatching ", plainTextFile)
	ctx := context.Background()
	r, err := loader.client.Bucket(loader.bucket).Object(plainTextFile).NewReader(ctx)
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
			loader, err := NewGCSLoader(bucket)
			if err == nil {
				applog.Info("fulltext.getLoader, using GCSLoader")
				return loader
			}
		}
	}
	if corpusDir, ok := os.LookupEnv("CORPUS_DIR"); ok {
		applog.Info("fulltext.getLoader, using LocalTextLoader, ", corpusDir)
		return LocalTextLoader{corpusDir}
	}
	applog.Info("fulltext.getLoader, using LocalTextLoader,default corpusDir")
	return LocalTextLoader{"../../../corpus"}
}

// Given the already retrieved text body, find the best match
func getMatch(txt string, queryTerms []string) MatchingText {
	if len(queryTerms) == 0 {
		return MatchingText{}
	}
	query := strings.Join(queryTerms, "")
	match := false
	snippet := ""
	longest := ""
	i := strings.Index(txt, query)
	if i > -1 {
		longest = query
		match = true
	} else {
		j := 1
		l := len(queryTerms)
		substr := ""
		maxLen := 0
		for ; j < l; j++ {
			substr = strings.Join(queryTerms[j:l], "")
			i = strings.Index(txt, substr)
			if i > -1 {
				longest = substr
				maxLen = l - j
				break
			}
		}
		k := -1
		for j = l - 1; j > 0; j-- {
			substr = strings.Join(queryTerms[0:j], "")
			k = strings.Index(txt, substr)
			if k > -1 {
				break
			}
		}
		if j > maxLen {
			i = k
			longest = substr
		}
	}
	if i > -1 {
		s := i - SNIPPET_LEN / 2
		if s < 0 {
			s = 0
		}
		e := i + SNIPPET_LEN / 2
		if e > (len(txt) - 1) {
			e = len(txt) - 1
		}
		start := 0
		end := e
		// Make sure that snippet falls on a proper unicode boundary
		for j, _ := range txt {
			if (start == 0) && (j > s) {
				start = j
			}
			if j > e {
				end = j
				break
			}
    	}
		snippet = txt[start:end]
	}
	mt := MatchingText{
		Snippet: 		snippet,
		LongestMatch:	longest,
		ExactMatch:		match,
	}
	return	mt
}

// Creates and initiates a new GCSLoader object
func NewGCSLoader(bucket string) (GCSLoader, error) {
	applog.Info("fulltext.NewGCSLoader, ", bucket)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
    	applog.Info("fulltext.NewGCSLoader error getting client ", err)
    	return GCSLoader{}, err 
	}
	return GCSLoader{bucket, client}, nil
}