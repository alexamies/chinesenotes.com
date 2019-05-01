/*
Library for Chinese vocabulary analysis
*/
package analysis

import (
	"bufio"
	"bytes"
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"cnreader/index"
	"cnreader/ngram"
	"cnreader/library"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
	"unicode/utf8"
)

// Maximum number of word frequency entries to output to the generated
// HTML file
const MAX_WF_OUTPUT = 500

// Maximum number of unknwon characters to output to the generated
// HTML file
const MAX_UNKOWN_OUTPUT = 50

// Max usage elements for a word
const MAX_USAGE = 25

// Max number of occurrences of same title in a list of word usages
const MAX_TITLE = 5

// Max number of keywords to display
const MAX_KEYWORDS = 10

// Holds vocabulary analysis for a corpus text
type AnalysisResults struct {
	Title                   string
	WC, UniqueWords, CCount int
	ProperNouns				dictionary.Headwords
	DocumentGlossary 		Glossary
	TopKeywords				dictionary.Headwords
	WordFrequencies         []WFResult
	LexicalWordFreq         []WFResult
	BigramFreqSorted        []ngram.BigramFreq
	UnkownnChars            []index.SortedWordItem
	DateUpdated             string
	MaxWFOutput             int
}

// The content for a corpus entry
type CorpusEntryContent struct {
	CorpusText, DateUpdated, CollectionURL, CollectionTitle, EntryTitle, AnalysisFile string
}

// Dictionary entry content struct used for writing a dictionary entry to HTML
type DictEntry struct {
	Headword     dictionary.HeadwordDef
	RelevantDocs []index.RetrievalResult
	Contains     []dictionary.HeadwordDef
	Collocations []ngram.BigramFreq
	UsageArr     []WordUsage
	DateUpdated  string
}

// HTML content for template
type HTMLContent struct {
	Content, DateUpdated, Title, FileName string
}

// Bundles up vocabulary analysis
type VocabAnalysis struct {
	UsageMap map[string]*[]WordUsage
	WFTotal map[*index.CorpusWord]index.CorpusWordFreq
	WCTotal map[string]int
	Collocations ngram.CollocationMap
}

// Word usage
type WordUsage struct {
	Freq                                      int
	RelFreq                                   float64
	Word, Example, File, EntryTitle, ColTitle string
}

// Vocabulary analysis entry for a single word
type WFResult struct {
	Freq, HeadwordId                int
	Chinese, Pinyin, English, Usage string
}

/* Break usage example text into links with highlight on headword
   Parameters:
      usageText - usage example or notes on lexical unit
      headword - Includes simplified and traditional Chinese text
   Return
      marked up text with links and highlight
*/
func decodeUsageExample(usageText string, headword dictionary.HeadwordDef) string {
	tokens, _ := ParseText(usageText, "", corpus.NewCorpusEntry())
	replacementText := ""
	for e := tokens.Front(); e != nil; e = e.Next() {
		word := e.Value.(string)
		if word == *headword.Simplified || word == *headword.Traditional {
			replacementText = replacementText +
				"<span class='usage-highlight'>" + word + "</span>"
		} else {
			ws, ok := dictionary.GetWord(word)
			if ok {
				replacementText = replacementText + hyperlink(ws, word)
			} else {
				replacementText = replacementText + word
			}
		}
	}
	return replacementText
}

// Breaks text into a list of CJK and non CJK strings
func GetChunks(text string) list.List {
	var chunks list.List
	cjk := ""
	noncjk := ""
	for _, character := range text {
		if dictionary.IsCJKChar(string(character)) {
			if noncjk != "" {
				chunks.PushBack(noncjk)
				noncjk = ""
			}
			cjk += string(character)
		} else if cjk != "" {
			chunks.PushBack(cjk)
			cjk = ""
			noncjk += string(character)
		} else {
			noncjk += string(character)
		}
	}
	if cjk != "" {
		chunks.PushBack(cjk)
	}
	if noncjk != "" {
		chunks.PushBack(noncjk)
	}
	return chunks
}

