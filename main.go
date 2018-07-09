package main

import (
	"net/http"
	"os"

	"google.golang.org/appengine" // Required external App Engine library
	"google.golang.org/appengine/log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	// Set Headers
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	log.Infof(ctx, "Endpoint reached "+r.URL.Path)
	// Check Endpoint for Secure Endpoint

	// Check Key
	if r.Header.Get("Authorization") != os.Getenv("SECURE_TOKEN") {
		http.Error(w, "Bad Shared Key", http.StatusForbidden)
		return
	} else {
		w.WriteHeader(200)
		w.Write([]byte("Success"))
		return
	}

}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main() // Starts the server to receive requests
}
