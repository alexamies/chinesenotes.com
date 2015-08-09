package main
// Command line utility to mark up HTML files with Chinese notes.

import analysis "cnreader/analysis"

func main() {
	analysis.ReadDict("../../../data/words.txt")
	filename := "../../../web/classical_chinese-raw.html"
	text := analysis.ReadText(filename)
	tokens, vocab := analysis.ParseText(text)
	outfile := "../../../web/classical_chinese.html"
	analysis.WriteDoc(tokens, vocab, outfile)
}
