package handlers

import (
	"net/http"
)

// NotFoundHandler handles not found requests
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	WriteErrorMessage(w, http.StatusNotFound, "could not find page")
}
