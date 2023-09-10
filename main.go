package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var invalidCredentialsFlagSignUp = ""
var invalidCredentialsFlagSignIn = ""
var emptyPostFlag = false
var emptyCommentFlag = false
var SESSION_ID = "SESSION_ID"
var filterCategories = []string{}
var SHAW_ALL = "Shaw All"
var MY_POSTS = "My Posts"
var MY_COMMENTS = "My Comments"
var MY_LIKES = "My Likes"
var currentMode = SHAW_ALL

func main() {
	dbLocal, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	db = dbLocal
	defer db.Close()
	createTables()
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/sign", signHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signout", signoutHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/savepost", savepostHandler)
	http.HandleFunc("/registerlike", registerlikeHandler)
	http.HandleFunc("/registercommentlike", registercommentlikeHandler)
	http.HandleFunc("/comment", commentHandler)
	http.HandleFunc("/commentsubmit", commentsubmitHandler)
	http.HandleFunc("/setfilter", setfilterHandler)
	http.HandleFunc("/removefilter", removefilterHandler)
	http.HandleFunc("/changemode", changemodeHandler)
	fmt.Println("Server start at port :9000")
	http.ListenAndServe(":9000", nil)
}
func createTables() {
	err := crerateUsersTable()
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
	err = creratePostsTable()
	if err != nil {
		log.Fatal(err)
	}
	err = crerateCommentsTable()
	if err != nil {
		log.Fatal(err)
	}
	err = creratePostLikesTable()
	if err != nil {
		log.Fatal(err)
	}
	err = crerateCommentLikesTable()
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
