/*
Package for scanning the corpora making up the library
*/
package library

import (
	"bufio"
	"cnreader/analysis"
	"cnreader/config"
	"cnreader/corpus"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type Corpus struct {
	Title, ShortName, Status, FileName string
}

type CorpusMeta struct {
	Title, Summary, DateUpdated string
	Collections []corpus.CollectionEntry
}

type LibraryMeta struct {
	Title, Summary, DateUpdated, TargetStatus string
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
		if len(row) < 4 {
			log.Fatal("library.loadLibrary: not enough rows in file ", i,
				      len(row), fname)
	  	}
		title := row[0]
		shortName := row[1]
		status := row[2]
		fileName := row[3]
		corpus := Corpus{title, shortName, status, fileName}
		corpora = append(corpora, corpus)
	}
	return corpora
}


// Writes a HTML files describing the corpora in the library, both public and
// for the translation portal (requiring login)
func writeLibraryFile(targetStatus, title, summary, outputFile string) {
	dateUpdated := time.Now().Format("2006-01-02")
	libraryMeta := LibraryMeta{title, summary,
				dateUpdated, targetStatus, corpora}
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

// Writes a HTML file describing the corpora in the library and for each corpus
// in the library
func WriteLibraryFiles() {
	title := "Library"
	summary := "Top level collection in the Library"
	libraryOutFile := config.ProjectHome() + "/web/library.html"
	writeLibraryFile("public", title, summary, libraryOutFile)
	portalDir := ""
	if config.GetVar("GoStaticDir") != "" {
		title2 := "Translator Portal Library"
		summary2 := "Access for Translators Only"
		portalDir = config.ProjectHome() + "/" + config.GetVar("GoStaticDir")
		_, err := os.Stat(portalDir)
		if err == nil {
			portalLibraryFile := portalDir + "/portal_library.html"
			writeLibraryFile("translator_portal", title2, summary2, portalLibraryFile)
		}
	}
	dateUpdated := time.Now().Format("2006-01-02")
	for _, c := range corpora {
		outputFile := ""
		baseDir := ""
		if c.Status == "public" {
			baseDir = config.ProjectHome() + "/web"
			outputFile = fmt.Sprintf("%s/web/%s.html", config.ProjectHome(),
					c.ShortName)
		} else if c.Status == "translator_portal" {
			baseDir = portalDir
			outputFile = fmt.Sprintf("%s/%s.html", portalDir,
					c.ShortName)
		} else {
			log.Printf("library.WriteLibraryFiles: not sure what to do with status",
				c.Status)
			continue
		}
		fName := fmt.Sprintf("data/corpus/%s", c.FileName)
		collections := corpus.CorpusCollections(fName)
		analysis.WriteCorpus(collections, baseDir)
		corpusMeta := CorpusMeta{c.Title, "", dateUpdated, collections}
		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatal("library.WriteLibraryFiles: could not open file", err)
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		templFile := config.TemplateDir() + "/corpus-list-template.html"
		tmpl:= template.Must(template.New(
					"corpus-list-template.html").ParseFiles(templFile))
		err = tmpl.Execute(w, corpusMeta)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}
}
