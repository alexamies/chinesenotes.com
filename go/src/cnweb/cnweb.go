/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/applog"
	"cnweb/find"
	"fmt"
	"encoding/json"
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
	results := find.FindDocuments(q)

	// If there is only one result, redirect to it
	if len(results.Collections) + len(results.Documents) == 1 {
		applog.Info("handler, unique result redirecting")
		if len(results.Collections) == 1 {
			url := "/" + results.Collections[0].GlossFile
			http.Redirect(response, request, url, http.StatusFound)
		} else {
			url := "/" + results.Documents[0].GlossFile
			http.Redirect(response, request, url, http.StatusFound)
		}
		return
	}

	// Otherwise send the results to the client in JSON form
	resultsJson, err := json.Marshal(results)
	if err != nil {
		applog.Error("main.handler error marshalling JSON, ", err)
	}
	applog.Info("handler, results returned: ", string(resultsJson))
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(response, string(resultsJson))
}

//Entry point for the web application
func main() {

	appLogFile := applog.Create()
	defer applog.Close(appLogFile)
	applog.Info("main.main Started cnweb")

	//index.LoadKeywordIndex()
	//documents := index.FindForKeyword("你")
	http.HandleFunc("/find/", handler)
	http.ListenAndServe(":8080", nil)
}
