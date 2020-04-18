package server

import (
	"fmt"
	"github.com/gorilla/handlers"
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
	AllowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost:8080"})
	externalRouter := buildExternalRouter()
	http.Handle("/", externalRouter)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), handlers.CORS(AllowedOrigins)(externalRouter)))
}

func buildExternalRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	addAPIs(r)
	return r
}

func addAPIs(router *mux.Router) {
	router.HandleFunc("/logs", api.Logs).Methods("POST")
	router.HandleFunc("/addLog", api.AddLog).Methods("POST")
	router.PathPrefix("/").HandlerFunc(api.NotFound).Methods("GET", "POST")
}
