package main

import (
	"database/sql"
	"fmt"
	"log"

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

// Creating a database
// func createUsersTable() error {
// 	users_table := `CREATE TABLE IF NOT EXISTS users (
//         id INTEGER PRIMARY KEY,
//         username TEXT NOT NULL UNIQUE,
//         email TEXT NOT NULL UNIQUE,
//         password TEXT NOT NULL,
// 		sessionID TEXT,
// 		)`
// 	//fmt.Println("Create user table...")
// 	statement, err := db.Prepare(users_table) // Prepare SQL Statement
// 	if err != nil {
// 		fmt.Println(err)
// 		//return err
// 	}
// 	defer statement.Close()
// 	statement.Exec() // Execute SQL Statements
// 	fmt.Println("Table created successfully!")

// 	return nil
// 	//defer db.Close()
// }

// Check if the email or username already exists in the "users" table
func isUserExists(email string, username string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? OR username = ?", email, username).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, username string, email string, password string) {
	log.Println("Inserting user record ...")
	insertUser := `INSERT INTO users(username, email, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUser) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err)
		//return
	}
	_, err = statement.Exec(username, email, password)
	if err != nil {
		fmt.Println(err)
		//return
	}
}

// getting data from table
func displayUser(db *sql.DB) {
	row, err := db.Query("SELECT * FROM users ORDER BY name")
	if err != nil {
		fmt.Println(err)
		//return
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var user User
		// var id int
		// var username string
		// var email string
		// var password string
		err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			fmt.Println(err)
			//row.Scan(&username, &email, &password)
			fmt.Println("User: ", user.Username, " ", user.Email, " ", user.Password)
		}
	}
}
