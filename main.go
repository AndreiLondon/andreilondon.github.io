package main

import (
	"fmt"
	"net/http"
)

func main() {
	const portNumber = ":8080"
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", indexHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// Start the HTTP server and listen on port 8080
	fmt.Printf("Starting application on port %s", portNumber)
	server.ListenAndServe()
}
