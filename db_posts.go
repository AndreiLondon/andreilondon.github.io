package main

import "fmt"

//      _________posts________________________________________________
//     |  id       |  userid   |  date     |  content  |  categories  |
//     |  INTEGER  |  INTEGER  |  INTEGER  |  TEXT     |  TEXT        |

//Create posts table
func creratePostsTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS posts(id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, date INTEGER NOT NULL, content TEXT NOT NULL, categories TEXT NOT NULL)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
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

func printPosts() {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return
	}
	defer rows.Close()
	post := Post{}
	for rows.Next() {
		var categories string
		err = rows.Scan(&(post.Id), &(post.Userid), &(post.Date), &(post.Content), &categories)
		if err != nil {
			return
		}
		post.Categories = stringToSlice(categories, ",")
		fmt.Println(post)
	}
	err = rows.Err()
	if err != nil {
		return
	}
}
