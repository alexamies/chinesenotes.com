package analysis

import (
	"fmt"
	"strings"
	"testing"
)

// Both traditional and simplified
func TestIsCJKChar1(t *testing.T) {
	fmt.Printf("TestIsCJKChar1: Begin unit tests\n")
	c := "中"
	result := IsCJKChar(c)
	if !result {
		fmt.Printf("TestIsCJKChar1: Testing Chinese %s\n", c)
		t.Error("Expected true, got ", result)
	}
}

// Non-Chinese
func TestIsCJKChar2(t *testing.T) {
	c := "a"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar2: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Simplified Chinese
func TestIsCJKChar3(t *testing.T) {
	c := "简"
	result := IsCJKChar(c)
	if !result {
		fmt.Printf("TestIsCJKChar3: Testing Chinese %s\n", c)
		t.Error("Expected true, got ", result)
	}
}

// Both traditional and simplified
func TestIsCJKChar4(t *testing.T) {
	c := "古"
	result := IsCJKChar(c)
	if !result {
		fmt.Printf("TestIsCJKChar4: Testing Chinese %s\n", c)
		t.Error("Expected true, got ", result)
	}
}

// Test for punctuation
func TestIsCJKChar5(t *testing.T) {
	c := "。"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar5: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Test for punctuation
func TestIsCJKChar6(t *testing.T) {
	c := "，"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar6: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

// Test for punctuation
func TestIsCJKChar7(t *testing.T) {
	c := "-"
	result := IsCJKChar(c)
	if result {
		fmt.Printf("TestIsCJKChar7: Testing Chinese %s\n", c)
		t.Error("Expected false, got ", result)
	}
}

func TestGetChunks1(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
	chunks := GetChunks("中文")
	if chunks.Len() != 1 {
		t.Error("Expected length of chunks 1, got ", chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "中文" {
		t.Error("Expected first element of chunk 中文, got ", chunk)
	}
}

func TestGetChunks2(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
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
	ReadDict("../testdata/testwords.txt")
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
	ReadDict("../testdata/testwords.txt")
	chunks := GetChunks("简体中文")
	if chunks.Len() != 1 {
		t.Error("Simplified Chinese 简体中文: expected length of chunks 1, got ",
				chunks.Len())
	}
	chunk := chunks.Front().Value.(string)
	if chunk != "简体中文" {
		for e := chunks.Front(); e != nil; e = e.Next() {
			fmt.Printf("TestGetChunks4: chunk: %s\n", e.Value.(string))
		}
		t.Error("Expected first element of chunk 简体中文 to be 简体中文, got ",
				chunk)
	}
}

func TestReadDict1(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
	ws, ok := GetWordSense("中文")
	if !ok {
		t.Error("Expected true, got ", ok)
	}
	if ws.Id != 1 {
		t.Error("Expected 1, got ", ws.Id)
	}
	if ws.Simplified != "中文" {
		t.Error("Expected 中文, got ", ws.Simplified)
	}
	if ws.Traditional != "\\N" {
		t.Error("Expected \\N, got ", ws.Traditional)
	}
	if ws.Pinyin != "Zhōngwén" {
		t.Error("Expected Zhōngwén, got ", ws.Pinyin)
	}
	if ws.English != "Chinese language" {
		t.Error("Expected Chinese language, got ", ws.English)
	}
	if ws.Grammar != "noun" {
		t.Error("Expected noun, got ", ws.Grammar)
	}
	if ws.Concept_cn != "\\N" {
		t.Error("Expected \\N, got ", ws.Concept_cn)
	}
	if ws.Concept_en != "\\N" {
		t.Error("Expected \\N, got ", ws.Concept_en)
	}
	if ws.Topic_cn != "语言" {
		t.Error("Expected 语言, got ", ws.Topic_cn)
	}
	if ws.Topic_en != "Language" {
		t.Error("Expected Language, got ", ws.Topic_en)
	}
	if ws.Parent_cn != "\\N" {
		t.Error("Expected \\N, got ", ws.Parent_cn)
	}
	if ws.Parent_en != "\\N" {
		t.Error("Expected \\N, got ", ws.Parent_en)
	}
	if ws.Image != "\\N" {
		t.Error("Expected \\N, got ", ws.Image)
	}
	if ws.Mp3 != "zhong1wen2.mp3" {
		t.Error("Expected zhong1wen2.mp3, got ", ws.Mp3)
	}
	if ws.Notes != "May refer to" {
		t.Error("Expected 'May refer to', got ", ws.Notes)
	}
}

func TestGetWord(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
	word, ok := GetWord("中")
	if !ok {
		t.Error("Expected true, got ", ok)
	}
	if len(word) != 2 {
		t.Error("Expected len(word) = 2, got ", len(word))
	}
}

func TestReadText1(t *testing.T) {
	text := ReadText("../testdata/sampletest.txt")
	expected := "繁體中文<br/>\n"
	//fmt.Printf("TestReadText1: Expected  '%s', got '%s'\n", expected, text)	
	if text != expected {
		t.Error("Expected ", expected, ", got ", text)
	}
}

func TestReadText2(t *testing.T) {
	text := ReadText("../testdata/test.html")
	if !strings.Contains(text, "繁體中文") {
		t.Error("Expected to contain '繁體中文', got ", text)
	}
}

func TestParseText1(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
	tokens, vocab, wc, _ := ParseText("繁體中文")
	if tokens.Len() != 2 {
		t.Error("Expected to get length 2, got ", tokens.Len())
		first := tokens.Front().Value.(string)
		if first != "繁體" {
			t.Error("Expected to get tokens.Front() 繁體, got ", first)
		}
	}
	if len(vocab) != 2 {
		t.Error("Expected to get len(vocab) = 2, got ", len(vocab))
		_, ok := wdict["繁體"]
		if !ok {
			t.Error("Expected to find vocab entry for 繁體 but did not.")
		}
	}
	if wc != 2 {
		t.Error("Expected to get wc = 2, got ", wc)
	}
}

func TestParseText2(t *testing.T) {
	//fmt.Printf("TestParseText2: Begin +++++++++++\n")
	ReadDict("../testdata/testwords.txt")
	tokens, vocab, wc, _ := ParseText("a繁體中文")
	if tokens.Len() != 3 {
		t.Error("Expected to get length 3, got ", tokens.Len())
	}
	if wc != 2 {
		t.Error("Expected to get wc 2, got ", wc)
	}
	first := tokens.Front().Value.(string)
	if first != "a" {
		t.Error("Expected to get tokens.Front() a, got ",first)
	}
	if len(vocab) != 2 {
		t.Error("Expected to get len(vocab) = 2, got ", len(vocab))
		_, ok := wdict["繁體"]
		if !ok {
			t.Error("Expected to find vocab entry for 繁體 but did not.")
		}
	}
}

func TestParseText3(t *testing.T) {
	//fmt.Printf("TestParseText2: Begin +++++++++++\n")
	ReadDict("../testdata/testwords.txt")
	tokens, _, wc, _ := ParseText("前不见古人")
	if tokens.Len() != 3 {
		t.Error("Expected to get length 3, got ", tokens.Len())
	}
	if wc != 3 {
		t.Error("Expected to get wc 3, got ", wc)
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
}

func TestWriteAnalysis(t *testing.T) {
	_, vocab, wc, _ := ParseText("繁")
	unknownChars := []string{"獢"}
	srcFile := "test.txt"
	WriteAnalysis(vocab, wc, unknownChars, srcFile, "Test Collection", "Test Doc")
}

func TestWriteCorpusDoc1(t *testing.T) {
	tokens, vocab, _, _ := ParseText("繁")
	outfile := "../testoutput/output.html"
	WriteCorpusDoc(tokens, vocab, outfile, "", "", "")
}

func TestWriteDoc1(t *testing.T) {
	tokens, vocab, _, _ := ParseText("繁")
	outfile := "../testoutput/output.html"
	WriteDoc(tokens, vocab, outfile)
}

func TestWriteDoc2(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
	text := ReadText("../testdata/test.html")
	tokens, vocab, _, _ := ParseText(text)
	if tokens.Len() != 4 {
		t.Error("Expected to get length 4, got ", tokens.Len())
	}
	outfile := "../testoutput/test-gloss.html"
	WriteDoc(tokens, vocab, outfile)
}

func TestWriteDoc3(t *testing.T) {
	ReadDict("../testdata/testwords.txt")
	text := ReadText("../testdata/test-simplified.html")
	tokens, vocab, _, _ := ParseText(text)
	if tokens.Len() != 6 {
		t.Error("Expected to get length 6, got ", tokens.Len())
	}
	outfile := "../testoutput/test-simplified-gloss.html"
	WriteDoc(tokens, vocab, outfile)
}