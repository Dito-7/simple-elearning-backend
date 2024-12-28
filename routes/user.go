package routes

import (
	"github.com/Dito-7/simple-elearning-backend/controllers"
	"github.com/gorilla/mux"
)

// Schedule Routes
func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	// Schedule Routes
	router.HandleFunc("", controllers.UserController{}.Index).Methods("GET")
	router.HandleFunc("", controllers.UserController{}.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", controllers.UserController{}.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", controllers.UserController{}.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", controllers.UserController{}.Delete).Methods("DELETE")
}
