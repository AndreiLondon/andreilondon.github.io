package main

import (
	//"github.com/gofrs/uuid"
	"github.com/google/uuid"
)

func generateSessionId() string {
	// u2, err := uuid.NewV4()
	// if err != nil {
	// 	milli := getCurrentMilli()
	// 	rand.Seed(milli)
	// 	sessionId := fmt.Sprintf("%d-%d", milli, rand.Intn(10000000))
	// 	return sessionId
	// }
	return uuid.New().String()
}
