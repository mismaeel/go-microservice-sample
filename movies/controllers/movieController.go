package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"io/ioutil"
   "bytes"
	"github.com/gorilla/mux"
	"github.com/mismaeel/moviesapp/movies/common"
	"github.com/mismaeel/moviesapp/movies/data"
	"github.com/mismaeel/moviesapp/movies/models"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/movies"
// Returns all Movie documents
func GetMovies(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("movies")
	repo := &data.MovieRepository{c}
	movies := repo.GetAll()
	j, err := json.Marshal(MoviesResource{Data: movies})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Post - "/movies"
// Insert a new Movie document
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var dataResourse MovieResource

	// Decode the incoming Movie json
	err := json.NewDecoder(r.Body).Decode(&dataResourse)

	if err != nil {
		common.DisplayAppError(w, err, "Invalid Movie data", 500)
		return
	}


	movie := &dataResourse.Data
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

 if len(movie.Title)< 3 || len(movie.Title) > 50 || IsLetter(movie.Title) == false {
  common.DisplayAppInfo(w,"Validation Error: Movie Title shoud by 3-50 letters only"  , 200)
	return
 }


	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("movies")
	// Insert a movie document
	repo := &data.MovieRepository{c}
	repo.Create(movie)
	j, err := json.Marshal(dataResourse)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


// Handler for HTTP Post - "/movies/addreview"
// Insert a new Movie document
func AddMovieReview(w http.ResponseWriter, r *http.Request) {
	var dataResourse MovieResource

	// Decode the incoming Movie json
	err := json.NewDecoder(r.Body).Decode(&dataResourse)

	if err != nil {
		common.DisplayAppError(w, err, "Invalid Movie data", 500)
		return
	}
	movie := &dataResourse.Data
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// replying to the http request
	common.DisplayAppInfo(w,"Review recieved it will be puplished after apprvoal", 200)
	// calling the approval service Async http call
	ch := make(chan string)
  go post(movie, ch)

	// recieving from chanel the response
  body := <-ch

    // updating the database if the review was approved
    if body == "approved"{

			// create new context
			context := NewContext()
			defer context.Close()
			c := context.DbCollection("movies")
			// Insert a movie document
			repo := &data.MovieRepository{c}
			repo.Update(id,movie)
		}

}

func post(movie *models.Movie, ch chan string) {
	url := "http://approvereviews:8080/approvereview"

	u := models.Movie{
		Id: movie.Id,
		Title: movie.Title,
		Rating: movie.Rating,
		Director: movie.Director,
		Actors: movie.Actors,
		CreatedOn: movie.CreatedOn,
		Review: movie.Review,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	req, err := http.NewRequest("POST", url,b)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- string(body)

}

// Handler for HTTP Get - "/movies/{id}"
// Get movie by id
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("movies")
	repo := &data.MovieRepository{c}

	// Get movie by id
	movie, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	j, err := json.Marshal(MovieResource{Data: movie})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/movies/{id}"
// Delete movie by id
func DeleteMovie(rw http.ResponseWriter, req *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(req)
	id := vars["id"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("movies")
	repo := &data.MovieRepository{c}

	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(rw, err, "An unexpected error has occurred", 500)
			return
		}
	}
	rw.WriteHeader(http.StatusNoContent)
}
