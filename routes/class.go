package routes

import (
	"github.com/Dito-7/simple-elearning-backend/controllers"
	"github.com/gorilla/mux"
)

func ClassRoutes(r *mux.Router) {
	router := r.PathPrefix("/classes").Subrouter()

	router.HandleFunc("", controllers.ClassController{}.Index).Methods("GET")
	router.HandleFunc("", controllers.ClassController{}.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", controllers.ClassController{}.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", controllers.ClassController{}.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", controllers.ClassController{}.Delete).Methods("DELETE")
}
