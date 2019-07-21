/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/applog"
	"cnweb/dictionary"
	"cnweb/find"
	"cnweb/identity"
	"cnweb/mail"
	"cnweb/media"
	"cnweb/webconfig"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	parser find.QueryParser
	wdict map[string]dictionary.Word
)

func init() {
	applog.Info("cnweb.main.init Initializing cnweb")
	var err error
	wdict, err = dictionary.LoadDict()
	if err != nil {
		applog.Error("main.init() unable to load dictionary: ", err)
	}
	parser = find.DictQueryParser{wdict}
}

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
			applog.Error("main.adminHandler: error parsing template", err)
		}
		if tmpl == nil {
			applog.Error("main.adminHandler: Template is nil")
		}
		if err != nil {
			applog.Error("main.adminHandler: error parsing template", err)
		}
		err = tmpl.Execute(w, vars)
		if err != nil {
			applog.Error("main.adminHandler: error rendering template", err)
		}
	} else {
		applog.Info("adminHandler, Not authorized: ", sessionInfo.User)
		http.Error(w, "Not authorized", http.StatusForbidden)
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
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	} else if tmpl == nil {
		applog.Error("displayPage: Template is nil")
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, content)
	if err != nil {
		applog.Error("displayPage: error rendering template", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}	
}

// HTML redirect to the index.html page, for healthchecks used by the load balancer.
// Do not expect a user to hit this.
func displayHome(w http.ResponseWriter, r *http.Request) {
	applog.Error("displayHome: r.URL", r.URL)
	page := `<!DOCTYPE html>
<html>
  <head>
   <meta http-equiv='refresh' content='0; url=index.html'>
  </head>
  <body>
   <p>Redirect to main page</p>
  </body>
</html>
`
	fmt.Fprintf(w, page)
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
			http.Error(w, "Not authorized", http.StatusForbidden)
			return sessionInfo
		}
	} else {
		applog.Info("enforceValidSession, Invalid session ", sessionInfo.User)
		http.Error(w, "Not authorized", http.StatusForbidden)
		return identity.InvalidSession()
	}
	return sessionInfo
}

// Finds documents matching the given query with search in text body
func findAdvanced(response http.ResponseWriter, request *http.Request) {
	applog.Info("main.findAdvanced, enter")
	findDocs(response, request, true)
}

// Finds documents matching the given query
func findDocs(response http.ResponseWriter, request *http.Request, advanced bool) {
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	} else {
		query := queryString["text"]
		if len(query) > 0 {
			q = query[0]
		}
	}

	var results find.QueryResults
	var err error

	c := queryString["collection"]
	if (len(c) > 0) && (c[0] != "") {
		results, err = find.FindDocumentsInCol(parser, q, c[0])
	} else {
		results, err = find.FindDocuments(parser, q, advanced)
	}

	if err != nil {
		applog.Error("main.findDocs Error searching docs, ", err)
		http.Error(response, "Error searching docs",
			http.StatusInternalServerError)
		return
	}
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.findDocs error marshalling JSON, ", err)
		http.Error(response, "Error marshalling results",
			http.StatusInternalServerError)
	} else {
		if (q != "hello" && q != "Eight" ) { // Health check monitoring probe
			applog.Info("main.findDocs, results: ", string(resultsJson))
		}
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(response, string(resultsJson))
	}
}

// Finds documents matching the given query
func findHandler(response http.ResponseWriter, request *http.Request) {
	applog.Info("main.findHandler, enter")
	findDocs(response, request, false)
}

// Finds terms matching the given query with a substring match
func findSubstring(response http.ResponseWriter, request *http.Request) {
	applog.Info("main.findSubstring, enter")
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	q := ""
	if len(query) > 0 {
		q = query[0]
	}
	topic := queryString["topic"]
	t := ""
	if len(topic) > 0 {
		t = topic[0]
	}
	results, err := dictionary.LookupSubstr(q, t)
	if err != nil {
		applog.Error("main.findSubstring Error looking up term, ", err)
		http.Error(response, "Error looking up term",
			http.StatusInternalServerError)
		return
	}
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.findSubstring error marshalling JSON, ", err)
		http.Error(response, "Error marshalling results",
			http.StatusInternalServerError)
	} else {
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(response, string(resultsJson))
	}
}

// Health check for monitoring or load balancing system, checks reachability
func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthcheck ok")
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
		http.Error(w, "Error checking login", http.StatusInternalServerError)
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

// Retrieves detail about media objects
func mediaDetailHandler(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	query := queryString["mediumResolution"]
	applog.Info("mediaDetailHandler: query: ", query)
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	}
	results, err := media.FindMedia(q)
	if err != nil {
		applog.Error("main.mediaDetailHandler Error retrieving media detail, ",
			err)
		http.Error(response, "Error retrieving media detail",
			http.StatusInternalServerError)
		return
	}
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.mediaDetailHandler error marshalling JSON, ", err)
		http.Error(response, "Error marshalling results",
			http.StatusInternalServerError)
	} else {
		response.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(response, string(resultsJson))
	}
}

// Starting point for the Translation Portal
func portalHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
	} else {
		applog.Info("portalHandler error getting cookie: ", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	user := sessionInfo.User
	if identity.IsAuthorized(user, "translation_portal") {
		displayPortalHome(w)
	} else {
		applog.Info("portalHandler with role not authorized for portal",
			user.UserName, user.Role)
		http.Error(w, "Not authorized", http.StatusForbidden)
	}
}

// Static handler for pages in the Translation Portal Library
func portalLibraryHandler(w http.ResponseWriter, r *http.Request) {
	applog.Info("portalLibraryHandler: url ", r.URL)
	sessionInfo := identity.InvalidSession()
	cookie, err := r.Cookie("session")
	if err == nil {
		sessionInfo = identity.CheckSession(cookie.Value)
	} else {
		applog.Info("portalLibraryHandler error getting cookie: ", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	user := sessionInfo.User
	if identity.IsAuthorized(user, "translation_portal") {
		portalLibHome := os.Getenv("PORTAL_LIB_HOME")
		filepart := r.URL.Path[len("/loggedin/portal_library/"):]
		filename := portalLibHome + "/" + filepart
		_, err := os.Stat(filename)
		if err != nil {
			applog.Info("portalLibraryHandler os.Stat error: ", err, filename)
			custom404(w, r, filename)
			return
		}
		applog.Info("portalLibraryHandler: serving file ", filename)
		http.ServeFile(w, r, filename)
	} else {
		applog.Info("portalLibraryHandler with role not authorized",
			user.UserName, user.Role)
		http.Error(w, "Not authorized", http.StatusForbidden)
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
	applog.Info("resetPasswordHandler enter")
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
		http.Error(w, "Error checking login", http.StatusInternalServerError)
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
	applog.Info("cnweb.main Starting cnweb")
	http.HandleFunc("/#", findHandler)
	http.HandleFunc("/find/", findHandler)
	http.HandleFunc("/findadvanced/", findAdvanced)
	http.HandleFunc("/findmedia", mediaDetailHandler)
	http.HandleFunc("/findsubstring", findSubstring)
	http.HandleFunc("/healthcheck/", healthcheck)
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
	http.HandleFunc("/", displayHome)
	portStr := ":" + strconv.Itoa(webconfig.GetPort())
	http.ListenAndServe(portStr, nil)
}
