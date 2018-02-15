/* 
Functions to load Chinese-English dictionary from database
*/
package find

import (
	"context"
	"cnweb/applog"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"cnweb/webconfig"
)

// Loads all words from the database
func LoadDict() (map[string]Word, error) {
	wdict := map[string]Word{}
	conString := webconfig.DBConfig()
	database, err := sql.Open("mysql", conString)
    if err != nil {
        applog.Error("find.load_dict connecting to database: ", err)
        return wdict, err
	}
	ctx := context.Background()
	stmt, err := database.PrepareContext(ctx, 
		"SELECT simplified, traditional, pinyin, english, headword FROM words")
    if err != nil {
        applog.Error("find.load_dict Error preparing stmt: ", err)
        return wdict, err
    }
	results, err := stmt.QueryContext(ctx)
	if err != nil {
		applog.Error("find.load_dict, Error for query: ", err)
        return wdict, err
	}
	for results.Next() {
		word := Word{}
		var hw sql.NullInt64
		var trad sql.NullString
		results.Scan(&word.Simplified, &trad, &word.Pinyin, &word.English, &hw)
		//applog.Info("find.load_dict, simplified, headword = ", word.Simplified, hw)
		if hw.Valid {
			word.HeadwordId = int(hw.Int64)
		}
		wdict[word.Simplified] = word
		if trad.Valid {
			word.Traditional = trad.String
			wdict[trad.String] = word
		}
	}
	return wdict, nil
}
