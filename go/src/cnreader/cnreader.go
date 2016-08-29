/* 
Command line utility to mark up HTML files with Chinese notes.
 */
package main

import (
	"flag"
	"cnreader/analysis"
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"log"
)

//Entry point for the chinesenotes command line tool.
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
	var wf = flag.Bool("wf", false, "Compute wf for all the corpus files " +
		"listed in data/corpus/collections.csv")
	flag.Parse()

	// Read in dictionary
	dictionary.ReadDict(config.LUFileNames())

	if (*collectionFile != "") {
		log.Printf("main: Analyzing collection %s\n", *collectionFile)
		analysis.WriteCorpusCol(*collectionFile)
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
			text := analysis.ReadText(src)
			tokens, results := analysis.ParseText(text, "",
				corpus.NewCorpusEntry())
			analysis.WriteDoc(tokens, results.Vocab, dest, conversion.Template,
				templateFile)
		}
	} else if *headwords {
		log.Printf("main: Write Headwords\n")
		dictionary.WriteHeadwords()
	} else if *wf {
		log.Printf("main: Computing word frequencies for whole corpus\n")
		analysis.WordFrequencies()
	} else if *hwFiles {
		log.Printf("main: Writing word entries for headwords\n")
		analysis.WriteHwFiles()
	} else {
		log.Printf("main: Writing out entire corpus\n")
		analysis.WriteCorpusAll()
	}
}
