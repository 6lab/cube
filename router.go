package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func routes() {
	// Routing
	router := mux.NewRouter()

	router.HandleFunc("/", appStart).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("")))

	// http.Handle("/", router)
	http.Handle("/", router)
}

func appStart(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
