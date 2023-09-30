package main

import (
	"net/http"
	"text/template"

	"github.com/ChristianMoesl/chat-server/log"
)

var logger = log.Logger

func getIndex(response http.ResponseWriter, request *http.Request) {
	logger.Info("Got request")

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(response, nil)
}

func main() {
	logger.Info("starting server...")

	http.HandleFunc("/", getIndex)

	logger.Fatal(http.ListenAndServe(":8080", nil))
}
