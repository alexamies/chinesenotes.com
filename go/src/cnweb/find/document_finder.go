/* 
Functions for finding collections by partial match on collection title
*/
package find

import (
	"cnweb/applog"
	"database/sql"
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"sort"
	"time"
	"cnweb/webconfig"
)

var (
	countColStmt, countDocStmt *sql.Stmt
	database *sql.DB
	colMap map[string]string
	docMap map[string]Document
	findAllTitlesStmt, findAllColTitlesStmt  *sql.Stmt
	findColStmt, findDocStmt, findWordStmt  *sql.Stmt
	simBitVector2Stmt, simBitVector3Stmt, simBitVector4Stmt *sql.Stmt
	simTFIDF2Stmt, simTFIDF3Stmt, simTFIDF4Stmt *sql.Stmt
	simBM252Stmt, simBM253Stmt, simBM254Stmt, simBM255Stmt, simBM256Stmt *sql.Stmt
	simBigram1Stmt, simBigram2Stmt, simBigram3Stmt, simBigram4Stmt *sql.Stmt
	simBigram5Stmt *sql.Stmt
)

type Collection struct {
	GlossFile, Title string
}

type DocSimilarity struct {
	Similarity float64
	Collection, Document string
}

type Document struct {
	GlossFile, Title, CollectionFile, CollectionTitle string
}

type QueryResults struct {
	NumCollections, NumDocuments int
	Collections []Collection
	Documents []Document
	Terms []TextSegment
}

// Structure remembering how similar a document is to another
type SimilarDoc struct {
	CollectionFile, CollectionTitle, GlossFile, Title string
	Similarity float64
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
		var gloss_file, title, col_gloss_file, col_title string
		results.Scan(&gloss_file, &title, &col_gloss_file, &col_title)
		document := Document{gloss_file, title, col_gloss_file, col_title}
		docMap[gloss_file] = document
	}
	applog.Info("cacheDocDetails, len(docMap) = ", len(docMap))
	return docMap
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

// Search the corpus for document bodies using bit vector similarity.
//  Param: terms - The decomposed query string with 1 < num elements < 5
func findBodyBitVector(terms []string) ([]DocSimilarity, error) {
	applog.Info("findBitVector, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findBodyBitVector, len(terms) < 2", len(terms))
		return []DocSimilarity{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		results, err = simBitVector2Stmt.QueryContext(ctx, terms[0], terms[1])
	} else if len(terms) == 3 {
		results, err = simBitVector3Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2])
	}  else {
		// Ignore arguments beyond the first four
		results, err = simBitVector4Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2], terms[3])
	}
	if err != nil {
		applog.Error("findBodyBitVector, Error for query: ", terms, err)
		return []DocSimilarity{}, err
	}
	simSlice := []DocSimilarity{}
	for results.Next() {
		docSim := DocSimilarity{}
		results.Scan(&docSim.Similarity, &docSim.Collection, &docSim.Document)
		applog.Info("findBodyBitVector, Similarity, Document = ",
			docSim.Similarity, docSim.Collection, docSim.Document)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

// Search the corpus for document bodies most similar using a BM25 model.
//  Param: terms - The decomposed query string with 1 < num elements < 7
func findBodyBM25(terms []string) ([]DocSimilarity, error) {
	applog.Info("findBodyBM25, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findBodyBM25, len(terms) < 2", len(terms))
		return []DocSimilarity{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		results, err = simBM252Stmt.QueryContext(ctx, terms[0], terms[1])
	} else if len(terms) == 3 {
		results, err = simBM253Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2])
	}  else if len(terms) == 4 {
		results, err = simBM254Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2], terms[3])
	}  else if len(terms) == 5 {
		results, err = simBM255Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2], terms[3], terms[4])
	}  else {
		// Ignore arguments beyond the first six
		results, err = simBM256Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2], terms[3], terms[4], terms[5])
	}
	if err != nil {
		applog.Error("findBodyBM25, Error for query: ", terms, err)
		return []DocSimilarity{}, err
	}
	simSlice := []DocSimilarity{}
	for results.Next() {
		docSim := DocSimilarity{}
		results.Scan(&docSim.Similarity, &docSim.Collection, &docSim.Document)
		applog.Info("findBodyBM25, Similarity, Document = ", docSim.Similarity,
			docSim.Collection, docSim.Document)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

// Search the corpus for document bodies most similar using bigrams with a BM25
// model.
//  Param: terms - The decomposed query string with 1 < num elements < 7
func findBodyBigram(terms []string) ([]DocSimilarity, error) {
	applog.Info("findBodyBM25, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findBodyBigram, len(terms) < 2", len(terms))
		return []DocSimilarity{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		bigram1 := terms[0] + terms[1]
		results, err = simBigram1Stmt.QueryContext(ctx, bigram1)
	} else if len(terms) == 3 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		results, err = simBigram2Stmt.QueryContext(ctx, bigram1, bigram2)
	}  else if len(terms) == 4 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		results, err = simBigram3Stmt.QueryContext(ctx, bigram1, bigram2,
			bigram3)
	}  else if len(terms) == 5 {
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		bigram4 := terms[3] + terms[4]
		results, err = simBigram4Stmt.QueryContext(ctx, bigram1, bigram2,
			bigram3, bigram4)
	}  else {
		// Ignore arguments beyond the first six
		bigram1 := terms[0] + terms[1]
		bigram2 := terms[1] + terms[2]
		bigram3 := terms[2] + terms[3]
		bigram4 := terms[3] + terms[4]
		bigram5 := terms[4] + terms[5]
		results, err = simBigram5Stmt.QueryContext(ctx, bigram1, bigram2,
			bigram3, bigram4, bigram5)
	}
	if err != nil {
		applog.Error("simBigram4Stmt, Error for query: ", terms, err)
		return []DocSimilarity{}, err
	}
	simSlice := []DocSimilarity{}
	for results.Next() {
		docSim := DocSimilarity{}
		results.Scan(&docSim.Similarity, &docSim.Collection, &docSim.Document)
		applog.Info("findBodyBM25, Similarity, Document = ", docSim.Similarity,
			docSim.Collection, docSim.Document)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
}

