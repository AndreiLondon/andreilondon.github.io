package main

import (
	"net/http"
)

// Function sets sessionId as a coockie
func setCookie(w http.ResponseWriter, sessionId string) {
	// create a new cookie
	cookie := http.Cookie{
		Name:   SESSION_ID,
		Value:  sessionId,
		MaxAge: 3600,
	}
	// set the cookie in the response header which will be sent to the client
	http.SetCookie(w, &cookie)
}

// Function gets sessionId from cookies
func getCookie(r *http.Request) string {
	cookieValue, err := r.Cookie(SESSION_ID)
	if err != nil {
		return ""
	}
	return cookieValue.Value
}
