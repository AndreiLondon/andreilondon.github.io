package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var invalidCredentialsFlagSignUp = false
var invalidCredentialsFlagSignIn = false

var posts = []Post{}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// t, err := template.ParseFiles("templates/signup.html", "templates/header.html", "templates/footer.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	if r.URL.Path != "/signup" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	username := strings.TrimSpace(r.FormValue("signup_username"))
	email := strings.TrimSpace(r.FormValue("signup_email"))
	password := strings.TrimSpace(r.FormValue("signup_password"))

	// if username == "" || email == "" || password == "" {
	// 	http.Error(w, "400 Bad Request.", http.StatusBadRequest)
	// 	return
	// }

	if username == "" || email == "" || password == "" || len(username) > 40 {
		invalidCredentialsFlagSignUp = true
		invalidCredentialsFlagSignIn = false
		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
		return
	}

	sessionId := generateSessionId()

	err := insertUser(username, email, encrypt(password), sessionId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
			invalidCredentialsFlagSignUp = true
			invalidCredentialsFlagSignIn = false
			http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
			return
		}
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}

	// t.ExecuteTemplate(w, "signup", nil)
	setCookie(w, sessionId)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func signHandler(w http.ResponseWriter, r *http.Request) {
	// templ, err := template.ParseFiles("templates/sign.html")
	// if err != nil {
	// 	showError(w, 500, "500 Internal Server Error")
	// 	return
	// }
	t, err := template.ParseFiles("templates/signup.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		// fmt.Println(err)
		showError(w, 500, "500 Internal Server Error")
		return
	}

	singCredentials := SingCredentials{}
	singCredentials.SignIn = invalidCredentialsFlagSignIn
	singCredentials.SignUp = invalidCredentialsFlagSignUp
	err = t.Execute(w, singCredentials)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}

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

// func signUpHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	userName := strings.TrimSpace(r.FormValue("username"))
// 	email := strings.TrimSpace(r.FormValue("email"))
// 	password := strings.TrimSpace(r.FormValue("password"))
// 	if userName == "" || email == "" || password == "" || len(userName) > 40 {
// 		invalidCredentialsFlagSignUp = true
// 		invalidCredentialsFlagSignIn = false
// 		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	if userName == "" || email == "" || password == "" {
// 		http.Error(w, "400 Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	sessionId := generateSessionId()

// 	// Perform user registration logic, e.g., save to database, generate session ID, etc.
// 	// Replace the placeholders with your actual implementation.
// 	sessionID, err := registerUser(userName, email, password)
// 	if err != nil {
// 		http.Error(w, "500 Internal Server Error. Error while registering user", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set session ID as a cookie for authenticated session
// 	setCookie(w, sessionID)

// 	// Redirect to homepage after successful registration
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }
