package main

import (
	"net/http"
	"log"
)

func startServer() {
	// HTTP
	err := http.ListenAndServe(":1630", nil)
	if err != nil {
		log.Println("HTTP failed:", err.Error())
	}
}
