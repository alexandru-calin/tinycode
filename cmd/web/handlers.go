package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to tinycode!")
}

func codeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Malformed ID", http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Viewing code snippet with ID %d", id)
	fmt.Fprintln(w, message)
}
