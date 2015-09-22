/*
Package for scanning the corpus collections
*/
package corpus

import (
	"bufio"
	"cnreader/config"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type CollectionEntry struct {
	CollectionFile, GlossFile, Title, DateUpdated string
	CorpusEntries []CorpusEntry
}

type CorpusEntry struct {
	RawFile, GlossFile, Title string
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
func Collections() []CollectionEntry {
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
	collections := make([]CollectionEntry, 0)
	for _, row := range rawCSVdata {
		corpusEntries := make([]CorpusEntry, 0)
		collections = append(collections, CollectionEntry{row[0], row[1],
			row[2], "", corpusEntries})
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
		corpusEntries = append(corpusEntries, CorpusEntry{row[0], row[1],
			row[2]})
	}
	return corpusEntries	
}

// Sets the public home directory, relative to the cnreader command line tool
func SetProjectHome(home string) {
	projectHome = home
}

// Writes a HTML file describing the collection
// Parameter
// collectionFile: The name of the file describing the collection
func WriteCollectionFile(collectionFile string) {
	fmt.Printf("WriteCollectionFile: Writing collection file.\n")
	collections := Collections()
	for _, entry := range collections {
		if entry.CollectionFile == collectionFile && entry.GlossFile != "\\N" {
			outputFile := config.ProjectHome() + "/data/corpus/" +collectionFile
			entry.CorpusEntries = CorpusEntries(outputFile)
			fmt.Printf("WriteCollectionFile: Writing collection file %s\n",
				outputFile)

			// Write to file
			f, err := os.Create(config.ProjectHome() + "/web/" +
				entry.GlossFile)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			w := bufio.NewWriter(f)
			entry.DateUpdated = time.Now().Format("2006-01-02")
			templFile := config.ProjectHome() +
					"/corpus/collection-template.html"
			fmt.Println("Home: ", config.ProjectHome())
			tmpl:= template.Must(template.New(
					"collection-template.html").ParseFiles(templFile))
			err = tmpl.Execute(w, entry)
			if err != nil { panic(err) }
			w.Flush()
		}
	}
}
