/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/applog"
	"cnweb/find"
	"cnweb/identity"
	"cnweb/mail"
	"cnweb/webconfig"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Starting point for the Administration Portal
func adminHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
	}
	if identity.IsAuthorized(sessionInfo.User, "admin_portal") {
		vars := webconfig.GetAll()
		tmpl, err := template.New("admin_portal.html").ParseFiles("templates/admin_portal.html")
		if err != nil {
			applog.Error("adminHandler: error parsing template", err)
		}
		if tmpl == nil {
			applog.Error("adminHandler: Template is nil")
		}
		if err != nil {
			applog.Error("adminHandler: error parsing template", err)
		}
		err = tmpl.Execute(w, vars)
		if err != nil {
			applog.Error("adminHandler: error rendering template", err)
		}
	} else {
		http.Error(w, "Not authorized", 403)
	}
}

// Process a change password request
func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := enforceValidSession(w, r)
	if sessionInfo.Authenticated == 1 {
		oldPassword := r.PostFormValue("OldPassword")
		password := r.PostFormValue("Password")
		result := identity.ChangePassword(sessionInfo.User, oldPassword,
			password)
    	if strings.Contains(r.Header.Get("Accept"), "application/json") {
    		sendJSON(w, result)
		} else {
			displayPage(w, "change_password_form.html", result)
		}
	}
}

// Display change password form
func changePasswordFormHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := enforceValidSession(w, r)
	if sessionInfo.Authenticated == 1 {
		// fresh form
		result := identity.ChangePasswordResult{false, false, true}
		displayPage(w, "change_password_form.html", result)
	}
}

// Custom 404 page handler
func custom404(w http.ResponseWriter, r *http.Request, url string) {
	applog.Error("custom404: sending 404 for ", url)
	displayPage(w, "404.html", nil)
}

func displayPage(w http.ResponseWriter, templateName string, content interface{}) {
	tmpl, err := template.New(templateName).ParseFiles("templates/" + templateName)
	if err != nil {
		applog.Error("displayPage: error parsing template", err)
		http.Error(w, "Server Error", 500)
		return
	} else if tmpl == nil {
		applog.Error("displayPage: Template is nil")
		http.Error(w, "Server Error", 500)
		return
	}
	err = tmpl.Execute(w, content)
	if err != nil {
		applog.Error("displayPage: error rendering template", err)
		http.Error(w, "Server Error", 500)
	}	
}

// Displays the translation portal home page
func displayPortalHome(w http.ResponseWriter) {
	vars := webconfig.GetAll()
	displayPage(w, "translation_portal.html", vars)
}

// Process a change password request
func enforceValidSession(w http.ResponseWriter, r *http.Request) identity.SessionInfo {
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
		if sessionInfo.Authenticated != 1 {
			http.Error(w, "Not authorized", 403)
			return sessionInfo
		}
	} else {
		http.Error(w, "Not authorized", 403)
		return identity.InvalidSession()
	}
	return sessionInfo
}

// Finds documents matching the given query
func findHandler(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	}
	results, err := find.FindDocuments(q)
	if err != nil {
		applog.Error("main.findHandler searching docs, ", err)
		http.Error(response, "Error searching docs", 500)
		return
	}
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.findHandler error marshalling JSON, ", err)
		http.Error(response, "Error marshalling results", 500)
	} else {
		//applog.Info("handler, results returned: ", string(resultsJson))
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(response, string(resultsJson))
	}
}

// Display login form for the Translation Portal
func loginFormHandler(w http.ResponseWriter, r *http.Request) {
	displayPage(w, "login_form.html", nil)
}

// Process a login request
func loginHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.InvalidSession()
	err := r.ParseForm()
	if err != nil {
		applog.Error("loginHandler: error parsing form", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.PostFormValue("UserName")
	applog.Info("loginHandler: username = ", username)
	password := r.PostFormValue("Password")
	users, err := identity.CheckLogin(username, password)
	if err != nil {
		applog.Error("main.loginHandler checking login, ", err)
		http.Error(w, "Error checking login", 500)
		return
	}
	if len(users) != 1 {
		applog.Error("loginHandler: user not found", username)
	} else {
		cookie, err := r.Cookie("session")
		if err == nil {
			applog.Info("loginHandler: updating session", cookie.Value)
			sessionInfo = identity.UpdateSession(cookie.Value, users[0], 1)
		}
		if (err != nil) || !sessionInfo.Valid {
			sessionid := identity.NewSessionId()
			//applog.Info("loginHandler: creating new session %v", sessionid)
			cookie := &http.Cookie{
        		Name: "session",
        		Value: sessionid,
        		Domain: webconfig.GetSiteDomain(),
        		Path: "/",
        		MaxAge: 86400*30, // One month
        	}
        	http.SetCookie(w, cookie)
        	sessionInfo = identity.SaveSession(sessionid, users[0], 1)
        }
    }
    if strings.Contains(r.Header.Get("Accept"), "application/json") {
    	sendJSON(w, sessionInfo)
	} else {
		if sessionInfo.Authenticated == 1 {
			displayPortalHome(w)
		} else {
			loginFormHandler(w, r)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		// OK, just don't show the contents that require a login
		applog.Error("logoutHandler: no cookie")
	} else {
		identity.Logout(cookie.Value)
	}
	message := "Please come back again"
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\"message\" :\"%s\"}", message)
}

// Starting point for the Translation Portal
func portalHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
	}
	if identity.IsAuthorized(sessionInfo.User, "translation_portal") {
		displayPortalHome(w)
	} else {
		http.Error(w, "Not authorized", 403)
	}
}

