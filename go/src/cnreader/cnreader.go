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
	"cnreader/ngram"
	"log"
)

//Entry point for the chinesenotes command line tool.
func main() {
	// Command line flags
	var analysisFile = flag.String("analysis", "testoutput/test-analysis.html",
		"Vocabulary Analysis file")
	var collectionFile = flag.String("collection", "", 
		"Enhance HTML markup and do vocabulary analysis for all the files " +
		"listed in given collection.")
	var headwords = flag.Bool("headwords", false, "Compute headword definitions " +
		" for all lexical units listed in data/words.txt, writing to the " +
		"headword.txt file.")
	var html = flag.Bool("html", false, "Enhance HTML markup for all files " +
		"listed in data/corpus/html-conversion.csv")
	var hwFiles = flag.Bool("hwfiles", false, "Compute and write " +
		"HTML entries for each headword, writing the files to the "+
		"web/words directory.")
	var infile = flag.String("infile", "", "Input file")
	var outfile = flag.String("outfile", "", "Output file")
	var wf = flag.Bool("wf", false, "Compute wf for all the corpus files " +
		"listed in data/corpus/collections.csv")
	flag.Parse()

	// Read in dictionary
	dictionary.ReadDict(config.LUFileName())

	if (*collectionFile != "") {
		log.Printf("main: Analyzing collection %s\n", *collectionFile)
		collectionEntry, err := corpus.GetCollectionEntry(*collectionFile)
		if err != nil {
			log.Fatalf("main: %v", err)
		}
		corpusEntries := corpus.CorpusEntries(config.CorpusDataDir() + "/" +
			*collectionFile)
		aResults := analysis.CollectionAResults{
			Vocab: map[string]int{},
			Usage: map[string]string{}, 
			BigramFrequencies: ngram.BigramFreqMap{},
			WC: 0,
			UnknownChars: map[string]int{},
		}
		for _, entry := range corpusEntries {
			src := config.CorpusDir() + "/" + entry.RawFile
			dest := config.WebDir() + "/" + entry.GlossFile
			log.Printf("main: input file: %s, output file: %s\n", src, dest)
			text := analysis.ReadText(src)
			tokens, results := analysis.ParseText(text, collectionEntry.Title,
				&entry)
			aFile := analysis.WriteAnalysis(results, entry.RawFile,
					collectionEntry.Title, entry.Title)
			analysis.WriteCorpusDoc(tokens, results.Vocab, dest,
				collectionEntry.GlossFile, collectionEntry.Title, aFile)
			aResults.AddResults(results)
		}
		aFile := analysis.WriteAnalysis(aResults, *collectionFile,
				collectionEntry.Title, "")
		corpus.WriteCollectionFile(*collectionFile, aFile)
	} else if *infile != "" {
		log.Printf("main: input file: %s, output file: %s, analysis file: %s\n",
			*infile, *outfile, *analysisFile)

		// Read text and perform vocabulary analysis
		text := analysis.ReadText(*infile)
		tokens, results := analysis.ParseText(text, "", corpus.NewCorpusEntry())
		analysis.WriteDoc(tokens, results.Vocab, *outfile, `\N`, `\N`)
		analysis.WriteAnalysis(results, *analysisFile,
			"To do: figure out the colleciton title",
			"To do: figure out the document title")
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
			tokens, results := analysis.ParseText(text, "", corpus.NewCorpusEntry())
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
		log.Printf("main: Nothing to do. Please enter a command\n")
	}
}