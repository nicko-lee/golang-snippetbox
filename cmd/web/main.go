package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flag. Params are flag name, default value, and brief description.
	// The value of the flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse CLI flag
	flag.Parse()


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

	// Spin up a new web server and pass it the servemux (router) we just created + a port number (user supplied via CLI flag)
	// If http.ListenAndServe() returns an error we use the log.Fatal() function to log the error message and exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}