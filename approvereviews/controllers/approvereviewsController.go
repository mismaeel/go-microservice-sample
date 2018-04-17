package controllers

import (
//	"encoding/json"
	"net/http"

	//"github.com/gorilla/mux"
	//"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/movies"
// Returns all Movie documents
func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("huston we are live !!!!!!!!!!!!!"))
}
