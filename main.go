package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/mrtzee/go-api-native/config"
	"github.com/mrtzee/go-api-native/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	r := mux.NewRouter()
	routes.RoutesIndex(r)

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