// Search the corpus for document bodies most similar using a TF-IDF model.
//  Param: terms - The decomposed query string with 1 < num elements < 5
func findBodyTFIDF(terms []string) ([]DocSimilarity, error) {
	applog.Info("findBodyTFIDF, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findBodyTFIDF, len(terms) < 2", len(terms))
		return []DocSimilarity{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		results, err = simTFIDF2Stmt.QueryContext(ctx, terms[0], terms[1])
	} else if len(terms) == 3 {
		results, err = simTFIDF3Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2])
	}  else {
		// Ignore arguments beyond the first four
		results, err = simTFIDF4Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2], terms[3])
	}
	if err != nil {
		applog.Error("findBodyTFIDF, Error for query: ", terms, err)
		return []DocSimilarity{}, err
	}
	simSlice := []DocSimilarity{}
	for results.Next() {
		docSim := DocSimilarity{}
		results.Scan(&docSim.Similarity, &docSim.Collection, &docSim.Document)
		applog.Info("findBodyTFIDF, Similarity, Document = ",
			docSim.Similarity, docSim.Collection, docSim.Document)
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
		results.Scan(&doc.Title, &doc.GlossFile)
		documents = append(documents, doc)
	}
	return documents, nil
}

// Find documents by both title and contents, and merge the lists
func findDocuments(query string, terms []TextSegment, advanced bool) ([]Document, error) {
	applog.Info("findDocuments, terms: ", terms)
	docs, err := findDocsByTitle(query)
	applog.Info("findDocuments, len(docs): ", len(docs))
	if err != nil {
		return nil, err
	}
	if len(terms) < 2 {
		return docs, nil
	}
	queryTerms := []string{}
	for _, term := range terms {
		queryTerms = append(queryTerms, term.QueryText)
	}
	if (!advanced) {
		return docs, nil
	}

	// For more than one term find docs that are similar body and merge
	docMap := toSimilarDocMap(docs) // similarity = 1.0
	//simDocs, err := findBodyBitVector(queryTerms)
	//simDocs, err := findBodyTFIDF(queryTerms)
	simDocs, err := findBodyBM25(queryTerms)
	if err != nil {
		return nil, err
	}
	mergedDocs := mergeBySimilarity(docMap, simDocs)
	moreDocs, err := findBodyBigram(queryTerms)
	if err != nil {
		return nil, err
	}
	mergedDocsMap := toSimilarDocMap(mergedDocs) // similarity = 1.0
	allMergedDocs := mergeBySimilarity(mergedDocsMap, moreDocs)
	// findBodyBigram
	applog.Info("findDocuments, len(allMergedDocs): ", len(allMergedDocs))
	return allMergedDocs, nil
}

