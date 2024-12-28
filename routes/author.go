package routes

import (
	"github.com/gorilla/mux"
	"github.com/mrtzee/go-api-native/controllers/authorcontroller"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/authors").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	router.HandleFunc("", authorcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", authorcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", authorcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", authorcontroller.Delete).Methods("DELETE")
}
