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
)

type CorpusData struct {
	Title, ShortName, Status, FileName string
}

type Corpus struct {
	Title, Summary, DateUpdated string
	Collections []corpus.CollectionEntry
}

// A Library is a set of corpora loaded using a LibraryLoader and metadata
type Library struct {
	Title, Summary, DateUpdated, TargetStatus string
	Loader LibraryLoader
}
// A LibraryData is a struct to output library metadata to a HTML file
type LibraryData struct {
	Title, Summary, DateUpdated, TargetStatus string
	Corpora []CorpusData
}

// A LibraryLoader loads teh corpora into the library
type LibraryLoader interface {

	// Method to load the corpora in the library
	LoadLibrary() []CorpusData

	// Method to load the collections in a corpus
	// Parameter:
	//  fName: the file name listing the collections
	LoadCorpus(fName string) []corpus.CollectionEntry
}

// A FileLibraryLoader loads the corpora from files
type FileLibraryLoader struct{FileName string}

func (loader FileLibraryLoader) LoadLibrary() []CorpusData {
	return loadLibrary(loader.FileName)
}

func (loader FileLibraryLoader) LoadCorpus(fName string) []corpus.CollectionEntry {
	return corpus.CorpusCollections(fName)
}

// The library file listing the corpora
const LibraryFile = "data/corpus/library.csv"

// Gets the list of source and destination files for HTML conversion
func loadLibrary(fname string) []CorpusData {
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
	corpora := []CorpusData{}
	for i, row := range rawCSVdata {
		if len(row) < 4 {
			log.Fatal("library.loadLibrary: not enough rows in file ", i,
				      len(row), fname)
	  	}
		title := row[0]
		shortName := row[1]
		status := row[2]
		fileName := row[3]
		corpus := CorpusData{title, shortName, status, fileName}
		corpora = append(corpora, corpus)
	}
	return corpora
}


// Writes a HTML files describing the corpora in the library, both public and
// for the translation portal (requiring login)
func writeLibraryFile(lib Library, corpora []CorpusData, outputFile string) {
	libData := LibraryData{
		Title: lib.Title,
		Summary: lib.Summary,
		DateUpdated: lib.DateUpdated,
		TargetStatus: lib.TargetStatus,
		Corpora: corpora,
	}
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("library.WriteLibraryFile: could not open file", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	templFile := config.TemplateDir() + "/library-template.html"
	tmpl:= template.Must(template.New(
					"library-template.html").ParseFiles(templFile))
	err = tmpl.Execute(w, libData)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()

}

// Writes a HTML file describing the corpora in the library and for each corpus
// in the library
func WriteLibraryFiles(lib Library) {
	corpora := lib.Loader.LoadLibrary()
	libraryOutFile := config.ProjectHome() + "/web/library.html"
	writeLibraryFile(lib, corpora, libraryOutFile)
	portalDir := ""
	if config.GetVar("GoStaticDir") != "" {
		portalDir = config.ProjectHome() + "/" + config.GetVar("GoStaticDir")
		_, err := os.Stat(portalDir)
		if err == nil {
			portalLibraryFile := portalDir + "/portal_library.html"
			writeLibraryFile(lib, corpora, portalLibraryFile)
		}
	}
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
		collections := lib.Loader.LoadCorpus(fName)
		analysis.WriteCorpus(collections, baseDir)
		corpus := Corpus{c.Title, "", lib.DateUpdated, collections}
		f, err := os.Create(outputFile)
		if err != nil {
			log.Fatal("library.WriteLibraryFiles: could not open file", err)
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		templFile := config.TemplateDir() + "/corpus-list-template.html"
		tmpl:= template.Must(template.New(
					"corpus-list-template.html").ParseFiles(templFile))
		err = tmpl.Execute(w, corpus)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}
}
