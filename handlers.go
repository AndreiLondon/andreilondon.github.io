package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if email is already taken
		db, err := getDB()
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&count)
		if err != nil {
			fmt.Println(err)
		}
		if count > 0 {
			http.Error(w, "Email already taken", http.StatusBadRequest)
			return
		}

		// Encrypt password
		encryptedPassword, err := encryptPassword(password)
		if err != nil {
			fmt.Println(err)
		}

		// Insert new user into database
		result, err := db.Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", email, username, encryptedPassword)
		if err != nil {
			fmt.Println(err)
		}
		userID, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}

		// Create a session for the registered user
		createSession(w, userID)

		// Redirect to home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Render the registration form
	tpl := template.Must(template.ParseFiles("register.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		// Check if email is present in the database
		db, err := getDB()
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		var user User
		err = db.QueryRow("SELECT id, email, username, password FROM users WHERE email=?", email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Compare password with stored password
		err = comparePasswords(user.Password, password)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Create a session for the logged in user
		createSession(w, int64(user.ID))

		// Render logged in user's username on the webpage
		tpl := template.Must(template.ParseFiles("loggedin.html"))
		err = tpl.Execute(w, user.Username)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Render the login form
	tpl := template.Must(template.ParseFiles("login.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Delete session cookie to logout user
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	http.SetCookie(w, cookie)

	// Redirect to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createSession(w http.ResponseWriter, userID int64) {
	// Generate a new UUID for the session
	sessionID := uuid.New().String()

	// Set the expiration time for the session
	expiration := time.Now().Add(time.Hour * 24) // 1 day

	// Insert the session into the database
	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO sessions (id, user_id, expiration) VALUES (?, ?, ?)", sessionID, userID, expiration)
	if err != nil {
		fmt.Println(err)
	}

	// Create a session cookie with the session ID
	cookie := &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func encryptPassword(password string) (string, error) {
	// TODO: Implement password encryption logic here
	// For demonstration purposes, return the password as is
	return password, nil
}

func comparePasswords(storedPassword, password string) error {
	// TODO: Implement password comparison logic here
	// For demonstration purposes, compare passwords as strings
	if storedPassword != password {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in
	authenticated, userID, err := authenticateSession(r)
	if err != nil {
		fmt.Println(err)
	}

	if authenticated {
		// Get user details from the database
		db, err := getDB()
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		var user User
		err = db.QueryRow("SELECT username FROM users WHERE id=?", userID).Scan(&user.Username)
		if err != nil {
			fmt.Println(err)
		}

		// Render home page with logged-in user's username
		tpl := template.Must(template.ParseFiles("home.html"))
		err = tpl.Execute(w, user.Username)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Render home page without logged-in user's username
		tpl := template.Must(template.ParseFiles("home.html"))
		err = tpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func authenticateSession(r *http.Request) (bool, int64, error) {
	// Get session ID from cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			return false, 0, nil // No session ID found
		}
		return false, 0, err
	}

	// Query the database to check if session ID exists and is not expired
	db, err := getDB()
	if err != nil {
		return false, 0, err
	}
	defer db.Close()

	var userID int64
	var expiration time.Time
	err = db.QueryRow("SELECT user_id, expiration FROM sessions WHERE session_id=?", cookie.Value).Scan(&userID, &expiration)
	if err != nil {
		return false, 0, err
	}

	// Check if session is expired
	if time.Now().After(expiration) {
		// Delete expired session from the database
		_, err = db.Exec("DELETE FROM sessions WHERE session_id=?", cookie.Value)
		if err != nil {
			return false, 0, err
		}

		return false, 0, nil // Session expired
	}

	return true, userID, nil // Session is valid
}
