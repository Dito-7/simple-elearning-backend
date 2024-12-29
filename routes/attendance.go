package routes

import (
	"simple-elearning-backend/controllers/attendancecontroller"

	"github.com/gorilla/mux"
)

func AttendanceRoutes(r *mux.Router) {
	router := r.PathPrefix("/attenedance").Subrouter()

	router.HandleFunc("", attendancecontroller.Index).Methods("GET")
	router.HandleFunc("", attendancecontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", attendancecontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", attendancecontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", attendancecontroller.Delete).Methods("DELETE")
}
