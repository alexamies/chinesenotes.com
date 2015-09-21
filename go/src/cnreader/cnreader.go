/* 
Command line utility to mark up HTML files with Chinese notes.
 */
package main

import (
	"flag"
	"fmt"
	"cnreader/analysis"
	"cnreader/config"
	"cnreader/corpus"
)

//Entry point for the chinesenotes command line tool.
func main() {
	// Command line flags
	var all = flag.Bool("all", false, "Enhance HTML markup for all the files " +
		"listed in data/corpus/html-conversion.csv")
	var analysisFile = flag.String("analysis", "testoutput/test-analysis.html",
		"Vocabulary Analysis file")
	var corpusFile = flag.String("corpus", "", 
		"Enhance HTML markup and do vocabulary analysis for all the files " +
		"listed in given corpus file.")
	var infile = flag.String("infile", "testdata/test.html", "Input file")
	var outfile = flag.String("outfile", "testoutput/test-gloss.html",
		"Output file")
	flag.Parse()

	// Set project home relative to the command line tool directory
	projectHome := "../../.."
	config.SetProjectHome(projectHome)
	webDir := projectHome + "/web"
	corpusDir := projectHome + "/corpus"
	corpusDataDir := projectHome + "/data/corpus"


	// Read in dictionary
	analysis.ReadDict("../../../data/words.txt")

	if (*corpusFile != "") {
		fmt.Printf("main: Analyzing corpus %s\n", *corpusFile)
		corpusEntries := corpus.CorpusEntries(corpusDataDir + "/" + *corpusFile)
		for _, entry := range corpusEntries {
			src := corpusDir + "/" + entry.RawFile
			dest := webDir + "/" + entry.GlossFile
			aFile := webDir + "/analysis/" + entry.GlossFile
			fmt.Printf("main: input file: %s, output file: %s\n", src, dest)
			text := analysis.ReadText(src)
			tokens, vocab, wc := analysis.ParseText(text)
			analysis.WriteCorpusDoc(tokens, vocab, dest)
			analysis.WriteAnalysis(vocab, aFile, wc)
		}
	} else if !*all {
		fmt.Printf("main: input file: %s, output file: %s, analysis file: %s\n",
			*infile, *outfile, *analysisFile)

		// Read text and perform vocabulary analysis
		text := analysis.ReadText(*infile)
		tokens, vocab, wc := analysis.ParseText(text)
		analysis.WriteDoc(tokens, vocab, *outfile)
		analysis.WriteAnalysis(vocab, *analysisFile, wc)
	} else {
		fmt.Printf("main: Converting all HTML files\n")
		conversions := config.GetHTMLConversions()
		for _, conversion := range conversions {
			src := webDir + "/" + conversion.SrcFile
			dest := webDir + "/" + conversion.DestFile
			fmt.Printf("main: input file: %s, output file: %s\n", src, dest)
			text := analysis.ReadText(src)
			tokens, vocab, _ := analysis.ParseText(text)
			analysis.WriteDoc(tokens, vocab, dest)
		}
	}
}
