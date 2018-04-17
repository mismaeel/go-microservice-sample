package routers

import (
	"github.com/gorilla/mux"
	"github.com/mismaeel/moviesapp/approvereviews/controllers"
)

func setApproveReviewsRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/approvereview", controllers.GetMovies).Methods("GET")
	
	return router
}
