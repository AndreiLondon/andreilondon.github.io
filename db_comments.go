package main

import "fmt"

//       ________comments___________________________________________
//      |  id       |  date     |  userid   |  postid   |  content  |
//      |  INTEGER  |  INTEGER  |  INTEGER  |  INTEGER  |  TEXT     |

// Create comments table
func crerateCommentsTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS comments(id INTEGER PRIMARY KEY, date INTEGER NOT NULL, userid INTEGER NOT NULL, postid INTEGER NOT NULL, content TEXT NOT NULL)")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

func getComments(user *User, post Post) ([]Comment, error) {
	comments := []Comment{}
	sql := "SELECT comments.id, date, userid, username, postid, content, (SELECT COUNT(*) FROM comment_likes WHERE status = 1 AND commentid = comments.id) AS likes, (SELECT COUNT(*) FROM comment_likes WHERE status = -1 AND commentid = comments.id) AS dislikes, (SELECT SUM(status) from comment_likes WHERE comment_likes.commentid = comments.id AND comment_likes.userid = ?) AS status FROM comments INNER JOIN users ON userid = users.id WHERE postid = ?"
	rows, err := db.Query(sql, user.Id, post.Id)
	if err != nil {
		return comments, err
	}
	var status interface{}
	for rows.Next() {
		comment := Comment{}
		err = rows.Scan(&(comment.Id), &(comment.Date), &(comment.Userid), &(comment.Username), &(comment.Postid), &(comment.Content), &(comment.Likes), &(comment.Dislikes), &status)
		if err != nil {
			return comments, err
		}
		dateFormat := formatMilli(comment.Date)
		comment.DateFormat = dateFormat
		s, ok := status.(int64)
		if ok {
			comment.Status = int(s)
		} else {
			comment.Status = 0
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		return comments, err
	}
	return comments, nil
}
func saveComment(user *User, postId int, comment string) error {
	statement, err := db.Prepare("INSERT INTO comments (date, userid, postid, content) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	date := getCurrentMilli()
	_, err = statement.Exec(date, user.Id, postId, comment)
	if err != nil {
		return err
	}
	return nil
}

func printComments() {
	rows, err := db.Query("SELECT * FROM comments")
	if err != nil {
		return
	}
	defer rows.Close()
	comment := Comment{}
	for rows.Next() {
		err = rows.Scan(&(comment.Id), &(comment.Date), &(comment.Userid), &(comment.Postid), &(comment.Content))
		if err != nil {
			return
		}
		fmt.Println(comment)
	}
	err = rows.Err()
	if err != nil {
		return
	}
}
