package router

import (
	"github.com/gorilla/mux"
	"github.com/von-salumbides/go/20-mongo/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("DELETE")
	router.HandleFunc("/api/delall", controller.GetAllMovies).Methods("DELETE")

	return router
}
