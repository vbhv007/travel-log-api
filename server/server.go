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
	r.Use(mux.CORSMethodMiddleware(r))
	addAPIs(r)
	return r
}

func addAPIs(router *mux.Router) {
	router.HandleFunc("/logs", api.Logs).Methods("POST")
	router.HandleFunc("/logs", api.EmptyResponse).Methods("OPTIONS")

	router.HandleFunc("/addLog", api.AddLog).Methods("POST")
	router.HandleFunc("/addLog", api.EmptyResponse).Methods("OPTIONS")

	router.PathPrefix("/").HandlerFunc(api.NotFound).Methods("GET", "POST")
	router.PathPrefix("/").HandlerFunc(api.EmptyResponse).Methods("OPTIONS")
}
