package main
// Command line utility to mark up HTML files with Chinese notes.

import (
	"flag"
	"fmt"
	analysis "cnreader/analysis"
)

/*
 * Entry point for the chinesenotes command line tool.
 */
func main() {
	// Command line flags
	var infile = flag.String("infile", "testdata/test.html", "Input file")
	var outfile = flag.String("outfile", "testoutput/test-gloss.html", "Output file")
	flag.Parse()
	fmt.Printf("main: input file: %s, output file: %s\n", *infile, *outfile)

	// Read in dictionary
	analysis.ReadDict("../../../data/words.txt")

	// Read text and perform vocabulary analysis
	text := analysis.ReadText(*infile)
	tokens, vocab := analysis.ParseText(text)
	analysis.WriteDoc(tokens, vocab, *outfile)
}
