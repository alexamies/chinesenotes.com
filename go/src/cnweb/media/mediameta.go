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
	applog.Info("media.init Initializing mediameta")
	err := initQuery()
	if err != nil {
		applog.Error("media/init: error preparing database statements, retrying",
			err)
		time.Sleep(60000 * time.Millisecond)
		err = initQuery()
		conString := webconfig.DBConfig()
		applog.Error("media/init: error preparing database statements, giving up",
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
	applog.Info("FindMedia: objectId (len) ", objectId, len(objectId))
	mediaMeta := MediaMetadata{}
	ctx := context.Background()
	results, err := findMediaStmt.QueryContext(ctx, objectId)
	results.Next()
	var medium, titleZhCn, titleEn, author, license sql.NullString
	results.Scan(&medium, &titleZhCn, &titleEn, &author, &license)
	if err != nil {
		applog.Error("FindMedia: Error for query: ", objectId, err)
		// Retry
		time.Sleep(200 * time.Millisecond)
		initQuery()
		results, err = findMediaStmt.QueryContext(ctx, objectId)
		results.Next()
		results.Scan(&medium, &titleZhCn, &titleEn, &author, &license)
		if err != nil {
			applog.Error("FindMedia: Retry failed: ", objectId, err)
		}
	}
	results.Close()
	if medium.Valid {
		mediaMeta.ObjectId = medium.String
		applog.Info("FindMedia: medium: ", medium)
	} else {
		applog.Error("FindMedia: ObjectId is not valid")
	}
	if titleZhCn.Valid {
		mediaMeta.TitleZhCn = titleZhCn.String
	}
	if titleEn.Valid {
		mediaMeta.TitleEn = titleEn.String
	}
	if author.Valid {
		mediaMeta.Author = author.String
	}
	if license.Valid {
		mediaMeta.License = license.String
	}
	applog.Info("FindMedia: mediaMeta ", mediaMeta)
	return mediaMeta, nil
}
