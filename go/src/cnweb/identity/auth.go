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
	changePasswordStmt *sql.Stmt
	checkSessionStmt *sql.Stmt
	getResetRequestStmt *sql.Stmt
	getUserStmt *sql.Stmt
	getUserByEmailStmt *sql.Stmt
	loginStmt *sql.Stmt
	logoutStmt *sql.Stmt
	requestResetStmt *sql.Stmt
	saveSessionStmt *sql.Stmt
	updateSessionStmt *sql.Stmt
	updateResetRequestStmt *sql.Stmt
)

type ChangePasswordResult struct {
	OldPasswordValid bool
	ChangeSuccessful bool
	ShowNewForm bool
}

type RequestResetResult struct {
	EmailValid bool
	RequestResetSuccess bool
	ShowNewForm bool
	User UserInfo
	Token string
}

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
	applog.Info("identity/init: Ready to go, ")
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
        applog.Error("auth.init() Error preparing stmt5: ", err)
        return err
    }
    updateSessionStmt = stmt5

	stmt6, err := database.PrepareContext(ctx,
		`UPDATE passwd SET
		Password = ?
		WHERE UserID = ?`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt6: ", err)
        return err
    }
    changePasswordStmt = stmt6

	stmt7, err := database.PrepareContext(ctx,
		`SELECT user.UserID, UserName, Email, FullName, Role 
		FROM user
		WHERE UserName = ? 
		LIMIT 1`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt7: ", err)
        return err
    }
    getUserStmt = stmt7

	stmt8, err := database.PrepareContext(ctx,
		`INSERT INTO
		passwdreset (Token, UserID)
		VALUES (?, ?)`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt8: ", err)
        return err
    }
    requestResetStmt = stmt8

	stmt9, err := database.PrepareContext(ctx,
		`SELECT user.UserID, UserName, Email, FullName, Role 
		FROM user
		WHERE Email = ? 
		LIMIT 1`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt9: ", err)
        return err
    }
    getUserByEmailStmt = stmt9

	stmt10, err := database.PrepareContext(ctx,
		`SELECT UserID
		FROM passwdreset
		WHERE Token = ?
		AND Valid = 1
		LIMIT 1`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt10: ", err)
        return err
    }
    getResetRequestStmt = stmt10

	stmt11, err := database.PrepareContext(ctx,
		`UPDATE passwdreset SET
		Valid = 0
		WHERE Token = ?`)
    if err != nil {
        applog.Error("auth.init() Error preparing stmt11: ", err)
        return err
    }
    updateResetRequestStmt = stmt11

    return nil
}

// Log a user in when they already have an unauthenticated session
func ChangePassword(userInfo UserInfo, oldPassword, password string) ChangePasswordResult {
	users, err := CheckLogin(userInfo.UserName, oldPassword)
	if err != nil {
		applog.Error("ChangePassword checking login, ", err)
		return ChangePasswordResult{true, false, false}
	}
	if len(users) != 1 {
		applog.Info("ChangePassword, user or password wrong: ",
			userInfo.UserName)
		return ChangePasswordResult{false, false, false}
	}
	ctx := context.Background()
	h := sha256.New()
	h.Write([]byte(password))
	hstr := fmt.Sprintf("%x", h.Sum(nil))
	result, err := changePasswordStmt.ExecContext(ctx, hstr, userInfo.UserID)
	if err != nil {
		applog.Error("ChangePassword, Error: ", err)
		return ChangePasswordResult{true, false, false}
	} 
	rowsAffected, _ := result.RowsAffected()
	applog.Info("ChangePassword, rows updated:", rowsAffected)
	return ChangePasswordResult{true, true, false}
}

