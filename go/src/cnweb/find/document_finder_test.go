// Unit tests for find functions
package find

import (
	"fmt"
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

func TestCombineByWeight(t *testing.T) {
	doc := Document{
		GlossFile: "f2.html",
		Title: "Very Good doc",
		SimTitle: 1.0,
		SimWords: 0.5,
		SimBigram: 1.5,
		SimBitVector: 1.0,
	}
	simDoc := combineByWeight(doc)
	if simDoc.Similarity == 0.0 {
		t.Error("TestCombineByWeight: simDoc.Similarity == 0.0")
	}
	fmt.Printf("TestCacheColDetails: simDoc %v\n", simDoc)
	similarity := INTERCEPT + WEIGHT[0] * doc.SimWords +
		WEIGHT[1] * doc.SimBigram + WEIGHT[2] * doc.SimBitVector
	expectedMin := 0.99 * similarity
	expectedMax := 1.01 * similarity
	if ((expectedMin > simDoc.Similarity) || 
		(simDoc.Similarity > expectedMax)) {
		t.Error("TestCombineByWeight: result out of expected range %v\n",
			simDoc)
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

func testFindBodyBM25(terms []string, t *testing.T) {
	docs, err := findBodyBM25(terms)
	if err != nil {
		t.Errorf("testFindBodyBM25: %s got an error, %v\n", terms, err)
	}
	fmt.Printf("testFindBodyBM25, len(docs) = %d\n", len(docs))
	if len(docs) > 0 {
		fmt.Printf("testFindBodyBM25, docs[0] = %v\n", docs[0])
	}
}

func TestFindBodyBM251(t *testing.T) {
	terms := []string{"后妃"}
	testFindBodyBM25(terms, t)
}

func TestFindBodyBM252(t *testing.T) {
	terms := []string{"后妃", "之"}
	testFindBodyBM25(terms, t)
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

func TestMergeDocList1(t *testing.T) {
	simDocMap := map[string]Document{}
	docList := []Document{}
	doc1 := Document{
		GlossFile: "f1.html",
		Title: "Good doc by title",
		SimTitle: 1.0,
	}
	simDocMap[doc1.GlossFile] = doc1
	doc2 := Document{
		GlossFile: "f2.html",
		Title: "Very Good doc",
		SimWords: 0.5,
		SimBigram: 1.5,
	}
	docList = append(docList, doc2)
	mergeDocList(simDocMap, docList)
	if len(simDocMap) != 2 {
		t.Error("TestMergeDocList1: len(simDocMap) != 2, ", len(simDocMap))
	}
	docs := toSortedDocList(simDocMap)
	if len(docs) != 2 {
		t.Error("TestMergeDocList1: len(docs) != 2, ", len(docs))
	}
	expected := doc2.GlossFile
	result := docs[0]
	if result.GlossFile != expected {
		t.Errorf("TestMergeDocList1: expected %s, got, %v, docs: %v", expected,
			result, docs)
	}
}

func TestMergeDocList2(t *testing.T) {
	simDocMap := map[string]Document{}
	docList := []Document{}
	doc1 := Document{
		GlossFile: "f1.html",
		Title: "SAme Very Good doc",
		SimTitle: 1.0,
	}
	simDocMap[doc1.GlossFile] = doc1
	doc2 := Document{
		GlossFile: "f2.html",
		Title: "Reasonable by word frequ",
		SimWords: 1.6,
	}
	doc3 := Document{
		GlossFile: "f1.html",
		Title: "Same Very Good doc",
		SimWords: 1.5,
		SimBigram: 1.5,
	}
	docList = append(docList, doc2)
	docList = append(docList, doc3)
	mergeDocList(simDocMap, docList)
	if len(simDocMap) != 2 {
		t.Error("TestMergeDocList2: len(simDocMap) != 2, ", len(simDocMap))
	}
	docs := toSortedDocList(simDocMap)
	if len(docs) != 2 {
		t.Error("TestMergeDocList2: len(docs) != 2, ", len(docs))
	}
	expected := doc1.GlossFile
	result := docs[0]
	if result.GlossFile != expected {
		t.Errorf("TestMergeDocList2: expected %s, got, %v, docs: %v", expected,
			result, docs)
	}
}

func TestSetContainsTerms1(t *testing.T) {
	terms := []string{"后妃"}
	doc := Document{
		ContainsWords: "后妃",
	}
	doc = setMatchDetails(doc, terms)
	expected0 := "后妃"
	result := doc.ContainsTerms
	if (len(result) != 1) {
		t.Errorf("TestSetContainsTerms1: expected len = 1, got, %d\n",
			len(result))
		return
	}
	if result[0] != expected0 {
		t.Errorf("TestSetContainsTerms1: expected %s, got, %v\n", expected0,
			result)
	}
}

func TestSetContainsTerms2(t *testing.T) {
	terms := []string{"后妃", "之"}
	doc := Document{
		ContainsWords: "之,后妃",
	}
	doc = setMatchDetails(doc, terms)
	expected0 := "后妃"
	expected1 := "之"
	result := doc.ContainsTerms
	if (len(result) != 2) {
		t.Errorf("TestSetContainsTerms2: expected len = 2, got, %d\n",
			len(result))
		return
	}
	if result[0] != expected0 {
		t.Errorf("TestSetContainsTerms2: expected0 %s, got, %s\n", expected0,
			result[0])
	}
	if result[1] != expected1 {
		t.Errorf("TestSetContainsTerms2: expected1 %s, got, %s\n", expected1,
			result[1])
	}
	fmt.Println("TestSetContainsTerms2: ", result)
}

func TestSetContainsTerms3(t *testing.T) {
	terms := []string{"后妃", "之"}
	doc := Document{
		ContainsWords: "之,后妃",
		ContainsBigrams: "后妃之",
	}
	doc = setMatchDetails(doc, terms)
	expected0 := "后妃之"
	result := doc.ContainsTerms
	if (len(result) != 1) {
		t.Errorf("TestSetContainsTerms3: expected len = 1, got, %d\n",
			len(result))
		return
	}
	if result[0] != expected0 {
		t.Errorf("TestSetContainsTerms3: expected0 %s, got, %s\n", expected0,
			result[0])
	}
	fmt.Println("TestSetContainsTerms3: ", result)
}

func TestSetContainsTerms4(t *testing.T) {
	terms := []string{"十年", "之", "計"}
	doc := Document{
		ContainsWords: "十年,之,計",
		ContainsBigrams: "十年之,之計",
	}
	doc = setMatchDetails(doc, terms)
	expected0 := "十年之"
	expected1 := "之計"
	result := doc.ContainsTerms
	if (len(result) != 2) {
		t.Errorf("TestSetContainsTerms4: expected len = 2, got, %d\n",
			len(result))
		return
	}
	if result[0] != expected0 {
		t.Errorf("TestSetContainsTerms4: expected0 %s, got, %s\n", expected0,
			result[0])
	}
	if result[1] != expected1 {
		t.Errorf("TestSetContainsTerms4: expected0 %s, got, %s\n", expected1,
			result[1])
	}
	fmt.Println("TestSetContainsTerms4: ", result)
}

func TestToRelevantDocList(t *testing.T) {
	similarDocMap := map[string]Document{}
	doc1 := Document{
		GlossFile: "f1.html",
		Title: "Good doc",
		Similarity: 1.0,
	}
	similarDocMap[doc1.GlossFile] = doc1
	doc2 := Document{
		GlossFile: "f2.html",
		Title: "Very Good doc",
		Similarity: 1.5,
	}
	similarDocMap[doc2.GlossFile] = doc2
	doc3 := Document{
		GlossFile: "f3.html",
		Title: "Irrelevant doc",
		Similarity: 0.2,
	}
	similarDocMap[doc3.GlossFile] = doc3
	docs := toSortedDocList(similarDocMap)
	queryTerms := []string{}
	docs = toRelevantDocList(docs, queryTerms)
	expected := 2
	result := len(docs)
	if result == expected {
		t.Errorf("TestToRelevantDocList: expected %s, got, %v", expected,
			result)
	}
}

func TestToSortedDocList1(t *testing.T) {
	similarDocMap := map[string]Document{}
	doc1 := Document{
		GlossFile: "f1.html",
		Title: "Good doc",
		SimWords: 1.0,
	}
	similarDocMap[doc1.GlossFile] = doc1
	doc2 := Document{
		GlossFile: "f2.html",
		Title: "Very Good doc",
		SimWords: 1.5,
	}
	similarDocMap[doc2.GlossFile] = doc2
	doc3 := Document{
		GlossFile: "f3.html",
		Title: "Reasonable doc",
		SimWords: 0.5,
	}
	similarDocMap[doc3.GlossFile] = doc3
	docs := toSortedDocList(similarDocMap)
	expected := doc2.GlossFile
	result := docs[0]
	if result.Similarity == 0.0 {
		t.Error("TestToSortedDocList1: result.Similarity == 0.0")
	}
	if result.GlossFile != expected {
		t.Errorf("TestToSortedDocList1: expected %s, got, %v", expected, result)
	}
}

func TestToSortedDocList2(t *testing.T) {
	similarDocMap := map[string]Document{}
	doc1 := Document{
		GlossFile: "f1.html",
		Title: "Good doc by title",
		SimTitle: 1.0,
	}
	similarDocMap[doc1.GlossFile] = doc1
	doc2 := Document{
		GlossFile: "f2.html",
		Title: "Very Good doc",
		SimWords: 0.5,
		SimBigram: 1.5,
	}
	similarDocMap[doc2.GlossFile] = doc2
	doc3 := Document{
		GlossFile: "f3.html",
		Title: "Reasonable doc",
		SimWords: 0.5,
	}
	similarDocMap[doc3.GlossFile] = doc3
	docs := toSortedDocList(similarDocMap)
	expected := doc2.GlossFile
	result := docs[0]
	if result.GlossFile != expected {
		t.Error("TestToSortedDocList2: expected %s, got, %v", expected, result)
	}
}

