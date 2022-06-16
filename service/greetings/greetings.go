package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// If no name was given, return an empty string.
	if name == "" {
		return ""
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(message(), name)
	return message
}

func message() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
