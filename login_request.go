package main

import (
	"fmt"
	"net/http"
)

// LoginHandler adds login HTTP requests to the job queue.
func LoginHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Runner) {
	// Make sure we can only be called with an HTTP POST request.
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Render success.
	w.WriteHeader(http.StatusCreated)

	uuid := r.FormValue("uuid")
	Log("Received request for UUID", uuid)
	fmt.Fprint(w, "UUID:", uuid)

}
