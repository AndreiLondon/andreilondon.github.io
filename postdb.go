package main

import (
	"fmt"
	"time"
)

//      _________posts________________________________________________
//     |  id       |  userid   |  date     |  content  |  categories  |
//     |  INTEGER  |  INTEGER  |  INTEGER  |  TEXT     |  TEXT        |

func createPostsTable() error {
	posts_table := `CREATE TABLE IF NOT EXISTS posts (
	id INTEGER PRIMARY KEY,
	userid TEXT NOT NULL,
	date INTEGER,
	content TEXT,
	categories TEXT,
	);`
	statement, err := db.Prepare(posts_table)
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

// We are passing db reference connection from main to our method with other parameters
// func insertPost(user User, postContent string, postCategories string) error {
func insertPost(postContent string, postCategories string) error {
	//log.Println("Inserting post record ...")
	insertPost := `INSERT INTO posts(userid, date, content, categories) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertPost) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err)
		return err
	}
	//defer statement.Close()
	date := getCurrentMillisecods()
	//_, err = statement.Exec(user.Id, date, postContent, postCategories)
	_, err = statement.Exec(date, postContent, postCategories)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
	// defer statement.Close()
}

// getCurrentMilliseconds retrieves the current time in milliseconds since the Unix epoch.
func getCurrentMillisecods() int64 {
	// Get the current time as a Time value
	now := time.Now()
	// Convert the Time value to milliseconds by dividing UnixNano() by 1e6
	return now.UnixNano() / 1e6
}

// formatMilliseconds formats a date in milliseconds since the Unix epoch as a string in "02-Jan-2006 15:04:05" format.
func formatMilliseconds(date int64) string {
	// Convert the input date from milliseconds to seconds by dividing by 1000, and create a Time value
	t := time.Unix(0, date*int64(time.Millisecond))
	// Format the Time value using the desired format string "02-Jan-2006 15:04:05"
	return t.Format("26-March-2020 13:30:30")
}
