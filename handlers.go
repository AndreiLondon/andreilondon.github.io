package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
)

var (
	db             *sql.DB
	cookieStore    *securecookie.SecureCookie
	sessionName    = "session"
	sessionMaxAge  = 60 * 60 * 24 * 7 // 1 week
	cookieHashKey  = []byte("your-secret-cookie-hash-key")
	cookieBlockKey = []byte("your-secret-cookie-block-key")
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/header.html", "templates/index.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "500 Internal error", http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "index", nil)
	//err = t.Execute(w, "index")
	if err != nil {
		http.Error(w, "500 Internal error", http.StatusInternalServerError)
		return
	}
}

/*
The template.ParseFiles() function loads three separate templates from
the templates directory: header.html, index.html, and footer.html.
These templates are represented by a single *template.Template value t.
If there is an error parsing any of the templates,
an HTTP 500 error response is returned to the client with the error message.
The t.ExecuteTemplate() function is used to execute the index.html template,
passing in an empty interface value (nil) as data.
The resulting HTML is written to the http.ResponseWriter object w.
If there is an error executing the template, an HTTP 500 error response
is returned to the client with the error message.

The difference between the two lines of code you provided is in the template
that is being executed and the data that is being passed to it.

In the first line of code, t.ExecuteTemplate(w, "index", nil),
the ExecuteTemplate() method is being called on the *template.Template object t.
This method executes a specific template identified by its name
(in this case, "index") and writes the result to the http.ResponseWriter object w.
The third argument of nil means that no data is being passed to the template.

In the second line of code, t.Execute(w, index), the Execute() method
is being called on the same *template.Template object t.
This method executes the first template that was parsed from the file(s)
passed to ParseFiles(), and writes the result to the http.ResponseWriter object w.
The second argument, index, is a data object that is being passed to the template.
In this case, index is a variable or constant containing some data
that will be used in the template to render the dynamic content.

In summary, ExecuteTemplate() is used to execute a specific named template,
while Execute() is used to execute the first template that was parsed from the file(s).
Additionally, Execute() can take a data object as a second argument,
 while ExecuteTemplate() requires that any data to be
 passed to the template be done so using a call to t.Execute().
*/

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if email is already taken
	var user User
	err := db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&user.ID)
	if err == nil {
		// Email is already taken
		http.Error(w, "Email is already taken", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword := hashPassword(password)

	// Insert new user into database
	result, err := db.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", email, username, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	id, _ := result.LastInsertId()
	fmt.Fprintf(w, "User %d registered successfully", id)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Retrieve user from database
	var user User
	err := db.QueryRow("SELECT id, email, username, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		// User not found
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	// Check if password is correct
	if !checkPassword(password, user.Password) {
		// Incorrect password
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return

	}
	// Create a session for the user and store it in a cookie
	sessionID, err := createSession(user.ID)
	if err != nil {
		http.Error(w, "Could not create session", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	// Redirect the user to the forum homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