// Compute word frequencies, collocations, and usage for the entire corpus
func GetWordFrequencies(libLoader library.LibraryLoader) VocabAnalysis {

	log.Printf("analysis.GetWordFrequencies: enter")

	// Overall word frequencies per corpus
	collocations := ngram.CollocationMap{}
	usageMap := map[string]*[]WordUsage{}
	ccount := 0 // character count
	wcTotal := map[string]int{}
	wfTotal := map[*index.CorpusWord]index.CorpusWordFreq{}
	corpusDir := config.ProjectHome() + "/corpus/"

	corpLoader := libLoader.GetCorpusLoader()
	collectionEntries := corpLoader.LoadCorpus(corpus.COLLECTIONS_FILE)
	for _, col := range collectionEntries {
		colFile := col.CollectionFile
		//log.Printf("GetWordFrequencies: input file: %s\n", colFile)
		corpusEntries := corpLoader.LoadCollection(colFile, col.Title)
		for _, entry := range corpusEntries {
			src := corpusDir + entry.RawFile
			text := corpLoader.ReadText(src)
			ccount += utf8.RuneCountInString(text)
			_, results := ParseText(text, col.Title, &entry)
			wcTotal[col.Corpus] += results.WC

			// Process collocations
			collocations.MergeCollocationMap(results.Collocations)

			// Find word usage for the given word
			for word, count := range results.Vocab {
				cw := &index.CorpusWord{col.Corpus, word}
				cwf := &index.CorpusWordFreq{col.Corpus, word, count}
				if cwfPrev, found := wfTotal[cw]; found {
					cwf.Freq += cwfPrev.Freq
				}
				wfTotal[cw] = *cwf
				rel_freq := 1000.0 * float64(count) / float64(results.WC)
				usage := WordUsage{cwf.Freq, rel_freq, word, results.Usage[word],
					entry.GlossFile, entry.Title, col.Title}
				usageArr, ok := usageMap[word]
				if !ok {
					usageArr = new([]WordUsage)
					usageMap[word] = usageArr
				}
				*usageArr = append(*usageArr, usage)
				//fmt.Fprintf(w, "%s\t%d\t%f\t%s\t%s\t%s\t%s\n", word, count, rel_freq,
				//	entry.GlossFile, col.Title, entry.Title, usage[word])
			}
		}
	}

	usageMap = sampleUsage(usageMap)

	// Print out totals for each corpus
	for corpus, count := range wcTotal {
		log.Printf("WordFrequencies: Total word count for corpus %s: %d\n",
			corpus, count)
	}
	log.Printf("WordFrequencies: len(collocations) = %d\n", len(collocations))
	log.Printf("WordFrequencies: character count = %d\n", ccount)

	return VocabAnalysis{usageMap, wfTotal, wcTotal, collocations}
}

// Constructs a hyperlink for a headword, including Pinyin and English in the
// title attribute for the link mouseover
func hyperlink(entries []*dictionary.WordSenseEntry, text string) string {
	classTxt := "vocabulary"
	if entries[0].IsProperNoun() {
		classTxt = classTxt + " propernoun"
	}
	pinyin := entries[0].Pinyin
	english := entries[0].English
	if len(entries) > 1 {
		english = ""
		for i, entry := range entries {
			english += fmt.Sprintf("%d. %s, ", i + 1, entry.English)
		}
		english = english[0:len(english)-2]
	}
	return fmt.Sprintf(config.VocabFormat(),
		pinyin, english, classTxt, entries[0].HeadwordId, text)
}

