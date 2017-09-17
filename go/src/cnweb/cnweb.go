/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/applog"
	"cnweb/find"
	"cnweb/identity"
	"cnweb/webconfig"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func findHandler(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	}
	results := find.FindDocuments(q)
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.handler error marshalling JSON, ", err)
	}
	//applog.Info("handler, results returned: ", string(resultsJson))
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, string(resultsJson))
}

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
	users := identity.CheckLogin(username, password)
	if len(users) != 1 {
		applog.Error("loginHandler: user not found", username)
	} else {
		cookie, err := r.Cookie("session")
		if err == nil {
			applog.Error("loginHandler: updating session", cookie.Value)
			sessionInfo = identity.UpdateSession(cookie.Value, users[0], 1)
		}
		if (err != nil) || !sessionInfo.Valid {
			sessionid := identity.NewSessionId()
			//applog.Info("loginHandler: creating new session %v", sessionid)
			cookie := &http.Cookie{
        		Name: "session",
        		Value: sessionid,
        		Domain: identity.GetSiteDomain(),
        		Path: "/",
        		MaxAge: 86400*30, // One month
        	}
        	http.SetCookie(w, cookie)
        	sessionInfo = identity.SaveSession(sessionid, users[0], 1)
        }
    }
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resultsJson, err := json.Marshal(sessionInfo)
	fmt.Fprintf(w, string(resultsJson))
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
		vars := webconfig.GetAll()
		tmpl, err := template.New("translation_portal.html").ParseFiles("templates/translation_portal.html")
		if err != nil {
			applog.Error("portalHandler: error parsing template", err)
		}
		if tmpl == nil {
			applog.Error("portalHandler: Template is nil")
		}
		if err != nil {
			applog.Error("portalHandler: error parsing template", err)
		}
		err = tmpl.Execute(w, vars)
		if err != nil {
			applog.Error("portalHandler: error rendering template", err)
		}
	} else {
		http.Error(w, "Not authorized", 403)
	}
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
        	Domain: identity.GetSiteDomain(),
        	Path: "/",
        	MaxAge: 86400, // One day
        }
        http.SetCookie(w, cookie)
        userInfo := identity.UserInfo{
			UserID: -1,
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
	http.HandleFunc("/loggedin/login", loginHandler)
	http.HandleFunc("/loggedin/logout", logoutHandler)
	http.HandleFunc("/loggedin/session", sessionHandler)
	http.HandleFunc("/loggedin/portal", portalHandler)
	http.ListenAndServe(":8080", nil)
}
