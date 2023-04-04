package database

import (
	"database/sql"
	"fmt"
)

func createUsersTable(db *sql.DB) /**NewUserDataBase*/ {
	users_table := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		sessionID TEXT,
		);`
	query, err := db.Prepare(users_table)
	if err != nil {
		fmt.Println(err)
		return
	}
	query.Exec()
	fmt.Println("Table created successfully!")
	//query.Close()
	//defer db.Close()
}
