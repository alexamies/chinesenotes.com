/* 
Command line utility to mark up HTML files with Chinese notes.
 */
package main

import (
	"flag"
	"cnreader/analysis"
	"cnreader/config"
	"cnreader/corpus"
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
	var html = flag.Bool("html", false, "Enhance HTML markup for all files " +
		"listed in data/corpus/html-conversion.csv")
	var infile = flag.String("infile", "", "Input file")
	var outfile = flag.String("outfile", "", "Output file")
	var wf = flag.Bool("wf", false, "Compute wf for all the corpus files " +
		"listed in data/corpus/collections.csv")
	flag.Parse()

	// Set project home relative to the command line tool directory
	projectHome := "../../.."
	config.SetProjectHome(projectHome)
	webDir := projectHome + "/web"
	corpusDir := projectHome + "/corpus"
	corpusDataDir := projectHome + "/data/corpus"
	dataDir := projectHome + "/data/"


	// Read in dictionary
	analysis.ReadDict(dataDir + "words.txt")

	if (*collectionFile != "") {
		log.Printf("main: Analyzing collection %s\n", *collectionFile)
		collectionEntry, err := corpus.GetCollectionEntry(*collectionFile)
		if err != nil {
			log.Fatalf("Could not find collection file %s\n", *collectionFile)
			return
		}
		corpus.WriteCollectionFile(*collectionFile)
		corpusEntries := corpus.CorpusEntries(corpusDataDir + "/" +
			*collectionFile)
		for _, entry := range corpusEntries {
			src := corpusDir + "/" + entry.RawFile
			dest := webDir + "/" + entry.GlossFile
			log.Printf("main: input file: %s, output file: %s\n", src, dest)
			text := analysis.ReadText(src)
			tokens, vocab, wc, unknownChars := analysis.ParseText(text)
			aFile := analysis.WriteAnalysis(vocab, wc, unknownChars,
				entry.RawFile, collectionEntry.Title, entry.Title)
			analysis.WriteCorpusDoc(tokens, vocab, dest,
				collectionEntry.GlossFile, collectionEntry.Title, aFile)
		}
	} else if *infile != "" {
		log.Printf("main: input file: %s, output file: %s, analysis file: %s\n",
			*infile, *outfile, *analysisFile)

		// Read text and perform vocabulary analysis
		text := analysis.ReadText(*infile)
		tokens, vocab, wc, unknownChars := analysis.ParseText(text)
		analysis.WriteDoc(tokens, vocab, *outfile)
		analysis.WriteAnalysis(vocab, wc, unknownChars, *analysisFile,
			"To do: figure out the colleciton title",
			"To do: figure out the document title")
	} else if *html {
		log.Printf("main: Converting all HTML files\n")
		conversions := config.GetHTMLConversions()
		for _, conversion := range conversions {
			src := webDir + "/" + conversion.SrcFile
			dest := webDir + "/" + conversion.DestFile
			log.Printf("main: input file: %s, output file: %s\n", src, dest)
			text := analysis.ReadText(src)
			tokens, vocab, _, _ := analysis.ParseText(text)
			analysis.WriteDoc(tokens, vocab, dest)
		}
	} else if *wf {
		log.Printf("main: Computing word frequencies for whole corpus\n")
		analysis.WordFrequencies()
	} else {
		log.Printf("main: Nothing to do. Please enter a command\n")
	}
}