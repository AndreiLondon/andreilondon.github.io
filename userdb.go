package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

//     _________users_________________________________________________
//     |  id      |  email    |  username  |  password  |  sessionId  |
//     |  INTEGER |  TEXT     |  TEXT      |  TEXT      |  TEXT       |

func createUsersTable() error {

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT NOT NULL UNIQUE, username TEXT NOT NULL UNIQUE, password TEXT NOT NULL, sessionId TEXT)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(username string, email string, password string, sessionId string) error {
	statement, err := db.Prepare("INSERT INTO users (email, username, password, sessionId) VALUES(?, ?, ?, ?)")
	// This is good to avoid SQL injections
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

//Check the user

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

//Display user

// getting data from table
func displayUsers() {
	//row, err := db.Query("SELECT * FROM users ORDER BY name")
	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
		//return
	}
	defer row.Close()
	/*
		The row variable is closed using the defer statement,
		which ensures that the row.Close() method is called at
		the end of the function to close the result cursor and
		free up resources.
	*/
	for row.Next() { // Iterate and fetch the records from result cursor
		user := User{}
		err = row.Scan(&(user.Id), &(user.Email), &(user.Username), &(user.Password), &(user.Sessionid))
		/*
		   Inside the loop, a User struct is created to store the data of the current row.
		   The row.Scan() method is used to scan the values from the current row into
		   the fields of the User struct, which represents the columns of the users table
		   in the same order. The &user.ID, &user.Username, &user.Email, and &user.Password
		   are pointers to the fields of the User struct, where the corresponding
		   column values from the current row are scanned.
		*/
		if err != nil {
			//log.Fatal(err)
			//row.Scan(&username, &email, &password)
			return
		}
		fmt.Println("User: ", user.Username, " ", user.Email, " ", user.Password)
	}
	err = row.Err()
	if err != nil {
		//log.Fatal(err)
		return
	}
}

//Set SesioniD

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

//Reset Session

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
