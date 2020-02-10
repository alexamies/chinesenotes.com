/* 
Command line utility to mark up HTML files with Chinese notes.
 */
package main

import (
	"flag"
	"github.com/alexamies/chinesenotes-go/fileloader"
	"github.com/alexamies/chinesenotes-go/tokenizer"
	"github.com/alexamies/cnreader/analysis"
	"github.com/alexamies/cnreader/config"
	"github.com/alexamies/cnreader/corpus"
	"github.com/alexamies/cnreader/dictionary"
	"github.com/alexamies/cnreader/library"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

// Entry point for the chinesenotes command line tool.
// Default action is to write out all corpus entries to HTML files
func main() {
	// Command line flags
	var collectionFile = flag.String("collection", "", 
		"Enhance HTML markup and do vocabulary analysis for all the files " +
		"listed in given collection.")
	var headwords = flag.Bool("headwords", false,
		"Compute headword definitions " +
		" for all lexical units listed in data/words.txt, writing to the " +
		"headword.txt file.")
	var html = flag.Bool("html", false, "Enhance HTML markup for all files " +
		"listed in data/corpus/html-conversion.csv")
	var hwFiles = flag.Bool("hwfiles", false, "Compute and write " +
		"HTML entries for each headword, writing the files to the "+
		"web/words directory.")
	var librarymeta = flag.Bool("librarymeta", false, "Top level " +
		"collection entries for the digital library.")
	var memprofile = flag.String("memprofile", "", "write memory profile to " +
				"this file")
	flag.Parse()

	// Read in dictionary
	dictionary.ReadDict(config.LUFileNames())

	// Setup loader for library
	fname := config.ProjectHome() + "/" + library.LibraryFile
	fileLibraryLoader := library.FileLibraryLoader{fname}

	wdict, err := fileloader.LoadDictFile(config.LUFileNames())
	if err != nil {
		log.Fatal("Error opening dictionary, ", err)
		os.Exit(1)
	}
	dictTokenizer := tokenizer.DictTokenizer{wdict}

	if (*collectionFile != "") {
		log.Printf("main: Analyzing collection %s\n", *collectionFile)
		analysis.WriteCorpusCol(*collectionFile, fileLibraryLoader)
	} else if *html {
		log.Printf("main: Converting all HTML files\n")
		conversions := config.GetHTMLConversions()
		for _, conversion := range conversions {
			src := config.WebDir() + "/" + conversion.SrcFile
			dest := config.WebDir() + "/" + conversion.DestFile
			templateFile := `\N`
			if conversion.Template != `\N` {
				templateFile = config.TemplateDir() + "/" + conversion.Template
			}
			log.Printf("main: input file: %s, output file: %s, template: %s\n",
				src, dest, templateFile)
			text := fileLibraryLoader.GetCorpusLoader().ReadText(src)
			tokens, results := analysis.ParseText(text, "",
				corpus.NewCorpusEntry(), dictTokenizer)
			analysis.WriteDoc(tokens, results.Vocab, dest, conversion.Template,
				templateFile, conversion.GlossChinese, conversion.Title)
		}
	} else if *headwords {
		log.Printf("main: Write Headwords\n")
		dictionary.WriteHeadwords()
	} else if *hwFiles {
		log.Printf("main: Writing word entries for headwords\n")
		analysis.WriteHwFiles(fileLibraryLoader, dictTokenizer)
	} else if *librarymeta {
		log.Printf("main: Writing digital library metadata\n")
		fname := config.ProjectHome() + "/" + library.LibraryFile
		fileLibraryLoader := library.FileLibraryLoader{fname}
		dateUpdated := time.Now().Format("2006-01-02")
		lib := library.Library{
			Title: "Library",
			Summary: "Top level collection in the Library",
			DateUpdated: dateUpdated,
			TargetStatus: "public",
			Loader: fileLibraryLoader,
		}
		analysis.WriteLibraryFiles(lib)
	} else {
		log.Printf("main: Writing out entire corpus\n")
		analysis.WriteCorpusAll(fileLibraryLoader)
	}

	// Memory profiling
	if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.WriteHeapProfile(f)
        f.Close()
    }
}