// Check password when the user logs in
func CheckLogin(username, password string) ([]UserInfo, error) {
	h := sha256.New()
	h.Write([]byte(password))
	hstr := fmt.Sprintf("%x", h.Sum(nil))
	//applog.Info("CheckLogin, username, hstr:", username, hstr)
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
	if len(users) == 0 {
		applog.Info("CheckLogin, user or password wrong: ", username)
		u, _ := GetUser(username)
		if len(u) == 0 {
			applog.Info("CheckLogin, user not found: ", username)
		}
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
	applog.Info("CheckSession, sessionid: ", sessionid)
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

// Get the user information
func GetUser(username string) ([]UserInfo, error) {
	applog.Info("getUser, username:", username)
	ctx := context.Background()
	results, err := getUserStmt.QueryContext(ctx, username)
	defer results.Close()
	if err != nil {
		applog.Error("getUser, Error for username: ", username, err)
		return []UserInfo{}, err
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

// Empty session struct for an unauthenticated session
func InvalidUser() UserInfo {
	return UserInfo{
		UserID: 1,
		UserName: "",
		Email: "",
		FullName: "",
		Role: "",
	}
}

// Generate a new session id after login
func IsAuthorized(user UserInfo, permission string) bool {
	if user.Role == "admin" || user.Role == "editor" || user.Role == "translator" {
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

// Old password does not match
func OldPasswordDoesNotMatch() ChangePasswordResult {
	return ChangePasswordResult{false, true, false}
}

// Request a password reset, to be sent by email
func RequestPasswordReset(email string) RequestResetResult {
	applog.Info("RequestPasswordReset, email:", email)
	b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        applog.Error("RequestPasswordReset, Error: ", err)
        return RequestResetResult{true, false, true, InvalidUser(), ""}
    }
    token, err := base64.URLEncoding.EncodeToString(b), err
	if err != nil {
		applog.Info("RequestPasswordReset, Error: ", err)
		return RequestResetResult{true, false, true, InvalidUser(), ""}
	}
	ctx := context.Background()
	results, err := getUserByEmailStmt.QueryContext(ctx, email)
	defer results.Close()
	if err != nil {
		applog.Error("RequestPasswordReset, Error for email: ", email, err)
		return RequestResetResult{true, false, true, InvalidUser(), ""}
	}
	users := []UserInfo{}
	for results.Next() {
		user := UserInfo{}
		results.Scan(&user.UserID, &user.UserName, &user.Email, &user.FullName,
			&user.Role)
		users = append(users, user)
	}

	if len(users) != 1 {
		applog.Error("RequestPasswordReset, No email: ", email)
		return RequestResetResult{false, false, true, InvalidUser(), ""}
	}

	result, err := requestResetStmt.ExecContext(ctx, token, users[0].UserID)
	if err != nil {
		applog.Info("RequestPasswordReset, Error for email: ", email, err)
		return RequestResetResult{true, false, true, InvalidUser(), ""}
	}
	rowsAffected, _ := result.RowsAffected()
	applog.Info("RequestPasswordReset, rows updated: ", rowsAffected)
	return RequestResetResult{true, true, false, users[0], token}
}

// Reset a password
func ResetPassword(token, password string) bool {
	applog.Info("ResetPassword, token:", token)
	ctx := context.Background()
	results, err := getResetRequestStmt.QueryContext(ctx, token)
	defer results.Close()
	if err != nil {
		applog.Error("ResetPassword, Error for token: ", token, err)
		return false
	}
	userIds := []string{}
	for results.Next() {
		userId := ""
		results.Scan(&userId)
		userIds = append(userIds, userId)
	}
	if len(userIds) != 1 {
		applog.Error("ResetPassword, No userId: ", token)
		return false
	}
	userId := userIds[0]

	// Change password
	h := sha256.New()
	h.Write([]byte(password))
	hstr := fmt.Sprintf("%x", h.Sum(nil))
	result, err := changePasswordStmt.ExecContext(ctx, hstr, userId)
	if err != nil {
		applog.Error("ResetPassword, Error setting password: ", err)
		return false
	} 
	rowsAffected, _ := result.RowsAffected()
	applog.Info("ResetPassword, rows updated for change pwd:", rowsAffected)

	// Update reset token so that it cannot be used again
	result, err = updateResetRequestStmt.ExecContext(ctx, token)
	if err != nil {
		applog.Error("ResetPassword, Error updating reset token: ", err)
	} 
	rowsAffected, _ = result.RowsAffected()
	applog.Info("ResetPassword, rows updated for token:", rowsAffected)

	return true
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