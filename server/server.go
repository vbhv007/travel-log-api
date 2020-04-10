package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vbhv007/travel-log-api/api"
	"github.com/vbhv007/travel-log-api/database"
	"log"
	"net/http"
)

const (
	PORT = 8080
)

func Start() {
	startExternalServer()
}

func startExternalServer() {
	fmt.Println("Starting server on port:", PORT)
	dbErr := database.LogEntityDao.Migrate()
	if dbErr != nil {
		panic("Unable to migrate DB")
	}
	externalRouter := buildExternalRouter()
	http.Handle("/", externalRouter)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), externalRouter))
}

func buildExternalRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	addAPIs(r)
	return r
}

func addAPIs(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(api.NotFound)
}
