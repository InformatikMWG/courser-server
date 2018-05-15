package main

import (
	"fmt"
	"net/http"
)

// SimpleLoginHandler prints the posted UUID to the console and client browser.
func SimpleLoginHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Runner) {

	// Make sure we can only be called with an HTTP GET request.
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Render success.
	w.WriteHeader(http.StatusCreated)

	// Get field UUID from request.
	uuid := r.FormValue("uuid")
	// Print UUID to log and to client browser.
	Log("Received request for UUID", uuid)
	fmt.Fprint(w, "UUID:", uuid)
}
