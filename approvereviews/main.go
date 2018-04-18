package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

type Movie struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Title     string        `json:"title"`
		Rating    float32       `json:"rating"`
		Director  string        `json:"director"`
		Actors    []string      `json:"actors"`
		CreatedOn time.Time     `json:"createdAt,omitempty"`
		Review    string      `json:"review,omitempty"`
	}

func approveMovieReview(w http.ResponseWriter, r *http.Request) {

		// Read body
		var m Movie

	        err := json.NewDecoder(r.Body).Decode(&m)
	        if err != nil {
	            http.Error(w, err.Error(), 400)
	            return
	        }


		// simple review processing
		time.Sleep(100 * time.Millisecond)
		if  len(m.Review) < 5 {
           fmt.Fprintf(w, "notapproved")
        }else {
         fmt.Fprintf(w, "approved")
    }



}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(false)

  myRouter.HandleFunc("/approvereview", approveMovieReview).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
