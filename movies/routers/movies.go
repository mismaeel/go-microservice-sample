package routers

import (
	"github.com/gorilla/mux"
	"github.com/mismaeel/moviesapp/movies/controllers"
)

func setMovieRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/addreview/{id}", controllers.AddMovieReview).Methods("PUT")
	return router
}
