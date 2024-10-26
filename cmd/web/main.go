package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	app := &application{}

	log.Println("Starting the server on port 3000")
	http.ListenAndServe(":3000", app.routes())
}
