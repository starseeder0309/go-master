package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	name := queries.Get("name")
	if name == "" {
		name = "World"
	}

	id, _ := strconv.Atoi(queries.Get("id"))

	fmt.Fprintf(w, "Hello, %s! id: %d", name, id)
}
