package main

import (
	"net/http"
)

var SESSION_ID = "SESSION_ID"

func setCookie(w http.ResponseWriter, sessionId string) {

	ck := http.Cookie{
		Name:   SESSION_ID,
		Value:  sessionId,
		MaxAge: 1000 * 60 * 5,
	}
	http.SetCookie(w, &ck)
}

func getCookie(r *http.Request) string {
	tokenCookie, err := r.Cookie(SESSION_ID)
	if err != nil {
		return ""
	}
	return tokenCookie.Value
}
