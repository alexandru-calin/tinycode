package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to tinycode!")
	})

	log.Println("Starting the server on port 3000")
	http.ListenAndServe(":3000", mux)
}
