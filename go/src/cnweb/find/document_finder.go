/* 
Functions for finding collections by partial match on collection title
*/
package find

import (
	"cnweb/applog"
	"database/sql"
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sort"
	"strings"
	"time"

	"cnweb/fulltext"
	"cnweb/webconfig"
)

const (
	MAX_RETURNED = 50
	MIN_SIMILARITY = 0.07
	AVG_DOC_LEN = 4497
	INTERCEPT = -4.75 // From logistic regression
)

var (
	countColStmt *sql.Stmt
	database *sql.DB
	colMap map[string]string
	docMap map[string]Document
	findAllTitlesStmt, findAllColTitlesStmt  *sql.Stmt
	findColStmt, findDocStmt, findDocInColStmt, findWordStmt  *sql.Stmt
	simBM251Stmt, simBM252Stmt, simBM253Stmt, simBM254Stmt *sql.Stmt
	simBM255Stmt, simBM256Stmt *sql.Stmt
	simBM25Col1Stmt, simBM25Col2Stmt, simBM25Col3Stmt, simBM25Col4Stmt *sql.Stmt
	simBM25Col5Stmt, simBM25Col6Stmt *sql.Stmt
	simBigram1Stmt, simBigram2Stmt, simBigram3Stmt, simBigram4Stmt *sql.Stmt
	simBigram5Stmt *sql.Stmt
	simBgCol1Stmt, simBgCol2Stmt, simBgCol3Stmt, simBgCol4Stmt *sql.Stmt
	simBgCol5Stmt *sql.Stmt
	//  From logistic regression
	WEIGHT = []float64{0.080, 2.327, 3.040} // [BM25 words, BM25 bigrams, bit vector]
    avdl int // The average document length
)

type Collection struct {
	GlossFile, Title string
}

type Document struct {
	GlossFile, Title, CollectionFile, CollectionTitle, ContainsWords string
	ContainsBigrams string
	SimTitle, SimWords, SimBigram, SimBitVector, Similarity float64
	ContainsTerms []string
	MatchDetails fulltext.MatchingText
}

type QueryResults struct {
	Query, CollectionFile string
	NumCollections, NumDocuments int
	Collections []Collection
	Documents []Document
	Terms []TextSegment
}

// Open database connection and prepare statements
func init() {
	err := initStatements()
	if err != nil {
		applog.Error("find/init: error preparing database statements, retrying",
			err)
		time.Sleep(60000 * time.Millisecond)
		err = initStatements()
		conString := webconfig.DBConfig()
		applog.Fatal("find/init: error preparing database statements, giving up",
			conString, err)
	}
	result := hello() 
	if !result {
		conString := webconfig.DBConfig()
		applog.Fatal("find/init: got error with findWords ", conString, err)
	}
	docMap = cacheDocDetails()
	colMap = cacheColDetails()

	avdl = webconfig.GetEnvIntValue("AVG_DOC_LEN", AVG_DOC_LEN)
}

// For printing out retrieved document metadata
func (doc Document) String() string {
    return fmt.Sprintf("%s, %s, SimTitle %f, SimWords %f, SimBigram %f, " +
    	"SimBitVector %f, Similarity %f, ContainsWords %s, ContainsBigrams %s" +
    	", MatchDetails %v", 
    	doc.GlossFile, doc.CollectionFile, doc.SimTitle, doc.SimWords,
    	doc.SimBigram, doc.SimBitVector, doc.Similarity, doc.ContainsWords,
    	doc.ContainsBigrams, doc.MatchDetails)
 }

// Cache the details of all collecitons by target file name
func cacheColDetails() map[string]string {
	colMap = map[string]string{}
	ctx := context.Background()
	results, err := findAllColTitlesStmt.QueryContext(ctx)
	if err != nil {
		applog.Error("cacheColDetails, Error for query: ", err)
		return colMap
	}
	defer results.Close()

	for results.Next() {
		var gloss_file, title string
		results.Scan(&gloss_file, &title)
		colMap[gloss_file] = title
	}
	applog.Info("cacheColDetails, len(colMap) = ", len(colMap))
	return colMap
}

