package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"simple-elearning-backend/config"
	"simple-elearning-backend/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	r := mux.NewRouter()
	routes.RoutesIndex(r)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
