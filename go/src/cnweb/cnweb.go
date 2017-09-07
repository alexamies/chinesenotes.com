/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/applog"
	"cnweb/find"
	"cnweb/identity"
	"fmt"
	"encoding/json"
	"log"
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
	err := r.ParseForm()
	if err != nil {
		log.Printf("loginHandler: error parsing form %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.PostFormValue("UserName")
	log.Printf("loginHandler: %s", username)
	password := r.PostFormValue("Password")
	users := identity.CheckLogin(username, password)
	message := ""
	if len(users) != 1 {
		message = "Sorry, your either your username is not found or password do not match."
	} else {
		userInfo := users[0]
		message = fmt.Sprintf("Hello, %s!", userInfo.FullName)
		sessionid := identity.NewSessionId()
		cookie := &http.Cookie{
        	Name: "session",
        	Value: sessionid,
        	Domain: "hsingyundl.org",
        	Path: "/",
        	MaxAge: 86400*30, // One month
        }
        http.SetCookie(w, cookie)
        identity.SaveSession(sessionid, username)
    }
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\"greeting\" :\"%s\"}", message)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		// OK, just don't show the contents that require a login
		log.Printf("logoutHandler: no cookie")
	} else {
		identity.Logout(cookie.Value)
	}
	message := "Please come back again"
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "{\"message\" :\"%s\"}", message)
}

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionInfo := identity.UnauthSession()
	cookie, err := r.Cookie("session")
	if err != nil {
		// OK, just don't show the contents that require a login
		log.Printf("sessionHandler: no cookie")
	} else {
		sessionInfo = identity.CheckSession(cookie.Value)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resultsJson, err := json.Marshal(sessionInfo)
	fmt.Fprintf(w, string(resultsJson))
}

//Entry point for the web application
func main() {

	appLogFile := applog.Create()
	defer applog.Close(appLogFile)
	applog.Info("main.main Started cnweb")

	//index.LoadKeywordIndex()
	//documents := index.FindForKeyword("你")
	http.HandleFunc("/find/", findHandler)
	http.HandleFunc("/loggedin/login", loginHandler)
	http.HandleFunc("/loggedin/logout", logoutHandler)
	http.HandleFunc("/loggedin/session", sessionHandler)
	http.ListenAndServe(":8080", nil)
}