// Parses a Chinese text into words
// Parameters:
//   text: the string to parse
//   ColTitle: Optional parameter used for tracing collocation usage
//   document: Optional parameter used for tracing collocation usage
// Returns:
//   tokens: the tokens for the parsed text
//   results: vocabulary analysis results
func ParseText(text string, colTitle string, document *corpus.CorpusEntry) (
		tokens list.List, results CollectionAResults) {
	vocab := map[string]int{}
	bigrams := map[string]int{}
	bigramMap := ngram.BigramFreqMap{}
	collocations := ngram.CollocationMap{}
	unknownChars := map[string]int{}
	usage := map[string]string{}
	wc := 0
	cc := 0
	chunks := GetChunks(text)
	wdict := dictionary.GetWDict()
	hwIdMap := dictionary.GetHwMap()
	lastHWPtr := new(dictionary.HeadwordDef)
	lastHW := *lastHWPtr
	lastHWText := ""
	//fmt.Printf("ParseText: For text %s got %d chunks\n", text, chunks.Len())
	for e := chunks.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("ParseText: chunk %s\n", chunk)
		characters := strings.Split(chunk, "")
		if !dictionary.IsCJKChar(characters[0]) || corpus.IsExcluded(chunk) {
			tokens.PushBack(chunk)
			lastHWPtr = new(dictionary.HeadwordDef)
			lastHW = *lastHWPtr
			lastHWText = ""
			continue
		}
		for i := 0; i < len(characters); i++ {
			for j := len(characters); j > 0; j-- {
				w := strings.Join(characters[i:j], "")
				//log.Printf("analysis.ParseText: i = %d, j = %d, w = %s\n", i, j, w)
				if wsArray, ok := wdict[w]; ok {
					//log.Printf("analysis.ParseText: found word %s, i = %d\n", w, i)
					tokens.PushBack(w)
					i = j - 1
					j = 0
					if !corpus.IsExcluded(w) {
						wc++
						cc += utf8.RuneCountInString(w)
						vocab[w]++
						if lastHWText != "" {
							bg := lastHWText + w
							bigrams[bg]++
						}
						lastHWText = w
						if _, ok := usage[w]; !ok {
							usage[w] = chunk
						}
						hwid := wsArray[0].HeadwordId
						hw := hwIdMap[hwid]
						if lastHW.Id != 0 {
							if hw.WordSenses == nil {
								log.Printf("ParseText: WordSenses nil for %s "+
									", id = %d, in %s, %s\n", w, hwid,
									document.Title, colTitle)
							}
							bigram, ok := bigramMap.GetBigramVal(lastHW.Id,
								wsArray[0].HeadwordId)
							if !ok {
								bigram = ngram.NewBigram(lastHW, hw, chunk,
									document.GlossFile, document.Title, colTitle)
							}
							bigramMap.PutBigram(bigram)
							collocations.PutBigram(bigram.HeadwordDef1.Id, bigram)
							collocations.PutBigram(bigram.HeadwordDef2.Id, bigram)
						}
						lastHW = hw
					}
				} else if utf8.RuneCountInString(w) == 1 {
					//log.Printf("ParseText: found unknown character %s\n", w)
					unknownChars[w]++
					tokens.PushBack(w)
					break
				}
			}
		}
	}
	//log.Printf("analysis.ParseText: %s found character count %d, vocab %d\n",
	//	document.RawFile, cc, len(vocab))
	dl := index.DocLength{document.GlossFile, wc}
	dlArray := []index.DocLength{dl}
	results = CollectionAResults{
		Vocab:				vocab,
		Bigrams:			bigrams,
		Usage:				usage,
		BigramFrequencies:	bigramMap,
		Collocations:		collocations,
		WC:					wc,
		CCount:				cc,
		UnknownChars:		unknownChars,
		DocLengthArray:		dlArray,
	}
	return tokens, results
}

// Sample word usage for usability making sure that the list of word usage
// samples is not dominated by any one title and truncating at MAX_USAGE
// examples.
func sampleUsage(usageMap map[string]*[]WordUsage) map[string]*[]WordUsage {
	for word, usagePtr := range usageMap {
		sampleMap := map[string]int{}
		usage := *usagePtr
		usageCapped := new([]WordUsage)
		j := 0
		for _, wu := range usage {
			count, _ := sampleMap[wu.ColTitle]
			if count < MAX_TITLE && j < MAX_USAGE {
				*usageCapped = append(*usageCapped, wu)
				sampleMap[wu.ColTitle]++
				j++
			}
		}
		usageMap[word] = usageCapped
	}
	return usageMap
}

// For the HTML template
func add(x, y int) int {
	return x + y
}

