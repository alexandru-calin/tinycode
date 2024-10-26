package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)

	log.Println("Starting the server on port 3000")
	http.ListenAndServe(":3000", mux)
}
