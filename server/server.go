package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vbhv007/travel-log-api/api"
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
	router.PathPrefix("/logs").HandlerFunc(api.Logs).Methods("POST")
	router.PathPrefix("/").HandlerFunc(api.NotFound).Methods("GET", "POST")
}