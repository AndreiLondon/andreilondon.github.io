package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//      _________posts________________________________________________
//     |  id       |  userid   |  date     |  content  |  categories  |
//     |  INTEGER  |  INTEGER  |  INTEGER  |  TEXT     |  TEXT        |

func createPostsTable() error {
	//statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, date INTEGER, content TEXT, categories TEXT)")
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts (id INTEGER PRIMARY KEY, content TEXT, categories TEXT)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

// We are passing db reference connection from main to our method with other parameters
// func insertPost(user User, postContent string, postCategories string) error {
func insertPost(title string, text string) error {
	//log.Println("Inserting post record ...")
	//statement, err := db.Prepare("INSERT INTO posts (userid, date, content, categories) VALUES(?, ?, ?, ?)") // Prepare statement.
	statement, err := db.Prepare("INSERT INTO posts (content, categories) VALUES(?, ?)") // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()
	// date := getCurrentMilliseconds()
	// _, err = statement.Exec(user.Id, date, postContent, postCategories)
	_, err = statement.Exec(title, text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
	// defer statement.Close()
}

// getCurrentMilliseconds retrieves the current time in milliseconds since the Unix epoch.
func getCurrentMilliseconds() int64 {
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

func printPosts() {
	row, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	posts = []Post{}
	//post := Post{}
	for row.Next() {
		var post Post
		//var categories string
		//err = rows.Scan(&(post.Id), &(post.Userid), &(post.Date), &(post.Content), &categories)
		err = row.Scan(&(post.Id), &(post.Title), &(post.Text))
		if err != nil {
			log.Fatal(err)
		}
		// post.Categories = stringToSlice(categories, ",")
		//fmt.Println("Post: ", post.Id, " ", post.Title, " ", post.Text)
		posts = append(posts, post)
	}
	err = row.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func stringToSlice(data string, sep string) []string {
	sl := strings.Split(data, sep)
	result := []string{}

	for _, str := range sl {
		s := strings.TrimSpace(str)
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
