package main

import (
	"github.com/google/uuid"
)

/*
A UUID – that's short for Universally Unique IDentifier,
by the way – is a 36-character alphanumeric string that can be used to identify information
(such as a table row).
*/

// generate a new session ID using either a UUID (if it is generated successfully)
// or a combination of the current time in milliseconds and a random integer (if the UUID generation fails)
func generateSessionId() string {
	return uuid.New().String()
}
