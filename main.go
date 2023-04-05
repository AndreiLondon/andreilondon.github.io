package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// initialize the database
// db, err := sql.Open("sqlite3", "./forum.db")
// if err != nil {
// 	log.Fatal(err)
// }

// defer db.Close()
func main() {
	dbLocal, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	db = dbLocal
	defer db.Close()

	createUsersTable()

	// Register the request handlers
	// http.HandleFunc("/register", registerHandler)
	// http.HandleFunc("/login", loginHandler)
	// http.HandleFunc("/logout", logoutHandler)

	// Serve the static files in the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// start the server
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// func createTables() {
// 	err := createUsersTable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