// Cache the details of all documents by target file name
func cacheDocDetails() map[string]Document {
	docMap = map[string]Document{}
	ctx := context.Background()
	results, err := findAllTitlesStmt.QueryContext(ctx)
	if err != nil {
		applog.Error("cacheDocDetails, Error for query: ", err)
		return docMap
	}
	defer results.Close()

	for results.Next() {
		doc := Document{}
		results.Scan(&doc.GlossFile, &doc.Title, &doc.CollectionFile,
			&doc.CollectionTitle)
		docMap[doc.GlossFile] = doc
	}
	applog.Info("cacheDocDetails, len(docMap) = ", len(docMap))
	return docMap
}

// Combine tHE similarity of title, word frequency, and bigrams into a single
// weighted measure
func combineByWeight(doc Document) Document {
	similarity := INTERCEPT + WEIGHT[0] * doc.SimWords +
		WEIGHT[1] * doc.SimBigram + WEIGHT[2] * doc.SimBitVector
	simDoc := Document{
		GlossFile: doc.GlossFile,
		Title: doc.Title,
		CollectionFile: doc.CollectionFile,
		CollectionTitle: doc.CollectionTitle,
		SimTitle: doc.SimTitle,
		SimWords: doc.SimWords,
		SimBigram: doc.SimBigram,
		SimBitVector: doc.SimBitVector,
		Similarity: similarity,
		ContainsWords: doc.ContainsWords,
		ContainsBigrams: doc.ContainsBigrams,
		ContainsTerms: doc.ContainsTerms,
		MatchDetails: doc.MatchDetails,
	}
	return simDoc
}

func countCollections(query string) int {
	var count int
	ctx := context.Background()
	results, err := countColStmt.QueryContext(ctx, "%" + query + "%")
	results.Next()
	results.Scan(&count)
	if err != nil {
		applog.Error("countCollections: Error for query: ", query, err)
	}
	results.Close()
	return count
}

