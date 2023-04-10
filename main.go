package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

//var db *sql.DB

// func dbRun() {
// 	createUsersTable()
// 	createPostsTable()

// }

func main() {

	// initialize the database

	dbLocal, err := sql.Open("sqlite3", "./forum.db")

	if err != nil {
		fmt.Println(err)
	}
	db = dbLocal
	defer db.Close()
	//dbRun()
	createUsersTable()
	// createPostsTable()

	// Serve the static files in the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/createpost", createpostHandler)
	http.HandleFunc("/savepost", savePostHanlder)

	// start the server
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}