// Writes out an analysis of the entire corpus, including word frequencies
// and other data. The output file is called 'corpus-analysis.html' in the
// web/analysis directory.
// Parameters:
//   results: The results of corpus analysis
//   docFreq: document frequency for terms
// Returns: the name of the file written to
func writeAnalysisCorpus(results CollectionAResults,
	docFreq index.DocumentFrequency) string {

	// If the web/analysis directory does not exist, then skip the analysis
	analysisDir := config.WebDir() + "/analysis/"
	_, err := os.Stat(analysisDir)
	if err != nil {
		return ""
	}

	// Parse template and organize template parameters
	sortedWords := index.SortedFreq(results.Vocab)
	wfResults := results.GetWordFreq(sortedWords)
	maxWf := len(wfResults)
	if maxWf > MAX_WF_OUTPUT {
		maxWf = MAX_WF_OUTPUT
	}

	lexicalWordFreq := results.GetLexicalWordFreq(sortedWords)
	maxLex := len(lexicalWordFreq)
	if maxLex > MAX_WF_OUTPUT {
		maxLex = MAX_WF_OUTPUT
	}

	sortedUnknownWords := index.SortedFreq(results.UnknownChars)
	maxUnknown := len(sortedUnknownWords)
	if maxUnknown > MAX_UNKOWN_OUTPUT {
		maxUnknown = MAX_UNKOWN_OUTPUT
	}

	// Bigrams, also truncated
	bFreq := ngram.SortedFreq(results.BigramFrequencies)
	maxBFOutput := len(bFreq)
	if maxBFOutput > MAX_WF_OUTPUT {
		maxBFOutput = MAX_WF_OUTPUT
	}

	dateUpdated := time.Now().Format("2006-01-02")
	title := "Terminology Extraction and Vocabulary Analysis"
	aResults := AnalysisResults{
		Title:            title,
		WC:               results.WC,
		CCount:			  results.CCount,
		ProperNouns:      dictionary.Headwords{},
		DocumentGlossary: MakeGlossary("", []dictionary.HeadwordDef{}),
		TopKeywords:	  dictionary.Headwords{},
		UniqueWords:      len(results.Vocab),
		WordFrequencies:  wfResults[:maxWf],
		LexicalWordFreq:  lexicalWordFreq[:maxLex],
		BigramFreqSorted: bFreq[:maxBFOutput],
		UnkownnChars:     sortedUnknownWords[:maxUnknown],
		DateUpdated:      dateUpdated,
		MaxWFOutput:      len(wfResults),
	}
	tmplFile := config.TemplateDir() + "/corpus-summary-analysis-template.html"
	funcs := template.FuncMap{
		"add": add,
		"Deref":   func(sp *string) string { return *sp },
		"DerefNe": func(sp *string, s string) bool { return *sp != s },
	}
	tmpl, err := template.New("corpus-summary-analysis-template.html").Funcs(funcs).ParseFiles(tmplFile)
	if (err != nil || tmpl == nil) {
		log.Printf("writeAnalysisCorpus: Error getting template %v)", tmplFile)
		return ""
	}
	basename := "corpus_analysis.html"
	filename := analysisDir + basename
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("writeAnalysisCorpus: error creating file %v", err)
		return ""
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, aResults)
	if err != nil {
		log.Printf("writeAnalysisCorpus: error executing template%v", err)
	}
	w.Flush()
	log.Printf("writeAnalysisCorpus: finished executing template%v", err)

	// Write results to plain text files
	index.WriteWFCorpus(sortedWords, sortedUnknownWords, bFreq, results.WC)

	return basename
}

