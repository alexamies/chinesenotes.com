/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/applog"
	"cnweb/find"
	"fmt"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	queryString := url.Query()
	query := queryString["query"]
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	}
	documents := find.FindDocuments(q)
	fmt.Fprintf(response, documents)
}

//Entry point for the web application
func main() {

	appLogFile := applog.Create()
	defer applog.Close(appLogFile)
	applog.Info("Started cnweb")


	//index.LoadKeywordIndex()
	//documents := index.FindForKeyword("你")
	http.HandleFunc("/find/", handler)
	http.ListenAndServe(":8080", nil)
}
