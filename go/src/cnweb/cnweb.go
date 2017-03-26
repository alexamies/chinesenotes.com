/*
Web application for finding documents in the corpus
*/
package main

import (
	"cnweb/find"
	"fmt"
	"log"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	queryString := url.Query()
	log.Println("cnweb.handler url: ", url)
	log.Println("cnweb.handler queryString length: ", len(queryString))
	query := queryString["query"]
	log.Println("cnweb.handler query len: ", len(query))
	q := "No Query"
	if len(query) > 0 {
		q = query[0]
	}
	documents := find.FindDocuments(q)
	fmt.Fprintf(response, documents)
}

//Entry point for the web application
func main() {
	log.Println("cnweb.main: Web app started")
	//index.LoadKeywordIndex()
	//documents := index.FindForKeyword("你")
	http.HandleFunc("/find", handler)
	http.ListenAndServe(":8080", nil)
}
