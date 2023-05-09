package main

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	Sessionid string
}
type Post struct {
	Id         int
	Userid     int
	Date       int
	DateFormat string
	Content    string
	Categories []string
	Comments   []Comment
	Username   string
	Likes      int
	Dislikes   int
	Status     int
}
type Comment struct {
	Id         int
	Date       int
	DateFormat string
	Userid     int
	Postid     int
	Content    string
	Username   string
	Likes      int
	Dislikes   int
	Status     int
}
type PostLike struct {
	Id     int
	Userid int
	Postid int
	Status int
}
type CommentLike struct {
	Id        int
	Userid    int
	Commentid int
	Status    int
}
type NewPostObject struct {
	User        *User
	Categories  []string
	IsEmptyPost bool
}
type NewCommentObject struct {
	User         *User
	Post         *Post
	EmptyComment bool
}
type IndexObject struct {
	User    *User
	Posts   []Post
	Filters []string
	Mode    string
}
type SingCredentials struct {
	SignIn string
	SignUp string
}
