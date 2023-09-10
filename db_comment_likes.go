package main

//     _________comment_likes__________________________
//    |  id       |  userid   |  commentid  |  status  |
//    |  INTEGER  |  INTEGER  |  INTEGER    |  INTEGER |

// Create comment_likes table
func crerateCommentLikesTable() error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS comment_likes(id INTEGER PRIMARY KEY, userid INTEGER NOT NULL, commentid INTEGER NOT NULL, status INTEGER NOT NULL CHECK(status = 1 OR status = 0 OR status = -1))")
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()
	return nil
}

func updatePostCommentLikes(user *User, commentId int, status int) {
	//Check if user tryes to like own comment
	rows, err := db.Query("SELECT * FROM comments WHERE id = ? AND userid = ?", commentId, user.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		return
	}
	err = rows.Err()
	if err != nil {
		return
	}
	//Try to update
	statement, err := db.Prepare("UPDATE comment_likes SET status = ? WHERE userid = ? AND commentid = ?")
	if err != nil {
		return
	}
	defer statement.Close()
	result, err := statement.Exec(status, user.Id, commentId)
	if err != nil {
		return
	}
	numbOfRows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if numbOfRows == 0 {
		statement1, err := db.Prepare("INSERT INTO comment_likes (userid, commentid, status) VALUES (?,?,?)")
		if err != nil {
			return
		}
		defer statement1.Close()
		statement1.Exec(user.Id, commentId, status)
	}
}
