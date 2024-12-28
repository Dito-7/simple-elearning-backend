package routes

import (
	"github.com/Dito-7/simple-elearning-backend/controllers"
	"github.com/gorilla/mux"
)

// Schedule Routes
func ScheduleRoutes(r *mux.Router) {
	router := r.PathPrefix("/schedules").Subrouter()

	// Schedule Routes
	router.HandleFunc("", controllers.ScheduleController{}.Index).Methods("GET")
	router.HandleFunc("", controllers.ScheduleController{}.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", controllers.ScheduleController{}.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", controllers.ScheduleController{}.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", controllers.ScheduleController{}.Delete).Methods("DELETE")
}
