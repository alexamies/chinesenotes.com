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

// Test package initialization, which requires a database connection
func TestCacheColDetails(t *testing.T) {
	cMap := cacheColDetails()
	title := cMap["wenxuan.html"]
	if title == "" {
		t.Error("TestCacheColDetails: got empty title, map size, ",
			len(cMap))
	}
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
	docSimilarity, err := findBodyBitVector(terms)
	if err == nil {
		t.Error("TestFindInBody0: expected an error, ", err)
		return
	}
	fmt.Printf("TestFindInBody0, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindBodyBitVector1(t *testing.T) {
	terms := []string{"后妃"}
	docSimilarity, err := findBodyBitVector(terms)
	if err == nil {
		t.Error("TestFindBodyBitVector1: expected an error, ", err)
		return
	}
	fmt.Printf("TestFindBodyBitVector1, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBitVector2(t *testing.T) {
	terms := []string{"后妃", "之"}
	docSimilarity, err := findBodyBitVector(terms)
	if err != nil {
		t.Error("TestFindBodyBitVector2: got error, ", err)
	}
	fmt.Printf("TestFindBodyBitVector2, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBitVector3(t *testing.T) {
	terms := []string{"后妃", "之", "德"}
	docSimilarity, err := findBodyBitVector(terms)
	if err != nil {
		t.Error("TestFindBodyBitVector3: got error, ", err)
	}
	fmt.Printf("TestFindBodyBitVector3, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBitVector4(t *testing.T) {
	terms := []string{"后妃", "之", "德", "也"}
	docSimilarity, err := findBodyBitVector(terms)
	if err != nil {
		t.Error("TestFindBodyBitVector4: got error, ", err)
	}
	fmt.Printf("TestFindBodyBitVector4, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBM251(t *testing.T) {
	terms := []string{"后妃"}
	docSimilarity, err := findBodyBM25(terms)
	if err == nil {
		t.Error("TestfindBodyBM251: expected an error, ", err)
		return
	}
	fmt.Printf("TestfindBodyBM251, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBM252(t *testing.T) {
	terms := []string{"后妃", "之"}
	docSimilarity, err := findBodyBM25(terms)
	if err != nil {
		t.Error("TestfindBodyBM251: got error, ", err)
	}
	fmt.Printf("TestfindBodyBM252, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindBodyBM253(t *testing.T) {
	terms := []string{"后妃", "之", "德"}
	docSimilarity, err := findBodyBM25(terms)
	if err != nil {
		t.Error("TestfindBodyBM251: got error, ", err)
	}
	fmt.Printf("TestfindBodyBM251, len(docSimilarity) = %d", len(docSimilarity))
}

func TestFindBodyBM254(t *testing.T) {
	terms := []string{"后妃", "之", "德", "也"}
	docSimilarity, err := findBodyBM25(terms)
	if err != nil {
		t.Error("TestfindBodyBM254: got error, ", err)
	}
	fmt.Printf("TestfindBodyBM254, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBigram1(t *testing.T) {
	terms := []string{"后妃"}
	docSimilarity, err := findBodyBigram(terms)
	if err == nil {
		t.Error("TestFindBodyBigram1: expected an error, ", err)
		return
	}
	fmt.Printf("TestFindBodyBigram1, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBigram2(t *testing.T) {
	terms := []string{"后妃", "之"}
	docSimilarity, err := findBodyBigram(terms)
	if err != nil {
		t.Error("TestFindBodyBigram2: got error, ", err)
	}
	fmt.Printf("TestFindBodyBigram2, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBigram3(t *testing.T) {
	terms := []string{"后妃", "之", "德"}
	docSimilarity, err := findBodyBigram(terms)
	if err != nil {
		t.Error("TestFindBodyBigram2: got error, ", err)
	}
	fmt.Printf("TestFindBodyBigram2, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBigram4(t *testing.T) {
	terms := []string{"后妃", "之", "德", "也"}
	docSimilarity, err := findBodyBigram(terms)
	if err != nil {
		t.Error("TestFindBodyBigram4: got error, ", err)
	}
	fmt.Printf("TestFindBodyBigram4, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBigram5(t *testing.T) {
	terms := []string{"箴", "也", "所以", "攻", "疾"}
	docSimilarity, err := findBodyBigram(terms)
	if err != nil {
		t.Error("TestFindBodyBigram5: got error, ", err)
	}
	fmt.Printf("TestFindBodyBigram5, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindBodyBigram6(t *testing.T) {
	terms := []string{"箴", "也", "所以", "攻", "疾", "防患"}
	docSimilarity, err := findBodyBigram(terms)
	if err != nil {
		t.Error("TestFindBodyBigram6: got error, ", err)
	}
	fmt.Printf("TestFindBodyBigram6, len(docSimilarity) = %d",
		len(docSimilarity))
}

func TestFindDocumentsInCol0(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	_, err := FindDocumentsInCol(parser, "", "wenxuan.html")
	if err == nil {
		t.Error("TestFindDocumentsInCol2: expected error for empty string")
	}
}

func TestFindDocumentsInCol1(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocumentsInCol(parser, "箴", "wenxuan.html")
	if err != nil {
		t.Error("TestFindDocumentsInCol1: got error, ", err)
	}
	if len(qr.Terms) != 1 {
		t.Error("TestFindDocumentsInCol1: len(qr.Terms) != 1, ", qr)
	}
}

func TestFindDocumentsInCol2(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocumentsInCol(parser, "箴也", "wenxuan.html")
	if err != nil {
		t.Error("TestFindDocumentsInCol2: got error, ", err)
	}
	if len(qr.Terms) != 2 {
		t.Error("TestFindDocumentsInCol2: len(qr.Terms) != 2, ", qr)
	}
}

func TestFindDocumentsInCol3(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocumentsInCol(parser, "箴也所", "wenxuan.html")
	if err != nil {
		t.Error("TestFindDocumentsInCol3: got error, ", err)
	}
	if len(qr.Terms) != 3 {
		t.Error("TestFindDocumentsInCol3: len(qr.Terms) != 3, ", qr)
	}
}

func TestFindDocumentsInCol4(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocumentsInCol(parser, "箴也所以", "wenxuan.html")
	if err != nil {
		t.Error("TestFindDocumentsInCol4: got error, ", err)
	}
	if len(qr.Terms) != 4 {
		t.Error("TestFindDocumentsInCol4: len(qr.Terms) != 4, ", qr)
	}
}

func TestFindDocumentsInCol5(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict} // 箴也所以攻疾防患
	qr, err := FindDocumentsInCol(parser, "箴也所以攻", "wenxuan.html")
	if err != nil {
		t.Error("TestFindDocumentsInCol5: got error, ", err)
	}
	if len(qr.Terms) != 5 {
		t.Error("TestFindDocumentsInCol5: len(qr.Terms) != 4, ", qr)
	}
}

func TestFindDocumentsInCol6(t *testing.T) {
	dict := map[string]Word{}
	parser := DictQueryParser{dict}
	qr, err := FindDocumentsInCol(parser, "箴也所以攻疾", "wenxuan.html")
	if err != nil {
		t.Error("TestFindDocumentsInCol6: got error, ", err)
	}
	if len(qr.Terms) != 6 {
		t.Error("TestFindDocumentsInCol6: len(qr.Terms) != 6, ", qr)
	}
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