// Search the corpus for document bodies most similar using a BM25 model.
//  Param: terms - The decomposed query string with 0 < num elements < 7
func findBodyBM25(terms []string) ([]Document, error) {
	applog.Info("findBodyBM25, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) == 1 {
		results, err = simBM251Stmt.QueryContext(ctx, avdl, terms[0])
	} else if len(terms) == 2 {
		results, err = simBM252Stmt.QueryContext(ctx, avdl, terms[0], terms[1])
	} else if len(terms) == 3 {
		results, err = simBM253Stmt.QueryContext(ctx, avdl, terms[0], terms[1],
			terms[2])
	}  else if len(terms) == 4 {
		results, err = simBM254Stmt.QueryContext(ctx, avdl, terms[0], terms[1],
			terms[2], terms[3])
	}  else if len(terms) == 5 {
		results, err = simBM255Stmt.QueryContext(ctx, avdl, terms[0], terms[1],
			terms[2], terms[3], terms[4])
	}  else {
		// Ignore arguments beyond the first six
		results, err = simBM256Stmt.QueryContext(ctx, avdl, terms[0], terms[1],
			terms[2], terms[3], terms[4], terms[5])
	}
	if err != nil {
		applog.Error("findBodyBM25, Error for query: ", terms, err)
		return []Document{}, err
	}
	simSlice := []Document{}
	for results.Next() {
		docSim := Document{}
		results.Scan(&docSim.SimWords, &docSim.SimBitVector,
			&docSim.ContainsWords, &docSim.CollectionFile, &docSim.GlossFile)
		//applog.Info("findBodyBM25, Similarity, Document = ", docSim)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

// Search the corpus for document bodies most similar using a BM25 model in a
// specific collection.
//  Param: terms - The decomposed query string with 1 < num elements < 7
func findBodyBM25InCol(terms []string,
		col_gloss_file string) ([]Document, error) {
	applog.Info("findBodyBM25InCol, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) == 1 {
		results, err = simBM25Col1Stmt.QueryContext(ctx, avdl, terms[0],
			col_gloss_file)
	} else if len(terms) == 2 {
		results, err = simBM25Col2Stmt.QueryContext(ctx, avdl, terms[0],
			terms[1], col_gloss_file)
	} else if len(terms) == 3 {
		results, err = simBM25Col3Stmt.QueryContext(ctx, avdl, terms[0],
			terms[1], terms[2], col_gloss_file)
	}  else if len(terms) == 4 {
		results, err = simBM25Col4Stmt.QueryContext(ctx, avdl, terms[0],
			terms[1], terms[2], terms[3], col_gloss_file)
	}  else if len(terms) == 5 {
		results, err = simBM25Col5Stmt.QueryContext(ctx, avdl, terms[0],
			terms[1], terms[2], terms[3], terms[4], col_gloss_file)
	}  else {
		// Ignore arguments beyond the first six
		results, err = simBM25Col6Stmt.QueryContext(ctx, avdl, terms[0],
			terms[1], terms[2], terms[3], terms[4], terms[5],
			col_gloss_file)
	}
	if err != nil {
		applog.Error("findBodyBM25InCol, Error for query: ", terms, err)
		return []Document{}, err
	}
	simSlice := []Document{}
	for results.Next() {
		docSim := Document{}
		docSim.CollectionFile = col_gloss_file
		results.Scan(&docSim.SimWords, &docSim.SimBitVector,
			&docSim.ContainsWords, &docSim.GlossFile)
		//applog.Info("findBodyBM25InCol, Similarity, Document = ", docSim)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

// Search the corpus for document bodies most similar using bigrams with a BM25
// model.
//  Param: terms - The decomposed query string with 1 < num elements < 7
func findBodyBigram(terms []string) ([]Document, error) {
	applog.Info("findBodyBigram, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findBodyBigram, len(terms) < 2", len(terms))
		return []Document{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		bigram1 := terms[0] + terms[1]
		results, err = simBigram1Stmt.QueryContext(ctx, avdl, bigram1)
	} else if len(terms) == 3 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		results, err = simBigram2Stmt.QueryContext(ctx, avdl, bigram1, bigram2)
	}  else if len(terms) == 4 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		results, err = simBigram3Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			bigram3)
	}  else if len(terms) == 5 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		bigram4 := terms[3] + terms[4]
		results, err = simBigram4Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			bigram3, bigram4)
	}  else {
		// Ignore arguments beyond the first six
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		bigram4 := terms[3] + terms[4]
		bigram5 := terms[4] + terms[5]
		results, err = simBigram5Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			bigram3, bigram4, bigram5)
	}
	if err != nil {
		applog.Error("findBodyBigram, Error for query: ", terms, err)
		return []Document{}, err
	}
	simSlice := []Document{}
	for results.Next() {
		docSim := Document{}
		results.Scan(&docSim.SimBigram, &docSim.ContainsBigrams,
			&docSim.CollectionFile, &docSim.GlossFile)
		//applog.Info("findBodyBigram, Similarity, Document = ", docSim)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

// Search the corpus for document bodies most similar using bigrams with a BM25
// model within a specific collection
//  Param: terms - The decomposed query string with 1 < num elements < 7
func findBodyBgInCol(terms []string,
		col_gloss_file string) ([]Document, error) {
	applog.Info("findBodyBgInCol, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findBodyBgInCol, len(terms) < 2", len(terms))
		return []Document{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		bigram1 := terms[0] + terms[1]
		results, err = simBgCol1Stmt.QueryContext(ctx, avdl, bigram1,
			col_gloss_file)
	} else if len(terms) == 3 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		results, err = simBgCol2Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			col_gloss_file)
	}  else if len(terms) == 4 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		results, err = simBgCol3Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			bigram3, col_gloss_file)
	}  else if len(terms) == 5 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		bigram4 := terms[3] + terms[4]
		results, err = simBgCol4Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			bigram3, bigram4, col_gloss_file)
	}  else {
		// Ignore arguments beyond the first six
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		bigram4 := terms[3] + terms[4]
		bigram5 := terms[4] + terms[5]
		results, err = simBgCol5Stmt.QueryContext(ctx, avdl, bigram1, bigram2,
			bigram3, bigram4, bigram5, col_gloss_file)
	}
	if err != nil {
		applog.Error("findBodyBgInCol, Error for query: ", terms, err)
		return []Document{}, err
	}
	simSlice := []Document{}
	for results.Next() {
		docSim := Document{}
		docSim.CollectionFile = col_gloss_file
		results.Scan(&docSim.SimBigram, &docSim.ContainsBigrams,
			&docSim.GlossFile)
		//applog.Info("findBodyBgInCol, Similarity, Document = ", docSim)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

func findCollections(query string) []Collection {
	ctx := context.Background()
	results, err := findColStmt.QueryContext(ctx, "%" + query + "%")
	if err != nil {
		applog.Error("findCollections, Error for query: ", query, err)
	}
	defer results.Close()

	collections := []Collection{}
	for results.Next() {
		col := Collection{}
		results.Scan(&col.Title, &col.GlossFile)
		collections = append(collections, col)
	}
	return collections
}

// Find documents based on a match in title
func findDocsByTitle(query string) ([]Document, error) {
	ctx := context.Background()
	results, err := findDocStmt.QueryContext(ctx, "%" + query + "%")
	if err != nil {
		applog.Error("findDocsByTitle, Error for query: ", query, err)
		return nil, err
	}
	defer results.Close()

	documents := []Document{}
	for results.Next() {
		doc := Document{}
		results.Scan(&doc.Title, &doc.GlossFile, &doc.CollectionFile,
			&doc.CollectionTitle)
		doc.SimTitle = 1.0
		documents = append(documents, doc)
	}
	return documents, nil
}

// Find documents based on a match in title within a specific collection
func findDocsByTitleInCol(query, col_gloss_file string) ([]Document, error) {
	ctx := context.Background()
	results, err := findDocInColStmt.QueryContext(ctx, "%" + query + "%",
		col_gloss_file)
	if err != nil {
		applog.Error("findDocsByTitleInCol, Error for query: ", query, err)
		return nil, err
	}
	defer results.Close()

	documents := []Document{}
	for results.Next() {
		doc := Document{}
		doc.CollectionFile = col_gloss_file
		results.Scan(&doc.Title, &doc.GlossFile, &doc.CollectionTitle)
		doc.SimTitle = 1.0
		//applog.Info("findDocsByTitleInCol, doc: ", doc)
		documents = append(documents, doc)
	}
	return documents, nil
}

// Find documents by both title and contents, and merge the lists
func findDocuments(query string, terms []TextSegment,
		advanced bool) ([]Document, error) {
	applog.Info("findDocuments, enter: ", query)
	docs, err := findDocsByTitle(query)
	if err != nil {
		return nil, err
	}
	applog.Info("findDocuments, len(docs): ", query, len(docs))
	queryTerms := toQueryTerms(terms)
	if (!advanced) {
		return docs, nil
	}

	// For more than one term find docs that are similar body and merge
	docMap := toSimilarDocMap(docs) // similarity = 1.0
	applog.Info("findDocuments, len(docMap): ", query, len(docMap))
	simDocs, err := findBodyBM25(queryTerms)
	if err != nil {
		return nil, err
	}
	mergeDocList(docMap, simDocs)

	// If less than 2 terms then do not need to check bigrams
	if len(terms) < 2 {
		sortedDocs := toSortedDocList(docMap)
		applog.Info("findDocuments, < 2 len(sortedDocs): ", query, 
			len(sortedDocs))
		relevantDocs := toRelevantDocList(sortedDocs, queryTerms)
		return relevantDocs, nil
	}
	moreDocs, err := findBodyBigram(queryTerms)
	if err != nil {
		return nil, err
	}
	mergeDocList(docMap, moreDocs)
	sortedDocs := toSortedDocList(docMap)
	applog.Info("findDocuments, len(sortedDocs): ", query, len(sortedDocs))
	relevantDocs := toRelevantDocList(sortedDocs, queryTerms)
	applog.Info("findDocuments, len(relevantDocs): ", query, len(relevantDocs))
	return relevantDocs, nil
}

// Find documents in a specific collection by both title and contents, and
// merge the lists
func findDocumentsInCol(query string, terms []TextSegment,
		col_gloss_file string) ([]Document, error) {
	applog.Info("findDocumentsInCol, col_gloss_file, terms: ", col_gloss_file,
		terms)
	docs, err := findDocsByTitleInCol(query, col_gloss_file)
	if err != nil {
		return nil, err
	}
	applog.Info("findDocumentsInCol, len(docs) by title: ", len(docs))
	//applog.Info("findDocumentsInCol, docs array by title: ", docs)
	queryTerms := toQueryTerms(terms)

	// For more than one term find docs that are similar body and merge
	docMap := toSimilarDocMap(docs) // similarity = 1.0
	//simDocs, err := findBodyBitVector(queryTerms)
	simDocs, err := findBodyBM25InCol(queryTerms, col_gloss_file)
	if err != nil {
		return nil, err
	}
	//applog.Info("findDocumentsInCol, len(simDocs) by word freq: ", len(simDocs))
	mergeDocList(docMap, simDocs)

	if len(terms) > 1 {
		// If there are 2 or more terms then check bigrams
		simBGDocs, err := findBodyBgInCol(queryTerms, col_gloss_file)
		//applog.Info("findDocumentsInCol, len(simBGDocs) ", len(simBGDocs))
		if err != nil {
			applog.Info("findDocumentsInCol, findBodyBgInCol error: ", err)
			return nil, err
		}
		mergeDocList(docMap, simBGDocs)
	}
	sortedDocs := toSortedDocList(docMap)
	applog.Info("findDocumentsInCol, len(sortedDocs): ", len(sortedDocs))
	relevantDocs := toRelevantDocList(sortedDocs, queryTerms)
	applog.Info("findDocuments, len(relevantDocs): ", query, len(relevantDocs))
	return relevantDocs, nil
}

// Returns a QueryResults object containing matching collections, documents,
// and dictionary words. For dictionary lookup, a text segment will
// contains the QueryText searched for and possibly a matching
// dictionary entry. There will only be matching dictionary entries for 
// Chinese words in the dictionary. If there are no Chinese words in the query
// then the Chinese word senses matching the English or Pinyin will be included
// in the TextSegment.Senses field.
func FindDocuments(parser QueryParser, query string,
		advanced bool) (QueryResults, error) {
	if query == "" {
		applog.Error("FindDocuments, Empty query string")
		return QueryResults{}, errors.New("Empty query string")
	}
	terms := parser.ParseQuery(query)
	if (len(terms) == 1) && (terms[0].DictEntry.HeadwordId == 0) {
	    applog.Info("FindDocuments,Query does not contain Chinese, look for " +
	    	"English and Pinyin matches: ", query)
		senses, err := findWordsByEnglish(terms[0].QueryText)
		if err != nil {
			return QueryResults{}, err
		} else {
			terms[0].Senses = senses
		}
	}
	nCol := countCollections(query)
	collections := findCollections(query)
	documents, err := findDocuments(query, terms, advanced)
	if err != nil {
		// Got an error, see if we can connect and try again
		if hello() {
			documents, err = findDocuments(query, terms, advanced)
		} // else do not try again, giveup and return the error
	}
	nDoc := len(documents)
	applog.Info("FindDocuments, query, nTerms, collection, doc count: ", query,
		len(terms), nCol, nDoc)
	return QueryResults{query, "", nCol, nDoc, collections, documents, terms}, err
}

// Returns a QueryResults object containing matching collections, documents,
// and dictionary words within a specific collecion.
// For dictionary lookup, a text segment will
// contains the QueryText searched for and possibly a matching
// dictionary entry. There will only be matching dictionary entries for 
// Chinese words in the dictionary. If there are no Chinese words in the query
// then the Chinese word senses matching the English or Pinyin will be included
// in the TextSegment.Senses field.
func FindDocumentsInCol(parser QueryParser, query,
		col_gloss_file string) (QueryResults, error) {
	if query == "" {
		applog.Error("FindDocumentsInCol, Empty query string")
		return QueryResults{}, errors.New("Empty query string")
	}
	terms := parser.ParseQuery(query)
	if (len(terms) == 1) && (terms[0].DictEntry.HeadwordId == 0) {
	    applog.Info("FindDocumentsInCol, Query does not contain Chinese, " +
	    	"look for English and Pinyin matches: ", query)
		senses, err := findWordsByEnglish(terms[0].QueryText)
		if err != nil {
			return QueryResults{}, err
		} else {
			terms[0].Senses = senses
		}
	}
	documents, err := findDocumentsInCol(query, terms, col_gloss_file)
	if err != nil {
		// Got an error, see if we can connect and try again
		if hello() {
			documents, err = findDocumentsInCol(query, terms, col_gloss_file)
		} // else do not try again, giveup and return the error
	}
	nDoc := len(documents)
	applog.Info("FindDocumentsInCol, query, nTerms, collection, doc count: ", query,
		len(terms), 1, nDoc)
	return QueryResults{query, col_gloss_file, 1, nDoc, []Collection{}, documents, terms}, err
}

// Returns the headword words in the query (only a single word based on Chinese
// query)
func findWords(query string) ([]Word, error) {
	ctx := context.Background()
	results, err := findWordStmt.QueryContext(ctx, query, query)
	if err != nil {
		applog.Error("findWords, Error for query: ", query, err)
		// Sleep for a while, reinitialize, and retry
		time.Sleep(2000 * time.Millisecond)
		initStatements()
		results, err = findWordStmt.QueryContext(ctx, query, query)
		if err != nil {
			applog.Error("findWords, Give up after retry: ", query, err)
			return []Word{}, err
		}
	}
	words := []Word{}
	for results.Next() {
		word := Word{}
		var hw sql.NullInt64
		var trad sql.NullString
		results.Scan(&word.Simplified, &trad, &word.Pinyin, &hw)
		applog.Info("findWords, simplified, headword = ", word.Simplified, hw)
		if trad.Valid {
			word.Traditional = trad.String
		}
		if hw.Valid {
			word.HeadwordId = int(hw.Int64)
		}
		words = append(words, word)
	}
	return words, nil
}

func hello() bool {
	words, err := findWords("你好")
	if err != nil {
		conString := webconfig.DBConfig()
		applog.Error("find/hello: got error with findWords ", conString, err)
		return false
	}
	if len(words) != 1 {
		applog.Error("find/hello: could not find my word ", len(words))
		return false
	} 
	applog.Info("find/hello: Ready to go")
	return true
}

func initStatements() error {
	conString := webconfig.DBConfig()
	db, err := sql.Open("mysql", conString)
	if err != nil {
		return err
	}
	database = db

	ctx := context.Background()
	findColStmt, err = database.PrepareContext(ctx,
		"SELECT title, gloss_file FROM collection WHERE title LIKE ? LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error preparing collection stmt: ",
        	err)
        return err
    }

	countColStmt, err = database.PrepareContext(ctx,
		"SELECT count(title) FROM collection WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.initStatements() Error preparing cstmt: ",err)
        return err
    }

    // Search documents by title substring
	findDocStmt, err = database.PrepareContext(ctx,
		"SELECT title, gloss_file, col_gloss_file, col_title " +
		"FROM document " +
		"WHERE col_plus_doc_title LIKE ? LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error preparing dstmt: ", err)
        return err
    }

    // Search documents by title substring within a collection
	findDocInColStmt, err = database.PrepareContext(ctx,
		"SELECT title, gloss_file, col_title " +
		"FROM document " +
		"WHERE col_plus_doc_title LIKE ? " +
		"AND col_gloss_file = ? " +
		"LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error preparing dstmt: ", err)
        return err
    }

	findWordStmt, err = database.PrepareContext(ctx, 
		"SELECT simplified, traditional, pinyin, headword FROM words WHERE " +
		"simplified = ? OR traditional = ? LIMIT 1")
    if err != nil {
        applog.Error("find.init() Error preparing fwstmt: ", err)
        return err
    }

    // Document similarity with BM25 using 1-6 terms, k = 1.5, b = 0.65
	simBM251Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" collection, document " +
		"FROM word_freq_doc " +
		"WHERE word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM251Stmt: ", err)
        return err
    }

	simBM252Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 2.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" collection, document " +
		"FROM word_freq_doc " +
		"WHERE word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM252Stmt: ", err)
        return err
    }

	simBM253Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 3.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" collection, document " +
		"FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM253Stmt: ", err)
        return err
    }

	simBM254Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 4.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" collection, document " +
		"FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM254Stmt: ", err)
        return err
    }


	simBM255Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 5.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" collection, document " +
		"FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM255Stmt: ", err)
        return err
    }

	simBM256Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 5.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" collection, document " +
		"FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? OR word = ? " +
		"OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM256Stmt: ", err)
        return err
    }

    // Document similarity with BM25 using 2-6 terms, for a specific collection
	simBM25Col1Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" SUM(2.5 * frequency * idf / (frequency + 1.5)) AS bm25, " +
		" COUNT(frequency) / 1.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" document " +
		"FROM word_freq_doc " +
		"WHERE " +
		" (word = ?) AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM25Col1Stmt: ", err)
        return err
    }

	simBM25Col2Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 2.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" document " +
		"FROM word_freq_doc " +
		"WHERE (word = ? OR word = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM252Stmt: ", err)
        return err
    }

	simBM25Col3Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 3.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" document " +
		"FROM word_freq_doc " +
		"WHERE (word = ? OR word = ? OR word = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM253Stmt: ", err)
        return err
    }

	simBM25Col4Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 4.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" document " +
		"FROM word_freq_doc " +
		"WHERE (word = ? OR word = ? OR word = ? OR word = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM254Stmt: ", err)
        return err
    }

	simBM25Col5Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 5.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" document " +
		"FROM word_freq_doc " +
		"WHERE (word = ? OR word = ? OR word = ? OR word = ? OR word = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM255Stmt: ", err)
        return err
    }

	simBM25Col6Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" COUNT(frequency) / 5.0 AS bitvector, " +
		" GROUP_CONCAT(word) AS contains_words, " +
		" document " +
		"FROM word_freq_doc " +
		"WHERE (word = ? OR word = ? OR word = ? OR word = ? OR word = ? " +
		"OR word = ?) " +
		"AND collection = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM256Stmt: ", err)
        return err
    }

    // Document similarity with Bigram using 1-6 bigrams, k = 1.5, b = 0
	simBigram1Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBigram1Stmt: ", err)
        return err
    }

	simBigram2Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBM252Stmt: ", err)
        return err
    }

	simBigram3Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBigram3Stmt: ", err)
        return err
    }

	simBigram4Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? OR bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBigram4Stmt: ", err)
        return err
    }

	simBigram5Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? OR bigram = ? " +
		"OR bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBigram5Stmt: ", err)
        return err
    }

    // Document similarity with Bigram using 1-6 bigrams, within a specific
    // collection
	simBgCol1Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBgCol1Stmt: ", err)
        return err
    }

	simBgCol2Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" document " +
		"FROM bigram_freq_doc " +
		"WHERE (bigram = ? OR bigram = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBgCol2Stmt: ", err)
        return err
    }

	simBgCol3Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBgCol3Stmt: ", err)
        return err
    }

	simBgCol4Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" document " +
		"FROM bigram_freq_doc " +
		"WHERE (bigram = ? OR bigram = ? OR bigram = ? OR bigram = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBgCol4Stmt: ", err)
        return err
    }

	simBgCol5Stmt, err = database.PrepareContext(ctx, 
		"SELECT " +
		" SUM((1.5 + 1) * frequency * idf / " +
		"  (frequency + 1.5 * (1 - 0.65 + 0.65 * (doc_len / ?)))) AS bm25, " +
		" GROUP_CONCAT(bigram) AS contains_bigrams, " +
		" document " +
		"FROM bigram_freq_doc " +
		"WHERE (bigram = ? OR bigram = ? OR bigram = ? OR bigram = ? " +
		"OR bigram = ?) " +
		"AND collection = ? " +
		"GROUP BY document " +
		"ORDER BY bm25 DESC LIMIT 500")
    if err != nil {
        applog.Error("find.initStatements() Error for simBgCol5Stmt: ", err)
        return err
    }

    // Find the titles of all documents
	findAllTitlesStmt, err = database.PrepareContext(ctx, 
		"SELECT gloss_file, title, col_gloss_file, col_title " +
		"FROM document LIMIT 5000000")
    if err != nil {
        applog.Error("find.initStatements() Error for findAllTitlesStmt: ", err)
        return err
    }

    // Find the titles of all documents
	findAllColTitlesStmt, err = database.PrepareContext(ctx, 
		"SELECT gloss_file, title FROM collection LIMIT 500000")
    if err != nil {
        applog.Error("find.initStatements() Error for findAllColTitlesStmt: ",
        	err)
        return err
    }

    return nil
}