// Writes a document with vocabulary analysis of the text. The name of the
// output file will be source file with '-analysis' appended, placed in the
// web/analysis directory
// results: The results of vocabulary analysis
// collectionTitle: The title of the whole colleciton
// docTitle: The title of this specific document
// Returns the name of the file written to
func writeAnalysis(results CollectionAResults, srcFile, glossFile,
		collectionTitle, docTitle string) string {

	//log.Printf("analysis.writeAnalysis: enter")
	analysisDir := config.WebDir() + "/analysis/"
	_, err := os.Stat(analysisDir)
	if err != nil {
		return ""
	}

	// Parse template and organize template parameters
	properNouns := makePNList(results.Vocab)

	domain_label := config.GetVar("Domain")
	//log.Printf("analysis.writeAnalysis: domain_label: %s\n", domain_label)
	glossary := MakeGlossary(domain_label, results.GetHeadwords())

	sortedWords := index.SortedFreq(results.Vocab)
	//log.Printf("analysis.writeAnalysis: found sortedWords for %s, count %d\n",
	//	srcFile, len(sortedWords))

	wfResults := results.GetWordFreq(sortedWords)
	maxWf := len(wfResults)
	if maxWf > MAX_WF_OUTPUT {
		maxWf = MAX_WF_OUTPUT
	}

	lexicalWordFreq := results.GetLexicalWordFreq(sortedWords)
	maxLex := len(lexicalWordFreq)
	if maxLex > MAX_WF_OUTPUT {
		maxLex = MAX_WF_OUTPUT
	}

	topKeywords := []dictionary.HeadwordDef{}
	if domain_label != "" {
		keywords := index.SortByWeight(results.Vocab)
		topKeywords = index.GetHeadwordArray(keywords)
		topKeywords = dictionary.FilterByDomain(topKeywords, domain_label)
		maxKeywords := len(topKeywords)
		if maxKeywords > MAX_KEYWORDS {
			maxKeywords = MAX_KEYWORDS
		}
		topKeywords = topKeywords[:maxKeywords]
	}

	//log.Printf("analysis.writeAnalysis: len topKeywords: %d\n", len(topKeywords))

	sortedUnknownWords := index.SortedFreq(results.UnknownChars)
	maxUnknown := len(sortedUnknownWords)
	if maxUnknown > MAX_UNKOWN_OUTPUT {
		maxUnknown = MAX_UNKOWN_OUTPUT
	}

	// Bigrams, also truncated
	bFreq := ngram.SortedFreq(results.BigramFrequencies)
	maxBFOutput := len(bFreq)
	if maxBFOutput > MAX_WF_OUTPUT {
		maxBFOutput = MAX_WF_OUTPUT
	}

	dateUpdated := time.Now().Format("2006-01-02")
	title := "Glossary and Vocabulary for " + collectionTitle
	if docTitle != "" {
		title += ", " + docTitle
	}

	aResults := AnalysisResults{
		Title:            title,
		WC:               results.WC,
		CCount:			  results.CCount,
		ProperNouns:      properNouns,
		DocumentGlossary: glossary,
		TopKeywords:	  topKeywords,
		UniqueWords:      len(results.Vocab),
		WordFrequencies:  wfResults[:maxWf],
		LexicalWordFreq:  lexicalWordFreq[:maxLex],
		BigramFreqSorted: bFreq[:maxBFOutput],
		UnkownnChars:     sortedUnknownWords[:maxUnknown],
		DateUpdated:      dateUpdated,
		MaxWFOutput:      len(wfResults),
	}
	tmplFile := config.TemplateDir() + "/corpus-analysis-template.html"
	funcs := template.FuncMap{
		"add": add,
		"Deref":   func(sp *string) string { return *sp },
		"DerefNe": func(sp *string, s string) bool { 
			if sp != nil {
				return *sp != s 
			}
			return false
		},
	}
	tmpl, err := template.New("corpus-analysis-template.html").Funcs(funcs).ParseFiles(tmplFile)
	if (err != nil ||  tmpl == nil) {
		log.Printf("analysiswriteAnalysis: Skipping document analysis (%v)\n",
			tmplFile)
		return ""
	}

	// Write output
	i := strings.Index(srcFile, ".txt")
	if i <= 0 {
		i = strings.Index(srcFile, ".html")
		if i <= 0 {
			i = strings.Index(srcFile, ".csv")
			if i <= 0 {
				log.Fatal("writeAnalysis: Bad name for source file: ", srcFile)
			}
		}
	}
	basename := srcFile[:i] + "_analysis.html"
	filename := analysisDir + basename
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("analysis.writeAnalysis", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, aResults)
	if err != nil {
		panic(err)
	}
	w.Flush()

	return basename
}

// Writes a corpus document collection to HTML, including all the entries
// contained in the collection
// collectionEntry: the CollectionEntry struct
// baseDir: The base directory to use
func writeCollection(collectionEntry corpus.CollectionEntry, baseDir string,
		libLoader library.LibraryLoader) CollectionAResults {

	//log.Printf("analysis.writeCollection: enter CollectionFile =" +
	//		collectionEntry.CollectionFile)
	corpLoader := libLoader.GetCorpusLoader()
	corpusEntries := corpLoader.LoadCollection(collectionEntry.CollectionFile,
		collectionEntry.Title)
	aResults := NewCollectionAResults()
	for _, entry := range corpusEntries {
		//log.Printf("analysis.writeCollection: entry.RawFile = " + entry.RawFile)
		src := config.CorpusDir() + "/" + entry.RawFile
		dest := baseDir + "/" + entry.GlossFile
		//if collectionEntry.Title == "" {
		//	log.Printf("analysis.writeCollection: collectionEntry.Title is " +
		//		"empty, input file: %s, output file: %s\n", src, dest)
		//}
		text := corpLoader.ReadText(src)
		tokens, results := ParseText(text, collectionEntry.Title, &entry)
		aFile := writeAnalysis(results, entry.RawFile, entry.GlossFile,
			collectionEntry.Title, entry.Title)
		sourceFormat := "TEXT"
		if strings.HasSuffix(entry.RawFile, ".html") {
			sourceFormat = "HTML"
		}
		writeCorpusDoc(tokens, results.Vocab, dest, collectionEntry.GlossFile,
			collectionEntry.Title, entry.Title,  aFile, sourceFormat)
		aResults.AddResults(results)
		aResults.DocFreq.AddVocabulary(results.Vocab)
		aResults.BigramDF.AddVocabulary(results.Bigrams)
		aResults.WFDocMap.AddWF(results.Vocab, collectionEntry.GlossFile,
			entry.GlossFile, results.WC)
		aResults.BigramDocMap.AddWF(results.Bigrams, collectionEntry.GlossFile,
			entry.GlossFile, results.WC)
	}
	aFile := writeAnalysis(aResults, collectionEntry.CollectionFile,
		collectionEntry.GlossFile, collectionEntry.Title, "")
	corpus.WriteCollectionFile(collectionEntry, aFile, baseDir)
	//log.Printf("analysis.writeCollection: exit\n")
	return aResults
}

