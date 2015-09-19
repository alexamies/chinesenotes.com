/*
Package for scanning the corpus collections
*/
package corpus

import (
	"encoding/csv"
	"log"
	"os"
)

type CorpusEntry struct {
	RawFile string
	GlossFile string
}

const collectionsFile = "data/corpus/collections.csv"

var projectHome string

func init() {
	projectHome = "../../../.."
}

// Gets the public home directory, relative to the cnreader command line tool
func ProjectHome() string {
	return projectHome
}

// Gets the list of source and destination files for HTML conversion
func Collections() []string {
	collectionsFile := projectHome + "/" + collectionsFile
	file, err := os.Open(collectionsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	collections := make([]string, 0)
	for _, row := range rawCSVdata {
		collections = append(collections, row[0])
	}
	return collections
}

// Get a list of files for a corpus
func CorpusEntries(collectionFile string) []CorpusEntry {
	file, err := os.Open(collectionFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.Comma = rune('\t')
	reader.Comment = rune('#')
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	corpusEntries := make([]CorpusEntry, 0)
	for _, row := range rawCSVdata {
		corpusEntries = append(corpusEntries, CorpusEntry{row[0], row[1]})
	}
	return corpusEntries	
}

// Sets the public home directory, relative to the cnreader command line tool
func SetProjectHome(home string) {
	projectHome = home
}