// Merge a list of documents with map of similar docs, adding the similarity
// for docs that are in both lists
func mergeDocList(simDocMap map[string]Document, docList []Document) {
	for _, simDoc := range docList {
		sDoc, ok := simDocMap[simDoc.GlossFile]
		if ok {
			sDoc.SimTitle += simDoc.SimTitle
			sDoc.SimWords += simDoc.SimWords
			sDoc.SimBigram += simDoc.SimBigram
			sDoc.SimBitVector += simDoc.SimBitVector
			if sDoc.ContainsWords == "" {
				sDoc.ContainsWords = simDoc.ContainsWords
			} else {
				sDoc.ContainsWords += "," + simDoc.ContainsWords
			}
			if sDoc.ContainsBigrams == "" {
				sDoc.ContainsBigrams = simDoc.ContainsBigrams
			} else {
				sDoc.ContainsBigrams += "," + simDoc.ContainsBigrams
			}
			simDocMap[simDoc.GlossFile] = sDoc
		} else {
			colTitle, ok1 := colMap[simDoc.CollectionFile]
			document, ok2 := docMap[simDoc.GlossFile]
			if (ok1 && ok2) {
				doc := Document{CollectionFile: simDoc.CollectionFile,
								CollectionTitle: colTitle, 
								GlossFile: simDoc.GlossFile,
								Title: document.Title, 
								SimTitle: simDoc.SimTitle,
								SimWords: simDoc.SimWords,
								SimBigram: simDoc.SimBigram,
								SimBitVector: simDoc.SimBitVector,
								Similarity: simDoc.Similarity,
								ContainsWords: simDoc.ContainsWords,
								ContainsBigrams: simDoc.ContainsBigrams,
							}
				simDocMap[simDoc.GlossFile] = doc
			} else if ok2 {
				applog.Info("mergeDocList, collection title not found: ",
					simDoc)
				doc := Document{CollectionFile: "",
								CollectionTitle: "", 
								GlossFile: simDoc.GlossFile,
								Title: document.Title, 
								SimTitle: simDoc.SimTitle,
								SimWords: simDoc.SimWords,
								SimBigram: simDoc.SimBigram,
								SimBitVector: simDoc.SimBitVector,
								Similarity: simDoc.Similarity,
								ContainsWords: simDoc.ContainsWords,
								ContainsBigrams: simDoc.ContainsBigrams,
							}
				simDocMap[simDoc.GlossFile] = doc
			} else {
				applog.Info("mergeDocList, doc title not found: ", simDoc)
				simDocMap[simDoc.GlossFile] = simDoc
			}
		}
	}
}