// Write all the collections in the given corpus
// collections: The set of collections to write to HTML
// baseDir: The base directory to use to write the files
func WriteCorpus(collections []corpus.CollectionEntry, baseDir string,
		libLoader library.LibraryLoader) {
	log.Printf("analysis.WriteCorpus: enter")
	index.Reset()
	wfDocMap := index.TermFreqDocMap{}
	bigramDocMap := index.TermFreqDocMap{}
	docFreq := index.NewDocumentFrequency() // used to accumulate doc frequencies
	bigramDF := index.NewDocumentFrequency()
	aResults := NewCollectionAResults()
	for _, collectionEntry := range collections {
		results := writeCollection(collectionEntry, baseDir, libLoader)
		aResults.AddResults(results)
		docFreq.AddDocFreq(results.DocFreq)
		bigramDF.AddDocFreq(results.BigramDF)
		wfDocMap.Merge(results.WFDocMap)
		bigramDocMap.Merge(results.BigramDocMap)
	}
	writeAnalysisCorpus(aResults, docFreq)
	docFreq.WriteToFile(index.DOC_FREQ_FILE)
	bigramDF.WriteToFile(index.BIGRAM_DOC_FREQ_FILE)
	wfDocMap.WriteToFile(docFreq, index.WF_DOC_FILE)
	bigramDocMap.WriteToFile(bigramDF, index.BF_DOC_FILE)
	index.WriteDocLengthToFile(aResults.DocLengthArray, index.DOC_LENGTH_FILE)
	index.BuildIndex()
	log.Printf("analysis.WriteCorpus: exit")
}

// Write all the collections in the default corpus (collections.csv file)
func WriteCorpusAll(libLoader library.LibraryLoader) {
	log.Printf("analysis.WriteCorpusAll: enter")
	corpLoader := libLoader.GetCorpusLoader()
	collections := corpLoader.LoadCorpus(corpus.COLLECTIONS_FILE)
	WriteCorpus(collections, config.WebDir(), libLoader)
}

// Writes a corpus document collection to HTML, including all the entries
// contained in the collection
// collectionFile: the name of the collection file
func WriteCorpusCol(collectionFile string,
			libLoader library.LibraryLoader) {
	collectionEntry, err := libLoader.GetCorpusLoader().GetCollectionEntry(collectionFile)
	if err != nil {
		log.Fatalf("analysis.WriteCorpusCol: fatal error %v", err)
	}
	writeCollection(collectionEntry, config.WebDir(), libLoader)
}

// Writes a corpus document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
// HTML template to use
// collectionURL: the URL of the collection that the corpus text belongs to
// collectionTitle: The collection title that the corpus entry belongs to
// aFile: The vocabulary analysis file written to or empty string for none
// sourceFormat: TEXT, or HTML used for formatting output
func writeCorpusDoc(tokens list.List, vocab map[string]int, filename string,
	collectionURL string, collectionTitle string, entryTitle string,
	aFile string, sourceFormat string) {

	var b bytes.Buffer
	replacer := strings.NewReplacer("\n", "<br/>")

	// Iterate over text chunks
	for e := tokens.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("WriteDoc: Word %s\n", word)
		if entries, ok := dictionary.GetWord(chunk); ok && !corpus.IsExcluded(chunk) {
			wordIds := ""
			for _, ws := range entries {
				if wordIds == "" {
					wordIds = fmt.Sprintf("%d", ws.Id)
				} else {
					wordIds = fmt.Sprintf("%s,%d", wordIds, ws.Id)
				}
			}
			fmt.Fprintf(&b, hyperlink(entries, chunk))
		} else {
			if sourceFormat != "HTML" {
				chunk = replacer.Replace(chunk)
			}
			b.WriteString(chunk)
		}
	}

	textContent := b.String()
	dateUpdated := time.Now().Format("2006-01-02")
	content := CorpusEntryContent{textContent, dateUpdated, collectionURL,
		collectionTitle, entryTitle, aFile}

	// Write to file
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	templFile := config.TemplateDir() + "/corpus-template.html"
	tmpl := template.Must(template.New("corpus-template.html").ParseFiles(templFile))
	err = tmpl.Execute(w, content)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// Writes a document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
