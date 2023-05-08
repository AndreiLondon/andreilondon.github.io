package main

import (
	"database/sql"
	"fmt"
	"html/template"
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
	createPostsTable()
	printPosts()

	// Serve the static files in the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/sign", signupHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/createpost", createpostHandler)
	http.HandleFunc("/savepost", savePostHanlder)

	// start the server
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}

func showError(w http.ResponseWriter, code int, message string) {
	templ, err := template.ParseFiles("templates/error.html")
	w.WriteHeader(code)
	if err != nil {
		fmt.Fprint(w, "500 Internal Server Error")
		return
	}
	templ.Execute(w, message)
}
