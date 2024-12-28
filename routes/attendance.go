package routes

import (
	"github.com/Dito-7/simple-elearning-backend/controllers"
	"github.com/gorilla/mux"
)

func AttendanceRoutes(r *mux.Router) {
	router := r.PathPrefix("/attendance").Subrouter()

	router.HandleFunc("", controllers.AttendanceController{}.Index).Methods("GET")
	router.HandleFunc("", controllers.AttendanceController{}.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", controllers.AttendanceController{}.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", controllers.AttendanceController{}.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", controllers.AttendanceController{}.Delete).Methods("DELETE")
}
