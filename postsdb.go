package main

import (
	"fmt"
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

func savePost(user User, postContent string, postCategories string) error {
	statement, err := db.Prepare("INSERT INTO posts (userid, date, content, categories) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	date := getCurrentMilli()
	_, err = statement.Exec(user.Id, date, postContent, postCategories)
	if err != nil {
		return err
	}
	return nil
}

func getPosts(user *User) ([]Post, error) {
	if user == nil {
		user = &User{Id: -1}
	}
	posts := []Post{}
	sql := "SELECT posts.id, userid, username, date, content, categories, (SELECT COUNT(*) FROM post_likes WHERE status = 1 AND postid = posts.id) AS likes, (SELECT COUNT(*) FROM post_likes WHERE status = -1 AND postid = posts.id) AS dislikes, (SELECT SUM(status) from post_likes WHERE post_likes.postid = posts.id AND post_likes.userid = ? LIMIT 1) AS status FROM posts INNER JOIN users ON userid = users.id ORDER BY date DESC"
	rows, err := db.Query(sql, user.Id)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		post := Post{}
		var categories string
		var status interface{}
		err = rows.Scan(&(post.Id), &(post.Userid), &(post.Username), &(post.Date), &(post.Content), &categories, &(post.Likes), &(post.Dislikes), &status)
		if err != nil {
			return posts, err
		}
		categoriesArr := stringToSlice(categories, ",")
		post.Categories = categoriesArr
		dateFormat := formatMilli(post.Date)
		post.DateFormat = dateFormat
		s, ok := status.(int64)
		if ok {
			post.Status = int(s)
		} else {
			post.Status = 0
		}
		comments, err := getComments(user, post)
		if err != nil {
			return posts, err
		}
		post.Comments = comments
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return posts, err
	}
	return posts, nil
}
func getPostById(postId int) (*Post, error) {
	rows, err := db.Query("SELECT posts.id, userid, username, date, content, categories FROM posts INNER JOIN users ON userid = users.id WHERE posts.id = ? LIMIT 1", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	post := Post{}
	for rows.Next() {
		var categories string
		err = rows.Scan(&(post.Id), &(post.Userid), &(post.Username), &(post.Date), &(post.Content), &categories)
		if err != nil {
			return nil, err
		}
		categoriesArr := stringToSlice(categories, ",")
		post.Categories = categoriesArr
		dateFormat := formatMilli(post.Date)
		post.DateFormat = dateFormat
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &post, nil
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
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		//log.Fatal(err)
		return
	}
	defer rows.Close()
	//posts = []Post{}
	post := Post{}
	for rows.Next() {
		//var post Post
		var categories string
		err = rows.Scan(&(post.Id), &(post.Userid), &(post.Date), &(post.Content), &categories)
		//err = row.Scan(&(post.Id), &(post.Title), &(post.Text))
		if err != nil {
			//log.Fatal(err)
			return
		}
		post.Categories = stringToSlice(categories, ",")
		//fmt.Println("Post: ", post.Id, " ", post.Title, " ", post.Text)
		//posts = append(posts, post)
		fmt.Println(post)
	}
	err = rows.Err()
	if err != nil {
		//log.Fatal(err)
		return
	}
}

// func stringToSlice(data string, sep string) []string {
// 	sl := strings.Split(data, sep)
// 	result := []string{}

// 	for _, str := range sl {
// 		s := strings.TrimSpace(str)
// 		if s != "" {
// 			result = append(result, s)
// 		}
// 	}
// 	return result
// }
