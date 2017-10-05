package identity

import (
	"cnweb/applog"
	"cnweb/webconfig"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	database *sql.DB
	domain *string
	checkSessionStmt *sql.Stmt
	loginStmt *sql.Stmt
	logoutStmt *sql.Stmt
	saveSessionStmt *sql.Stmt
	updateSessionStmt *sql.Stmt
)

type SessionInfo struct {
	Authenticated int
	Valid bool
	User UserInfo
}

type UserInfo struct {
	UserID int
	UserName, Email, FullName, Role string
}

func init() {
	err := initStatements()
	if err != nil {
		applog.Error("identity/init: error preparing database statements, retrying",
			err)
		time.Sleep(60000 * time.Millisecond)
		err = initStatements()
		conString := webconfig.DBConfig()
		applog.Fatal("identity/init: error preparing database statements, giving up",
			conString, err)
	}
	applog.Info("identity/init: Ready to go")
}

// Open database connection and prepare statements
func initStatements() error {
	conString := webconfig.DBConfig()
	db, err := sql.Open("mysql", conString)
	if err != nil {
		applog.Fatal("FATAL: could not connect to the database, ",
			err)
		panic(err.Error())
	}
	database = db

	ctx := context.Background()
	stmt1, err := database.PrepareContext(ctx,
		`SELECT user.UserID, UserName, Email, FullName, Role 
		FROM user, passwd 
		WHERE UserName = ? 
		AND user.UserID = passwd.UserID
		AND Password = ?
		LIMIT 1`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt1: ", err)
        return err
    }
    loginStmt = stmt1

	stmt2, err := database.PrepareContext(ctx,
		`INSERT INTO
		  session (SessionID, UserID, Authenticated)
		VALUES (?, ?, ?)`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt2: ", err)
        return err
    }
    saveSessionStmt = stmt2

    // Need to fix use of username in session table. Should be UserId
	stmt3, err := database.PrepareContext(ctx,
		`SELECT user.UserID, UserName, Email, FullName, Role, Authenticated
		FROM user, session 
		WHERE SessionID = ? 
		AND user.UserID = session.UserID
		LIMIT 1`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt3: ", err)
        return err
    }
    checkSessionStmt = stmt3

	stmt4, err := database.PrepareContext(ctx,
		`UPDATE session SET
		Authenticated = 0
		WHERE SessionID = ?`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt4: ", err)
        return err
    }
    logoutStmt = stmt4

	stmt5, err := database.PrepareContext(ctx,
		`UPDATE session SET
		Authenticated = ?,
		UserID = ?
		WHERE SessionID = ?`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt4: ", err)
        return err
    }
    updateSessionStmt = stmt5
    return nil
}

// Check password when the user logs in
func CheckLogin(username, password string) ([]UserInfo, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hstr := fmt.Sprintf("%x", h.Sum(nil))
	applog.Info("CheckLogin, username, hstr:", username, hstr)
	ctx := context.Background()
	results, err := loginStmt.QueryContext(ctx, username, hstr)
	defer results.Close()
	if err != nil {
		applog.Error("CheckLogin, Error for username: ", username, err)
		// Sleep for a while, reinitialize, and retry
		time.Sleep(2000 * time.Millisecond)
		initStatements()
		results, err = loginStmt.QueryContext(ctx, username, hstr)
		if err != nil {
			applog.Error("CheckLogin, Give up after retry: ", username, err)
			return []UserInfo{}, err
		}
	}

	users := []UserInfo{}
	for results.Next() {
		user := UserInfo{}
		results.Scan(&user.UserID, &user.UserName, &user.Email, &user.FullName,
			&user.Role)
		users = append(users, user)
	}
	return users, nil
}

// Check session when the user requests a page
func CheckSession(sessionid string) SessionInfo {
	sessions := checkSessionStore(sessionid)
	if len(sessions) != 1 {
		return InvalidSession()
	}
	applog.Info("CheckSession, Authenticated =", sessions[0].Authenticated)
	return sessions[0]
}

// Check session when the user requests a page
func checkSessionStore(sessionid string) []SessionInfo {
	applog.Info("CheckSession, sessionid: %s", sessionid)
	ctx := context.Background()
	results, err := checkSessionStmt.QueryContext(ctx, sessionid)
	if err != nil {
		applog.Error("checkSessionStore, Error: ", err)
	}
	defer results.Close()

	sessions := []SessionInfo{}
	for results.Next() {
		user := UserInfo{}
		session := SessionInfo{}
		results.Scan(&user.UserID, &user.UserName, &user.Email, &user.FullName,
			&user.Role, &session.Authenticated)
		session.User = user
		session.Valid = true
		sessions = append(sessions, session)
	}
	applog.Info("checkSessionStore, sessions found: ", len(sessions))
	return sessions
}

// Generate a new session id after login
func IsAuthorized(user UserInfo, permission string) bool {
	if user.Role == "admin" {
	  return true
	}
	return false
}

// Log the user out of the current session
func Logout(sessionid string) {
	ctx := context.Background()
	result, err := logoutStmt.ExecContext(ctx, sessionid)
	if err != nil {
		applog.Error("Logout, Error: ", err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		applog.Info("Logout, rows updated:", rowsAffected)
	}
}

// Generate a new session id after login
func NewSessionId() string {
	value := "invalid"
	b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        applog.Error("NewSessionId, Error: ", err)
        return value
    }
    val, err := base64.URLEncoding.EncodeToString(b), err
	if err != nil {
		applog.Info("NewSessionId, Error: ", err)
		return value
	}
	return val
}

// Save an authenticated session to the database
func SaveSession(sessionid string, userInfo UserInfo, authenticated int) SessionInfo {
	applog.Info("SaveSession, sessionid:", sessionid)
	ctx := context.Background()
	result, err := saveSessionStmt.ExecContext(ctx, sessionid, userInfo.UserID,
		authenticated)
	if err != nil {
		applog.Info("SaveSession, Error for username: ", userInfo.UserID, err)
		return InvalidSession()
	}
	rowsAffected, _ := result.RowsAffected()
	applog.Info("SaveSession, rows updated: ", rowsAffected)
	return SessionInfo{
		Authenticated: authenticated,
		Valid: true,
		User: userInfo,
	}
}

// Empty session struct for an unauthenticated session
func InvalidSession() SessionInfo {
	userInfo := UserInfo{
		UserID: 1,
		UserName: "",
		Email: "",
		FullName: "",
		Role: "",
	}
	return SessionInfo{
		Authenticated: 0,
		Valid: false,
		User: userInfo,
	}
}

// Log a user in when they already have an unauthenticated session
func UpdateSession(sessionid string, userInfo UserInfo, authenticated int) SessionInfo {
	ctx := context.Background()
	result, err := updateSessionStmt.ExecContext(ctx, authenticated,
		userInfo.UserID, sessionid)
	if err != nil {
		applog.Error("UpdateSession, Error: ", err)
		return InvalidSession()
	} 
	rowsAffected, _ := result.RowsAffected()
	applog.Info("UpdateSession, rows updated:", rowsAffected)
	return SessionInfo{
		Authenticated: authenticated,
		User: userInfo,
	}
}