// Organizes the contains terms found of the document in a form that helps
// the user.
// 
// doc.ContainsWords is a contained list of terms found in the query and doc
// doc.ContainsBigrams is a contained list of bigrams found in the query and doc
// doc.ContainsTerms is a list of terms found both in the query and the doc
// sorted in the same order as the query terms with words merged to bigrams
func setMatchDetails(doc Document, terms []string) Document {
	fmt.Println("sortContainsWords: ", terms)
	containsTems := []string{}
	for i, term := range terms {
		fmt.Printf("sortContainsWords: i = %d\n", i)
		bigram := ""
		if (i > 0) {
			bigram = terms[i - 1] + terms[i]
		}
		if (i > 0) && strings.Contains(doc.ContainsBigrams, bigram) {
			j := len(containsTems)
			if (j > 0) && strings.Contains(bigram, containsTems[j - 1]) {
				containsTems[j-1] = bigram
			} else {
				containsTems = append(containsTems, bigram)
			}
		} else if strings.Contains(doc.ContainsWords, term) {
			containsTems = append(containsTems, term)
		}
	}
	doc.ContainsTerms = containsTems
	doc.MatchDetails = fulltext.GetMatch(doc.GlossFile, terms)
	return doc
}

// Filter documents that are not similar
func toRelevantDocList(docs []Document, terms []string) []Document {
	if len(docs) < 1 {
		return docs
	}
	max := docs[0].Similarity
	relDocs := []Document{}
	for _, doc  := range docs {
		scaledSimilarity := doc.Similarity / max
		// If doc is less than min scaled similarity then we are done
		applog.Info("toRelevantDocList, check: ", scaledSimilarity, 
			MIN_SIMILARITY)
		doc = setMatchDetails(doc, terms)
		if scaledSimilarity < MIN_SIMILARITY {
			return relDocs
		}
		relDocs = append(relDocs, doc)
	}
	return relDocs
}

