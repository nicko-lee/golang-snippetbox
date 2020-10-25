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

	// Create a fileserver which serves files out of the "./ui/static" directory.
	// Register the file server as the handler for all URL paths that start with "/static/"
	// For matching paths, strip the "/static" prefix before the request reaches the file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Spin up a new web server and pass it the servemux (router) we just created + a port number
	// If http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}