// Returns a QueryResults object containing matching collections, documents,
// and dictionary words. For dictionary lookup, a text segment will
// contains the QueryText searched for and possibly a matching
// dictionary entry. There will only be matching dictionary entries for 
// Chinese words in the dictionary. If there are no Chinese words in the query
// then the Chinese word senses matching the English or Pinyin will be included
// in the TextSegment.Senses field.
func FindDocuments(parser QueryParser, query string, advanced bool) (QueryResults, error) {
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
	nDoc := len(documents)
	if err != nil {
		// Got an error, see if we can connect and try again
		if hello() {
			documents, err = findDocuments(query, terms, advanced)
		} // else do not try again, giveup and return the error
	}
	applog.Info("FindDocuments, query, nTerms, collection, doc count: ", query,
		len(terms), nCol, nDoc)
	return QueryResults{nCol, nDoc, collections, documents, terms}, err
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
	stmt, err := database.PrepareContext(ctx,
		"SELECT title, gloss_file FROM collection WHERE title LIKE ? LIMIT 50")
    if err != nil {
        applog.Error("find.initStatements() Error preparing collection stmt: ",
        	err)
        return err
    }
    findColStmt = stmt

	cstmt, err := database.PrepareContext(ctx,
		"SELECT count(title) FROM collection WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.initStatements() Error preparing cstmt: ",err)
        return err
    }
    countColStmt = cstmt

	dstmt, err := database.PrepareContext(ctx,
		"SELECT title, gloss_file " +
		"FROM document " +
		"WHERE col_plus_doc_title LIKE ? LIMIT 50")
    if err != nil {
        applog.Error("find.initStatements() Error preparing dstmt: ", err)
        return err
    }
    findDocStmt = dstmt

	cdstmt, err := database.PrepareContext(ctx,
		"SELECT count(title) FROM document WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.initStatements() Error preparing cDocStmt: ", err)
        return err
    }
    countDocStmt = cdstmt    

	fwstmt, err := database.PrepareContext(ctx, 
		"SELECT simplified, traditional, pinyin, headword FROM words WHERE " +
		"simplified = ? OR traditional = ? LIMIT 1")
    if err != nil {
        applog.Error("find.init() Error preparing fwstmt: ", err)
        return err
    }
    findWordStmt = fwstmt

    // For a query with two terms in the query string decomposition
	sim2Stmt, err := database.PrepareContext(ctx, 
		"SELECT COUNT(frequency) / 2.0 AS similarity, collection, document " +
		"FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBitVector2Stmt: ",
        	err)
        return err
    }
    simBitVector2Stmt = sim2Stmt

    // For a query with three terms in the query string decomposition
	sim3Stmt, err := database.PrepareContext(ctx, 
		"SELECT COUNT(frequency) / 3.0 AS similarity, collection, document " +
		"FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBitVector3Stmt: ",
        	err)
        return err
    }
    simBitVector3Stmt = sim3Stmt

    // For a query with four terms in the query string decomposition
	sim4Stmt, err := database.PrepareContext(ctx, 
		"SELECT COUNT(frequency) / 4.0 AS similarity, collection, document " +
		"FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBitVector4Stmt: ",
        	err)
        return err
    }
    simBitVector4Stmt = sim4Stmt

    // For a query with two terms in the query string decomposition
	simTF2Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(frequency * idf) AS similarity, collection, document " +
		"FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simTFIDF2Stmt: ",
        	err)
        return err
    }
    simTFIDF2Stmt = simTF2Stmt

    // For a query with three terms in the query string decomposition
	simTF3Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(frequency * idf) AS similarity, collection, document " +
		"FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simTFIDF3Stmt: ",
        	err)
        return err
    }
    simTFIDF3Stmt = simTF3Stmt

    // For a query with four terms in the query string decomposition
	simTF4Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(frequency * idf) AS similarity, collection, document " +
		"FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simTFIDF4Stmt: ",
        	err)
        return err
    }
    simTFIDF4Stmt = simTF4Stmt

    // Document similarity with BM25 using 2-6 terms, k = 1.5, b = 0
	simBM2Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document FROM word_freq_doc " +
		"WHERE word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBM252Stmt: ", err)
        return err
    }
    simBM252Stmt = simBM2Stmt

	simBM3Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBM253Stmt: ", err)
        return err
    }
    simBM253Stmt = simBM3Stmt

	simBM4Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBM254Stmt: ", err)
        return err
    }
    simBM254Stmt = simBM4Stmt


	simBM5Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBM255Stmt: ", err)
        return err
    }
    simBM255Stmt = simBM5Stmt

	simBM6Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document FROM word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? OR word = ? " +
		"OR word = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBM256Stmt: ", err)
        return err
    }
    simBM256Stmt = simBM6Stmt

    // Document similarity with Bigram using 1-6 bigrams, k = 1.5, b = 0
	simBg1Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBigram1Stmt: ",
        	err)
        return err
    }
    simBigram1Stmt = simBg1Stmt

	simBg2Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBM252Stmt: ", err)
        return err
    }
    simBigram2Stmt = simBg2Stmt

	simBg3Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBigram3Stmt: ",
        	err)
        return err
    }
    simBigram3Stmt = simBg3Stmt

	simBg4Stmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? OR bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBigram4Stmt: ",
        	err)
        return err
    }
    simBigram4Stmt = simBg4Stmt


	simBgStmt, err := database.PrepareContext(ctx, 
		"SELECT SUM(2.5 * frequency * idf / (frequency + 1.5)) AS similarity, " +
		"collection, document " +
		"FROM bigram_freq_doc " +
		"WHERE bigram = ? OR bigram = ? OR bigram = ? OR bigram = ? " +
		"OR bigram = ? " +
		"GROUP BY collection, document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.initStatements() Error preparing simBigram5Stmt: ",
        	err)
        return err
    }
    simBigram5Stmt = simBgStmt

    // Find the titles of all documents
	fAllTitlesStmt, err := database.PrepareContext(ctx, 
		"SELECT gloss_file, title, col_gloss_file, col_title " +
		"FROM document LIMIT 1000000")
    if err != nil {
        applog.Error("find.initStatements() Error preparing findAllTitlesStmt: ",
        	err)
        return err
    }
    findAllTitlesStmt = fAllTitlesStmt

    // Find the titles of all documents
	fAllColTitlesStmt, err := database.PrepareContext(ctx, 
		"SELECT gloss_file, title FROM collection LIMIT 100000")
    if err != nil {
        applog.Error("find.initStatements() Error preparing findAllColTitlesStmt: ",
        	err)
        return err
    }
    findAllColTitlesStmt = fAllColTitlesStmt

    return nil
}

