package analysis

import (
	"cnreader/config"
	"cnreader/corpus"
	"cnreader/dictionary"
	"log"
	"strings"
	"testing"
)

func TestDecodeUsageExample1(t *testing.T) {
	hw := dictionary.HeadwordDef{
		Id: 1,
		Simplified: "海",
		Traditional: "\\N",
		Pinyin: []string{"hǎi"},
		WordSenses: []dictionary.WordSenseEntry{},
	}
	highlighted := decodeUsageExample("海", hw)
	expected := "<span class='usage-highlight'>海</span>"
	if highlighted != expected {
		t.Error("TestDecodeUsageExample1: Expected ", expected, ", got",
			highlighted)
	}
}

func TestDecodeUsageExample2(t *testing.T) {
	hw := dictionary.HeadwordDef{
		Id: 1,
		Simplified: "海",
		Traditional: "\\N",
		Pinyin: []string{"hǎi"},
		WordSenses: []dictionary.WordSenseEntry{},
	}
	highlighted := decodeUsageExample("banana", hw)
	expected := "banana"
	if highlighted != expected {
		t.Error("TestDecodeUsageExample2: Expected ", expected, ", got",
			highlighted)
	}
}

func TestDecodeUsageExample3(t *testing.T) {
	hw := dictionary.HeadwordDef{
		Id: 1,
		Simplified: "国",
		Traditional: "國",
		Pinyin: []string{"guó"},
		WordSenses: []dictionary.WordSenseEntry{},
	}
	highlighted := decodeUsageExample("國", hw)
	expected := "<span class='usage-highlight'>國</span>"
	if highlighted != expected {
		t.Error("TestDecodeUsageExample3: Expected ", expected, ", got",
			highlighted)
	}
}
func TestGetChunks1(t *testing.T) {
	dictionary.ReadDict("../testdata/testwords.txt")
	chunks := GetChunks("中文")
	if chunks.Len() != 1 {
		t.Error("TestGetChunks1: Expected length of chunks 1, got ",
			chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "中文" {
		t.Error("TestGetChunks1: Expected first element of chunk 中文, got ",
			chunk)
	}
}

func TestGetChunks2(t *testing.T) {
	dictionary.ReadDict("../testdata/testwords.txt")
	chunks := GetChunks("a中文")
	if chunks.Len() != 2 {
		t.Error("Expected length of chunks 2, got ", chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "a" {
		t.Error("Expected first element of chunk a, got ", chunk)
	}
}

func TestGetChunks3(t *testing.T) {
	dictionary.ReadDict("../testdata/testwords.txt")
	chunks := GetChunks("a中文b")
	if chunks.Len() != 3 {
		t.Error("Expected length of chunks 3, got ", chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "a" {
		t.Error("Expected first element of chunk a, got ", chunk)
	}
}

// Simplified Chinese
func TestGetChunks4(t *testing.T) {
	dictionary.ReadDict("../testdata/testwords.txt")
	chunks := GetChunks("简体中文")
	if chunks.Len() != 1 {
		t.Error("Simplified Chinese 简体中文: expected length of chunks 1, got ",
				chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "简体中文" {
		for e := chunks.Front(); e != nil; e = e.Next() {
			log.Printf("TestGetChunks4: chunk: %s\n", e.Value.(string))
		}
		t.Error("Expected first element of chunk 简体中文 to be 简体中文, got ",
				chunk)
	}
}

func TestReadText1(t *testing.T) {
	//log.Printf("TestReadText1: Begin ******** \n")	
	text := ReadText("../testdata/sampletest.txt")
	expected := "繁體中文<br/>\n"
	//log.Printf("TestReadText1: Expected  '%s', got '%s'\n", expected, text)	
	if text != expected {
		t.Error("Expected ", expected, ", got ", text)
	}
	//log.Printf("TestReadText1: End ******** \n")	
}

func TestReadText2(t *testing.T) {
	//log.Printf("TestReadText2: Begin ******** \n")	
	text := ReadText("../testdata/test.html")
	if !strings.Contains(text, "繁體中文") {
		t.Error("Expected to contain '繁體中文', got ", text)
	}
	//log.Printf("TestReadText2: End ******** \n")	
}

func TestParseText1(t *testing.T) {
	//log.Printf("TestParseText1: Begin ******** \n")	
	dictionary.ReadDict("../testdata/testwords.txt")
	tokens, results := ParseText("繁體中文", "", corpus.NewCorpusEntry())
	if tokens.Len() != 2 {
		t.Error("Expected to get length 2, got ", tokens.Len())
		first := tokens.Front().Value.(string)
		if first != "繁體" {
			t.Error("Expected to get tokens.Front() 繁體, got ", first)
		}
	}
	if len(results.Vocab) != 2 {
		t.Error("Expected to get len(vocab) = 2, got ", len(results.Vocab))
	}
	if results.WC != 2 {
		t.Error("Expected to get wc = 2, got ", results.WC)
	}
	//log.Printf("TestParseText1: End ******** \n")	
}

func TestParseText2(t *testing.T) {
	//log.Printf("TestParseText2: Begin ******** \n")	
	dictionary.ReadDict("../testdata/testwords.txt")
	tokens, results := ParseText("a繁體中文", "", corpus.NewCorpusEntry())
	if tokens.Len() != 3 {
		t.Error("Expected to get length 3, got ", tokens.Len())
	}
	if results.WC != 2 {
		t.Error("Expected to get wc 2, got ", results.WC)
	}
	first := tokens.Front().Value.(string)
	if first != "a" {
		t.Error("Expected to get tokens.Front() a, got ", first)
	}
	if len(results.Vocab) != 2 {
		t.Error("Expected to get len(vocab) = 2, got ", len(results.Vocab))
	}
	//log.Printf("TestParseText2: End ******** \n")	
}

func TestParseText3(t *testing.T) {
	//log.Printf("TestParseText3: Begin +++++++++++\n")
	dictionary.ReadDict("../testdata/testwords.txt")
	tokens, results := ParseText("前不见古人", "", corpus.NewCorpusEntry())
	if tokens.Len() != 3 {
		t.Error("Expected to get length 3, got ", tokens.Len())
	}
	if results.WC != 3 {
		t.Error("Expected to get wc 3, got ", results.WC)
	}
	expected := []string{"前", "不见", "古人"}
	i := 0
	for e := tokens.Front(); e != nil; e = e.Next() {
		word := e.Value.(string)
		if expected[i] != e.Value.(string) {
			t.Error("Failed to get expected word", expected[i], word, i)
		}
		i++
	}
	//log.Printf("TestParseText3: End +++++++++++\n")
}

// Basic test with no data
func TestSampleUsage1(t *testing.T) {
	//log.Printf("TestSampleUsage1: Begin +++++++++++\n")
	usageMap := map[string]*[]WordUsage{}
	usageMap = sampleUsage(usageMap)
	l := len(usageMap)
	expected := 0
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

// Basic test with minimal data
func TestSampleUsage2(t *testing.T) {
	wu := WordUsage{
		Freq: 1,
		RelFreq: 0.01,
		Word: "大",
		Example: "大蛇",
		File: "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle: "A Big Snake",
	}
	wuArray := []WordUsage{wu}
	usageMap := map[string]*[]WordUsage{"大": &wuArray}
	usageMap = sampleUsage(usageMap)
	l := len(usageMap)
	expected := 1
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

// Basic test with more data
func TestSampleUsage3(t *testing.T) {
	log.Printf("TestSampleUsage3: Begin +++++++++++\n")
	wu1 := WordUsage{
		Freq: 1,
		RelFreq: 0.01,
		Word: "蛇",
		Example: "大蛇",
		File: "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle: "Some Snakes",
	}
	wu2 := WordUsage{
		Freq: 1,
		RelFreq: 0.01,
		Word: "蛇",
		Example: "小蛇",
		File: "afile.txt",
		EntryTitle: "Scroll 2",
		ColTitle: "Some Snakes",
	}
	wuArray := []WordUsage{wu1, wu2}
	usageMap := map[string]*[]WordUsage{"蛇": &wuArray}
	usageMap = sampleUsage(usageMap)
	l := len(*usageMap["蛇"])
	expected := 2
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

// Basic test with more data
func TestSampleUsage4(t *testing.T) {
	log.Printf("TestSampleUsage4: Begin +++++++++++\n")
	wu1 := WordUsage{
		Freq: 1,
		RelFreq: 0.01,
		Word: "大",
		Example: "大蛇",
		File: "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle: "Some Big Animals",
	}
	wu2 := WordUsage{
		Freq: 1,
		RelFreq: 0.01,
		Word: "大",
		Example: "大老虎",
		File: "afile.txt",
		EntryTitle: "Scroll 2",
		ColTitle: "Some Big Animals",
	}
	wu3 := WordUsage{
		Freq: 1,
		RelFreq: 0.01,
		Word: "大",
		Example: "大树",
		File: "afile.txt",
		EntryTitle: "Scroll 1",
		ColTitle: "Some Big Trees",
	}
	wuArray := []WordUsage{wu1, wu2, wu3}
	usageMap := map[string]*[]WordUsage{"大": &wuArray}
	usageMap = sampleUsage(usageMap)
	l := len(*usageMap["大"])
	expected := 3
	if l != expected {
		t.Error("Expected to get length ", expected, ", got ", l)
	}
}

func TestWriteAnalysis(t *testing.T) {
	//log.Printf("TestWriteAnalysis: Begin +++++++++++\n")
	_, results := ParseText("繁", "", corpus.NewCorpusEntry())
	srcFile := "test.txt"
	writeAnalysis(results, srcFile, "Test Collection", "Test Doc")
	//log.Printf("TestWriteAnalysis: End +++++++++++\n")
}

func TestWriteCorpusAll(t *testing.T) {
	//log.Printf("TestWriteCorpusAll: Begin +++++++++++\n")
	WriteCorpusAll()
	//log.Printf("TestWriteCorpusAll: End +++++++++++\n")
}

func TestWriteCorpusCol(t *testing.T) {
	log.Printf("TestWriteCorpusCol: Begin +++++++++++\n")
	WriteCorpusCol("lunyu.csv")
	log.Printf("TestWriteCorpusCol: End +++++++++++\n")
}

func TestWriteCorpusDoc1(t *testing.T) {
	//log.Printf("TestWriteCorpusDoc1: Begin +++++++++++\n")
	tokens, results := ParseText("繁", "", corpus.NewCorpusEntry())
	outfile := "../testoutput/output.html"
	writeCorpusDoc(tokens, results.Vocab, outfile, "", "", "")
	//log.Printf("TestWriteCorpusDoc1: End +++++++++++\n")
}

func TestWriteDoc1(t *testing.T) {
	//log.Printf("TestWriteDoc1: Begin +++++++++++\n")
	tokens, results := ParseText("繁", "", corpus.NewCorpusEntry())
	outfile := "../testoutput/output.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`)
	//log.Printf("TestWriteDoc1: End +++++++++++\n")
}

func TestWriteDoc2(t *testing.T) {
	dictionary.ReadDict("../testdata/testwords.txt")
	text := ReadText("../testdata/test.html")
	tokens, results := ParseText(text, "", corpus.NewCorpusEntry())
	if tokens.Len() != 4 {
		t.Error("Expected to get length 4, got ", tokens.Len())
	}
	outfile := "../testoutput/test-gloss.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`)
}

func TestWriteDoc3(t *testing.T) {
	dictionary.ReadDict("../testdata/testwords.txt")
	text := ReadText("../testdata/test-simplified.html")
	tokens, results := ParseText(text, "", corpus.NewCorpusEntry())
	if tokens.Len() != 6 {
		t.Error("Expected to get length 6, got ", tokens.Len())
	}
	outfile := "../testoutput/test-simplified-gloss.html"
	WriteDoc(tokens, results.Vocab, outfile, `\N`, `\N`)
}


// Test that WriteHwFiles() does not explode
func TestWriteHwFiles(t *testing.T) {
	log.Printf("TestWriteHwFiles: Begin +++++++++++\n")
	log.Printf("TestWriteHwFiles: LUFileName: ", config.LUFileName())
	dictionary.ReadDict(config.LUFileName())
	WriteHwFiles()
	log.Printf("TestWriteHwFiles: End +++++++++++\n")
}

// Test that WordFrequencies() does not explode
func TestWordFrequencies(t *testing.T) {
	//log.Printf("TestWordFrequencies: Begin +++++++++++\n")
	dictionary.ReadDict(config.LUFileName())
	WordFrequencies()
	//log.Printf("TestWordFrequencies: End +++++++++++\n")
}