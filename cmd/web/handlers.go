package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"html/template"
)

// Home handler function which writes a byte slice containing "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Check if current request URL path exactly matches "/". If it doesn't, use the http.NotFound()
	// function to send a 404 response to the client. Importantly don't forget to return from the handler.
	// Else it would keep executing and also write the "Hello from Snippetbox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return 
	}

	// Initialize a slice with paths to the two files. Order is important
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}


	// Read in template file. If there is an error while reading, log the detailed error message and use
	// the http.Error() function to send a 500 Internal Server Error response to the user.
	// Can pass the slice of file paths as a variadic param
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// If no issue with template parsing, then we execute the template by using the Execute() method on the
	// template set to write the  template content as the response body. The last parameter to Execute()
	// lets us pass in any dynamic data we want to pass to the template
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// Handler function for /snippet route
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract value from id parameter from query string and try to convert to integer for the next step
	// which is validation of user input
	// If can't be converted to int or if value is less than 1, we return 404
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use fmt. Fprintf() to interpolate the id value with our response and write it to 
	// http.ResponseWriter. This is easier to use than w.Write([]byte("Hello"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Handler function for /snippet/create route
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Check whether the request is using POST (we only want to respond to POST methods on this route).
	// Note that http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodPost) // tell the user we only accept POST here
		http.Error(w, "Method Not Allowed", 405)
		return // so it never runs the line below
	}

	w.Write([]byte("Create a new snippet..."))
}

