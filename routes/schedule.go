package routes

import (
	"simple-elearning-backend/controllers/schedulecontroller"

	"github.com/gorilla/mux"
)

func ScheduleRoutes(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.HandleFunc("", schedulecontroller.Index).Methods("GET")
	router.HandleFunc("", schedulecontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", schedulecontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", schedulecontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", schedulecontroller.Delete).Methods("DELETE")
}
