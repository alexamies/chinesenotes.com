// Unit tests for query parsing functions
package find

import (
	"cnweb/dictionary"
	"log"
	"testing"
)

// Test trivial query with empty dictionary
func TestParseChinese0(t *testing.T) {
	log.Printf("TestParseChinese: Begin unit tests\n")
	dict := map[string]dictionary.Word{}
	parser := DictQueryParser{dict}
	s1 := "小"
	query := s1
	terms := parser.parse_chinese(query)
	if len(terms) != 1 {
		t.Error("TestParseChinese0: len(terms) != 1: ", terms)
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseChinese0: terms[0] != s1: ", s1, terms)
		return
	}
}

// Test simple query with empty dictionary
func TestParseChinese1(t *testing.T) {
	log.Printf("TestParseChinese: Begin unit tests\n")
	dict := map[string]dictionary.Word{}
	parser := DictQueryParser{dict}
	s1 := "小"
	s2 := "王"
	query := s1 + s2
	terms := parser.parse_chinese(query)
	if len(terms) != 2 {
		t.Error("TestParseChinese1: len(terms) != 2: ", terms)
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseChinese1: terms[0] != s1: ", s1, terms)
		return
	}
	if terms[1].QueryText != s2 {
		t.Error("TestParseChinese1: terms[1] != s2: ", s2, terms)
		return
	}
}

// Test simple query with non-empty dictionary
func TestParseChinese2(t *testing.T) {
	log.Printf("TestParseChinese: Begin unit tests\n")
	dict := map[string]dictionary.Word{}
	s1 := "小"
	w := dictionary.Word{}
	w.Simplified = s1
	w.Traditional = "\\N"
	w.Pinyin = "xiǎo"
	w.HeadwordId = 42
	dict["小"] = w
	parser := DictQueryParser{dict}
	s2 := "王"
	query := s1 + s2
	terms := parser.parse_chinese(query)
	if len(terms) != 2 {
		t.Error("TestParseChinese2: len(terms) != 2: ", terms)
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseChinese2: terms[0] != s1: ", s1, terms)
		return
	}
	if terms[1].QueryText != s2 {
		t.Error("TestParseChinese2: terms[1] != s2: ", s2, terms)
		return
	}
}

// Test less simple query with non-empty dictionary
func TestParseChinese3(t *testing.T) {
	log.Printf("TestParseChinese: Begin unit tests\n")
	dict := map[string]dictionary.Word{}
	s1 := "你好"
	w := dictionary.Word{}
	w.Simplified = s1
	w.Traditional = "\\N"
	w.Pinyin = "nǐhǎo"
	w.HeadwordId = 42
	dict["你好"] = w
	parser := DictQueryParser{dict}
	s2 := "小"
	s3 := "王"
	query := s1 + s2 + s3
	terms := parser.parse_chinese(query)
	if len(terms) != 3 {
		t.Error("TestParseChinese2: len(terms) != 2: ", terms)
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseChinese2: terms[0] != s1: ", s1, terms)
		return
	}
	if terms[1].QueryText != s2 {
		t.Error("TestParseChinese2: terms[1] != s2: ", s2, terms)
		return
	}
}

// Test less simple query, including punctuation, with non-empty dictionary
func TestParseChinese4(t *testing.T) {
	log.Printf("TestParseChinese: Begin unit tests\n")
	dict := map[string]dictionary.Word{}
	s1 := "你好"
	w := dictionary.Word{}
	w.Simplified = s1
	w.Traditional = "\\N"
	w.Pinyin = "nǐhǎo"
	w.HeadwordId = 42
	dict["你好"] = w
	parser := DictQueryParser{dict}
	s2 := "，"
	s3 := "小"
	s4 := "王"
	s5 := "！"
	query := s1 + s2 + s3 + s4 + s5
	terms := parser.parse_chinese(query)
	if len(terms) != 5 {
		t.Error("TestParseChinese2: len(terms) != 2: ", terms)
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseChinese2: terms[0] != s1: ", s1, terms)
		return
	}
	if terms[1].QueryText != s2 {
		t.Error("TestParseChinese2: terms[1] != s2: ", s2, terms)
		return
	}
}

// Test empty query
func TestParseQuery0(t *testing.T) {
	dict := map[string]dictionary.Word{}
	parser := DictQueryParser{dict}
	terms := parser.ParseQuery("")
	if len(terms) != 0 {
		t.Error("TestParseQuery0: len(terms) != 0: ", len(terms))
	}
}

// Test simple English query
func TestParseQuery1(t *testing.T) {
	query := "hello"
	dict := map[string]dictionary.Word{}
	parser := DictQueryParser{dict}
	terms := parser.ParseQuery(query)
	if len(terms) != 1 {
		t.Error("TestParseQuery1: len(terms) != 1: ", len(terms))
		return
	}
	if terms[0].QueryText != query {
		t.Error("TestParseQuery1: terms[0] != query: ", query, terms[0])
		return
	}
}

// Test simple English query
func TestParseQuery2(t *testing.T) {
	s1 := "Hello"
	s2 := "王"
	query := s1 + s2
	dict := map[string]dictionary.Word{}
	parser := DictQueryParser{dict}
	terms := parser.ParseQuery(query)
	if len(terms) != 2 {
		t.Error("TestParseQuery2: len(terms) != 2: ", len(terms))
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseQuery2: terms[0] != s1: ", s1, terms)
		return
	}
	if terms[1].QueryText != s2 {
		t.Error("TestParseQuery2: terms[1] != s2: ", s2, terms)
		return
	}
}

// Test simple English query
func TestParseQuery3(t *testing.T) {
	s1 := "Hello"
	s2 := "小"
	s3 := "王"
	query := s1 + s2 + s3
	dict := map[string]dictionary.Word{}
	parser := DictQueryParser{dict}
	terms := parser.ParseQuery(query)
	if len(terms) != 3 {
		t.Error("TestParseQuery3: len(terms) != 3: ", terms)
		return
	}
	if terms[0].QueryText != s1 {
		t.Error("TestParseQuery3: terms[0] != s1: ", s1, terms[0])
		return
	}
	if terms[1].QueryText != s2 {
		t.Error("TestParseQuery3: terms[1] != s2: ", s2, terms[1])
		return
	}
	if terms[2].QueryText != s3 {
		t.Error("TestParseQuery3: terms[1] != s2: ", s2, terms[2])
		return
	}
}