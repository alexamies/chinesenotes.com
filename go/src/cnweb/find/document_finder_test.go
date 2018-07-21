// Unit tests for find functions
package find

import (
	"fmt"
	"sort"
	"testing"
)

// Test package initialization, which requires a database connection
func TestInit(t *testing.T) {
	fmt.Printf("TestInit: Begin unit tests\n")
}

func TestFindDocuments1(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocuments(parser, "Assembly", false)
	if err != nil {
		t.Error("TestFindDocuments1: got error, ", err)
	}
	if len(qr.Terms) != 1 {
		t.Error("TestFindDocuments1: len(qr.Terms) != 1, ", qr)
	}
}

func TestFindDocuments2(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	_, err := FindDocuments(parser, "", false)
	if err == nil {
		t.Error("TestFindDocuments2: expected error for empty string")
	}
}

func TestFindDocuments3(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocuments(parser, "hello", false)
	if err != nil {
		t.Error("TestFindDocuments3: got error, ", err)
	}
	if len(qr.Terms) != 1 {
		t.Error("TestFindDocuments3: len(qr.Terms) != 1, ", qr)
	}
	if len(qr.Terms[0].Senses) == 0 {
		t.Error("TestFindDocuments3: len(qr.Terms.Senses) == 0, ", qr)
	}
}

func TestFindInBody0(t *testing.T) {
	terms := []string{}
	docSimilarity, err := findInBody(terms)
	if err == nil {
		t.Error("TestFindInBody0: expected an error, ", err)
		return
	}
	fmt.Printf("TestFindInBody0, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody1(t *testing.T) {
	terms := []string{"后妃"}
	docSimilarity, err := findInBody(terms)
	if err == nil {
		t.Error("TestFindInBody0: expected an error, ", err)
		return
	}
	fmt.Printf("TestFindInBody1, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody2(t *testing.T) {
	terms := []string{"后妃", "之"}
	docSimilarity, err := findInBody(terms)
	if err != nil {
		t.Error("TestFindInBody: got error, ", err)
	}
	fmt.Printf("TestFindInBody2, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody3(t *testing.T) {
	terms := []string{"后妃", "之", "德"}
	docSimilarity, err := findInBody(terms)
	if err != nil {
		t.Error("TestFindInBody: got error, ", err)
	}
	fmt.Printf("TestFindInBody3, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindInBody4(t *testing.T) {
	terms := []string{"后妃", "之", "德", "也"}
	docSimilarity, err := findInBody(terms)
	if err != nil {
		t.Error("TestFindInBody: got error, ", err)
	}
	fmt.Printf("TestFindInBody4, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindWords1(t *testing.T) {
	words, err := findWords("Assembly")
	if err != nil {
		t.Error("TestFindWords1: got error, ", err)
	}
	if len(words) != 0 {
		t.Error("TestFindWords1: len(words) != 0, ", len(words))
	}
}

func TestFindWords2(t *testing.T) {
	words, err := findWords("金剛")
	if err != nil {
		t.Error("TestFindWords2: got error, ", err)
	}
	if len(words) != 1 {
		t.Error("TestFindWords2: len(words) != 1, ", len(words))
	}
}

func TestSimilarDocSorting(t *testing.T) {
	doc1 := SimilarDoc{
		GlossFile: "f1.html",
		Title: "Good doc",
		Similarity: 1.0,
	}
	doc2 := SimilarDoc{
		GlossFile: "f2.html",
		Title: "Very Good doc",
		Similarity: 1.5,
	}
	doc3 := SimilarDoc{
		GlossFile: "f3.html",
		Title: "Reasonable doc",
		Similarity: 0.5,
	}
	docs := []SimilarDoc{doc1, doc2, doc3}
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].Similarity > docs[j].Similarity
	})
	fmt.Println(docs)

}

