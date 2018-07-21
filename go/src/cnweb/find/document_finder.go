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
	docMap map[string]string
	findAllTitlesStmt, findColStmt, findDocStmt, findWordStmt  *sql.Stmt
	similarity2Stmt, similarity3Stmt, similarity4Stmt *sql.Stmt
)

type Collection struct {
	GlossFile, Title string
}

type DocSimilarity struct {
	Similarity float64
	Document string
}

type Document struct {
	GlossFile, Title string
}

type QueryResults struct {
	NumCollections, NumDocuments int
	Collections []Collection
	Documents []Document
	Terms []TextSegment
}

// Structure remembering how similar a document is to another
type SimilarDoc struct {
	GlossFile, Title string
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
}

// Cache the details of documents by target file name
func cacheDocDetails() map[string]string {
	docMap = map[string]string{}
	ctx := context.Background()
	results, err := findAllTitlesStmt.QueryContext(ctx)
	if err != nil {
		applog.Error("cacheDocDetails, Error for query: ", err)
		return docMap
	}
	defer results.Close()

	for results.Next() {
		var gloss_file, title string
		results.Scan(&gloss_file, &title)
		docMap[gloss_file] = title
	}
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
	simDocs, err := findInBody(queryTerms)
	applog.Info("findDocuments, len(simDocs): ", len(simDocs))
	return mergeBySimilarity(docMap, simDocs), nil
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

// Search the corpus for document bodies most similar to the expanded query
// given.
//  Param: terms - The decomposed query string with 1 < num elements
//  Param: retry - if true and this query fails due to an eror then retry
func findInBody(terms []string) ([]DocSimilarity, error) {
	applog.Info("findInBody, terms = ", terms)
	ctx := context.Background()
	var results *sql.Rows
	var err error
	if len(terms) < 2 {
		applog.Error("findInBody, len(terms) < 2", len(terms))
		return []DocSimilarity{}, errors.New("Too few arguments")
	} else if len(terms) == 2 {
		results, err = similarity2Stmt.QueryContext(ctx, terms[0], terms[1])
	} else if len(terms) == 3 {
		results, err = similarity3Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2])
	}  else {
		// Ignore arguments beyond the first four
		results, err = similarity4Stmt.QueryContext(ctx, terms[0], terms[1],
			terms[2], terms[3])
	}
	if err != nil {
		applog.Error("findInBody, Error for query: ", terms, err)
		return []DocSimilarity{}, err
	}
	simSlice := []DocSimilarity{}
	for results.Next() {
		docSim := DocSimilarity{}
		results.Scan(&docSim.Similarity, &docSim.Document)
		applog.Info("findInBody, Similarity, Document = ", docSim.Similarity,
			docSim.Document)
		simSlice = append(simSlice, docSim)
	}
	return simSlice, nil
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
        applog.Error("find.init() Error preparing collection stmt: ", err)
        return err
    }
    findColStmt = stmt

	cstmt, err := database.PrepareContext(ctx,
		"SELECT count(title) FROM collection WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.init() Error preparing cstmt: ", err)
        return err
    }
    countColStmt = cstmt

	dstmt, err := database.PrepareContext(ctx,
		"SELECT title, gloss_file FROM document WHERE title LIKE ? LIMIT 50")
    if err != nil {
        applog.Error("find.init() Error preparing dstmt: ", err)
        return err
    }
    findDocStmt = dstmt

	cdstmt, err := database.PrepareContext(ctx,
		"SELECT count(title) FROM document WHERE title LIKE ?")
    if err != nil {
        applog.Error("find.init() Error preparing cDocStmt: ", err)
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
		"SELECT COUNT(frequency)/2.0 AS similarity, document FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? GROUP BY document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.init() Error preparing similarity2Stmt: ", err)
        return err
    }
    similarity2Stmt = sim2Stmt

    // For a query with three terms in the query string decomposition
	sim3Stmt, err := database.PrepareContext(ctx, 
		"SELECT COUNT(frequency)/3.0 AS similarity, document FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? GROUP BY document " +
		"ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.init() Error preparing similarity3Stmt: ", err)
        return err
    }
    similarity3Stmt = sim3Stmt

    // For a query with four terms in the query string decomposition
	sim4Stmt, err := database.PrepareContext(ctx, 
		"SELECT COUNT(frequency)/4.0 AS similarity, document FROM  word_freq_doc " +
		"WHERE word = ? OR word = ? OR word = ? OR word = ? GROUP BY document" +
		" ORDER BY similarity DESC LIMIT 20")
    if err != nil {
        applog.Error("find.init() Error preparing similarity4Stmt: ", err)
        return err
    }
    similarity4Stmt = sim4Stmt

    // Find the titles of all documents
	fAllTitlesStmt, err := database.PrepareContext(ctx, 
		"SELECT gloss_file, title FROM document LIMIT 1000000")
    if err != nil {
        applog.Error("find.init() Error preparing findAllTitlesStmt: ", err)
        return err
    }
    findAllTitlesStmt = fAllTitlesStmt

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
			title, ok := docMap[simDoc.Document]
			if ok {
				doc := SimilarDoc{simDoc.Document, title, simDoc.Similarity}
				simDocMap[simDoc.Document] = doc
			} else {
				applog.Info("mergeBySimilarity, doc not found: ", simDoc.Document)
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
		doc := Document{similarDoc.GlossFile, similarDoc.Title}
		docs = append(docs, doc)
	}
	return docs
}