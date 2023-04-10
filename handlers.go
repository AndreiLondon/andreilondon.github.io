package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/signup.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "signup", nil)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "login", nil)

}

func createpostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createpost.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "createpost", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "index", nil)
}
func savePostHanlder(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/post" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	title := r.FormValue("title")
	text := r.FormValue("text")

	// insertPost(*user, postContent, postCategories)
	//insertPost(*user, title, text)
	insertPost(title, text)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
