package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
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

func main() {
	// Initialize database connection and cookie store
	db, _ = sql.Open("sqlite3", "forum.db")
	cookieStore = securecookie.New(cookieHashKey, cookieBlockKey)

	// Initialize HTTP router
	r := mux.NewRouter()

	// Routes for user registration and login
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	// Routes for forum posts and comments (authentication required)
	r.HandleFunc("/posts", postsHandler).Methods("GET")
	r.HandleFunc("/posts", addPostHandler).Methods("POST")
	r.HandleFunc("/posts/{id}/comments", commentsHandler).Methods("GET")
	r.HandleFunc("/posts/{id}/comments", addCommentHandler).Methods("POST")

	// Start HTTP server
	http.ListenAndServe(":8080", r)
}