// Merge a list of documents with map of similar docs, adding the similarity
// for docs that are in both lists
func mergeBySimilarity(simDocMap map[string]SimilarDoc, docList []DocSimilarity) []Document {
	for _, simDoc := range docList {
		sDoc, ok := simDocMap[simDoc.Document]
		if ok {
			sDoc.Similarity += simDoc.Similarity
		} else {
			colTitle, ok1 := colMap[simDoc.Collection]
			document, ok2 := docMap[simDoc.Document]
			if (ok1 && ok2) {
				doc := SimilarDoc{simDoc.Collection, colTitle, simDoc.Document,
					document.Title, simDoc.Similarity}
				simDocMap[simDoc.Document] = doc
			} else {
				applog.Info("mergeBySimilarity, col or doc title not found: ",
					ok1, ok2, simDoc.Collection,simDoc.Document)
			}
		}
	}
	return toSortedDocList(simDocMap)
}

// Convert list to a map of similar docs with similarity set to 1.0
func toSimilarDocMap(docs []Document) map[string]SimilarDoc {
	similarDocMap := map[string]SimilarDoc{}
	for _, doc  := range docs {
		simDoc := SimilarDoc{
			GlossFile: doc.GlossFile,
			Title: doc.Title,
			Similarity: 1.0,
		}
		similarDocMap[doc.GlossFile] = simDoc
	}
	return similarDocMap
}

// Convert a map of similar docs into a sorted list
func toSortedDocList(similarDocMap map[string]SimilarDoc) []Document {
	similarDocs := []SimilarDoc{}
	for _, similarDoc  := range similarDocMap {
		similarDocs = append(similarDocs, similarDoc)
	}
	sort.Slice(similarDocs, func(i, j int) bool {
		return similarDocs[i].Similarity > similarDocs[j].Similarity
	})
	docs := []Document{}
	for _, similarDoc := range similarDocs {
		doc := Document{similarDoc.GlossFile, similarDoc.Title,
			similarDoc.CollectionFile, similarDoc.CollectionTitle}
		docs = append(docs, doc)
	}
	return docs
}