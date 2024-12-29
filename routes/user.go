package routes

import (
	"simple-elearning-backend/controllers/usercontroller"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.HandleFunc("", usercontroller.Index).Methods("GET")
	router.HandleFunc("", usercontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", usercontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", usercontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", usercontroller.Delete).Methods("DELETE")
}