// GlossChinese: whether to convert the Chinese text in the file to hyperlinks
func WriteDoc(tokens list.List, vocab map[string]int, filename,
	templateName, templateFile string, glossChinese bool, title string) {
	if templateFile != `\N` {
		writeHTMLDoc(tokens, vocab, filename, templateName, templateFile,
			glossChinese, title)
		return
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	// Iterate over text chunks
	for e := tokens.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("WriteDoc: Word %s\n", word)
		if entries, ok := dictionary.GetWord(chunk); ok && !corpus.IsExcluded(chunk) {
			wordIds := ""
			for _, ws := range entries {
				if wordIds == "" {
					wordIds = fmt.Sprintf("%d", ws.Id)
				} else {
					wordIds = fmt.Sprintf("%s,%d", wordIds, ws.Id)
				}
			}
			fmt.Fprintf(w, "<span title='%s' data-wordid='%s'"+
				" class='dict-entry' data-toggle='popover'>%s</span>",
				chunk, wordIds, chunk)
		} else {
			w.WriteString(chunk)
		}
	}
	w.Flush()
}

// Writes a document with markup for the array of tokens
// tokens: A list of tokens forming the document
// vocab: A list of word id's in the document
// filename: The file name to write to
// GlossChinese: whether to convert the Chinese text in the file to hyperlinks
func writeHTMLDoc(tokens list.List, vocab map[string]int, filename,
	templateName, templateFile string, glossChinese bool, title string) {
	var b bytes.Buffer

	// Iterate over text chunks
	for e := tokens.Front(); e != nil; e = e.Next() {
		chunk := e.Value.(string)
		//fmt.Printf("WriteDoc: Word %s\n", word)
		if !glossChinese {
			fmt.Fprintf(&b, chunk)
		} else if entries, ok := dictionary.GetWord(chunk); ok {
			// Regular HTML link
			mouseover := fmt.Sprintf("%s | %s", entries[0].Pinyin,
				entries[0].English)
			link := fmt.Sprintf("/words/%d.html", entries[0].HeadwordId)
			fmt.Fprintf(&b, "<a title='%s' href='%s'>%s</a>", mouseover, link,
				chunk)
		} else {
			fmt.Fprintf(&b, chunk)
		}
	}
	dateUpdated := time.Now().Format("2006-01-02")
	fnameParts := strings.Split(filename, "/")
	fname := fnameParts[len(fnameParts) - 1]
	content := HTMLContent{b.String(), dateUpdated, title, fname}

	// Prepare template
	tmpl := template.Must(template.New(templateName).ParseFiles(templateFile))
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	err = tmpl.Execute(w, content)
	if err != nil {
		panic(err)
	}
	w.Flush()
	f.Close()

}

