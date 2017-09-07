package identity

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	database *sql.DB
	domain *string
	checkSessionStmt *sql.Stmt
	loginStmt *sql.Stmt
	logoutStmt *sql.Stmt
	saveSessionStmt *sql.Stmt
)

type SessionInfo struct {
	Authenticated bool
	User UserInfo
}

type UserInfo struct {
	UserID int
	UserName, Email, FullName, Role string
}

// Open database connection and prepare statements
func init() {

	localhost := "localhost"
	domain = &localhost
	site_domain := os.Getenv("SITEDOMAIN")
	if site_domain != "" {
		domain = &site_domain
	}
	dbhost := "mariadb"
	host := os.Getenv("DBDBHOST")
	if host != "" {
		dbhost = host
	}
	dbport := "3306"
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASSWORD")
	dbname := "hsingyundl"
	conString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost,
		dbport, dbname)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		log.Fatal("FATAL: could not connect to the database, ",
			err)
		panic(err.Error())
	}
	database = db

	stmt1, err := database.Prepare(
		`SELECT user.UserID, UserName, Email, FullName, Role 
		FROM user, passwd 
		WHERE UserName = ? 
		AND user.UserID = passwd.UserID
		AND Password = ?
		LIMIT 1`)
    if err != nil {
        log.Fatal("auth.init() Error preparing stmt1: ", err)
    }
    loginStmt = stmt1

	stmt2, err := database.Prepare(
		`INSERT INTO
		  session (SessionID, UserID)
		VALUES (?, ?)`)
    if err != nil {
        log.Fatal("auth.init() Error preparing stmt2: ", err)
    }
    saveSessionStmt = stmt2

	stmt3, err := database.Prepare(
		`SELECT user.UserID, UserName, Email, FullName, Role 
		FROM user, session 
		WHERE SessionID = ? 
		AND user.UserID = session.UserID
		LIMIT 1`)
    if err != nil {
        log.Fatal("auth.init() Error preparing stmt1: ", err)
    }
    checkSessionStmt = stmt3

	stmt4, err := database.Prepare(
		`UPDATE session SET
		Authenticated = 0
		WHERE SessionID = ?`)
    if err != nil {
        log.Fatal("auth.init() Error preparing stmt4: ", err)
    }
    logoutStmt = stmt4

}

// Check password when the user logs in
func CheckLogin(username, password string) []UserInfo {
	h := sha256.New()
	h.Write([]byte(password))
	hstr := fmt.Sprintf("%x", h.Sum(nil))
	//log.Printf("CheckLogin, username: %s, hstr: %s", username, hstr)
	results, err := loginStmt.Query(username, hstr)
	if err != nil {
		log.Printf("CheckLogin, Error for username: ", username, err)
	}
	defer results.Close()

	users := []UserInfo{}
	for results.Next() {
		user := UserInfo{}
		results.Scan(&user.UserID, &user.UserName, &user.Email, &user.FullName,
			&user.Role)
		users = append(users, user)
	}
	return users
}

// Check session when the user requests a page
func CheckSession(sessionid string) SessionInfo {
	users := checkSessionStore(sessionid)
	if len(users) != 1 {
		return UnauthSession()
	}
	return SessionInfo{
		Authenticated: true,
		User: users[0],
	}
}

// Check session when the user requests a page
func checkSessionStore(sessionid string) []UserInfo {
	log.Printf("CheckSession, sessionid: %s", sessionid)
	results, err := checkSessionStmt.Query(sessionid)
	if err != nil {
		log.Printf("CheckSession, Error: ", err)
	}
	defer results.Close()

	users := []UserInfo{}
	for results.Next() {
		user := UserInfo{}
		results.Scan(&user.UserID, &user.UserName, &user.Email, &user.FullName,
			&user.Role)
		users = append(users, user)
	}
	return users
}

// Generate a new session id after login
func GetSiteDomain() string {
	return *domain
}

// Log the user out of the current session
func Logout(sessionid string) {
	result, err := logoutStmt.Exec(sessionid)
	if err != nil {
		log.Printf("Logout, Error: ", err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("Logout, rows updated: %d", rowsAffected)
	}
}

// Generate a new session id after login
func NewSessionId() string {
	value := "invalid"
	b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        log.Printf("NewSessionId, Error: ", err)
        return value
    }
    val, err := base64.URLEncoding.EncodeToString(b), err
	if err != nil {
		log.Printf("NewSessionId, Error: ", err)
		return value
	}
	return val
}

// Save the session id to the database
func SaveSession(sessionid, username string) {
	//log.Printf("SaveSession, sessionid: %s\n", sessionid)
	result, err := saveSessionStmt.Exec(sessionid, username)
	if err != nil {
		log.Printf("SaveSession, Error for username: ", username, err)
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Printf("SaveSession, rows updated: %d", rowsAffected)
	}
}

// Empty session struct for an unauthenticated session
func UnauthSession() SessionInfo {
	userInfo := UserInfo{
		UserID: -1,
		UserName: "",
		Email: "",
		FullName: "",
		Role: "",
	}
	return SessionInfo{
		Authenticated: false,
		User: userInfo,
	}
}
