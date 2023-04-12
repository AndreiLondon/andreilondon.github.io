package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var posts = []Post{}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/signup.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Println(err)
	}

	if r.URL.Path != "/signup" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if username == "" || email == "" || password == "" {
		http.Error(w, "400 Bad Request.", http.StatusBadRequest)
		return
	}

	t.ExecuteTemplate(w, "signup", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "login", nil)

}

func createpostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createpost.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "createpost", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	printPosts()

	if err != nil {
		fmt.Println(err)
	}

	t.ExecuteTemplate(w, "index", posts)
}
func savePostHanlder(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/post" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	title := r.FormValue("title")
	text := r.FormValue("text")

	if title == "" || text == "" {
		fmt.Fprintf(w, "Complete all rows")
	} else {

		// insertPost(*user, postContent, postCategories)
		//insertPost(*user, title, text)
		insertPost(title, text)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userName := strings.TrimSpace(r.FormValue("username"))
	email := strings.TrimSpace(r.FormValue("email"))
	password := strings.TrimSpace(r.FormValue("password"))

	if userName == "" || email == "" || password == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	// Perform user registration logic, e.g., save to database, generate session ID, etc.
	// Replace the placeholders with your actual implementation.
	sessionID, err := registerUser(userName, email, password)
	if err != nil {
		http.Error(w, "500 Internal Server Error. Error while registering user", http.StatusInternalServerError)
		return
	}

	// Set session ID as a cookie for authenticated session
	setCookie(w, sessionID)

	// Redirect to homepage after successful registration
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
