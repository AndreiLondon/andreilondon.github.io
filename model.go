package main

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	SessionId string
}

// type Post struct {
// 	Id         int
// 	Userid     int
// 	Date       int
// 	DateFormat string
// 	Content    string
// 	Categories []string
// 	Comments   []Comment
// 	Username   string
// 	Likes      int
// 	Dislikes   int
// 	Status     int
// }

type Post struct {
	Id    int
	Title string
	Text  string
}
