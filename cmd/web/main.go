package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize new servemux (router)
	// Then register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Spin up a new web server and pass it the servemux (router) we just created + a port number
	// If http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}