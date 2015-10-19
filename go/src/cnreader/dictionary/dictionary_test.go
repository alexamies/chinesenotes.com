package dictionary

import (
	"cnreader/config"
	"fmt"
	"testing"
)

func init() {
	config.SetProjectHome("../../../..")
}

func TestWriteHeadwords(t *testing.T) {
	fmt.Printf("TestWriteHeadwords: Begin +++++++++++\n")
	ReadDict(config.ProjectHome() + "/data/words.txt")
	WriteHeadwords()
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