// Static handler for pages in the Translation Portal Library
func portalLibraryHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
	}
	if identity.IsAuthorized(sessionInfo.User, "translation_portal") {
		portalLibHome := os.Getenv("PORTAL_LIB_HOME")
		filepart := r.URL.Path[len("/loggedin/portal_library/"):]
		filename := portalLibHome + "/" + filepart
		_, err := os.Stat(filename)
		if err != nil {
			custom404(w, r, filename)
			return
		}
		applog.Info("portalLibraryHandler: serving file ", filename)
		http.ServeFile(w, r, filename)
	} else {
		http.Error(w, "Not authorized", 403)
	}
}

// Display form to request a password reset
func requestResetFormHandler(w http.ResponseWriter, r *http.Request) {
	content := identity.RequestResetResult{true, false, true,
		identity.InvalidUser(), ""}
	displayPage(w, "request_reset_form.html", content)
}

// Process a request for password reset
func requestResetHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("Email")
	result := identity.RequestPasswordReset(email)
	if result.RequestResetSuccess {
		err := mail.SendPasswordReset(result.User, result.Token)
		if err != nil {
			result.RequestResetSuccess = false
		}
	}
    if strings.Contains(r.Header.Get("Accept"), "application/json") {
    	sendJSON(w, result)
	} else {
		displayPage(w, "request_reset_form.html", result)
	}
}

func resetPasswordFormHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	token := queryString["token"]
	content := make(map[string]string)
	if len(token) == 1 {
		content["Token"] = token[0]
	} else {
		content["Token"] = ""
	}
	displayPage(w, "reset_password_form.html", content)
}

func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	token := r.PostFormValue("Token")
	newPassword := r.PostFormValue("NewPassword")
	result := identity.ResetPassword(token, newPassword)
	content := make(map[string]bool)
	if result {
		content["ResetPasswordSuccessful"] = true
	}
    if strings.Contains(r.Header.Get("Accept"), "application/json") {
    	sendJSON(w, result)
	} else {
		displayPage(w, "reset_password_confirmation.html", content)
	}
}

func sendJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resultsJson, err := json.Marshal(obj)
	if err != nil {
		applog.Error("changePasswordHandler: error marshalling json", err)
		http.Error(w, "Error checking login", 500)
		return
	}
	fmt.Fprintf(w, string(resultsJson))
}

// Check to see if the user has a session
func sessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
	}
	if (err != nil) || (!sessionInfo.Valid) {
		// OK, just don't show the contents that don't require a login
		applog.Info("sessionHandler: creating a new cookie")
		sessionid := identity.NewSessionId()
		cookie := &http.Cookie{
        	Name: "session",
        	Value: sessionid,
        	Domain: webconfig.GetSiteDomain(),
        	Path: "/",
        	MaxAge: 86400, // One day
        }
        http.SetCookie(w, cookie)
        userInfo := identity.UserInfo{
			UserID: 1,
			UserName: "",
			Email: "",
			FullName: "",
			Role: "",
		}
        identity.SaveSession(sessionid, userInfo, 0)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resultsJson, err := json.Marshal(sessionInfo)
	fmt.Fprintf(w, string(resultsJson))
}

//Entry point for the web application
func main() {
	applog.Info("main.main Started cnweb")

	//index.LoadKeywordIndex()
	//documents := index.FindForKeyword("你")
	http.HandleFunc("/find/", findHandler)
	http.HandleFunc("/loggedin/admin", adminHandler)
	http.HandleFunc("/loggedin/changepassword", changePasswordFormHandler)
	http.HandleFunc("/loggedin/login", loginHandler)
	http.HandleFunc("/loggedin/login_form", loginFormHandler)
	http.HandleFunc("/loggedin/logout", logoutHandler)
	http.HandleFunc("/loggedin/session", sessionHandler)
	http.HandleFunc("/loggedin/portal", portalHandler)
	http.HandleFunc("/loggedin/portal_library/", portalLibraryHandler)
	http.HandleFunc("/loggedin/request_reset", requestResetHandler)
	http.HandleFunc("/loggedin/request_reset_form", requestResetFormHandler)
	http.HandleFunc("/loggedin/reset_password", resetPasswordFormHandler)
	http.HandleFunc("/loggedin/reset_password_submit", resetPasswordHandler)
	http.HandleFunc("/loggedin/submitcpwd", changePasswordHandler)
	http.ListenAndServe(":8080", nil)
}
