package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// var invalidCredentialsFlagSignUp = false
// var invalidCredentialsFlagSignIn = false
// var posts = []Post{}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		showError(w, 400, "400 Bad Request")
		return
	}
	// invalidCredentialsFlagSignUp = false
	// invalidCredentialsFlagSignIn = false
	invalidCredentialsFlagSignUp = ""
	invalidCredentialsFlagSignIn = ""
	//emptyPostFlag = false
	emptyCommentFlag = false
	indexObject := IndexObject{}
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	indexObject.User = user
	if user == nil {
		currentMode = SHAW_ALL
	}
	posts, err := getPosts(indexObject.User)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	posts = filterByCategories(posts, filterCategories)
	posts = filterByMode(posts, currentMode, user)
	indexObject.Filters = filterCategories
	indexObject.Posts = posts
	indexObject.Mode = currentMode
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	//templ, err := template.ParseFiles("templates/index.html")
	err = t.Execute(w, indexObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	displayUsers()
	printPosts()
	printComments()
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	username := strings.TrimSpace(r.FormValue("signup_username"))
	email := strings.TrimSpace(r.FormValue("signup_email"))
	password := strings.TrimSpace(r.FormValue("signup_password"))
	// invalidCredentialsFlagSignUp = true
	// invalidCredentialsFlagSignIn = false
	if username == "" || email == "" || password == "" || len(username) > 40 || len(password) < 3 {
		invalidCredentialsFlagSignUp = "Invalid username or password"
		invalidCredentialsFlagSignIn = ""
		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
		return
	}
	sessionId := generateSessionId()
	err := insertUser(username, email, encrypt(password), sessionId)
	if err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
			// invalidCredentialsFlagSignUp = true
			// invalidCredentialsFlagSignIn = false
			invalidCredentialsFlagSignUp = "User name or email alreay in use"
			invalidCredentialsFlagSignIn = ""
			http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
			return
		}
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}
	// t.ExecuteTemplate(w, "signup", nil)
	setCookie(w, sessionId)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	email := r.FormValue("login_email")
	password := r.FormValue("login_password")
	user, err := checkUser(email, password)
	if err != nil {
		showError(w, 500, "500 Internal Server Error. Error while working with database")
		return
	}

	if err == nil && user == nil {
		// invalidCredentialsFlagSignIn = true
		// invalidCredentialsFlagSignUp = false
		invalidCredentialsFlagSignIn = "Invalid email or password"
		invalidCredentialsFlagSignUp = ""
		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
		return
	}
	if user != nil {
		sessionId := generateSessionId()
		setCookie(w, sessionId)
		err := setSessionId(user, sessionId)
		if err != nil {
			showError(w, 500, "500 Internal Server Error")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

}

func signHandler(w http.ResponseWriter, r *http.Request) {
	// templ, err := template.ParseFiles("templates/sign.html")
	// if err != nil {
	// 	showError(w, 500, "500 Internal Server Error")
	// 	return
	// }
	t, err := template.ParseFiles("templates/sign.html")
	if err != nil {
		// fmt.Println(err)
		showError(w, 500, "500 Internal Server Error")
		return
	}

	singCredentials := SingCredentials{}
	singCredentials.SignIn = invalidCredentialsFlagSignIn
	singCredentials.SignUp = invalidCredentialsFlagSignUp
	err = t.Execute(w, singCredentials)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}

}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	if sessionId != "" {
		err := resetSessionId(sessionId)
		if err != nil {
			showError(w, 500, "500 Internal Server Error. Error while working with database")
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func createpostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createpost.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "createpost", nil)

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	t, err := template.ParseFiles("templates/post.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newPostObject := NewPostObject{}
	newPostObject.User = user
	categories, err := getCategories()
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newPostObject.Categories = categories
	newPostObject.IsEmptyPost = emptyPostFlag
	err = t.Execute(w, newPostObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}

}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	postContent := strings.TrimSpace(r.FormValue("post_content"))
	postCategories := r.FormValue("categories")
	if postContent == "" {
		emptyPostFlag = true
		http.Redirect(w, r, "/post", http.StatusTemporaryRedirect)
		return
	}
	err := savePost(*user, postContent, postCategories)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func registerlikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sessionId := getCookie(r)
		user := getUserBySessionId(sessionId)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		postIdStr := r.FormValue("postId")
		statusStr := r.FormValue("status")
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			return
		}
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			return
		}
		updatePostLikes(user, postId, status)
	}

}

func registercommentlikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sessionId := getCookie(r)
		user := getUserBySessionId(sessionId)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		commentIdStr := r.FormValue("commentId")
		statusStr := r.FormValue("status")
		commentId, err := strconv.Atoi(commentIdStr)
		if err != nil {
			return
		}
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			return
		}
		updatePostCommentLikes(user, commentId, status)
	}

}

func commentHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	t, err := template.ParseFiles("templates/comment.html")
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	postIdStr := r.FormValue("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newCommentObject := NewCommentObject{}
	newCommentObject.User = user
	post, err := getPostById(postId)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	newCommentObject.Post = post
	newCommentObject.EmptyComment = emptyCommentFlag
	err = t.Execute(w, newCommentObject)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
}

func commentsubmitHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getCookie(r)
	user := getUserBySessionId(sessionId)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	comment := strings.TrimSpace(r.FormValue("comment"))
	if comment == "" {
		emptyCommentFlag = true
		http.Redirect(w, r, "/comment", http.StatusTemporaryRedirect)
		return
	}
	postIdStr := r.FormValue("postId")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	err = saveComment(user, postId, comment)
	if err != nil {
		showError(w, 500, "500 Internal Server Error")
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func setfilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	filterCategory := r.FormValue("filterCategory")
	if !contains(filterCategories, filterCategory) {
		filterCategories = append(filterCategories, filterCategory)
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func removefilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	filterCategory := r.FormValue("filterCategory")
	for i, cat := range filterCategories {
		if cat == filterCategory {
			filterCategories = append(filterCategories[:i], filterCategories[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func changemodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		showError(w, 405, "405 Method Not Allowed")
		return
	}
	currentMode = r.FormValue("mode")
	if currentMode != MY_POSTS && currentMode != MY_COMMENTS && currentMode != MY_LIKES {
		currentMode = SHAW_ALL
	}
}

// func signUpHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	userName := strings.TrimSpace(r.FormValue("username"))
// 	email := strings.TrimSpace(r.FormValue("email"))
// 	password := strings.TrimSpace(r.FormValue("password"))
// 	if userName == "" || email == "" || password == "" || len(userName) > 40 {
// 		invalidCredentialsFlagSignUp = true
// 		invalidCredentialsFlagSignIn = false
// 		http.Redirect(w, r, "/sign", http.StatusTemporaryRedirect)
// 		return
// 	}

// 	if userName == "" || email == "" || password == "" {
// 		http.Error(w, "400 Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	sessionId := generateSessionId()

// 	// Perform user registration logic, e.g., save to database, generate session ID, etc.
// 	// Replace the placeholders with your actual implementation.
// 	sessionID, err := registerUser(userName, email, password)
// 	if err != nil {
// 		http.Error(w, "500 Internal Server Error. Error while registering user", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set session ID as a cookie for authenticated session
// 	setCookie(w, sessionID)

// 	// Redirect to homepage after successful registration
// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }
