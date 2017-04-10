
package dictionary

import (
	"sort"
	"testing"
)

// make example data
func makeHW0() HeadwordDef {
	simp := "国"
	trad := "國"
	pinyin := "guó"
	wsArray := []WordSenseEntry{}
	return HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// make example data
func makeHW1() HeadwordDef {
	simp := "严净"
	trad := "嚴淨"
	pinyin := "yán jìng"
	wsArray := []WordSenseEntry{}
	return HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// make example data
func makeHW2() HeadwordDef {
	simp := "缘缘"
	trad := "緣緣"
	pinyin := "yuányuán"
	wsArray := []WordSenseEntry{}
	return HeadwordDef{
		Id: 1,
		Simplified: &simp,
		Traditional: &trad,
		Pinyin: []string{pinyin},
		WordSenses: &wsArray,
	}
}

// Trival test for headword sorting
func TestHeadwords0(t *testing.T) {
	hws := Headwords{}
	sort.Sort(hws)
}

// Easy test for headword sorting
func TestHeadwords1(t *testing.T) {
	hw0 := makeHW0()
	hw1 := makeHW1()
	hws := Headwords{hw1, hw0}
	sort.Sort(hws)
	firstWord := hws[0].Pinyin[0]
	pinyinExpected := hw0.Pinyin[0]
	if pinyinExpected != firstWord {
		t.Error("dictionary.TestHeadwords1: Expected pinyin ", pinyinExpected,
			", got", firstWord)
	}
}

// Better test for headword sorting
func TestHeadwords2(t *testing.T) {
	hw0 := makeHW0()
	hw1 := makeHW1()
	hw2 := makeHW2()
	hws := Headwords{hw2, hw1, hw0}
	sort.Sort(hws)
	firstWord := hws[0].Pinyin[0]
	pinyinExpected := hw0.Pinyin[0]
	if pinyinExpected != firstWord {
		t.Error("dictionary.TestHeadwords2: Expected pinyin ", pinyinExpected,
			", got", firstWord)
	}
	secondWord := hws[1].Pinyin[0]
	secondExpected := hw1.Pinyin[0]
	if secondExpected != secondWord {
		t.Error("dictionary.TestHeadwords2: 2nd expected pinyin ",
			secondExpected,	", got", secondWord)
	}
}

// Test removal of tones from Pinyin
func TestNormalizePinyin0(t *testing.T) {
	pinyin := "guó"
	noTones := normalizePinyin(pinyin)
	expected := "guo"
	if expected != noTones {
		t.Error("dictionary.TestNormalizePinyin0: expected noTones ",
			expected, ", got", noTones)
	}
}

// Test removal of tones from Pinyin
func TestNormalizePinyin1(t *testing.T) {
	pinyin := "Sān Bǎo"
	noTones := normalizePinyin(pinyin)
	expected := "san bao"
	if expected != noTones {
		t.Error("dictionary.TestNormalizePinyin1: expected noTones ",
			expected, ", got", noTones)
	}
}

// Test removal of tones from Pinyin
func TestNormalizePinyin2(t *testing.T) {
	pinyin := "Ēmítuó"
	noTones := normalizePinyin(pinyin)
	expected := "emituo"
	if expected != noTones {
		t.Error("dictionary.TestNormalizePinyin1: expected noTones ",
			expected, ", got", noTones)
	}
}