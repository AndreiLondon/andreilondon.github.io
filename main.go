package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//var db *sql.DB

// func dbRun() {
// 	createUsersTable()
// 	createPostsTable()

// }

var invalidCredentialsFlagSignUp = ""
var invalidCredentialsFlagSignIn = ""
var emptyPostFlag = false
var emptyCommentFlag = false
var SHAW_ALL = "Shaw All"
var currentMode = SHAW_ALL
var SESSION_ID = "SESSION_ID"
var filterCategories = []string{}
var MY_LIKES = "My Likes"
var MY_POSTS = "My Posts"
var MY_COMMENTS = "My Comments"

func main() {

	// initialize the database

	dbLocal, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	db = dbLocal
	defer db.Close()
	//dbRun()
	// createUsersTable()
	// createPostsTable()
	// printPosts()
	createTables()
	// Serve the static files in the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/sign", signHandler)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signout", signOutHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/savepost", savePostHandler)
	http.HandleFunc("/createpost", createpostHandler)
	http.HandleFunc("/registerlike", registerlikeHandler)
	http.HandleFunc("/registercommentlike", registercommentlikeHandler)
	http.HandleFunc("/comment", commentHandler)
	http.HandleFunc("/commentsubmit", commentsubmitHandler)
	http.HandleFunc("/setfilter", setfilterHandler)
	http.HandleFunc("/removefilter", removefilterHandler)
	http.HandleFunc("/changemode", changemodeHandler)

	// start the server
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func createTables() {
	err := createUsersTable()
	if err != nil {
		log.Fatal(err)
	}
	err = crerateCategoriesTable()
	if err != nil {
		log.Fatal(err)
	}
	err = insertCategories([]string{"C++", "C#", "Java", "JavaScript", "HTML", "CSS", "PHP", "Go", "Rust", "Node"})
	if err != nil {
		if !strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
			log.Fatal(err)
		}
	}
	err = createPostsTable()
	if err != nil {
		log.Fatal(err)
	}
	err = createCommentsTable()
	if err != nil {
		log.Fatal(err)
	}
	err = createPostLikesTable()
	if err != nil {
		log.Fatal(err)
	}
	err = createCommentLikesTable()
	if err != nil {
		log.Fatal(err)
	}
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