// Writes dictionary headword entries
func WriteHwFiles(loader library.LibraryLoader) {
	log.Printf("analysis.WriteHwFiles: Begin +++++++++++\n")
	index.BuildIndex()
	log.Printf("analysis.WriteHwFiles: Get headwords\n")
	hwArray := dictionary.GetHeadwords()
	vocabAnalysis := GetWordFrequencies(loader)
	usageMap := vocabAnalysis.UsageMap
	collocations := vocabAnalysis.Collocations
	corpusEntryMap := loader.GetCorpusLoader().LoadAll(corpus.COLLECTIONS_FILE)
	outfileMap := corpus.GetOutfileMap(corpusEntryMap)
	dateUpdated := time.Now().Format("2006-01-02")

	// Prepare template
	log.Printf("analysis.WriteHwFiles: Prepare template\n")
	templFile := config.TemplateDir() + "/headword-template.html"
	fm := template.FuncMap{
		"Deref":   func(sp *string) string { return *sp },
		"DerefNe": func(sp *string, s string) bool { return *sp != s },
	}
	tmpl := template.Must(template.New("headword-template.html").Funcs(fm).ParseFiles(templFile))

	i := 0
	for _, hw := range hwArray {

		if i%1000 == 0 {
			log.Printf("analysis.WriteHwFiles: wrote %d words\n", i)
		}

		// Look for different writings of traditional form
		tradVariants := []dictionary.WordSenseEntry{}
		for _, ws := range *hw.WordSenses {
			if hw.Id != ws.HeadwordId {
				tradVariants = append(tradVariants, ws)
				//log.Printf("analysis.WriteHwFiles: hw.Id != ws.HeadwordId: %d, %d\n",
				//	hw.Id, ws.HeadwordId)
			}
		}

		//if hw.Id == 873 {
		//	log.Printf("analysis.WriteHwFiles: hw.Id %d, image: %s\n", hw.Id, hw.WordSenses[0].Image)
		//}

		// Words that contain this word
		contains := dictionary.ContainsWord(*hw.Simplified, hwArray)

		// Sorted array of collocations
		wordCollocations := collocations.SortedCollocations(hw.Id)

		// Combine usage arrays for both simplified and traditional characters
		usageArrPtr, ok := usageMap[*hw.Simplified]
		if !ok {
			usageArrPtr, ok = usageMap[*hw.Traditional]
			if !ok {
				//log.Printf("WriteHwFiles: no usage for %s", hw.Simplified)
				usageArrPtr = &[]WordUsage{}
			}
		} else {
			usageArrTradPtr, ok := usageMap[*hw.Traditional]
			if ok {
				usageArr := *usageArrPtr
				usageArrTrad := *usageArrTradPtr
				for j, _ := range usageArrTrad {
					usageArr = append(usageArr, usageArrTrad[j])
				}
				usageArrPtr = &usageArr
			}
		}

		// Decorate useage text
		hlUsageArr := []WordUsage{}
		for _, wu := range *usageArrPtr {
			hlText := decodeUsageExample(wu.Example, hw)
			hlWU := WordUsage{
				Freq:       wu.Freq,
				RelFreq:    wu.RelFreq,
				Word:       wu.Word,
				Example:    hlText,
				File:       wu.File,
				EntryTitle: wu.EntryTitle,
				ColTitle:   wu.ColTitle,
			}
			hlUsageArr = append(hlUsageArr, hlWU)
		}

		dictEntry := DictEntry {
			Headword:     hw,
			RelevantDocs: index.FindDocsForKeyword(hw, outfileMap),
			Contains:     contains,
			Collocations: wordCollocations,
			UsageArr:     hlUsageArr,
			DateUpdated:  dateUpdated,
		}
		filename := fmt.Sprintf("%s%s%d%s", config.WebDir(), "/words/",
			hw.Id, ".html")
		f, err := os.Create(filename)
		if err != nil {
			log.Printf("WriteHwFiles: Error creating file for hw.Id %d, "+
				"Simplified %s", hw.Id, hw.Simplified)
			log.Fatal(err)
		}
		w := bufio.NewWriter(f)
		err = tmpl.Execute(w, dictEntry)
		if err != nil {
			log.Printf("analysis.WriteHwFiles: error executing template for hw.Id: %d,"+
				" filename: %s, Simplified: %s", hw.Id, filename, hw.Simplified)
			log.Fatal(err)
		}
		w.Flush()
		f.Close()
		i++
	}
}

// Writes a HTML files describing the corpora in the library, both public and
// for the translation portal (requiring login)
func writeLibraryFile(lib library.Library, corpora []library.CorpusData,
		outputFile string) {
	log.Printf("analysis.writeLibraryFile: with %d corpora, outputFile = %s, " +
			"TargetStatus = %s", len(corpora), outputFile, lib.TargetStatus)
	libData := library.LibraryData{
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
func WriteLibraryFiles(lib library.Library) {
	corpora := lib.Loader.LoadLibrary()
	libraryOutFile := config.WebDir() + "/library.html"
	writeLibraryFile(lib, corpora, libraryOutFile)
	portalDir := ""
	if config.GetVar("GoStaticDir") != "" {
		portalDir = config.ProjectHome() + "/" + config.GetVar("GoStaticDir")
		_, err := os.Stat(portalDir)
		lib.TargetStatus = "translator_portal"
		if err == nil {
			portalLibraryFile := portalDir + "/portal_library.html"
			writeLibraryFile(lib, corpora, portalLibraryFile)
		}
	}
	for _, c := range corpora {
		outputFile := ""
		baseDir := ""
		if c.Status == "public" {
			baseDir = config.WebDir()
			outputFile = fmt.Sprintf("%s/%s.html", config.WebDir(),
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
		fName := fmt.Sprintf(c.FileName)
		collections := lib.Loader.GetCorpusLoader().LoadCorpus(fName)
		WriteCorpus(collections, baseDir, lib.Loader)
		corpus := library.Corpus{c.Title, "", lib.DateUpdated, collections}
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
