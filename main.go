package main

import (
	"encoding/json"
	"net/http"
)

var urls = make(map[string]string)

func AddURL(w http.ResponseWriter, r *http.Request) {
	// Parse the form so that we have access to all the POST
	// params.
	err := r.ParseForm()
	if err != nil {
		// To handle HTTP errors, you can call http.Error
		http.Error(w, "Internal Server Error", 500)
		// Always return after handling HTTP errors! Otherwise, the program
		// keeps evaluating, which could have unintended effects
		return
	}
	url := r.Form.Get("url")
	short := r.Form.Get("short")
	// TODO 1. Blacklist short urls "add" and "list", since they are routes in our app.
	// Return a 500 and a useful error message.

	// Now we can register the variables to our map, essentially "saving" the URL
	urls[short] = url
	w.Write([]byte("ok"))
}

func ShowURL(w http.ResponseWriter, r *http.Request) {
	// Display URLs in the "urls" map
	// HINT: make requests to /list to see your map when debugging
	jsonString, _ := json.Marshal(urls)
	w.Write(jsonString)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// TODO 1: Parse the request and save the shortcode

	// Assuming it's a "valid" shortcode, we need to make sure it exists
	//
	// TODO 2: Verify that the shortcode is in our map. If it's not,
	// return a 404 Not Found error message.
	//
	// HINT: Accessing a map returns 2 variables, one for the value if it
	// exists and one for a boolean if it was found.

	// TODO 3: Return a 302 redirect to the correct URL
}

func main() {
	http.HandleFunc("/add", AddURL)
	http.HandleFunc("/list", ShowURL)
	http.HandleFunc("/", RedirectHandler)
	http.ListenAndServe(":8080", nil)
}
