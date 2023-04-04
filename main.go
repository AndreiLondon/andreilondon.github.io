package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a handler function for the "/" route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the HTTP server and listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
