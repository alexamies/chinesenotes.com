/*
Package for scanning the corpora making up the library
*/
package corpus

import (
	"bufio"
	"cnreader/config"
	"encoding/csv"
	"log"
	"os"
	"text/template"
	"time"
)

type Corpus struct {
	Title, ShortName string
}

type LibraryMeta struct {
	Title, Summary, DateUpdated string
	Corpora []Corpus
}

const libraryFile = "data/corpus/library.csv"

// Contains the name of the corpora in the library
var corpora []Corpus

func init() {
	loadLibrary()
}

func Library() []Corpus {
	return corpora
}

// Gets the list of source and destination files for HTML conversion
func loadLibrary() []Corpus {
	fname := config.ProjectHome() + "/" + libraryFile
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal("library.loadLibrary: Error opening library file.", err)
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
	corpora = make([]Corpus, 0)
	for i, row := range rawCSVdata {
		if len(row) < 2 {
			log.Fatal("library.loadLibrary: not enough rows in file ", i,
				      len(row), fname)
	  	}
		title := row[0]
		shortName := row[1]
		corpus := Corpus{title, shortName}
		corpora = append(corpora, corpus)
	}
	return corpora
}

// Writes a HTML file describing the corpora in the library
func WriteLibraryFile() {
	dateUpdated := time.Now().Format("2006-01-02")
	libraryMeta := LibraryMeta{"Library", "Corpora in the Library", dateUpdated,
							corpora}
	outputFile := config.ProjectHome() + "/web/library.html"
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("library.WriteLibraryFile: could not open file", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	templFile := config.TemplateDir() + "/library-template.html"
	tmpl:= template.Must(template.New(
					"library-template.html").ParseFiles(templFile))
	err = tmpl.Execute(w, libraryMeta)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