// Convert list to a map of similar docs with similarity set to 1.0
func toSimilarDocMap(docs []Document) map[string]Document {
	similarDocMap := map[string]Document{}
	for _, doc  := range docs {
		simDoc := Document{
			GlossFile: doc.GlossFile,
			Title: doc.Title,
			CollectionFile: doc.CollectionFile,
			CollectionTitle: doc.CollectionTitle,
			SimTitle: doc.SimTitle,
			SimWords: doc.SimWords,
			SimBigram: doc.SimBigram,
			SimBitVector: doc.SimBitVector,
			ContainsWords: doc.ContainsWords,
			ContainsBigrams: doc.ContainsBigrams,
			Similarity: doc.Similarity,
		}
		similarDocMap[doc.GlossFile] = simDoc
	}
	return similarDocMap
}

// Convert a map of similar docs into a sorted list, and truncate
func toSortedDocList(similarDocMap map[string]Document) []Document {
	docs := []Document{}
	for _, similarDoc  := range similarDocMap {
		simDoc := combineByWeight(similarDoc)
		docs = append(docs, simDoc)
	}
	sort.Slice(docs, func(i, j int) bool {
		return docs[i].Similarity > docs[j].Similarity
	})
	if len(docs) > MAX_RETURNED {
		return docs[:MAX_RETURNED]
	}
	return docs
}