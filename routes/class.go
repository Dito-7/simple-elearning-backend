package routes

import (
	"simple-elearning-backend/controllers/classcontroller"

	"github.com/gorilla/mux"
)

func ClassRoutes(r *mux.Router) {
	router := r.PathPrefix("/classes").Subrouter()

	router.HandleFunc("", classcontroller.Index).Methods("GET")
	router.HandleFunc("", classcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", classcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", classcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", classcontroller.Delete).Methods("DELETE")
}
