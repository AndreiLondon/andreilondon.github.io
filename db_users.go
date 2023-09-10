package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//     _________users_________________________________________________
//     |  id      |  email    |  username  |  password  |  sessionId  |
//     |  INTEGER |  TEXT     |  TEXT      |  TEXT      |  TEXT       |
func crerateUsersTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT NOT NULL UNIQUE, username TEXT NOT NULL UNIQUE, password TEXT NOT NULL, sessionId TEXT)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}
func saveUser(username string, email string, password string, sessionId string) error {
	statement, err := db.Prepare("INSERT INTO users (email, username, password, sessionId) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(strings.ToLower(email), username, password, sessionId)
	if err != nil {
		return err
	}
	return nil
}
func getUserBySessionId(sessionId string) *User {
	if strings.TrimSpace(sessionId) == "" {
		return nil
	}
	rows, err := db.Query("SELECT * FROM users WHERE sessionId = ? LIMIT 1", sessionId)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var user *User = nil
	for rows.Next() {
		user = &User{}
		err = rows.Scan(&(user.Id), &(user.Email), &(user.Username), &(user.Password), &(user.Sessionid))
		if err != nil {
			return nil
		}
	}
	err = rows.Err()
	if err != nil {
		return nil
	}
	return user
}
func resetSessionId(sessionId string) error {
	statement, err := db.Prepare("UPDATE users SET sessionId = ? WHERE sessionId = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec("", sessionId)
	if err != nil {
		return err
	}
	return nil
}
func setSessionId(user *User, sessionId string) error {
	statement, err := db.Prepare("UPDATE users SET sessionId = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(sessionId, user.Id)
	if err != nil {
		return err
	}
	return nil
}
func checkUser(email string, password string) (*User, error) {
	rows, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&(user.Id), &(user.Email), &(user.Username), &(user.Password), &(user.Sessionid))
		if err != nil {
			return nil, err
		}
		if compairPasswords(user.Password, password) {
			return &user, nil
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func printUsers() {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&(user.Id), &(user.Email), &(user.Username), &(user.Password), &(user.Sessionid))
		if err != nil {
			return
		}
		fmt.Println(user)
	}
	err = rows.Err()
	if err != nil {
		return
	}
}
