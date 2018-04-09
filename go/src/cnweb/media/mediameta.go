// Package for media metadata
package media

import (
	"context"
	"cnweb/applog"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"cnweb/webconfig"
	"time"
)

var (
	findMediaStmt *sql.Stmt
)

// A media object structure for media metadata
type MediaMetadata struct {
	ObjectId, TitleZhCn, TitleEn, Author, License string
}

// Open database connection and prepare query
func init() {
	err := initQuery()
	if err != nil {
		applog.Error("media/init: error preparing database statements, retrying",
			err)
		time.Sleep(60000 * time.Millisecond)
		err = initQuery()
		conString := webconfig.DBConfig()
		applog.Fatal("media/init: error preparing database statements, giving up",
			conString, err)
	}
}

func initQuery() error {
	conString := webconfig.DBConfig()
	db, err := sql.Open("mysql", conString)
	if err != nil {
		return err
	}
	ctx := context.Background()
	fwstmt, err := db.PrepareContext(ctx, 
`SELECT medium_resolution, title_zh_cn, title_en, author, license
FROM illustrations
WHERE medium_resolution = ?
LIMIT 1`)
    if err != nil {
        applog.Error("media.initQuery() Error preparing fwstmt: ", err)
        return err
    }
    findMediaStmt = fwstmt
    return nil
}

// Looks up media metadata by object ID
func FindMedia(objectId string) (MediaMetadata, error) {
	mediaMeta := MediaMetadata{}
	ctx := context.Background()
	results, err := findMediaStmt.QueryContext(ctx, objectId)
	results.Next()
	results.Scan(&mediaMeta)
	if err != nil {
		applog.Error("FindMedia: Error for query: ", objectId, err)
		// Retry
		time.Sleep(200 * time.Millisecond)
		initQuery()
		results, err = findMediaStmt.QueryContext(ctx, objectId)
		results.Next()
		results.Scan(&mediaMeta)
		if err != nil {
			applog.Error("FindMedia: Retry failed: ", objectId, err)
		}
	}
	results.Close()
	return mediaMeta, nil
}
